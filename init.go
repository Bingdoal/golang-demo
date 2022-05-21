package main

import (
	"go-demo/api"
	"go-demo/config"
	"go-demo/config/db"
	"go-demo/internal/model/dao"
	"go-demo/internal/model/entity"
	"go-demo/internal/service"
	"go-demo/internal/util/logger"
)

var envPath = "_assets/"

func Initialization() {
	config.InitConfig(envPath)
	logger.InitLogger()
	db.InitDB()
	dao.InitDao()
	service.InitService()
	api.InitApiInstance()
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
		dao.UserDao.Create(&admin)
	}
}
