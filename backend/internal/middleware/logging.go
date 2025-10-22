package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/global"
	"go.uber.org/zap"
)

// RequestLogger logs each HTTP request with zap, including latency and request id.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		ip := c.ClientIP()
		ua := c.Request.UserAgent()
		rid := c.GetString("request_id")

		if raw != "" {
			path = path + "?" + raw
		}

		if global.Log != nil {
			fields := []zap.Field{
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("path", path),
				zap.Duration("latency", latency),
				zap.String("ip", ip),
				zap.String("user_agent", ua),
			}
			if rid != "" {
				fields = append(fields, zap.String("request_id", rid))
			}
			if len(c.Errors) > 0 {
				for _, e := range c.Errors {
					global.Log.Error("request error", append(fields, zap.String("error", e.Error()))...)
				}
			} else {
				global.Log.Info("request completed", fields...)
			}
		}
	}
}
