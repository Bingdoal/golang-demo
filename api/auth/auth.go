package auth

import (
	"fmt"
	"go-demo/internal/dto"
	"go-demo/internal/enum"
	"go-demo/internal/middleware"
	"go-demo/internal/model"
	"go-demo/internal/service/jwt_service"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct{}

func (a *AuthRoute) AddRoute(route *gin.RouterGroup, preMiddleware ...gin.HandlerFunc) (group *gin.RouterGroup) {
	group = route.Group("/auth")
	group.Use(preMiddleware...)

	group.POST("/login", login)
	group.POST("/refresh", middleware.AuthHandler, refresh)
	group.POST("/logout", middleware.AuthHandler, logout)
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
	token := jwt_service.GenerateToken(fmt.Sprint(user.ID), map[string]string{
		"name":  user.Name,
		"email": user.Email,
	})
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func refresh(ctx *gin.Context) {
	subject := ctx.GetString("subject")
	claims := ctx.GetStringMapString("claims")
	token := jwt_service.GenerateToken(subject, claims)
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "logout",
	})
}
