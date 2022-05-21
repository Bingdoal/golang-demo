package middleware

import (
	"fmt"
	"go-demo/api/common"
	"go-demo/internal/dto/basic"
	"go-demo/internal/enum"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err interface{}) {
		switch v := err.(type) {
		default:
			c.JSON(500, basic.RespDto{
				Message: enum.MessageType(enum.Error),
				Err:     fmt.Sprintf("%T: %s", v, err),
			})
		case common.StatusError:
			statusError, _ := err.(common.StatusError)
			c.JSON(statusError.Status, basic.RespDto{
				Message: enum.MessageType(enum.Error),
				Err:     statusError.Message,
			})
		}
		c.Abort()
	}
}
