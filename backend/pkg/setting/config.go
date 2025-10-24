package setting

// Config holds all configuration settings for the application
type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	Database DatabaseSetting `mapstructure:"database"`
	Log      LogSetting      `mapstructure:"log"`
	Mail     MailSetting     `mapstructure:"mail"`
}

// ServerSetting holds server configuration
type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

// DatabaseSetting holds database configuration
type DatabaseSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Name            string `mapstructure:"name"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// LogSetting holds logging configuration
type LogSetting struct {
	Level  string `mapstructure:"level"`
	AppEnv string `mapstructure:"app_env"`
}

// MailSetting holds mail service configuration
type MailSetting struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SMTPHost string `mapstructure:"smtp_host"`
	SMTPPort int    `mapstructure:"smtp_port"`
}
