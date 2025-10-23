package config

import (
	"fmt"
	"os"
)

// LoadDatabaseConfig loads database configuration from environment variables
func LoadDatabaseConfig() (*DatabaseConfig, error) {
	config := &DatabaseConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	// Set defaults
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "3306"
	}

	// Validate required fields
	if config.Username == "" {
		return nil, fmt.Errorf("DB_USERNAME environment variable is required")
	}
	if config.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD environment variable is required")
	}
	if config.Name == "" {
		return nil, fmt.Errorf("DB_NAME environment variable is required")
	}

	return config, nil
}

// GetDSN constructs the database DSN from config
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		c.Username, c.Password, c.Host, c.Port, c.Name)
}
