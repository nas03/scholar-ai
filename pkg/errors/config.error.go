package errors

import "errors"

var (
	// Mail related errors
	ErrMailConfigMissing    = errors.New("mail configuration is missing")
	ErrMailUsernameMissing  = errors.New("mail username is required")
	ErrMailPasswordMissing  = errors.New("mail password is required")
	ErrMailClientCreation   = errors.New("failed to create mail client")
	ErrMailConnectionFailed = errors.New("failed to connect to mail server")
	ErrMailSendFailed       = errors.New("failed to send email")
)
