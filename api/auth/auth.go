package auth

import (
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/model"

	"github.com/gin-gonic/gin"
)

func AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup) {
	group = route.Group("/auth")
	group.Use(preMiddleware...)

	group.POST("/login", login)
	group.POST("/register", register)
	group.POST("/logout", logout)
	return
}

func login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	if err := ctx.BindJSON(&loginDto); err != nil {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     err.Error(),
		})
		return
	}

	user := model.User{
		Name:     loginDto.Username,
		Password: loginDto.Password,
	}
	success := user.Login()
	if !success {
		ctx.JSON(400, dto.RespDto{
			Message: enum.MessageType(enum.Error),
			Err:     "username or password error.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}
