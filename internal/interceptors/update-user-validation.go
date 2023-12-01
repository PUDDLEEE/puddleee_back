package interceptors

import (
	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/gin-gonic/gin"
	"log"
)

func UpdateUserInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var requestBody dto.UpdateUserDTO

		err := ctx.ShouldBind(&requestBody)
		if err != nil {
			internalError := errors.INTERNAL_ERROR
			internalError.Data = err.Error()
			_ = ctx.Error(internalError)
			return
		}
		if err := requestBody.Validate(); err != nil {
			emailError := errors.INVALID_EMAIL
			emailError.Data = err.Error()
			_ = ctx.Error(emailError)
			return
		}
		ctx.Set("body", requestBody)
		log.Println(ctx.Value("body"))
		ctx.Next()
	}
}
