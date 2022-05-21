package middleware

import (
	"go-demo/api/common"
	"go-demo/internal/service/jwt_service"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthHandler(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		panic(common.StatusError{
			Status:  401,
			Message: "Unauthorized",
		})
	}
	token := strings.Split(authHeader, "Bearer ")[1]
	subject, claims, err := jwt_service.ValidateToken(token)
	if err != nil {
		panic(common.StatusError{
			Status:  401,
			Message: "Unauthorized",
		})
	}
	ctx.Set("claims", claims)
	ctx.Set("subject", subject)
	ctx.Next()
}
