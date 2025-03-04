package routes

import (
	"github.com/gin-gonic/gin"
	"library-management-api/controllers"
	"library-management-api/middleware"
)

// BookRoutes registers book-related routes
func BookRoutes(router *gin.Engine) {
	bookGroup := router.Group("/books")
	{
		// ✅ Public Route (Anyone Can View Books)

		bookGroup.Use(middleware.AuthMiddleware())
		bookGroup.GET("/:isbn", controllers.GetBookByISBN)
		bookGroup.GET("/lib", controllers.GetBooksByLibrary)

		bookGroup.Use(middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}))
		bookGroup.POST("/", controllers.AddBook)
		bookGroup.PUT("/:isbn", controllers.UpdateBook)
		bookGroup.DELETE("/:isbn", controllers.DeleteBook)
	}
}
