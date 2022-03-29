package user

import "github.com/gin-gonic/gin"

func AddRoute(route *gin.Engine) (group *gin.RouterGroup) {
	group = route.Group("/user")

	group.GET("/", getUsers)
	group.GET("/:id", getOneUser)
	group.GET("/:id/post", getUserPosts)
	group.POST("/", createUser)
	group.PUT("/:id", updateUser)
	group.DELETE("/:id", deleteUser)
	return
}

func getUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func getOneUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func getUserPosts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func createUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func updateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func deleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}
