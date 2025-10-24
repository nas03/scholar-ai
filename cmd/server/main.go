package main

import (
	"log"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/initialize"
)

func main() {
	// Bootstrap all services
	if err := initialize.Bootstrap(); err != nil {
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
