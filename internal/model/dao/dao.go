package dao

import (
	"go-demo/internal/dto/basic"

	"gorm.io/gorm"
)

func InitDao() {
	initPostDao()
	initUserDao()
}

func SetPagination(tx *gorm.DB, pagination basic.Pagination) *gorm.DB {
	offset := (pagination.Page - 1) * pagination.PageSize
	return tx.Offset(int(offset)).Limit(int(pagination.PageSize))
}
