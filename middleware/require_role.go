package middleware

import (
	"net/http"

	"github.com/Hisyam/freepass-2026/models"
	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCtx, exists := c.Get("currentUser")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		user := userCtx.(models.User)

		if string(user.Role) != allowedRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden - Access denied"})
			return
		}

		c.Next()
	}
}