package initialize

import (
	"fmt"
	"log"
	"time"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDSN constructs the database DSN from DatabaseSetting
func GetDSN(config *setting.DatabaseSetting) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		config.Username, config.Password, config.Host, config.Port, config.Name)
}

// InitGorm initializes and returns a GORM database connection
func InitGorm() {
	// Use database configuration directly from global config
	dbConfig := &global.Config.Database

	// Validate required fields
	if dbConfig.Username == "" {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Database username is required")
		} else {
			log.Printf("Database username is required")
		}
		return
	}
	if dbConfig.Password == "" {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Database password is required")
		} else {
			log.Printf("Database password is required")
		}
		return
	}
	if dbConfig.Name == "" {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Database name is required")
		} else {
			log.Printf("Database name is required")
		}
		return
	}

	// Configure GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Open database connection
	db, err := gorm.Open(mysql.Open(GetDSN(dbConfig)), gormConfig)
	if err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Failed to open database connection", "error", err)
		} else {
			log.Printf("Failed to open database connection: %v", err)
		}
		return
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Failed to get underlying sql.DB", "error", err)
		} else {
			log.Printf("Failed to get underlying sql.DB: %v", err)
		}
		return
	}

	// Configure connection pool using values from config
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	if dbConfig.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Second)
	}

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Failed to ping database", "error", err)
		} else {
			log.Printf("Failed to ping database: %v", err)
		}
		return
	}

	// Set global database instance
	global.Mdb = db

	if global.Log != nil {
		global.Log.Info("Database connection established successfully")
	} else {
		log.Println("Database connection established successfully")
	}
}
