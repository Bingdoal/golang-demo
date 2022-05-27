package user

import (
	"go-demo/api/common"
	"go-demo/config"
	"go-demo/internal/dto"
	"go-demo/internal/dto/basic"
	"go-demo/internal/enum"
	"go-demo/internal/model/base"
	"go-demo/internal/model/dao"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TypeUserApi struct {
	userDao interfaces.IUserDao
	postDao interfaces.IPostDao
}

type Options struct {
	UserDao interfaces.IUserDao
	PostDao interfaces.IPostDao
}

var UserApi common.IApiRoute

func Init() {
	UserApi = NewUserApi(Options{
		UserDao: dao.UserDao,
		PostDao: dao.PostDao,
	})
}

func NewUserApi(opt Options) *TypeUserApi {
	return &TypeUserApi{
		userDao: opt.UserDao,
		postDao: opt.PostDao,
	}
}

func (u TypeUserApi) AddRoute(route *gin.RouterGroup) (group *gin.RouterGroup) {
	group = route.Group("/user")

	group.GET("", u.getUsers)
	group.GET("/:id", u.getOneUser)
	group.POST("", u.createUser)
	group.PUT("/:id", u.updateUser)
	group.DELETE("/:id", u.deleteUser)
	return
}

func (u TypeUserApi) getUsers(ctx *gin.Context) {
	pagination, err := common.GetPagination(ctx)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	var filter entity.User
	if err := ctx.BindQuery(&filter); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	var users entity.Users
	pagination.Total, err = u.userDao.FindAll(filter, pagination, &users)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(200, basic.RespDto{
			Message:    enum.MessageType(enum.Success),
			Data:       users,
			Pagination: &pagination,
		})
	}
}

func (u TypeUserApi) getOneUser(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := entity.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	err = u.userDao.FindOne(&user)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	ctx.JSON(200, basic.RespDto{
		Message: enum.MessageType(enum.Success),
		Data:    user,
	})
}

func (u TypeUserApi) createUser(ctx *gin.Context) {
	userDto := dto.UserDto{}
	if err := ctx.BindJSON(&userDto); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	user := entity.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	if err := u.userDao.Create(&user); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})

	}
	ctx.JSON(201, basic.RespDto{
		Message: enum.MessageType(enum.Success),
		Data: basic.CreatedDto{
			ID: user.ID,
		},
	})
}

func (u TypeUserApi) updateUser(ctx *gin.Context) {
	var id, _ = ctx.Params.Get("id")
	var err error
	user := entity.User{}
	user.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	if err := u.userDao.FindOne(&user); err != nil {
		panic(common.StatusError{
			Status:  404,
			Message: err.Error(),
		})
	}

	userDto := dto.UserDto{}
	if err := ctx.BindJSON(&userDto); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}

	user.Email = userDto.Email

	if err := u.userDao.Update(&user); err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	ctx.Status(204)
}

func (u TypeUserApi) deleteUser(ctx *gin.Context) {
	idStr, ok := ctx.Params.Get("id")
	if !ok {
		panic(common.StatusError{
			Status:  400,
			Message: "id is required",
		})
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	user := entity.User{BaseModel: base.BaseModel{ID: id}}
	if err := u.userDao.FindOne(&user); err != nil {
		panic(common.StatusError{
			Status:  404,
			Message: err.Error(),
		})
	}
	if user.Name == config.Env.GetString("features.admin.username") {
		panic(common.StatusError{
			Status:  400,
			Message: "Can't delete admin user.",
		})
	}
	err = u.userDao.Delete(id)
	if err != nil {
		panic(common.StatusError{
			Status:  400,
			Message: err.Error(),
		})
	}
	ctx.Status(204)
}
