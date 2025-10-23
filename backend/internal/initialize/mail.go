package initialize

import (
	"fmt"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/config"
	mailErrors "github.com/nas03/scholar-ai/backend/pkg/errors"
	mail "github.com/wneessen/go-mail"
)

func InitGoMail() error {
	config, err := config.LoadMailConfig()
	if err != nil {
		global.Log.Sugar().Fatalw(mailErrors.ErrMailConfigMissing.Error(), "error", err)
		return fmt.Errorf("failed to load mail config: %w", err)
	}

	// Create mail client
	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(config.Username),
		mail.WithPassword(config.Password),
	)
	if err != nil {
		global.Log.Sugar().Errorw(mailErrors.ErrMailClientCreation.Error(), "error", err, "username", config.Username)
		return fmt.Errorf("failed to create mail client: %w", err)
	}

	// Store the client globally
	global.Mail = client
	global.Log.Sugar().Infow("Mail service initialized successfully", "username", config.Username)
	return nil
}
