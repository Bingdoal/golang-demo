package entity

import (
	"go-demo/internal/model/base"
)

type Posts []Post

type Post struct {
	base.BaseModel
	Content  string `json:"content" form:"content"`
	AuthorID uint64 `json:"authorId" form:"authorId"`
}
