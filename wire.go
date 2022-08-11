//go:build wireinject
// +build wireinject

package main

import (
	"go-demo/api"
	"go-demo/api/actuator"
	"go-demo/api/auth"
	"go-demo/api/post"
	"go-demo/api/user"
	"go-demo/internal/model/dao"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB) *Application {
	wire.Build(
		wire.Struct(new(user.Options), "*"),
		dao.NewPostDao,
		dao.NewUserDao,
		user.NewUserApi,
		auth.NewAuthApi,
		post.NewPostApi,
		actuator.NewActuatorApi,
		api.NewRest,
		NewApp,
	)
	return nil
}
