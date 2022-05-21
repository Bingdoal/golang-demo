package entity

import (
	"go-demo/internal/model/base"
)

type Users []User

type User struct {
	base.BaseModel
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-"`
}
