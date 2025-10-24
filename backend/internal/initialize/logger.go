package initialize

import (
	"log"

	"github.com/nas03/scholar-ai/backend/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes a global zap logger based on configuration
func InitLogger() {
	env := global.Config.Log.AppEnv
	levelStr := global.Config.Log.Level

	var cfg zap.Config
	if env == "dev" || env == "development" { // human-friendly console logs in dev
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.000")
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		cfg.EncoderConfig.ConsoleSeparator = " | "
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
		log.Printf("Failed to initialize logger: %v", err)
		return
	}

	global.Log = logger
	log.Println("Logger initialized successfully")
}

// SyncLogger flushes any buffered log entries.
func SyncLogger() {
	if global.Log != nil {
		_ = global.Log.Sync()
	}
}
