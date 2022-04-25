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

type RouteInterface interface {
	AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup)
}

type Rest struct {
	server *gin.Engine
}

func (r *Rest) Add(root string, routes ...RouteInterface) *Rest {
	group := r.server.Group(root)
	for _, route := range routes {
		route.AddRoute(group)
	}
	return r
}

func (r *Rest) AddWithMiddleware(root string,
	middleware gin.HandlerFunc,
	routes ...RouteInterface) *Rest {
	group := r.server.Group(root)
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
		server: ginServer,
	}

	return rest
}

func (r *Rest) Run() {
	fmt.Printf("\n============ Start [%s] version:%s on:%s ============\n",
		config.Env.GetString("name"),
		config.Env.GetString("version"),
		config.Env.GetString("server.port"))
	r.server.Run(":" + config.Env.GetString("server.port"))
}

func SetUpRoute() *Rest {
	rest := NewRest()
	rest.Add("/v1", &auth.AuthRoute{})
	rest.AddWithMiddleware("/v1", middleware.AuthHandler,
		&user.UserRoute{}, &post.PostRoute{})
	return rest
}
