package interfaces

import "go-demo/internal/model/entity"

type IPostDao interface {
	FindByUser(authorId uint64, dest *entity.Posts) error
	FindOne(dest *entity.Post) error
	FindAll(dest *entity.Posts) error
	Create(src *entity.Post) error
	Update(src *entity.Post) error
	Delete(id uint64) error
}
