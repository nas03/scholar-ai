package config

import "os"

// LoadServerConfig loads server configuration from environment variables
func LoadServerConfig() *ServerConfig {
	config := &ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
		Host: os.Getenv("SERVER_HOST"),
	}

	// Set defaults
	if config.Port == "" {
		config.Port = "8080"
	}
	if config.Host == "" {
		config.Host = ""
	}

	return config
}

// GetAddress constructs the server address from config
func (c *ServerConfig) GetAddress() string {
	if c.Host == "" {
		return ":" + c.Port
	}
	return c.Host + ":" + c.Port
}
