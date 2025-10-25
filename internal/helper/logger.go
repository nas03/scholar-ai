package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/utils/requestid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogWithRequestID logs a message with requestId from context
func LogWithRequestID(c *gin.Context, level zapcore.Level, msg string, fields ...zap.Field) {
	requestID := requestid.GetRequestIDFromContext(c)
	if requestID != "" {
		fields = append(fields, zap.String("requestId", requestID))
	}

	global.Log.Check(level, msg).Write(fields...)
}

// LogInfo logs an info message with requestId
func LogInfo(c *gin.Context, msg string, fields ...zap.Field) {
	LogWithRequestID(c, zap.InfoLevel, msg, fields...)
}

// LogWarn logs a warning message with requestId
func LogWarn(c *gin.Context, msg string, fields ...zap.Field) {
	LogWithRequestID(c, zap.WarnLevel, msg, fields...)
}

// LogError logs an error message with requestId
func LogError(c *gin.Context, msg string, fields ...zap.Field) {
	LogWithRequestID(c, zap.ErrorLevel, msg, fields...)
}

// LogDebug logs a debug message with requestId
func LogDebug(c *gin.Context, msg string, fields ...zap.Field) {
	LogWithRequestID(c, zap.DebugLevel, msg, fields...)
}

// GetRequestID is a convenience function to get requestId from context
func GetRequestID(c *gin.Context) string {
	return requestid.GetRequestIDFromContext(c)
}
