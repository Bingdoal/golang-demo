package model

import (
	"go-demo/config/db/postgres"

	_ "gorm.io/gorm"
)

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var userModel = postgres.DB.Model(&User{})

func (user *User) Create() error {
	return userModel.Create(user).Error
}

func (user *User) Update() error {
	return userModel.Save(user).Error
}

func (user *User) FindOne(id uint) error {
	return userModel.Where("id = ?", id).First(user).Error
}

func (user *User) FindAll() ([]User, error) {
	result := []User{}
	err := userModel.Find(&result).Error
	return result, err
}

func (user *User) Delete() error {
	return userModel.Delete(user).Error
}
