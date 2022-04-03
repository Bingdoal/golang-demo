package model

import (
	"go-demo/config/db/postgres"
	"go-demo/internal/model/base"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

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

func (post *Post) FindAll() ([]Post, error) {
	result := []Post{}
	err := post.model().Find(&result).Error
	return result, err
}

func (post *Post) FindByUser() ([]Post, error) {
	result := []Post{}
	err := post.model().Where("author_id = ?", post.AuthorID).Find(&result).Error
	return result, err
}

func (post *Post) Create() error {
	err := v.ValidateStruct(post,
		v.Field(&post.Content, v.Required, v.Min(2)),
		v.Field(&post.AuthorID, v.Required),
	)
	if err != nil {
		return err
	}

	return post.model().Create(post).Error
}

func (post *Post) Update() error {
	err := v.ValidateStruct(post,
		v.Field(&post.Content, v.Required, v.Min(2)),
		v.Field(&post.AuthorID, v.Required),
	)
	if err != nil {
		return err
	}

	return post.model().Save(post).Error
}

func (post *Post) Delete() error {
	return post.model().Delete(post).Error
}
