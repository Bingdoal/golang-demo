package post

import (
	"go-demo/api/common"
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/model"

	"github.com/gin-gonic/gin"
)

type PostRoute struct{}

func (p *PostRoute) AddRoute(route *gin.RouterGroup) (group *gin.RouterGroup) {
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
		common.RespError(ctx, 400, err.Error())
		return
	} else {
		ctx.JSON(200, dto.RespDto{
			Message: enum.MessageType(enum.Success),
			Data:    posts,
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
