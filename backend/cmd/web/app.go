package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/config"
	"github.com/nas03/scholar-ai/backend/internal/initialize"
)

// App represents the application instance
type App struct {
	Router       *gin.Engine
	ServerConfig *config.ServerConfig
}

// NewApp creates and initializes a new application instance
func NewApp() (*App, error) {
	// Load .env file (ignore error if not found - env vars may be set elsewhere)
	_ = godotenv.Load()

	// Enable Gin console colors
	gin.ForceConsoleColor()

	// Initialize router with middleware and routes
	router := initialize.InitRouter()

	// Load server configuration
	serverConfig := config.LoadServerConfig()

	return &App{
		Router:       router,
		ServerConfig: serverConfig,
	}, nil
}

// Run starts the application server
func (a *App) Run() error {
	address := a.ServerConfig.GetAddress()

	// Log server startup
	if global.Log != nil {
		global.Log.Sugar().Infow("Starting server", "address", address, "pid", os.Getpid())
	} else {
		log.Printf("Starting server on %s", address)
	}

	// Start server
	return a.Router.Run(address)
}
