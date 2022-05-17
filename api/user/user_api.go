package user

import (
	"go-demo/api/common"
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/model/dao"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userApi struct {
	userDao interfaces.IUserDao
	postDao interfaces.IPostDao
}

func NewUserApi(userDao interfaces.IUserDao, postDao interfaces.IPostDao) userApi {
	return userApi{
		userDao: userDao,
		postDao: postDao,
	}
}

var UserApi = NewUserApi(dao.UserDao, dao.PostDao)

func (u userApi) AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup) {
	group = route.Group("/user")
	group.Use(preMiddleware...)

	group.GET("/", u.getUsers)
	group.GET("/:id", u.getOneUser)
	group.GET("/:id/post", u.getUserPosts)
	group.POST("/", u.createUser)
	group.PUT("/:id", u.updateUser)
	group.DELETE("/:id", u.deleteUser)
	return
}

func (u userApi) getUsers(ctx *gin.Context) {
	var users entity.Users
	err := u.userDao.FindAll(&users)
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

func (u userApi) getOneUser(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := entity.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	err = u.userDao.FindOne(&user)
	if err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}
	ctx.JSON(200, dto.RespDto{
		Message: enum.MessageType(enum.Success),
		Data:    user,
	})
}

func (u userApi) getUserPosts(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := entity.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := u.userDao.FindOne(&user); err != nil {
		common.RespError(ctx, 404, err.Error())
		return
	}

	var posts entity.Posts
	err = u.postDao.FindByUser(user.ID, &posts)
	if err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}
	ctx.JSON(200, dto.RespDto{
		Message: enum.MessageType(enum.Success),
		Data:    posts,
	})
}

func (u userApi) createUser(ctx *gin.Context) {
	userDto := dto.UserDto{}
	if err := ctx.BindJSON(&userDto); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}

	user := entity.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	if err := u.userDao.Create(&user); err != nil {
		common.RespError(ctx, 400, err.Error())
		return

	}
	ctx.JSON(201, dto.RespDto{
		Message: enum.MessageType(enum.Success),
	})
}

func (u userApi) updateUser(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := entity.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	if err := u.userDao.FindOne(&user); err != nil {
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
	if err := u.userDao.Update(&user); err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}
	ctx.Status(204)
}

func (u userApi) deleteUser(ctx *gin.Context) {
	idStr, ok := ctx.Params.Get("id")
	if !ok {
		common.RespError(ctx, 400, "id is required.")
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.RespError(ctx, 400, "id must be uint.")
		return
	}
	err = u.userDao.Delete(id)
	if err != nil {
		common.RespError(ctx, 400, err.Error())
		return
	}
	ctx.Status(204)
}
