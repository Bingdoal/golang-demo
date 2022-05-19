package util

import (
	"fmt"
	"go-demo/config"
	_ "go-demo/config/db"
	"go-demo/internal/model/dao"
	"go-demo/internal/model/entity"
	"go-demo/internal/util/logger"
)

func Init() {
	initAdminUser()
}

func initAdminUser() {
	admin := entity.User{
		Name: config.Env.GetString("features.admin.username"),
	}
	if err := dao.UserDao.FindOne(&admin); err != nil {
		logger.Debug.Println("Create admin user.")
		admin.Password = config.Env.GetString("features.admin.password")
		admin.Email = config.Env.GetString("features.admin.email")
		fmt.Printf("admin: %v\n", admin)
		dao.UserDao.Create(&admin)
	}
}
