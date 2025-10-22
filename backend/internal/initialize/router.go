package initialize

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nas03/scholar-ai/backend/docs"
	"github.com/nas03/scholar-ai/backend/internal/middleware"
	"github.com/nas03/scholar-ai/backend/internal/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initializes the main router with middleware and routes
func InitRouter() *gin.Engine {
	// Create Gin engine
	r := gin.New()

	// Add global middleware
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.RequestID())
	r.Use(middleware.RequestLogger())

	// Setup API routes
	apiV1 := r.Group("/api/v1")
	{
		// Register user routes
		router.SetupUserRoutes(apiV1)

		// Add other route groups here as needed
		// router.SetupProductRoutes(apiV1)
		// router.SetupOrderRoutes(apiV1)
	}

	// Swagger documentation (dev only recommended)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
