package main

import (
	"basicMicroservice/config"
	"basicMicroservice/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	port := config.GetEnv("PORT", "8080")
	router := gin.Default()
	routes.RegisterUserRoutes(router)

	log.Printf("Server is running at: %s ", port)
	router.Run(":" + port)
}
