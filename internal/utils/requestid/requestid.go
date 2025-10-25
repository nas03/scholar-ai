package requestid

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/internal/consts"
	"github.com/nas03/scholar-ai/backend/internal/utils"
)

var (
	requestIDRegex = regexp.MustCompile(consts.RequestIDPattern)
)

// ValidateRequestID validates the format and length of requestId
func ValidateRequestID(requestID string) (bool, string) {
	// Check length
	if len(requestID) < consts.RequestIDMinLength {
		return false, fmt.Sprintf("Request ID too short, minimum %d characters required", consts.RequestIDMinLength)
	}

	if len(requestID) > consts.RequestIDMaxLength {
		return false, fmt.Sprintf("Request ID too long, maximum %d characters allowed", consts.RequestIDMaxLength)
	}

	// Check format (only lowercase hex characters)
	if !requestIDRegex.MatchString(requestID) {
		return false, "Request ID must contain only lowercase hexadecimal characters (a-f, 0-9)"
	}

	return true, ""
}

// GenerateRequestID generates a new unique requestId
func GenerateRequestID() string {
	// Generate 16 random bytes and convert to hex
	bytes := make([]byte, 16)
	rand.Read(bytes)

	// Convert to hex string
	hexString := fmt.Sprintf("%x", bytes)

	// Add timestamp prefix to ensure uniqueness
	timestamp := time.Now().UnixNano()
	timestampHex := fmt.Sprintf("%x", timestamp)

	// Combine timestamp (first 8 chars) + random bytes (32 chars)
	return timestampHex[:8] + hexString
}

// IsRequestIDDuplicate checks if requestId already exists in Redis
func IsRequestIDDuplicate(ctx *gin.Context, requestID string) (bool, error) {
	key := consts.RedisKeyRequestIDPrefix + requestID

	// Check if requestId exists using Redis
	exists, err := utils.NewRedisCache().Get(ctx, key)
	if err != nil {
		// If key not found, it means NOT duplicate
		if err.Error() == "redis: nil" {
			return false, nil
		}
		// If other Redis error, we can't determine if it's duplicate
		return false, err
	}

	// If exists is not empty, it means requestId was used before
	return exists != "", nil
}

// StoreRequestID stores requestId in Redis to prevent duplicates
func StoreRequestID(ctx *gin.Context, requestID string) error {
	key := consts.RedisKeyRequestIDPrefix + requestID

	// Store with expiration time
	return utils.NewRedisCache().SetEx(ctx, key, "1", consts.RedisKeyRequestIDExpiry)
}

// CleanupRequestID removes requestId from Redis (optional cleanup)
func CleanupRequestID(ctx *gin.Context, requestID string) error {
	// Note: The current Redis cache doesn't have Del method, so we'll skip cleanup for now
	// In production, you might want to add a Del method to the Redis cache interface
	_ = requestID // Suppress unused parameter warning
	return nil
}

// NormalizeRequestID normalizes requestId to lowercase and removes invalid characters
func NormalizeRequestID(requestID string) string {
	// Convert to lowercase
	normalized := strings.ToLower(requestID)

	// Remove any non-hex characters
	normalized = requestIDRegex.FindString(normalized)

	// Ensure minimum length
	if len(normalized) < consts.RequestIDMinLength {
		// Pad with zeros if too short
		normalized = strings.Repeat("0", consts.RequestIDMinLength-len(normalized)) + normalized
	}

	// Truncate if too long
	if len(normalized) > consts.RequestIDMaxLength {
		normalized = normalized[:consts.RequestIDMaxLength]
	}

	return normalized
}

// GetRequestIDFromContext gets requestId from gin context
func GetRequestIDFromContext(ctx *gin.Context) string {
	if requestID, exists := ctx.Get(consts.RequestIDContextKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}
