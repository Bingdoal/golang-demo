package api

import (
	"go-demo/api/actuator"
	"go-demo/api/auth"
	"go-demo/api/post"
	"go-demo/api/user"
	"go-demo/config"
	"go-demo/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApiInstance() {
	auth.Init()
	post.Init()
	user.Init()
}

func NewRest(authApi *auth.TypeAuthApi,
	userApi *user.TypeUserApi,
	postApi *post.TypePostApi,
	actuatorApi *actuator.TypeActuatorApi) *Rest {
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

	rest.Add("/v1", authApi)
	rest.Add("/actuator", actuatorApi)
	rest.Middleware(middleware.AuthHandler).Add("/v1", userApi, postApi)
	return rest
}

func SetUpRoute() *Rest {
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

	rest.Add("/v1", auth.AuthApi)
	rest.Middleware(middleware.AuthHandler).Add("/v1", user.UserApi, post.PostApi)
	return rest
}
