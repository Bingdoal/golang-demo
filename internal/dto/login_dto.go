package dto

type LoginDto struct {
	Username string `json:"username" binding:"required,min=2"`
	Password string `json:"password" binding:"required"`
}
