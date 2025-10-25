package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/consts"
	"github.com/nas03/scholar-ai/backend/internal/helper"
	"github.com/nas03/scholar-ai/backend/internal/models"
	repo "github.com/nas03/scholar-ai/backend/internal/repositories"
	"github.com/nas03/scholar-ai/backend/internal/utils"
	errMessage "github.com/nas03/scholar-ai/backend/pkg/errors"
	"github.com/nas03/scholar-ai/backend/pkg/response"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	CreateUser(ctx context.Context, username, password, email string) int
	GetUserByEmail(ctx context.Context, email string) (*models.User, int)
	GetUserByID(ctx context.Context, userID string) (*models.User, int)
	UpdateUserAccountStatus(ctx context.Context, userID string, status int8) int
	UpdateUserPassword(ctx context.Context, userID, password string) int
	UpdateUserVerification(ctx context.Context, userID string, isEmailVerified, isPhoneVerified bool) int
	// UpdateUserInfo(email, phoneNumber string) int
	VerifyUserEmail(ctx context.Context, otp, email string) int
}

type UserService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepository repo.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username, password, email string) int {
	// Validate input parameters
	if username == "" {
		global.Log.Warn(errMessage.ErrInvalidUsername.Error(), zap.String("username", username))
		return response.CodeInvalidUsername
	}
	if email == "" {
		global.Log.Warn(errMessage.ErrInvalidEmail.Error(), zap.String("email", email))
		return response.CodeInvalidEmail
	}
	if password == "" {
		global.Log.Warn(errMessage.ErrEmptyPassword.Error(), zap.String("password", password))
		return response.CodeEmptyPassword
	}

	// Generate user's uuid
	userUUID, err := uuid.NewRandom()
	if err != nil {
		global.Log.Error("Error creating UUID", zap.Error(err))
		return response.CodeRegisterInternalError
	}

	// Hash user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		global.Log.Error("Error generating hashedPassword", zap.Error(err))
		return response.CodeRegisterInternalError
	}

	user := &models.User{
		UserID:   userUUID.String(),
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	// Create user
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			global.Log.Warn(errMessage.ErrUserAlreadyExists.Error(), zap.String("email", email), zap.String("username", username))
			return response.CodeUserAlreadyExists
		}

		global.Log.Error("Error creating new user", zap.Error(err))
		return response.CodeRegisterInternalError
	}

	// Send OTP verify user's email
	otp := utils.GenerateSixDigitOtp()
	redisKey := fmt.Sprintf(consts.REDIS_KEY_URS_OTP_PREFIX, email)
	if err := utils.NewRedisCache().SetEx(ctx, redisKey, otp, consts.REDIS_OTP_EXPIRATION); err != nil {
		global.Log.Error("Failed to store otp in redis", zap.Error(err))
		return response.CodeRegisterInternalError
	}

	// TODO: Should save mailID, email, email type to DB
	_, err = helper.NewMailHelper().SendMail(
		ctx,
		email,
		fmt.Sprintf("ScholarAI Verification Code %d", otp),
		fmt.Sprintf("<p>%d</p>", otp),
	)
	if err != nil {
		global.Log.Error("Failed to send verification email", zap.String("email", email), zap.Error(err))
		return response.CodeMailSendFailed
	}
	global.Log.Info("Success creating new user", zap.String("userID", userUUID.String()))
	return response.CodeSuccess
}

// GetUserByEmail retrieves a user by email with proper error handling
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, int) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Warn(errMessage.ErrUserNotFound.Error(), zap.String("email", email))
			return nil, response.CodeUserNotFound
		}

		global.Log.Error("Error getting user by email", zap.Error(err), zap.String("email", email))
		return nil, response.CodeFailedGetUser
	}

	return user, response.CodeSuccess
}

// GetUserByID retrieves a user by ID with proper error handling
func (s *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, int) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Warn(errMessage.ErrUserNotFound.Error(), zap.String("userID", userID))
			return nil, response.CodeUserNotFound
		}

		global.Log.Error("Error getting user by ID", zap.Error(err), zap.String("userID", userID))
		return nil, response.CodeFailedGetUser
	}

	return user, response.CodeSuccess
}

// UpdateUserAccountStatus updates user account status with proper error handling
func (s *UserService) UpdateUserAccountStatus(ctx context.Context, userID string, status int8) int {
	err := s.userRepo.UpdateUserAccountStatus(ctx, userID, status)
	if err != nil {
		// Check if it's a validation error from repository
		if errors.Is(err, errMessage.ErrInvalidStatus) {
			global.Log.Warn(errMessage.ErrInvalidStatus.Error(), zap.String("userID", userID), zap.Int8("status", status))
			return response.CodeInvalidInput
		}

		global.Log.Error("Error updating user account status", zap.Error(err), zap.String("userID", userID), zap.Int8("status", status))
		return response.CodeFailedUpdateUser
	}

	global.Log.Info("Success updating user account status", zap.String("userID", userID), zap.Int8("status", status))
	return response.CodeSuccess
}

// UpdateUserPassword updates user password with proper error handling
func (s *UserService) UpdateUserPassword(ctx context.Context, userID, password string) int {
	// Validate password at service level
	if password == "" {
		global.Log.Warn(errMessage.ErrEmptyPassword.Error(), zap.String("userID", userID))
		return response.CodeEmptyPassword
	}

	err := s.userRepo.UpdateUserPassword(ctx, userID, password)
	if err != nil {
		global.Log.Error("Error updating user password", zap.Error(err), zap.String("userID", userID))
		return response.CodeFailedUpdateUser
	}

	global.Log.Info("Success updating user password", zap.String("userID", userID))
	return response.CodeSuccess
}

// UpdateUserVerification updates user verification status with proper error handling
func (s *UserService) UpdateUserVerification(ctx context.Context, userID string, isEmailVerified, isPhoneVerified bool) int {
	err := s.userRepo.UpdateUserVerification(ctx, userID, isEmailVerified, isPhoneVerified)
	if err != nil {
		global.Log.Error("Error updating user verification", zap.Error(err), zap.String("userID", userID))
		return response.CodeFailedUpdateUser
	}

	global.Log.Info("Success updating user verification", zap.String("userID", userID), zap.Bool("emailVerified", isEmailVerified), zap.Bool("phoneVerified", isPhoneVerified))
	return response.CodeSuccess
}

func (s *UserService) VerifyUserEmail(ctx context.Context, otp, email string) int {
	// Validate input parameters
	if otp == "" {
		global.Log.Warn(errMessage.ErrInvalidOTP.Error(), zap.String("otp", otp))
		return response.CodeInvalidOTP
	}
	if email == "" {
		global.Log.Warn(errMessage.ErrInvalidEmail.Error(), zap.String("email", email))
		return response.CodeInvalidEmail
	}

	_, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Warn(errMessage.ErrUserNotFound.Error(), zap.String("email", email))
			return response.CodeUserNotFound
		}

		global.Log.Error("Error getting user by email", zap.Error(err), zap.String("email", email))
		return response.CodeFailedGetUser
	}

	// TODO: Implement OTP verification logic
	// This would typically involve:
	// 1. Getting user by email
	// 2. Checking if OTP matches and is not expired
	// 3. Updating user verification status
	// 4. Clearing the OTP from storage

	global.Log.Info("Email verification successful", zap.String("email", email))
	return response.CodeSuccess
}
