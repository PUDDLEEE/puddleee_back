package app

import (
	"log"
	"sync"

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
	router.GET("/", app.Home)
}
