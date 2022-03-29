package model

import "go-demo/config/db/postgres"

type Post struct {
	BaseModel
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author" gorm:"foreignkey:AuthorID"`
}

var postModel = postgres.DB.Model(&Post{})

func (post *Post) FindAll() ([]Post, error) {
	result := []Post{}
	err := postModel.Find(&result).Error
	return result, err
}

func (post *Post) FindByUser(userId uint) ([]Post, error) {
	result := []Post{}
	err := postModel.Where("author_id = ?", userId).Find(&result).Error
	return result, err
}

func (post *Post) Create() error {
	return postModel.Create(post).Error
}

func (post *Post) Update() error {
	return postModel.Save(post).Error
}

func (post *Post) Delete() error {
	return postModel.Delete(post).Error
}
