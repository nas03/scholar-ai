package main

import (
	"log"

	"github.com/nas03/scholar-ai/backend/internal/config"
	"github.com/nas03/scholar-ai/backend/internal/initialize"
)

func main() {
	// Initialize database connection
	if err := initialize.InitGorm(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize router with middleware and routes
	r := initialize.InitRouter()

	// Load server configuration
	serverConfig := config.LoadServerConfig()
	address := serverConfig.GetAddress()

	// Start server
	log.Printf("Starting server on %s", address)
	if err := r.Run(address); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
