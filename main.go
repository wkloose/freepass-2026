package main

import (
	"fmt"
	"os"

	"github.com/Hisyam/freepass-2026/controllers"
	"github.com/Hisyam/freepass-2026/initializers"
	"github.com/Hisyam/freepass-2026/repositories"
	"github.com/Hisyam/freepass-2026/services"
	"github.com/gin-gonic/gin"
	"strings"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("Starting Server...")

	db := initializers.DB

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
    	port = "3000"
	}
	if !strings.HasPrefix(port, ":") {
    	port = ":" + port
	}
	r.Run(port)
}