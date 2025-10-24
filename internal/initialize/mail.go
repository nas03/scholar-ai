package initialize

import (
	"fmt"
	"log"

	"github.com/nas03/scholar-ai/backend/global"
	mailErrors "github.com/nas03/scholar-ai/backend/pkg/errors"
	"github.com/nas03/scholar-ai/backend/pkg/setting"
	mail "github.com/wneessen/go-mail"
)

// MailConfig holds mail service configuration

// LoadMailConfig loads mail configuration from global config
func LoadMailConfig() (*setting.MailSetting, error) {
	config := &setting.MailSetting{
		Username: global.Config.Mail.Username,
		Password: global.Config.Mail.Password,
	}

	// Validate required fields
	if config.Username == "" {
		return nil, fmt.Errorf("mail username is required")
	}
	if config.Password == "" {
		return nil, fmt.Errorf("mail password is required")
	}

	return config, nil
}

func InitGoMail() {
	config, err := LoadMailConfig()
	if err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw(mailErrors.ErrMailConfigMissing.Error(), "error", err)
		} else {
			log.Printf("Failed to load mail config: %v", err)
		}
		return
	}

	// Create mail client
	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(config.Username),
		mail.WithPassword(config.Password),
	)
	if err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw(mailErrors.ErrMailClientCreation.Error(), "error", err, "username", config.Username)
		} else {
			log.Printf("Failed to create mail client: %v", err)
		}
		return
	}

	// Store the client globally
	global.Mail = client
	if global.Log != nil {
		global.Log.Sugar().Infow("Mail service initialized successfully", "username", config.Username)
	} else {
		log.Printf("Mail service initialized successfully for user: %s", config.Username)
	}
}
