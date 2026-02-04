package controllers

import (
	"net/http"

	"github.com/Hisyam/freepass-2026/services"
	"github.com/Hisyam/freepass-2026/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (ctrl *UserController) Register(c *gin.Context) {
	var input services.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	user, err := ctrl.userService.Register(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    user,
	})
}

func (ctrl *UserController) Login(c *gin.Context) {
	var input services.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	tokenString, err := ctrl.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	expiryHour := utils.GetEnvAsInt("JWT_EXPIRY_HOUR", 12)

	c.SetCookie(
		"Authorization",
		tokenString,
		3600*expiryHour,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}