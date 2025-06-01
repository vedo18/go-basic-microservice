package routes

import (
	"basicMicroservice/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", handlers.GetAllUsers)
		userRoutes.GET("/:id", handlers.GetUserById)
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.PATCH("/:id", handlers.UpdateUser)
	}
}
