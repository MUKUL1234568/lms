package routes

import (
	"library-management-api/controllers"

	"github.com/gin-gonic/gin"
)

// BookRoutes registers book-related routes
func BookRoutes(router *gin.Engine) {
	bookGroup := router.Group("/books")
	{
		bookGroup.POST("/", controllers.AddBook)                        // ✅ Add a book
		bookGroup.GET("/library/:libID", controllers.GetBooksByLibrary) // ✅ Get books by library ID
		bookGroup.GET("/:isbn", controllers.GetBookByISBN)              // ✅ Get book by ISBN
		bookGroup.PUT("/:isbn", controllers.UpdateBook)                 // ✅ Update book details
		bookGroup.DELETE("/:isbn", controllers.DeleteBook)              // ✅ Delete a book
	}
}
