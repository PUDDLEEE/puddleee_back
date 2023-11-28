package user

import (
	"fmt"
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

func (c *UserController) ViewProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("View user id with %s", id),
	})
}

func NewController(service UserService) UserController {
	return UserController{userService: service}
}
