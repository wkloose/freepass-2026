package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Hisyam/freepass-2026/controllers"
	"github.com/Hisyam/freepass-2026/initializers"
	"github.com/Hisyam/freepass-2026/repositories"
	"github.com/Hisyam/freepass-2026/routes"
	"github.com/Hisyam/freepass-2026/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	initializers.SeedAdmin()
}

func main() {
	fmt.Println("Starting Server...")

	db := initializers.DB

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	adminService := services.NewAdminService(userRepo)
	adminController := controllers.NewAdminController(adminService)

	r := gin.Default()

	routes.SetupRoutes(r, userController, adminController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	r.Run(port)
}
