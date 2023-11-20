package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserGroup(router *gin.RouterGroup) *gin.RouterGroup {
	users := router.Group("/users")
	users.GET("", ViewUserList)
	users.GET("/signup", Signup)
	users.GET("/:id", ViewProfile)
	users.PUT("/:id", EditProfile)
	return users
}

func Signup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is signup test controller",
	})
}

func ViewProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is ViewProfile controller",
	})
}

func EditProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is EditProfile controller",
	})
}

func ViewUserList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is ViewUserList controller",
	})
}
