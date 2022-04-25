package model

import (
	"go-demo/config/db/postgres"
	"go-demo/internal/model/base"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users []User

func (users *Users) model() *gorm.DB { return postgres.DB.Model(users) }

func (users *Users) FindAll() error {
	return users.model().Find(&users).Error
}

type User struct {
	base.BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) model() *gorm.DB { return postgres.DB.Model(user) }

func (user *User) Create() error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)

	return user.model().Create(user).Error
}

func (user *User) Login() bool {
	loginPassword := user.Password
	if err := user.model().Where("name = ?", user.Name).First(user).Error; err != nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPassword))
	return err == nil
}

func (user *User) Update() error {
	return user.model().Updates(user).Error
}

func (user *User) FindOne() error {
	return user.model().Where("id = ?", user.ID).First(user).Error
}

func (user *User) Delete() error {
	return user.model().Delete(user).Error
}
