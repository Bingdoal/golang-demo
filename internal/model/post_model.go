package model

import (
	"go-demo/config/db/postgres"
	"go-demo/internal/model/base"

	"gorm.io/gorm"
)

type Posts []Post

func (posts *Posts) model() *gorm.DB { return postgres.DB.Model(posts) }

func (posts *Posts) FindAll() error {
	return posts.model().Find(&posts).Error
}

func (posts *Posts) FindByUser(authorId uint64) error {
	return posts.model().Where("author_id = ?", authorId).Find(&posts).Error
}

type Post struct {
	base.BaseModel
	Content  string `json:"content"`
	AuthorID uint64 `json:"authorId"`
	Author   User   `json:"author" gorm:"foreignkey:AuthorID"`
}

func (post *Post) model() *gorm.DB { return postgres.DB.Model(post) }

func (post *Post) FindOne() error {
	return post.model().Where("id = ?", post.ID).First(post).Error
}

func (post *Post) Create() error {

	return post.model().Create(post).Error
}

func (post *Post) Update() error {

	return post.model().Updates(post).Error
}

func (post *Post) Delete() error {
	return post.model().Delete(post).Error
}
