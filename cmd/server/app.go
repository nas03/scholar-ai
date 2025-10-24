package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/initialize"
)

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// GetAddress constructs the server address from config
func (c *ServerConfig) GetAddress() string {
	if c.Host == "" {
		return ":" + c.Port
	}
	return c.Host + ":" + c.Port
}

// LoadServerConfig loads server configuration from global config
func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: fmt.Sprintf("%d", global.Config.Server.Port),
		Host: global.Config.Server.Host,
	}
}

// App represents the application instance
type App struct {
	Router       *gin.Engine
	ServerConfig *ServerConfig
}

// NewApp creates and initializes a new application instance
func NewApp() (*App, error) {
	// Enable Gin console colors
	gin.ForceConsoleColor()

	// Initialize router with middleware and routes
	router := initialize.InitRouter()

	// Load server configuration
	serverConfig := LoadServerConfig()

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
