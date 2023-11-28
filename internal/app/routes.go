package app

import (
	"context"
	"log"
	"sync"

	"github.com/PUDDLEEE/puddleee_back/internal/middlewares"
	"github.com/PUDDLEEE/puddleee_back/internal/user"
	"github.com/gin-gonic/gin"
)

var once sync.Once
var routes Routes

func (app *Application) initRoutes() {
	if routes.Router == nil {
		once.Do(func() {
			app.setRoutes()
		})
	} else {
		log.Fatal("Router already configure")
	}
	app.Routes = routes
}

func (app *Application) setRoutes() {
	routes.Router = gin.Default()
	routes.Router.Use(middlewares.ErrorHandler())

	v1 := routes.Router.Group("/api/v1")
	userController := user.InitializeController(context.Background(), app.Client)
	routes.addUser(v1, userController)
}

func (r *Routes) addUser(rg *gin.RouterGroup, controller user.UserController) {
	userGroup := rg.Group("/user")

	userGroup.GET("", controller.ViewUser)
	userGroup.GET("/:id", controller.ViewProfile)
	userGroup.
		POST("", middlewares.HashPasswordHandler(), controller.CreateUser)
}
