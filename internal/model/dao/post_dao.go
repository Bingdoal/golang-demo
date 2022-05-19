package dao

import (
	"go-demo/config/db"
	"go-demo/internal/model/base"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
	"go-demo/internal/util"

	"gorm.io/gorm"
)

type postDao struct {
	db *gorm.DB
}

// Create implements interfaces.IPostDao
func (dao postDao) Create(src *entity.Post) error {
	return dao.db.Create(src).Error
}

// Delete implements interfaces.IPostDao
func (dao postDao) Delete(id uint64) error {
	return dao.db.Delete(&entity.Post{
		BaseModel: base.BaseModel{
			ID: id,
		},
	}).Error
}

// FindAll implements interfaces.IPostDao
func (dao postDao) FindAll(dest *entity.Posts) error {
	return dao.db.Find(dest).Error
}

func (dao postDao) FindOne(dest *entity.Post) error {
	return dao.db.First(dest).Error
}

// FindByUser implements interfaces.IPostDao
func (dao postDao) FindByUser(authorId uint64, dest *entity.Posts) error {
	return dao.db.Where(&entity.Post{
		AuthorID: authorId,
	}).Find(dest).Error
}

// Update implements interfaces.IPostDao
func (dao postDao) Update(src *entity.Post) error {
	return dao.db.Updates(src).Error
}

func NewPostDao(db *gorm.DB) interfaces.IPostDao {
	util.IfNilPanic(db)
	return postDao{
		db: db,
	}
}

// 事先宣告為 interface 才能在 compile time 進行檢查
var PostDao interfaces.IPostDao

func initPostDao() {
	PostDao = NewPostDao(db.DB)
}
