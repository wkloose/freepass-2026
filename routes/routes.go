package routes

import (
	"github.com/Hisyam/freepass-2026/controllers"
	"github.com/Hisyam/freepass-2026/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth)
	{
		protected.GET("/me", userController.Me)
		protected.PUT("/me", userController.UpdateProfile)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.RequireAuth, middleware.RequireRole("ADMIN"))
	{
		admin.POST("/users", userController.CreateUser)
	}
}