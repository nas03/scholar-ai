package initialize

import (
	"log"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/spf13/viper"
)

// LoadConfig loads configuration from YAML files using viper
func LoadConfig() {
	v := viper.New()

	// Set config file path and name
	v.AddConfigPath("./config")
	v.SetConfigName("development") // Default to development
	v.SetConfigType("yaml")

	// Allow environment variables to override config values
	v.AutomaticEnv()

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Failed to read config file", "error", err)
		} else {
			log.Printf("Failed to read config file: %v", err)
		}
		return
	}

	// Unmarshal config into global Config struct
	if err := v.Unmarshal(&global.Config); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Failed to unmarshal config", "error", err)
		} else {
			log.Printf("Failed to unmarshal config: %v", err)
		}
		return
	}

	// Log successful config loading
	if global.Log != nil {
		global.Log.Sugar().Infow("Configuration loaded successfully",
			"server_port", global.Config.Server.Port,
			"database_host", global.Config.Database.Host,
			"log_level", global.Config.Log.Level)
	} else {
		log.Printf("Configuration loaded successfully - Server: %d, DB: %s, Log: %s",
			global.Config.Server.Port, global.Config.Database.Host, global.Config.Log.Level)
	}
}
