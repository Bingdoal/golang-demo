package entity

import (
	"go-demo/internal/model/base"
)

type Posts []Post

type Post struct {
	base.BaseModel
	Content  string `json:"content"`
	AuthorID uint64 `json:"authorId"`
	Author   User   `json:"author" gorm:"foreignkey:AuthorID"`
}
