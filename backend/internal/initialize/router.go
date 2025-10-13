package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/internal/middleware"
	"github.com/nas03/scholar-ai/backend/internal/router"
)

// InitRouter initializes the main router with middleware and routes
func InitRouter() *gin.Engine {
	// Create Gin engine
	r := gin.Default()

	// Add global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.RequestID())

	// Setup API routes
	apiV1 := r.Group("/api/v1")
	{
		// Register user routes
		router.SetupUserRoutes(apiV1)

		// Add other route groups here as needed
		// router.SetupProductRoutes(apiV1)
		// router.SetupOrderRoutes(apiV1)
	}

	return r
}
