package main

import (
	"log"
	"os"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/config"
	"github.com/nas03/scholar-ai/backend/internal/initialize"
)

// @title Scholar AI Backend API
// @version 0.1.0
// @description REST API for Scholar AI backend services.
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token as: Bearer <token>
func main() {
	// Initialize logger
	if err := initialize.InitLogger(); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer initialize.SyncLogger()

	// Initialize database connection
	if err := initialize.InitGorm(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to initialize database", "error", err)
		}
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize router with middleware and routes
	r := initialize.InitRouter()

	// Load server configuration
	serverConfig := config.LoadServerConfig()
	address := serverConfig.GetAddress()

	// Start server
	if global.Log != nil {
		global.Log.Sugar().Infow("Starting server", "address", address, "pid", os.Getpid())
	} else {
		log.Printf("Starting server on %s", address)
	}
	if err := r.Run(address); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to start server", "error", err)
		}
		log.Fatal("Failed to start server:", err)
	}
}
