package initialize

import (
	"os"

	"github.com/nas03/scholar-ai/backend/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes a global zap logger based on environment variables.
// APP_ENV: dev|prod (default: prod)
// LOG_LEVEL: debug|info|warn|error (default: info)
func InitLogger() error {
	env := os.Getenv("APP_ENV")
	levelStr := os.Getenv("LOG_LEVEL")

	var cfg zap.Config
	if env == "dev" || env == "development" { // human-friendly console logs in dev
		cfg = zap.NewDevelopmentConfig()
		// Lower noise: timestamps in ISO8601
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else { // production defaults to JSON
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = "ts"
	}

	if levelStr != "" {
		var lvl zapcore.Level
		if err := lvl.UnmarshalText([]byte(levelStr)); err == nil {
			cfg.Level = zap.NewAtomicLevelAt(lvl)
		}
	}

	logger, err := cfg.Build()
	if err != nil {
		return err
	}

	global.Log = logger
	return nil
}

// SyncLogger flushes any buffered log entries.
func SyncLogger() {
	if global.Log != nil {
		_ = global.Log.Sync()
	}
}
