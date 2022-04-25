package post

import (
	"go-demo/api/common"
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostRoute struct{}

func (p *PostRoute) AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup) {
	group = route.Group("/post")
	group.Use(preMiddleware...)

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
	postDto := dto.PostDto{}
	if err := ctx.BindJSON(&postDto); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	post := model.Post{
		Content:  postDto.Content,
		AuthorID: postDto.AuthorID,
	}

	if err := post.Create(); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	ctx.JSON(201, post)
}

func updatePost(ctx *gin.Context) {
	var id = ctx.Param("id")
	var err error
	post := model.Post{}
	post.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := post.FindOne(); err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}
	postDto := dto.PostDto{}

	if err := ctx.BindJSON(&postDto); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	post.Content = postDto.Content

	ctx.Status(204)
}

func deletePost(ctx *gin.Context) {
	var id = ctx.Param("id")
	var err error
	post := model.Post{}
	post.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := post.FindOne(); err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}

	if err := post.Delete(); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	ctx.Status(204)
}
