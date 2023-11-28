package user

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	requestBody := ctx.Value("body").(dto.CreateUserDTO)
	user, err := c.userService.createUser(requestBody)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "email"):
			emailInUseError := errors.EMAIL_EXISTED
			emailInUseError.Data = err.Error()
			ctx.Error(emailInUseError)
			return
		case strings.Contains(err.Error(), "username"):
			usernameInUseError := errors.USERNAME_EXISTED
			usernameInUseError.Data = err.Error()
			ctx.Error(usernameInUseError)
			return
		}
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
