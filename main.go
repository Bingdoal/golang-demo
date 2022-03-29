package main

import (
	"fmt"
	"go-demo/api"
	"go-demo/config"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func init() {
	route = api.GetRoute()
}

func main() {
	fmt.Printf("\n============ Start [%s] version:%s on:%s ============\n",
		config.Env.GetString("name"),
		config.Env.GetString("version"),
		config.Env.GetString("server.port"))
	route.Run(":" + config.Env.GetString("server.port"))
}
