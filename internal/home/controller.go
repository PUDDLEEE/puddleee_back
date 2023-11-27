package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "home",
	})
}
