package interfaces

import (
	"go-demo/internal/dto/basic"
	"go-demo/internal/model/entity"
)

type IPostDao interface {
	FindOne(dest *entity.Post) error
	FindAll(condi entity.Post, pagination basic.Pagination, dest *entity.Posts) (int64, error)
	Create(src *entity.Post) error
	Update(src *entity.Post) error
	Delete(id uint64) error
}
