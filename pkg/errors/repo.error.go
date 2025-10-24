package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidStatus     = errors.New("invalid account status")
	ErrInvalidInput      = errors.New("invalid input parameters")
	ErrDatabaseError     = errors.New("database operation failed")
	ErrEmptyPassword     = errors.New("password cannot be empty")
	ErrInvalidOTP        = errors.New("invalid OTP provided")
	ErrOTPExpired        = errors.New("OTP has expired")
	ErrEmailNotVerified  = errors.New("email not verified")
	ErrPhoneNotVerified  = errors.New("phone number not verified")
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrInvalidUsername   = errors.New("invalid username format")
)
