package initialize

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/nas03/scholar-ai/backend/docs"
	"github.com/nas03/scholar-ai/backend/internal/middleware"
	"github.com/nas03/scholar-ai/backend/internal/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initializes the main router with middleware and routes
func InitRouter() *gin.Engine {
	if gin.Mode() != gin.ReleaseMode {
		gin.ForceConsoleColor()

		// Enable custom colored debug output if GIN_COLOR_DEBUG env is set (e.g., "1", "true", "yes")
		if enableColoredDebug() {
			// Customize how Gin prints route registrations to make them easier to scan
			gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
				c := methodColor(httpMethod)
				reset := "\033[0m"
				// Example: [GIN-debug] GET     | /api/v1/users            | controllers.UserController.List (3 handlers)
				fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] %s%-7s%s | %-40s | %s (%d handlers)\n",
					c, httpMethod, reset, absolutePath, handlerName, nuHandlers,
				)
			}
		}
	}

	// Create Gin engine
	r := gin.New()

	// Add global middleware
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.LoggerMiddleware()) // Simple, proven logging from fidecwalletserver

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

// enableColoredDebug checks if the GIN_COLOR_DEBUG env var is set to a truthy value.
// Examples: GIN_COLOR_DEBUG=1, true, yes, on (case-insensitive) all enable it.
// Defaults to true in dev if not set.
func enableColoredDebug() bool {
	val := strings.ToLower(strings.TrimSpace(os.Getenv("GIN_COLOR_DEBUG")))
	if val == "" {
		// Default to enabled in non-release mode
		return true
	}
	return val == "1" || val == "true" || val == "yes" || val == "on"
}

// methodColor returns an ANSI color escape code by HTTP method for nicer debug output.
func methodColor(method string) string {
	switch method {
	case "GET":
		return "\033[32m" // green
	case "POST":
		return "\033[36m" // cyan
	case "PUT":
		return "\033[34m" // blue
	case "DELETE":
		return "\033[31m" // red
	case "PATCH":
		return "\033[35m" // magenta
	case "OPTIONS":
		return "\033[33m" // yellow
	case "HEAD":
		return "\033[37m" // white
	default:
		return "\033[0m" // reset/default
	}
}
