package setting

// Config holds all configuration settings for the application
type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	Database DatabaseSetting `mapstructure:"database"`
	Log      LogSetting      `mapstructure:"log"`
	Redis    RedisSetting    `mapstructure:"redis"`
	Resend   ResendSetting   `mapstructure:"resend"`
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
type ResendSetting struct {
	ApiKey string `mapstructure:"api_key"`
	From   string `mapstructure:"from"`
}

// RedisSetting holds redis configuration
type RedisSetting struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}
