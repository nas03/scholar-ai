package config

import (
	"fmt"
	"os"
)

// LoadMailConfig loads mail configuration from environment variables
func LoadMailConfig() (*MailConfig, error) {
	config := &MailConfig{
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}

	// Validate required fields
	if config.Username == "" {
		return nil, fmt.Errorf("MAIL_USERNAME environment variable is required")
	}
	if config.Password == "" {
		return nil, fmt.Errorf("MAIL_PASSWORD environment variable is required")
	}

	return config, nil
}

// GetMailConfig returns mail configuration (deprecated - use LoadMailConfig instead)
func GetMailConfig() *MailConfig {
	config := &MailConfig{
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}

	return config
}
