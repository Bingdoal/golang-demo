package main

import (
	"go-demo/api"
	"go-demo/config"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
	"go-demo/internal/util/logger"
)

type Application struct {
	rest    *api.Rest
	userDao interfaces.IUserDao
}

func NewApp(rest *api.Rest, userDao interfaces.IUserDao) *Application {
	return &Application{
		rest:    rest,
		userDao: userDao,
	}
}

func (app Application) Run() {
	app.initAdminUser()
	app.rest.Run()
}

func (app Application) initAdminUser() {
	admin := entity.User{
		Name: config.Env.GetString("features.admin.username"),
	}
	if err := app.userDao.FindOne(&admin); err != nil {
		logger.Debug.Println("Create admin user.")
		admin.Password = config.Env.GetString("features.admin.password")
		admin.Email = config.Env.GetString("features.admin.email")
		app.userDao.Create(&admin)
	}
}
