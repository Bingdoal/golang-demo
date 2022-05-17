package api

import (
	"fmt"
	"go-demo/api/auth"
	"go-demo/api/post"
	"go-demo/api/user"
	"go-demo/config"
	"go-demo/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IApiRoute interface {
	AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup)
}

type Rest struct {
	Server *gin.Engine
}

func (r *Rest) Add(root string, routes ...IApiRoute) *Rest {
	group := r.Server.Group(root)
	for _, route := range routes {
		route.AddRoute(group)
	}
	return r
}

func (r *Rest) AddWithMiddleware(root string,
	middleware gin.HandlerFunc,
	routes ...IApiRoute) *Rest {
	group := r.Server.Group(root)
	for _, route := range routes {
		route.AddRoute(group, middleware)
	}
	return r
}

func NewRest() *Rest {
	if config.Env.GetString("mode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginServer := gin.New()
	ginServer.Use(gin.Logger(),
		gin.CustomRecovery(middleware.ErrorHandler()),
		cors.Default())
	rest := &Rest{
		Server: ginServer,
	}

	return rest
}

func (r *Rest) Run() {
	fmt.Printf("\n============ Start [%s] version:%s on:%s ============\n",
		config.Env.GetString("name"),
		config.Env.GetString("version"),
		config.Env.GetString("server.port"))
	r.Server.Run(":" + config.Env.GetString("server.port"))
}

func SetUpRoute() *Rest {
	rest := NewRest()
	rest.Add("/v1", &auth.AuthApi)
	rest.AddWithMiddleware("/v1", middleware.AuthHandler,
		&user.UserApi, &post.PostApi)
	return rest
}
