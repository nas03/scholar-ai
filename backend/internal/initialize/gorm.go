package initialize

import (
	"fmt"
	"log"

	"github.com/nas03/scholar-ai/backend/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

// GetDSN constructs the database DSN from config
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		c.Username, c.Password, c.Host, c.Port, c.Name)
}

// LoadDatabaseConfig loads database configuration from global config
func LoadDatabaseConfig() (*DatabaseConfig, error) {
	config := &DatabaseConfig{
		Username: global.Config.Database.Username,
		Password: global.Config.Database.Password,
		Host:     global.Config.Database.Host,
		Port:     fmt.Sprintf("%d", global.Config.Database.Port),
		Name:     global.Config.Database.Name,
	}

	// Validate required fields
	if config.Username == "" {
		return nil, fmt.Errorf("database username is required")
	}
	if config.Password == "" {
		return nil, fmt.Errorf("database password is required")
	}
	if config.Name == "" {
		return nil, fmt.Errorf("database name is required")
	}

	return config, nil
}

// InitGorm initializes and returns a GORM database connection
func InitGorm() {
	// Load database configuration
	dbConfig, err := LoadDatabaseConfig()
	if err != nil {
		if global.Log != nil {
			global.Log.Sugar().Errorw("Failed to load database config", "error", err)
		} else {
			log.Printf("Failed to load database config: %v", err)
		}
		return
	}

	// Configure GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Open database connection
	db, err := gorm.Open(mysql.Open(dbConfig.GetDSN()), gormConfig)
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

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

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
