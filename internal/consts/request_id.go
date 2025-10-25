package consts

import "time"

// RequestIDHeader is the header key for request ID
const RequestIDHeader = "X-Request-ID"

// RequestIDContextKey is the context key for storing request ID
const RequestIDContextKey = "requestId"

// RequestID validation rules
const (
	RequestIDMinLength = 16            // Minimum length for requestId
	RequestIDMaxLength = 64            // Maximum length for requestId
	RequestIDPattern   = `^[a-f0-9]+$` // Only lowercase hex characters allowed
)

// Redis keys for requestId tracking
const (
	RedisKeyRequestIDPrefix = "request_id:"
	RedisKeyRequestIDExpiry = 3600 * time.Second // 3600 seconds
)
