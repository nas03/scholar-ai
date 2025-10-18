package repositories

import (
	"fmt"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(userID, username, password, email string) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		db: global.Mdb,
	}
}

func (r *UserRepository) CreateUser(userID, username, password, email string) error {
	user := &models.User{
		UserID:   userID,
		Username: username,
		Password: password,
		Email:    email,
	}

	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("error creating new user: %w", err)
	}
	return nil
}
