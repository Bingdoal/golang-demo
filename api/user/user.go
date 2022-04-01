package user

import (
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRoute(route *gin.RouterGroup) (group *gin.RouterGroup) {
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
	user := model.User{}
	users, err := user.FindAll()
	if err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
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
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     "id must be uint.",
		})
		return
	}
	err = user.FindOne()
	if err != nil {
		ctx.JSON(404, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
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
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     "id must be uint.",
		})
		return
	}
	if err := user.FindOne(); err != nil {
		ctx.JSON(404, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}

	post := model.Post{}
	post.AuthorID = user.ID
	posts, err := post.FindByUser()
	if err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
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
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}

	user := model.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	if err := user.Create(); err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
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
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     "id must be uint.",
		})
		return
	}
	if err := user.FindOne(); err != nil {
		ctx.JSON(404, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}

	userDto := dto.UserDto{}
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}

	if userDto.Name != "" {
		user.Name = userDto.Name
	}
	if userDto.Email != "" {
		user.Email = userDto.Email
	}
	if err := user.Update(); err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}
	ctx.Status(204)
}

func deleteUser(ctx *gin.Context) {
	var err error
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     "id is required.",
		})
		return
	}
	user := model.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     "id must be uint.",
		})
		return
	}
	err = user.Delete()
	if err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}
	ctx.Status(204)
}
