package app

import (
	"context"
	"log"
	"sync"

	"github.com/PUDDLEEE/puddleee_back/internal/user"
	"github.com/gin-gonic/gin"
)

var once sync.Once
var router *gin.Engine

func (app *Application) initRoutes() {
	if router == nil {
		once.Do(func() {
			app.setRoutes()
		})
	} else {
		log.Fatal("Router already configure")
	}
	app.Router = router
}

func (app *Application) setRoutes() {
	router = gin.Default()
	userController := user.InitializeController(context.Background(), app.Client)
	router.GET("/", userController.ViewUser)
}
