package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/models"
	repo "github.com/nas03/scholar-ai/backend/internal/repositories"
	"github.com/nas03/scholar-ai/backend/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(ctx context.Context, username, password, email string) int
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
	// Generate user's uuid
	userUUID, err := uuid.NewRandom()
	if err != nil {
		global.Log.Sugar().Errorw("Error creating UUID", "error", err)
		return response.CodeFailedCreateUser
	}

	// Hash user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		global.Log.Sugar().Errorw("Error generating hashedPassword", "error", err)
		return response.CodeFailedCreateUser
	}

	user := &models.User{
		UserID:   userUUID.String(),
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		global.Log.Sugar().Errorw("Error creating new user", "error", err)
		return response.CodeFailedCreateUser
	}

	global.Log.Sugar().Infow("Success creating new user", "userID", userUUID)
	return response.CodeSuccess
}

func (s *UserService) VerifyUserEmail(ctx context.Context, otp, email string) int {

	return response.CodeSuccess
}
