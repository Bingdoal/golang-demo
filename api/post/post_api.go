package post

import (
	"fmt"
	"go-demo/api/common"
	"go-demo/internal/dto"
	"go-demo/internal/dto/basic"
	"go-demo/internal/enum"
	"go-demo/internal/model/dao"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
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
	pagination, err := common.GetPagination(ctx)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	var filter entity.Post
	if err := ctx.ShouldBind(&filter); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	var posts entity.Posts
	fmt.Printf("filter: %+v\n", filter)
	pagination.Total, err = p.postDao.FindAll(filter, pagination, &posts)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(200, basic.RespDto{
			Message:    enum.MessageType(enum.Success),
			Data:       posts,
			Pagination: &pagination,
		})
	}
}

func (p postApi) createPost(ctx *gin.Context) {
	postDto := dto.PostDto{}
	if err := ctx.BindJSON(&postDto); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	post := entity.Post{
		Content:  postDto.Content,
		AuthorID: postDto.AuthorID,
	}

	if err := p.postDao.Create(&post); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	ctx.JSON(201, basic.RespDto{
		Message: enum.MessageType(enum.Success),
		Data: basic.CreatedDto{
			ID: post.ID,
		},
	})
}

func (p postApi) updatePost(ctx *gin.Context) {
	var id = ctx.Param("id")
	var err error
	post := entity.Post{}
	post.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	if err := p.postDao.FindOne(&post); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	var postDto dto.PostDto

	if err := ctx.BindJSON(&postDto); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	post.Content = postDto.Content
	if err := p.postDao.Update(&post); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	ctx.Status(204)
}

func (p postApi) deletePost(ctx *gin.Context) {
	var id = ctx.Param("id")
	var err error
	post := entity.Post{}
	post.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	if err := p.postDao.FindOne(&post); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	if err := p.postDao.Delete(post.ID); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	ctx.Status(204)
}
