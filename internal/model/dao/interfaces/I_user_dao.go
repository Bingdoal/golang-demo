package interfaces

import (
	"go-demo/internal/dto/basic"
	"go-demo/internal/model/entity"
)

type IUserDao interface {
	FindAll(condi entity.User, pagination basic.Pagination, dest *entity.Users) (int64, error)
	FindOne(dest *entity.User) error
	Login(name string, password string) error
	Delete(id uint64) error
	Create(src *entity.User) error
	Update(src *entity.User) error
}
