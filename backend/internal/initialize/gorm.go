package initialize

import (
	"log"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitGorm initializes and returns a GORM database connection
func InitGorm() error {
	// Load database configuration
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		return err
	}

	// Configure GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Open database connection
	db, err := gorm.Open(mysql.Open(dbConfig.GetDSN()), gormConfig)
	if err != nil {
		return err
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	// Set global database instance
	global.Mdb = db

	log.Println("Database connection established successfully")
	return nil
}
