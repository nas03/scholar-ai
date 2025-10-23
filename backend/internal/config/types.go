package config

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// MailConfig holds mail service configuration
type MailConfig struct {
	Username string
	Password string
}
