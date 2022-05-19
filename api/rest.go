package api

import (
	"fmt"
	"go-demo/api/common"
	"go-demo/config"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	Server *gin.Engine
}

func (r *Rest) Add(root string, routes ...common.IApiRoute) *Rest {
	group := r.Server.Group(root)
	for _, route := range routes {
		route.AddRoute(group)
	}
	return r
}

func (r *Rest) AddWithMiddleware(root string,
	middleware gin.HandlerFunc,
	routes ...common.IApiRoute) *Rest {
	group := r.Server.Group(root)
	for _, route := range routes {
		route.AddRoute(group, middleware)
	}
	return r
}

func (r *Rest) Run() {
	fmt.Printf("\n============ Start [%s] version:%s on:%s ============\n",
		config.Env.GetString("name"),
		config.Env.GetString("version"),
		config.Env.GetString("server.port"))
	r.Server.Run(":" + config.Env.GetString("server.port"))
}
