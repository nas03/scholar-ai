package services

import (
	"github.com/google/uuid"
	repo "github.com/nas03/scholar-ai/backend/internal/repositories"
	"github.com/nas03/scholar-ai/backend/pkg/response"
)

type IUserService interface {
	CreateUser(username, password, email string) int
	VerifyUserEmail(otp, email string) int
}

type UserService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepository repo.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepository,
	}
}

func (s *UserService) CreateUser(username, password, email string) int {
	userUUID, err := uuid.NewRandom()
	if err != nil {
		return response.CodeFailedCreateUser
	}

	err = s.userRepo.CreateUser(userUUID.String(), username, password, email)
	if err != nil {
		return response.CodeFailedCreateUser
	}
	return response.CodeSuccess
}

func (s *UserService) VerifyUserEmail(otp, email string) int {

	return response.CodeSuccess
}
