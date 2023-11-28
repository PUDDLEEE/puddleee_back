package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var requestBody dto.CreateUserDTO

	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		log.Fatal(err)
	}
	if err := requestBody.Validate(); err != nil {
		emailError := errors.INVALID_EMAIL
		emailError.Data = err.Error()
		ctx.Error(emailError)
		return
	}
	user, err := c.userService.createUser(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, user)
}
func (c *UserController) ViewUser(ctx *gin.Context) {

}

func (c *UserController) ViewProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	user, err := c.userService.viewOneUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
	}
	ctx.JSON(http.StatusOK, user)
}

func NewController(service UserService) UserController {
	return UserController{userService: service}
}
