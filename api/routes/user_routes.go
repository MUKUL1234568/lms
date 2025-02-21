package routes

import (
	"library-management-api/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes registers user-related routes
func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser) // ✅ Register User
		userGroup.POST("/login", controllers.LoginUser)       // ✅ Login User
		userGroup.GET("/:id", controllers.GetUser)            // ✅ Get User by ID
	}
}
