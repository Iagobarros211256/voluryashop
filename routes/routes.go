package routes

import (
	"github.com/Iagobarros211256/voluryashop/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.POST("/", handlers.CreateUser)
	}

	return r
}
