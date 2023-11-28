package middlewares

import (
	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/gin-gonic/gin"
)

func HashPasswordHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody dto.CreateUserDTO

		err := ctx.ShouldBind(&requestBody)
		if err != nil {
			internalError := errors.INTERNAL_ERROR
			internalError.Data = err.Error()
			ctx.Error(internalError)
			return
		}
		if err := requestBody.Validate(); err != nil {
			emailError := errors.INVALID_EMAIL
			emailError.Data = err.Error()
			ctx.Error(emailError)
			return
		}
		requestBody.Password = "Custom Password"
		ctx.Set("body", requestBody)
		ctx.Next()
	}
}
