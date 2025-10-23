package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/nas03/scholar-ai/backend/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)

	// UpdateUserAccountStatus(ctx context.Context, status int)
}

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository with the given database connection.
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database.
// Returns an error if the username or email already exists.
func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		// Check for duplicate key error using GORM's built-in error type
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("user with this username or email already exists")
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUserByEmail retrieves a user by email address.
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// GetUserByID retrieves a user by ID.
func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
