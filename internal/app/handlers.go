package app

import (
	"github.com/PUDDLEEE/puddleee_back/internal/home"
	"github.com/gin-gonic/gin"
)

func (app *Application) Home(ctx *gin.Context) {
	home.GetHome(ctx)
}
