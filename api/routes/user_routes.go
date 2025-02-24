package routes

import (
	"library-management-api/controllers"
	"library-management-api/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoutes registers user-related routes
func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser) // ✅ Public (Anyone Can Register)

		// Protected User Routes
		userGroup.Use(middleware.AuthMiddleware())
		userGroup.GET("/", middleware.RoleMiddlewareMultiple([]string{"Owner", "LibraryAdmin"}), controllers.GetUsersByLibrary)
		// ✅ Requires Authentication
		userGroup.GET("/:id", controllers.GetUser) // ✅ Only Authenticated Users Can Fetch Their Profile
	}
}

// {"isbn": "978-0134190440",
// "lib_id": 1,
// "title": "The Go Programming Language",
// "authors": "Alan A. A. Donovan, Brian W. Kernighan",
// "publisher": "Addison-Wesley",
// "version": "1st Edition",
// "total_copies": 5,
// "available_copies": 7
// }
