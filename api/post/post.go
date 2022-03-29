package post

import (
	"go-demo/internal/model"

	"github.com/gin-gonic/gin"
)

func AddRoute(route *gin.RouterGroup) (group *gin.RouterGroup) {
	group = route.Group("/post")

	group.GET("/", getPosts)
	group.POST("/", createPost)
	group.PUT("/:id", updatePost)
	group.DELETE("/:id", deletePost)
	return
}

func getPosts(ctx *gin.Context) {
	post := model.Post{}
	posts, err := post.FindAll()
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "error",
			"err":     err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    posts,
		})
	}
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
