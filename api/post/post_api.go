package post

import (
	"go-demo/api/common"
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/model/dao"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
	"go-demo/internal/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postApi struct {
	postDao interfaces.IPostDao
}

var PostApi common.IApiRoute

func Init() {
	PostApi = NewPostApi(dao.PostDao)
}

func NewPostApi(postDao interfaces.IPostDao) common.IApiRoute {
	util.IfNilPanic(postDao)
	return &postApi{
		postDao: postDao,
	}
}

func (p postApi) AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup) {
	group = route.Group("/post")
	group.Use(preMiddleware...)

	group.GET("/", p.getPosts)
	group.POST("/", p.createPost)
	group.PUT("/:id", p.updatePost)
	group.DELETE("/:id", p.deletePost)
	return
}

func (p postApi) getPosts(ctx *gin.Context) {
	var posts entity.Posts
	err := p.postDao.FindAll(&posts)
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

func (p postApi) createPost(ctx *gin.Context) {
	postDto := dto.PostDto{}
	if err := ctx.BindJSON(&postDto); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	post := entity.Post{
		Content:  postDto.Content,
		AuthorID: postDto.AuthorID,
	}

	if err := p.postDao.Create(&post); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	ctx.JSON(201, post)
}

func (p postApi) updatePost(ctx *gin.Context) {
	var id = ctx.Param("id")
	var err error
	post := entity.Post{}
	post.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := p.postDao.FindOne(&post); err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}
	var postDto dto.PostDto

	if err := ctx.BindJSON(&postDto); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	post.Content = postDto.Content
	if err := p.postDao.Update(&post); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}
	ctx.Status(204)
}

func (p postApi) deletePost(ctx *gin.Context) {
	var id = ctx.Param("id")
	var err error
	post := entity.Post{}
	post.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := p.postDao.FindOne(&post); err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}

	if err := p.postDao.Delete(post.ID); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	ctx.Status(204)
}
