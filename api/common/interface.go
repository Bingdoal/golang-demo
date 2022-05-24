package common

import "github.com/gin-gonic/gin"

type IApiRoute interface {
	AddRoute(route *gin.RouterGroup) (group *gin.RouterGroup)
}
