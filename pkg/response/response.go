package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
	Error   interface{} `json:"error"`
}

// GetMessageByCode returns the appropriate message for a given response code
func GetMessageByCode(code int) string {
	if message, exists := msg[code]; exists {
		return message
	}
	return "Unknown error occurred"
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	message := GetMessageByCode(code)
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Content: data,
		Error:   nil,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	// If message is empty, use default message
	if message == "" {
		message = GetMessageByCode(code)
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Content: nil,
		Error:   nil,
	})
}
