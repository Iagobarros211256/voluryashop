package rotes

import (
	"os/user"

	"github.com/Iagobarros211256/voluryashop/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		user.Group.POST("/", handlers.CreateUser)
	}

	return r
}
