package user

import (
	"go-demo/internal/model"

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
		ctx.JSON(400, gin.H{
			"message": "error",
			"err":     err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    users,
		})
	}
}

func getOneUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func getUserPosts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func createUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func updateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}

func deleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}
