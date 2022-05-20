package common

import (
	"go-demo/internal/dto/basic"
	"go-demo/internal/enum"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RespError(ctx *gin.Context, status int, errorMsg string) {
	ctx.JSON(status, basic.RespDto{
		Message: enum.MessageType(enum.Error),
		Err:     errorMsg,
	})
}

func GetPagination(ctx *gin.Context) (pagination basic.Pagination, err error) {
	pagination = basic.Pagination{}
	pagination.Page, err = strconv.ParseUint(ctx.Query("page"), 10, 64)
	if err != nil {
		return
	}
	pagination.PageSize, err = strconv.ParseUint(ctx.Query("pageSize"), 10, 64)
	if err != nil {
		return
	}
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	}
	return
}
