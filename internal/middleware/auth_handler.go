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
		common.RespError(ctx, 401, "Unauthorized")
		ctx.Abort()
		return
	}
	token := strings.Split(authHeader, "Bearer ")[1]
	subject, claims, err := jwt_service.ValidateToken(token)
	if err != nil {
		common.RespError(ctx, 401, "Unauthorized")
		ctx.Abort()
		return
	}
	ctx.Set("claims", claims)
	ctx.Set("subject", subject)
	ctx.Next()
}
