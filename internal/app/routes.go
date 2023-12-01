package app

import (
	"context"
	"github.com/PUDDLEEE/puddleee_back/internal/interceptors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	routes.addSwagger(v1)
	routes.addUser(v1, userController)
}

func (r *Routes) addSwagger(rg *gin.RouterGroup) {
	swaggerGroup := rg.Group("/swagger")
	swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *Routes) addUser(rg *gin.RouterGroup, controller user.UserController) {

	createUserInterceptors := []gin.HandlerFunc{
		interceptors.CreateUserInterceptor(),
		interceptors.HashPasswordInterceptor(),
	}
	createUserInterceptors = append(createUserInterceptors, controller.CreateUser)

	updateUserInterceptors := []gin.HandlerFunc{
		interceptors.UpdateUserInterceptor(),
		interceptors.HashPasswordInterceptor(),
	}

	updateUserInterceptors = append(updateUserInterceptors, controller.UpdateProfile)

	userGroup := rg.Group("/user")

	userGroup.GET("", controller.ViewUser)
	userGroup.GET("/:id", controller.ViewProfile)

	userGroup.
		POST("", createUserInterceptors...)

	userGroup.PATCH("/:id", updateUserInterceptors...)

	userGroup.DELETE("/:id", controller.DeleteProfile)
}
