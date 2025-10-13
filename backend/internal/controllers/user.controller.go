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
	err := ctx.ShouldBindJSON(payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Required field is missing",
		})
	}
	code := c.userService.CreateUser(payload.Username, payload.Password, payload.Email)
	if code != response.CodeSuccess {
		response.ErrorResponse(ctx, code, "")
		return
	}

	response.SuccessResponse(ctx, code, "", nil)

}
