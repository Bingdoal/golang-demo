package model

import (
	"go-demo/config/db/postgres"
	"go-demo/internal/model/base"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	base.BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) model() *gorm.DB { return postgres.DB.Model(user) }

func (user *User) Create() error {
	err := v.ValidateStruct(user,
		v.Field(&user.Name, v.Required, v.Min(2)),
		v.Field(&user.Email, v.Required, v.Min(5), is.Email),
		v.Field(&user.Password, v.Required, v.Min(6)),
	)
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)

	return user.model().Create(user).Error
}

func (user *User) Login() bool {
	err := v.ValidateStruct(user,
		v.Field(&user.Name, v.Required, v.Min(2)),
		v.Field(&user.Password, v.Required),
	)
	if err != nil {
		return false
	}
	loginPassword := user.Password
	if err := user.model().Where("name = ?", user.Name).First(user).Error; err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPassword))
	return err == nil
}

func (user *User) Update() error {
	err := v.ValidateStruct(user,
		v.Field(&user.Name, v.Required, v.Min(2)),
		v.Field(&user.Email, v.Required, v.Min(5), is.Email),
		v.Field(&user.Password, v.Required),
	)
	if err != nil {
		return err
	}
	return user.model().Updates(user).Error
}

func (user *User) FindOne() error {
	return user.model().Where("id = ?", user.ID).First(user).Error
}

func (user *User) FindAll() ([]User, error) {
	result := []User{}
	err := user.model().Find(&result).Error
	return result, err
}

func (user *User) Delete() error {
	return user.model().Delete(user).Error
}
