package middlewares

import (
	"net/http"

	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, e := range c.Errors {
			err := e.Err
			if myErr, ok := err.(*errors.CustomError); ok {
				c.JSON(myErr.Code, myErr)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  errors.INTERNAL_ERROR,
					"data": err.Error(),
				})
			}
		}
	}
}
