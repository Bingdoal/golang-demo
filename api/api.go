package api

import (
	"go-demo/api/auth"
	"go-demo/api/post"
	"go-demo/api/user"
	"go-demo/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRoute() (route *gin.Engine) {
	if config.Env.GetString("mode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	route = gin.Default()
	route.Use(cors.Default())
	route.GET("/", hello)

	v1 := route.Group("/v1")
	user.AddRoute(v1)
	post.AddRoute(v1)
	auth.AddRoute(v1)

	return
}

func hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}
