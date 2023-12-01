package user

import (
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
	"net/http"
	"strconv"
	"strings"

	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService interfaces.IUserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	body := ctx.Value("body")
	if body != nil {
		requestBody := body.(dto.CreateUserDTO)
		user, err := c.userService.CreateUser(requestBody)
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
	existedUser, err := c.userService.FindOneUser(id)
	if err != nil {
		notfoundError := errors.NOT_FOUND
		notfoundError.Data = err.Error()
		ctx.Error(notfoundError)
		return
	}
	ctx.JSON(http.StatusOK, existedUser)
}

func NewController(service interfaces.IUserService) *UserController {
	if userService, ok := service.(*UserService); ok {
		return &UserController{userService: userService}
	}
	if mockUserService, ok := service.(*mocks.IUserService); ok {
		return &UserController{userService: mockUserService}
	}
	return &UserController{}
}
