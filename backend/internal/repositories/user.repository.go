package repositories

import (
	"fmt"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/models"
)

type IUserRepository interface {
	CreateUser(userID, username, password, email string) error
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) CreateUser(userID, username, password, email string) error {
	user := &models.User{
		UserID:   userID,
		Username: username,
		Password: password,
		Email:    email,
	}

	if err := global.Mdb.Create(user).Error; err != nil {
		return fmt.Errorf("error creating new user: %w", err)
	}
	return nil
}
