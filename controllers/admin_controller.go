package controllers

import (
	"net/http"

	"github.com/Hisyam/freepass-2026/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/Hisyam/freepass-2026/models"
)

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) *AdminController {
	return &AdminController{adminService}
}

func (ctrl *AdminController) CreateUser(c *gin.Context) {
	var input services.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.adminService.CreateUser(input)
	if err != nil {
		if err.Error() == "email has been used by another user" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    user,
	})
}

func (ctrl *AdminController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	var input services.UpdateUserByAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := ctrl.adminService.UpdateUser(userID, input)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "email has been used by another user" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data":    updatedUser,
	})
}

func (ctrl *AdminController) DeleteUser(c *gin.Context) {
    idParam := c.Param("id")
    userID, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
        return
    }

    currentUserCtx, _ := c.Get("currentUser")
    currentUser := currentUserCtx.(models.User)
    
    if currentUser.ID == userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - Cannot delete own account"})
        return
    }

    err = ctrl.adminService.DeleteUser(userID)
    if err != nil {
        if err.Error() == "user not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "User deleted successfully",
    })
}