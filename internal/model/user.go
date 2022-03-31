package model

import (
	"go-demo/config/db/postgres"
	"time"

	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func userModel() *gorm.DB { return postgres.DB.Model(&User{}) }

func (user *User) Create() error {
	return userModel().Create(user).Error
}

func (user *User) Update() error {
	return userModel().Save(user).Error
}

func (user *User) FindOne() error {
	return userModel().Where("id = ?", user.ID).First(user).Error
}

func (user *User) FindAll() ([]User, error) {
	result := []User{}
	err := userModel().Find(&result).Error
	return result, err
}

func (user *User) Delete() error {
	return userModel().Delete(user).Error
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.CreationTime = time.Now()
	user.ModificationTime = time.Now()
	return
}

func (user *User) BeforeUpdate(db *gorm.DB) (err error) {
	user.ModificationTime = time.Now()
	return
}
