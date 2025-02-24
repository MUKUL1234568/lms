package routes

import (
	"library-management-api/controllers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes registers authentication-related routes
func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login) // ✅ User Login
	}
}
