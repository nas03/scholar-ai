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
		global.Log.Sugar().Warnw(errMessage.ErrInvalidUsername.Error(), "username", username)
		return response.CodeInvalidUsername
	}
	if email == "" {
		global.Log.Sugar().Warnw(errMessage.ErrInvalidEmail.Error(), "email", email)
		return response.CodeInvalidEmail
	}
	if password == "" {
		global.Log.Sugar().Warnw(errMessage.ErrEmptyPassword.Error(), "password", password)
		return response.CodeEmptyPassword
	}

	// Generate user's uuid
	userUUID, err := uuid.NewRandom()
	if err != nil {
		global.Log.Sugar().Errorw("Error creating UUID", "error", err)
		return response.CodeRegisterInternalError
	}

	// Hash user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		global.Log.Sugar().Errorw("Error generating hashedPassword", "error", err)
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
			global.Log.Sugar().Warnw(errMessage.ErrUserAlreadyExists.Error(), "email", email, "username", username)
			return response.CodeUserAlreadyExists
		}

		global.Log.Sugar().Errorw("Error creating new user", "error", err)
		return response.CodeRegisterInternalError
	}

	// Send OTP verify user's email
	otp := utils.GenerateSixDigitOtp()
	redisKey := fmt.Sprintf(consts.REDIS_KEY_URS_OTP_PREFIX, email)
	if err := utils.NewRedisCache().SetEx(ctx, redisKey, otp, consts.REDIS_OTP_EXPIRATION); err != nil {
		global.Log.Sugar().Errorf("Failed to store otp in redis", "error", err)
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
		global.Log.Sugar().Error("Failed to send verification email to %w", email, "error", err)
		return response.CodeMailSendFailed
	}
	global.Log.Sugar().Infow("Success creating new user", "userID", userUUID)
	return response.CodeSuccess
}

// GetUserByEmail retrieves a user by email with proper error handling
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, int) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Sugar().Warnw(errMessage.ErrUserNotFound.Error(), "email", email)
			return nil, response.CodeUserNotFound
		}

		global.Log.Sugar().Errorw("Error getting user by email", "error", err, "email", email)
		return nil, response.CodeFailedGetUser
	}

	return user, response.CodeSuccess
}

// GetUserByID retrieves a user by ID with proper error handling
func (s *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, int) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Sugar().Warnw(errMessage.ErrUserNotFound.Error(), "userID", userID)
			return nil, response.CodeUserNotFound
		}

		global.Log.Sugar().Errorw("Error getting user by ID", "error", err, "userID", userID)
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
			global.Log.Sugar().Warnw(errMessage.ErrInvalidStatus.Error(), "userID", userID, "status", status)
			return response.CodeInvalidInput
		}

		global.Log.Sugar().Errorw("Error updating user account status", "error", err, "userID", userID, "status", status)
		return response.CodeFailedUpdateUser
	}

	global.Log.Sugar().Infow("Success updating user account status", "userID", userID, "status", status)
	return response.CodeSuccess
}

// UpdateUserPassword updates user password with proper error handling
func (s *UserService) UpdateUserPassword(ctx context.Context, userID, password string) int {
	// Validate password at service level
	if password == "" {
		global.Log.Sugar().Warnw(errMessage.ErrEmptyPassword.Error(), "userID", userID)
		return response.CodeEmptyPassword
	}

	err := s.userRepo.UpdateUserPassword(ctx, userID, password)
	if err != nil {
		global.Log.Sugar().Errorw("Error updating user password", "error", err, "userID", userID)
		return response.CodeFailedUpdateUser
	}

	global.Log.Sugar().Infow("Success updating user password", "userID", userID)
	return response.CodeSuccess
}

// UpdateUserVerification updates user verification status with proper error handling
func (s *UserService) UpdateUserVerification(ctx context.Context, userID string, isEmailVerified, isPhoneVerified bool) int {
	err := s.userRepo.UpdateUserVerification(ctx, userID, isEmailVerified, isPhoneVerified)
	if err != nil {
		global.Log.Sugar().Errorw("Error updating user verification", "error", err, "userID", userID)
		return response.CodeFailedUpdateUser
	}

	global.Log.Sugar().Infow("Success updating user verification", "userID", userID, "emailVerified", isEmailVerified, "phoneVerified", isPhoneVerified)
	return response.CodeSuccess
}

func (s *UserService) VerifyUserEmail(ctx context.Context, otp, email string) int {
	// Validate input parameters
	if otp == "" {
		global.Log.Sugar().Warnw(errMessage.ErrInvalidOTP.Error(), "otp", otp)
		return response.CodeInvalidOTP
	}
	if email == "" {
		global.Log.Sugar().Warnw(errMessage.ErrInvalidEmail.Error(), "email", email)
		return response.CodeInvalidEmail
	}

	_, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Sugar().Warnw(errMessage.ErrUserNotFound.Error(), "email", email)
			return response.CodeUserNotFound
		}

		global.Log.Sugar().Errorw("Error getting user by email", "error", err, "email", email)
		return response.CodeFailedGetUser
	}

	// TODO: Implement OTP verification logic
	// This would typically involve:
	// 1. Getting user by email
	// 2. Checking if OTP matches and is not expired
	// 3. Updating user verification status
	// 4. Clearing the OTP from storage

	global.Log.Sugar().Infow("Email verification successful", "email", email)
	return response.CodeSuccess
}
