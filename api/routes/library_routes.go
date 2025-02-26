package routes

import (
	"library-management-api/controllers"

	"github.com/gin-gonic/gin"
)

// LibraryRoutes registers routes for Library
func LibraryRoutes(router *gin.Engine) {
	libGroup := router.Group("/libraries")
	{
		libGroup.POST("/", controllers.CreateLibrary) // ✅ Create Library with Owner
		libGroup.GET("/", controllers.GetLibraries)   // ✅ Add this in your routes

	}
}
