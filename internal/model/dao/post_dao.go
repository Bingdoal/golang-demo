package dao

import (
	"go-demo/config/db"
	"go-demo/internal/dto/basic"
	"go-demo/internal/model/base"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"

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
func (dao postDao) FindAll(condition entity.Post, pagination basic.Pagination, dest *entity.Posts) (count int64, err error) {
	tx := dao.db.Model(condition).Where(condition)
	tx = SetPagination(tx, pagination)
	err = tx.Find(dest).Error
	if err != nil {
		return
	}
	err = tx.Count(&count).Error
	return
}

func (dao postDao) FindOne(dest *entity.Post) error {
	return dao.db.First(dest).Error
}

// Update implements interfaces.IPostDao
func (dao postDao) Update(src *entity.Post) error {
	return dao.db.Updates(src).Error
}

func NewPostDao(db *gorm.DB) interfaces.IPostDao {
	return postDao{
		db: db,
	}
}

// 事先宣告為 interface 才能在 compile time 進行檢查
var PostDao interfaces.IPostDao

func initPostDao() {
	PostDao = NewPostDao(db.DB)
}
