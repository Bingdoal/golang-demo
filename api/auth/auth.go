package auth

import (
	"fmt"
	"go-demo/api/common"
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/middleware"
	"go-demo/internal/model/dao"
	"go-demo/internal/model/dao/interfaces"
	"go-demo/internal/model/entity"
	"go-demo/internal/service/jwt_service"

	"github.com/gin-gonic/gin"
)

type authApi struct {
	userDao interfaces.IUserDao
}

var AuthApi common.IApiRoute

func Init() {
	AuthApi = NewAuthApi(dao.UserDao)
}

func NewAuthApi(userDao interfaces.IUserDao) common.IApiRoute {
	return &authApi{
		userDao: userDao,
	}
}

func (a authApi) AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup) {
	group = route.Group("/auth")
	group.Use(preMiddleware...)

	group.POST("/login", a.login)
	group.POST("/refresh", middleware.AuthHandler, a.refresh)
	group.POST("/logout", middleware.AuthHandler, a.logout)
	return
}

func (a authApi) login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	if err := ctx.BindJSON(&loginDto); err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}

	user := entity.User{
		Name:     loginDto.Username,
		Password: loginDto.Password,
	}

	if err := a.userDao.Login(loginDto.Username, loginDto.Password); err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}

	token := jwt_service.GenerateToken(fmt.Sprint(user.ID), map[string]string{
		"name":  user.Name,
		"email": user.Email,
	})
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func (a authApi) refresh(ctx *gin.Context) {
	subject := ctx.GetString("subject")
	claims := ctx.GetStringMapString("claims")
	token := jwt_service.GenerateToken(subject, claims)
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func (a authApi) logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "logout",
	})
}
