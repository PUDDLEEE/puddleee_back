package app

import (
	"context"
	"log"
	"sync"

	"github.com/PUDDLEEE/puddleee_back/internal/auth"
	"github.com/PUDDLEEE/puddleee_back/internal/mail"
	"github.com/PUDDLEEE/puddleee_back/internal/room"

	"github.com/PUDDLEEE/puddleee_back/internal/interceptors"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
	routes.Router.Use(cors.Default())
	v1 := routes.Router.Group("/api/v1")
	userController := user.InitializeController(context.Background(), app.Client)
	roomController := room.InitializeController(context.Background(), app.Client)

	mailParams := &mail.MailServiceParams{}
	mailParams.From = app.Config.Mail.From
	mailParams.Port = app.Config.Mail.Port
	mailParams.Server = app.Config.Mail.Server
	mailParams.Password = app.Config.Mail.Password

	authController := auth.InitializeController(context.Background(), app.Client, app.Config.Jwt.SecretKey, *mailParams)
	routes.addSwagger(v1)
	routes.addUser(v1, userController)
	routes.addRoom(v1, roomController)
	routes.addAuth(v1, authController)
}

func (r *Routes) addSwagger(rg *gin.RouterGroup) {
	swaggerGroup := rg.Group("/swagger")
	swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *Routes) addUser(rg *gin.RouterGroup, controller *user.UserController) {

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

	userGroup := rg.Group("/users")

	userGroup.GET("", controller.ViewUser)
	userGroup.GET("/:id", controller.ViewProfile)

	userGroup.
		POST("", createUserInterceptors...)

	userGroup.PATCH("/:id", updateUserInterceptors...)

	userGroup.DELETE("/:id", controller.DeleteProfile)
}

func (r *Routes) addRoom(rg *gin.RouterGroup, controller *room.RoomController) {
	roomGroup := rg.Group("/rooms")

	roomGroup.GET("", controller.ViewRooms)
	roomGroup.GET("/:id", controller.ViewOneRoom)
	roomGroup.POST("", controller.CreateRoom)
	roomGroup.DELETE("/:id", controller.DeleteRoom)
	roomGroup.PATCH("/:id", controller.UpdateRoomInfo)
}

func (r *Routes) addAuth(rg *gin.RouterGroup, controller *auth.AuthController) {
	authGroup := rg.Group("/auth")

	authGroup.POST("/signin", controller.Signin)
	authGroup.POST("/signup", controller.Signup)
	authGroup.POST("/mail", controller.SendConfirmationMail)
	authGroup.POST("/confirm", controller.VerifyConfirmationMail)
}
