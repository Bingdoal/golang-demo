package common

import (
	"go-demo/internal/dto"
	"go-demo/internal/enum"

	"github.com/gin-gonic/gin"
)

func RespError(ctx *gin.Context, status int, errorMsg string) {
	ctx.JSON(status, dto.RespDto{
		Message: enum.MessageType(enum.Error),
		Err:     errorMsg,
	})
}
