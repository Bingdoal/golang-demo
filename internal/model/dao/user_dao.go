package dao

import (
	"errors"
	"go-demo/config/db"
	"go-demo/internal/dto/basic"
	"go-demo/internal/model/base"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
	"go-demo/internal/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userDao struct {
	db *gorm.DB
}

// Create implements interfaces.IUserDao
func (dao userDao) Create(src *entity.User) error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(src.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	src.Password = string(hashPwd)
	return dao.db.Create(src).Error
}

// Delete implements interfaces.IUserDao
func (dao userDao) Delete(id uint64) error {
	return dao.db.Delete(&entity.User{
		BaseModel: base.BaseModel{
			ID: id,
		},
	}).Error
}

// FindAll implements interfaces.IUserDao
func (dao userDao) FindAll(condition entity.User, pagination basic.Pagination, dest *entity.Users) (count int64, err error) {
	tx := dao.db.Model(condition).Where(condition)
	tx = SetPagination(tx, pagination)
	err = tx.Find(dest).Error
	if err != nil {
		return
	}
	err = tx.Count(&count).Error
	return
}

// FindOne implements interfaces.IUserDao
func (dao userDao) FindOne(dest *entity.User) error {
	return dao.db.Where(dest).First(dest).Error
}

// Login implements interfaces.IUserDao
func (dao userDao) Login(name string, password string) error {
	var user entity.User
	if err := dao.db.Where(&entity.User{
		Name: name,
	}).First(&user).Error; err != nil {
		return errors.New("invalid username or password")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid username or password")
	}
	return nil
}

// Update implements interfaces.IUserDao
func (dao userDao) Update(src *entity.User) error {
	return dao.db.Updates(src).Error
}

func NewUserDao(db *gorm.DB) interfaces.IUserDao {
	util.IfNilPanic(db)
	return &userDao{
		db: db,
	}
}

// 事先宣告為 interface 才能在 compile time 進行檢查
var UserDao interfaces.IUserDao

func initUserDao() {
	UserDao = NewUserDao(db.DB)
}
