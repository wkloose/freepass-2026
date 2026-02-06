package controllers

import (
	"net/http"

	"github.com/Hisyam/freepass-2026/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.Register(input)
	if err != nil {
		if err.Error() == "email is registered" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
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

func (ctrl *UserController) Me(c *gin.Context) {
	userCtx, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
		return
	}

	currentUser := userCtx.(models.User)

	user, err := ctrl.userService.GetProfile(currentUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl *UserController) UpdateProfile(c *gin.Context) {
    userCtx, exists := c.Get("currentUser")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
        return
    }
    currentUser := userCtx.(models.User)

    var input services.UpdateProfileInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedUser, err := ctrl.userService.UpdateProfile(currentUser.ID, input)
    if err != nil {
        if err.Error() == "email has been used by another user" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Profile updated successfully",
        "data":    updatedUser,
    })
}

