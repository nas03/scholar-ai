package middleware

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/consts"
	"github.com/nas03/scholar-ai/backend/internal/utils/requestid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// responseWriter wraps gin.ResponseWriter to capture response body
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// ANSI color codes for beautiful terminal output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorGray   = "\033[90m"
	ColorBold   = "\033[1m"
)

// getMethodColor returns the appropriate color for each HTTP method
func getMethodColor(method string) string {
	switch strings.ToUpper(method) {
	case "GET":
		return ColorGreen // Green for safe operations
	case "POST":
		return ColorCyan // Cyan for creation
	case "PUT":
		return ColorBlue // Blue for updates
	case "PATCH":
		return ColorPurple // Purple for partial updates
	case "DELETE":
		return ColorRed // Red for destructive operations
	case "HEAD":
		return ColorGray // Gray for metadata requests
	case "OPTIONS":
		return ColorYellow // Yellow for preflight requests
	default:
		return ColorWhite // White for unknown methods
	}
}

// getStatusColor returns the appropriate color for HTTP status codes
func getStatusColor(status int) string {
	switch {
	case status >= 200 && status < 300:
		return ColorGreen // Green for success
	case status >= 300 && status < 400:
		return ColorBlue // Blue for redirects
	case status >= 400 && status < 500:
		return ColorYellow // Yellow for client errors
	case status >= 500:
		return ColorRed // Red for server errors
	default:
		return ColorWhite // White for unknown
	}
}

// getLogLevelColor returns the appropriate color for log levels
func getLogLevelColor(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return ColorGray
	case zapcore.InfoLevel:
		return ColorGreen
	case zapcore.WarnLevel:
		return ColorYellow
	case zapcore.ErrorLevel:
		return ColorRed
	default:
		return ColorWhite
	}
}

// formatDuration formats duration in a human-readable way
func formatDuration(duration time.Duration) string {
	if duration < time.Millisecond {
		return fmt.Sprintf("%.0fμs", float64(duration.Nanoseconds())/1000)
	} else if duration < time.Second {
		return fmt.Sprintf("%.2fms", float64(duration.Nanoseconds())/1e6)
	} else {
		return fmt.Sprintf("%.3fs", duration.Seconds())
	}
}

// supportsColor checks if the terminal supports color output
func supportsColor() bool {
	// Check if we're in development mode (more likely to have colors)
	if global.Config.Log.AppEnv == "dev" || global.Config.Log.AppEnv == "development" {
		return true
	}
	return false
}

// LoggerMiddleware creates a logger middleware with optional requestId from client
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Get requestId from client (optional)
		requestID := c.GetHeader(consts.RequestIDHeader)
		isClientProvided := requestID != ""

		// If no request ID provided, generate one for internal tracking
		if !isClientProvided {
			requestID = requestid.GenerateRequestID()
		}

		// If client provided request ID, validate it
		if isClientProvided {
			originalRequestID := requestID
			needNewRequestID := false
			reason := ""

			// Check format validity
			isValid, validationError := requestid.ValidateRequestID(requestID)
			if !isValid {
				needNewRequestID = true
				reason = "invalid_format: " + validationError
			}

			// Check for duplicate (only if format is valid)
			if isValid {
				isDuplicate, err := requestid.IsRequestIDDuplicate(c, requestID)
				if err != nil {
					global.Log.Warn("Failed to check requestId duplicate, proceeding anyway",
						zap.String("method", c.Request.Method),
						zap.String("path", c.Request.URL.Path),
						zap.String("request_id", requestID),
						zap.Error(err),
					)
				} else if isDuplicate {
					needNewRequestID = true
					reason = "duplicate_request_id"
				}
			}

			// Generate new requestId if needed
			if needNewRequestID {
				global.Log.Warn("Request ID needs replacement, generating new one",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("client_ip", c.ClientIP()),
					zap.String("user_agent", c.Request.UserAgent()),
					zap.String("client_request_id", originalRequestID),
					zap.String("reason", reason),
					zap.String("action", "generate_new_request_id"),
				)

				requestID = requestid.GenerateRequestID()
				isClientProvided = false // Mark as server-generated

				global.Log.Info("Generated new Request ID",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("client_request_id", originalRequestID),
					zap.String("server_request_id", requestID),
					zap.String("reason", reason),
				)
			}

			// Store requestId in Redis to prevent future duplicates (only if client provided)
			if err := requestid.StoreRequestID(c, requestID); err != nil {
				global.Log.Warn("Failed to store requestId in Redis",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("request_id", requestID),
					zap.Error(err),
				)
			}
		}

		// Add requestId to context
		c.Set(consts.RequestIDContextKey, requestID)

		// Capture request body
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Wrap response writer to capture response body
		responseWriter := &responseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = responseWriter

		// No request log - we'll show everything in the response log

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Get response status
		status := c.Writer.Status()

		// Determine log level and create beautiful response message
		logLevel := zap.InfoLevel
		if status >= 400 {
			logLevel = zap.WarnLevel
		}
		if status >= 500 {
			logLevel = zap.ErrorLevel
		}

		// Create beautiful formatted response message
		var responseMessage string
		if supportsColor() {
			methodColor := getMethodColor(c.Request.Method)
			statusColor := getStatusColor(status)
			formattedDuration := formatDuration(duration)

			responseMessage = fmt.Sprintf("%s%s%s %s%s%s %s%d%s %s%s%s %s%s%s",
				methodColor, strings.ToUpper(c.Request.Method), ColorReset,
				ColorWhite, c.Request.URL.Path, ColorReset,
				statusColor, status, ColorReset,
				ColorGray, formattedDuration, ColorReset,
				ColorGray, requestID[:8], ColorReset)
		} else {
			responseMessage = fmt.Sprintf("%s %s %d %s %s",
				strings.ToUpper(c.Request.Method),
				c.Request.URL.Path,
				status,
				formatDuration(duration),
				requestID[:8])
		}

		// Convert gin errors to strings
		var errorStrings []string
		for _, err := range c.Errors {
			errorStrings = append(errorStrings, err.Error())
		}

		// Log response with beautiful formatting - only add errors if they exist
		if len(errorStrings) > 0 {
			global.Log.Check(logLevel, responseMessage).Write(
				zap.Strings("errors", errorStrings),
			)
		} else {
			global.Log.Check(logLevel, responseMessage).Write()
		}
	}
}

// GetRequestID retrieves requestId from gin context
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(consts.RequestIDContextKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}
