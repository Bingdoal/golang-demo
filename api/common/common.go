package common

import (
	"errors"
	"go-demo/internal/dto/basic"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPagination(ctx *gin.Context) (pagination basic.Pagination, err error) {
	pagination = basic.Pagination{}
	pagination.Page, err = strconv.ParseUint(ctx.Query("page"), 10, 64)
	if err != nil {
		err = errors.New("Parameter [page] is invalid.")
		return
	}
	pagination.PageSize, err = strconv.ParseUint(ctx.Query("pageSize"), 10, 64)
	if err != nil {
		err = errors.New("Parameter [pageSize] is invalid.")
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
