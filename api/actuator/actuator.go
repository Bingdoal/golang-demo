package actuator

import (
	"go-demo/config"
	"net/http"

	"github.com/gin-gonic/gin"
	actuator "github.com/sinhashubham95/go-actuator"
)

type TypeActuatorApi struct {
	ActuatorHandler http.HandlerFunc
}

func NewActuatorApi() *TypeActuatorApi {
	config := &actuator.Config{
		Endpoints: []int{
			actuator.Info,
			actuator.Metrics,
			actuator.Ping,
			actuator.ThreadDump,
		},
		Env:     config.Env.GetString("mode"),
		Name:    config.Env.GetString("name"),
		Port:    config.Env.GetInt("server.port"),
		Version: config.Env.GetString("version"),
	}
	actuatorHandler := actuator.GetActuatorHandler(config)
	return &TypeActuatorApi{
		ActuatorHandler: actuatorHandler,
	}
}

func (api TypeActuatorApi) AddRoute(route *gin.RouterGroup) (group *gin.RouterGroup) {
	group = route.Group("/")
	group.GET("/*endpoint", api.ginActuatorHandler)
	return
}

func (api TypeActuatorApi) ginActuatorHandler(ctx *gin.Context) {
	api.ActuatorHandler(ctx.Writer, ctx.Request)
}
