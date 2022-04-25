package user

import (
	"go-demo/api/common"
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (u *UserRoute) AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup) {
	group = route.Group("/user")
	group.Use(preMiddleware...)

	group.GET("/", getUsers)
	group.GET("/:id", getOneUser)
	group.GET("/:id/post", getUserPosts)
	group.POST("/", createUser)
	group.PUT("/:id", updateUser)
	group.DELETE("/:id", deleteUser)
	return
}

func getUsers(ctx *gin.Context) {
	user := model.User{}
	users, err := user.FindAll()
	if err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	} else {
		ctx.JSON(200, dto.RespDto{
			Message: enum.MessageType(enum.Success),
			Data:    users,
		})
	}
}

func getOneUser(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := model.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	err = user.FindOne()
	if err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}
	ctx.JSON(200, dto.RespDto{
		Message: enum.MessageType(enum.Success),
		Data:    user,
	})
}

func getUserPosts(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := model.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := user.FindOne(); err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}

	post := model.Post{}
	post.AuthorID = user.ID
	posts, err := post.FindByUser()
	if err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}
	ctx.JSON(200, dto.RespDto{
		Message: enum.MessageType(enum.Success),
		Data:    posts,
	})
}

func createUser(ctx *gin.Context) {
	userDto := dto.UserDto{}
	if err := ctx.BindJSON(&userDto); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	user := model.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	if err := user.Create(); err != nil {
		common.RespError(ctx, 400, err.Error())
		return

	}
	ctx.JSON(201, dto.RespDto{
		Message: enum.MessageType(enum.Success),
	})
}

func updateUser(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := model.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := user.FindOne(); err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}

	userDto := dto.UserDto{}
	if err := ctx.BindJSON(&userDto); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	if userDto.Name != "" {
		user.Name = userDto.Name
	}
	if userDto.Email != "" {
		user.Email = userDto.Email
	}
	if err := user.Update(); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}
	ctx.Status(204)
}

func deleteUser(ctx *gin.Context) {
	var err error
	id, ok := ctx.Params.Get("id")
	if !ok {
		common.RespError(ctx, 400, "id is required.")
		return
	}
	user := model.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	err = user.Delete()
	if err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}
	ctx.Status(204)
}
