package routes

import (
	"library-management-api/controllers"
	"library-management-api/middleware"

	// "library-management-api/middleware"

	"github.com/gin-gonic/gin"
)

// LibraryRoutes registers routes for Library
func LibraryRoutes(router *gin.Engine) {
	libGroup := router.Group("/libraries")
	{
		libGroup.POST("/", controllers.CreateLibrary) // ✅ Create Library with Owner
		libGroup.GET("/", controllers.GetLibraries)   // ✅ Add this in your routes
		libGroup.GET("/states", controllers.GetStates)
		libGroup.GET("/states/id", middleware.AuthMiddleware(), middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}), controllers.GetStatesBylib)
	}
}
