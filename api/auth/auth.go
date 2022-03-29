package auth

import "github.com/gin-gonic/gin"

func AddRoute(route *gin.RouterGroup) (group *gin.RouterGroup) {
	group = route.Group("/auth")

	group.POST("/login", login)
	group.POST("/register", register)
	group.POST("/logout", logout)
	return
}

func login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}
