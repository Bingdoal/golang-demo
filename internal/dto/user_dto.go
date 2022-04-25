package dto

type UserDto struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,min=5,email"`
	Password string `json:"password" binding:"required,min=6"`
}
