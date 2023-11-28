package user

import (
	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	userService UserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	body := ctx.Value("body")
	if body != nil {
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

}
func (c *UserController) ViewUser(ctx *gin.Context) {

}

func (c *UserController) ViewProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		internalError := errors.INTERNAL_ERROR
		internalError.Data = err.Error()
		ctx.Error(internalError)
		return
	}
	existedUser, err := c.userService.findOneUser(id)
	if err != nil {
		notfoundError := errors.NOT_FOUND
		notfoundError.Data = err.Error()
		ctx.Error(notfoundError)
		return
	}
	ctx.JSON(http.StatusOK, existedUser)
}

func NewController(service UserService) UserController {
	return UserController{userService: service}
}
