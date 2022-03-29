package post

import "github.com/gin-gonic/gin"

func AddRoute(route *gin.Engine) (group *gin.RouterGroup) {
	group = route.Group("/post")

	group.GET("/", getPosts)
	group.POST("/", createPost)
	group.PUT("/:id", updatePost)
	group.DELETE("/:id", deletePost)
	return
}

func getPosts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func createPost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func updatePost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func deletePost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}
