package main

import (
	"go-demo/config"
	"go-demo/config/db"
	"go-demo/internal/util/logger"
)

func main() {
	config.InitConfig("./_assets")
	logger.InitLogger()
	postgres := db.NewDB(config.Env)
	app := InitApp(postgres)
	app.Run()
}