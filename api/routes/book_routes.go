package routes

import (
	"library-management-api/controllers"
	"library-management-api/middleware"

	"github.com/gin-gonic/gin"
)

// BookRoutes registers book-related routes
func BookRoutes(router *gin.Engine) {
	bookGroup := router.Group("/books")
	{
		// ✅ Public Route (Anyone Can View Books)
		bookGroup.GET("/:isbn", controllers.GetBookByISBN)
		bookGroup.GET("/lib/:lib_id", controllers.GetBooksByLibrary)
		bookGroup.POST("/filter", controllers.SearchByFilter)

		// ✅ Protected Routes (Only LibraryAdmins)
		protected := bookGroup.Group("/")
		protected.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("LibraryAdmin"))
		protected.POST("/", controllers.AddBook)
		protected.PUT("/:isbn", controllers.UpdateBook)
		protected.DELETE("/:isbn", controllers.DeleteBook)
	}
}
