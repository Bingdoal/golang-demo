package interfaces

import "go-demo/internal/model/entity"

type IUserDao interface {
	FindAll(dest *entity.Users) error
	FindOne(dest *entity.User) error
	Login(name string, password string) error
	Delete(id uint64) error
	Create(src *entity.User) error
	Update(src *entity.User) error
}
