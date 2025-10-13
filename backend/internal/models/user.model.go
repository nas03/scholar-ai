package models

import "time"

type User struct {
	UserID      string    `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username    string    `gorm:"uniqueIndex;not null" json:"username"`
	Email       string    `gorm:"uniqueIndex;not null" json:"email"`
	PhoneNumber string    `gorm:"uniqueIndex" json:"phone_number"`
	Password    string    `gorm:"not null" json:"password"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
