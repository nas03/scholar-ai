package repositories

import (
	"context"
	"time"

	"github.com/nas03/scholar-ai/backend/internal/models"
	"gorm.io/gorm"
)

// TransactionManager provides simple transaction utilities
type TransactionManager struct {
	db *gorm.DB
}

// NewTransactionManager creates a new transaction manager
func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

// WithTransaction executes a function within a database transaction
func (tm *TransactionManager) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

// WithTimeoutTransaction executes a function within a transaction with timeout
func (tm *TransactionManager) WithTimeoutTransaction(ctx context.Context, timeout time.Duration, fn func(tx *gorm.DB) error) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

// Example usage functions for common transaction patterns

// CreateUserWithProfile creates a user with additional profile information in a transaction
func CreateUserWithProfile(db *gorm.DB, ctx context.Context, user *models.User, profileData map[string]interface{}) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create user
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// Add any additional profile data here
		// For example, create user preferences, settings, etc.
		if profileData != nil {
			// Example: Create user preferences
			// preferences := &models.UserPreferences{
			//     UserID: user.UserID,
			//     // ... other fields from profileData
			// }
			// if err := tx.Create(preferences).Error; err != nil {
			//     return err
			// }
		}

		return nil
	})
}

// UpdateUserWithVerification updates user and verification status in a transaction
func UpdateUserWithVerification(db *gorm.DB, ctx context.Context, userID string, updates map[string]interface{}, isEmailVerified, isPhoneVerified bool) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Update user fields
		if len(updates) > 0 {
			if err := tx.Model(&models.User{}).Where("user_id = ?", userID).Updates(updates).Error; err != nil {
				return err
			}
		}

		// Update verification status
		if err := tx.Model(&models.User{}).Where("user_id = ?", userID).
			Updates(map[string]interface{}{
				"is_email_verified": isEmailVerified,
				"is_phone_verified": isPhoneVerified,
			}).Error; err != nil {
			return err
		}

		return nil
	})
}
