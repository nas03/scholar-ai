package main

import (
	"log"

	"github.com/nas03/scholar-ai/backend/global"
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
	// Bootstrap all services
	if err := Bootstrap(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to bootstrap application", "error", err)
		}
		log.Fatal("Failed to bootstrap application:", err)
	}

	// Create and run the application
	app, err := NewApp()
	if err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to create application", "error", err)
		}
		log.Fatal("Failed to create application:", err)
	}

	// Start the server
	if err := app.Run(); err != nil {
		if global.Log != nil {
			global.Log.Sugar().Fatalw("Failed to start server", "error", err)
		}
		log.Fatal("Failed to start server:", err)
	}
}
