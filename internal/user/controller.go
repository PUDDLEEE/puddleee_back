package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
}

func (c *UserController) ViewUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "View User",
	})
}

func NewController(service UserService) UserController {
	return UserController{userService: service}
}
