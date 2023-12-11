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

			if requestBody, ok := body.(userdto.CreateUserDTO); ok {
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

			if requestBody, ok := body.(userdto.UpdateUserDTO); ok {
				if requestBody.Password != nil {
					bytes, err := bcrypt.GenerateFromPassword([]byte(*requestBody.Password), 12)
					if err != nil {
						internalError := errors.INTERNAL_ERROR
						internalError.Data = err.Error()
						ctx.Error(internalError)
					}
					hashedPassword := string(bytes)
					requestBody.Password = &hashedPassword
					ctx.Set("body", requestBody)
					ctx.Next()
				}
			}

		}

	}
}
