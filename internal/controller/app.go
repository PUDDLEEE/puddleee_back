package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
