package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func SuccessResponse(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
