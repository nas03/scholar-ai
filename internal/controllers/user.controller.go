package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"github.com/nas03/scholar-ai/backend/internal/services"
	"github.com/nas03/scholar-ai/backend/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var payload models.CreateUserRequest

	// Validate JSON binding
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ErrorResponse(ctx, response.CodeInvalidInput, err.Error())
		return
	}

	// Call service to create user
	code := c.userService.CreateUser(ctx, payload.Username, payload.Password, payload.Email)

	// Handle response based on service result
	if code == response.CodeSuccess {
		data := map[string]interface{}{"requiresOtp": true}
		response.SuccessResponse(ctx, code, data)
	} else {
		response.ErrorResponse(ctx, code, "")
	}
}
