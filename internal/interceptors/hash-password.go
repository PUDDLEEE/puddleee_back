package interceptors

import (
	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		body := ctx.Value("body")
		if body != nil {
			requestBody := body.(dto.CreateUserDTO)
			bytes, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 12)
			if err != nil {
				internalError := errors.INTERNAL_ERROR
				internalError.Data = err.Error()
				ctx.Error(internalError)
			}
			requestBody.Password = string(bytes)
			ctx.Set("body", requestBody)
			ctx.Next()
		}

	}
}
