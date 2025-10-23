package main

import (
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/initialize"
)

// Bootstrap initializes all required services for the application
func Bootstrap() error {
	// Initialize logger first
	if err := initialize.InitLogger(); err != nil {
		return err
	}
	defer initialize.SyncLogger()

	// Initialize database connection
	if err := initializeDatabase(); err != nil {
		return err
	}

	// Initialize mail service
	if err := initializeMail(); err != nil {
		return err
	}

	return nil
}

// initializeDatabase initializes the database connection
func initializeDatabase() error {
	if err := initialize.InitGorm(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to initialize database", "error", err)
		}
		return err
	}
	return nil
}

// initializeMail initializes the mail service
func initializeMail() error {
	if err := initialize.InitGoMail(); err != nil {
		global.Log.Sugar().Fatalw("Failed to initialize go-mail:", "error", err)
		return err
	}

	global.Log.Sugar().Infow("Success config mail service")
	return nil
}
