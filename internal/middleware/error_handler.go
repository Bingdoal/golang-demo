package middleware

import (
	"fmt"
	"go-demo/internal/dto"
	"go-demo/internal/enum"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err interface{}) {
		c.JSON(500, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     fmt.Sprint(err),
		})
		c.Abort()
	}
}
