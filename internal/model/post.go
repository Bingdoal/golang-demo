package model

import (
	"go-demo/config/db/postgres"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	BaseModel
	Content  string `json:"content"`
	AuthorID uint64 `json:"authorId"`
	Author   User   `json:"author" gorm:"foreignkey:AuthorID"`
}

func postModel() *gorm.DB { return postgres.DB.Model(&Post{}) }

func (post *Post) FindAll() ([]Post, error) {
	result := []Post{}
	err := postModel().Find(&result).Error
	return result, err
}

func (post *Post) FindByUser() ([]Post, error) {
	result := []Post{}
	err := postModel().Where("author_id = ?", post.AuthorID).Find(&result).Error
	return result, err
}

func (post *Post) Create() error {
	return postModel().Create(post).Error
}

func (post *Post) Update() error {
	return postModel().Save(post).Error
}

func (post *Post) Delete() error {
	return postModel().Delete(post).Error
}

func (post *Post) BeforeCreate(db *gorm.DB) (err error) {
	post.CreationTime = time.Now()
	post.ModificationTime = time.Now()
	return
}

func (post *Post) BeforeUpdate(db *gorm.DB) (err error) {
	post.ModificationTime = time.Now()
	return
}
