package entity

import (
	"go-demo/internal/model/base"
)

type Users []User

type User struct {
	base.BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
