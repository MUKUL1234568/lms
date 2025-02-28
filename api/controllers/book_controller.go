package controllers

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"library-management-api/models"
	"library-management-api/services"
	"net/http"
)

// AddBook handles adding a new book to the library
// func AddBook(c *gin.Context) {
// 	// Get user making the request (Assuming middleware sets user info)
// 	user, exists := c.Get("user") // Fetch the user from the context
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	// Type assert user to models.User
// 	userModel, ok := user.(models.User)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
// 		return
// 	}

// 	// Only LibraryAdmin & Owner can add books
// 	if userModel.Role != "LibraryAdmin" && userModel.Role != "Owner" {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
// 		return
// 	}

// 	var request struct {
// 		ISBN            string `json:"isbn" binding:"required"`
// 		Title           string `json:"title" binding:"required"`
// 		Authors         string `json:"authors" binding:"required"`
// 		Publisher       string `json:"publisher" binding:"required"`
// 		Version         string `json:"version"`
// 		TotalCopies     int    `json:"total_copies" binding:"required"`
// 		AvailableCopies int    `json:"available_copies" binding:"required"`
// 	}

// 	// Bind JSON request
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Create Book with `LibID` from the user's profile
// 	book := models.Book{
// 		ISBN:            request.ISBN,
// 		LibID:           userModel.LibID, // ✅ Assign LibID from user's profile
// 		Title:           request.Title,
// 		Authors:         request.Authors,
// 		Publisher:       request.Publisher,
// 		Version:         request.Version,
// 		TotalCopies:     request.TotalCopies,
// 		AvailableCopies: request.AvailableCopies,
// 	}

// 	// Call service to add book
// 	err := services.AddBook(&book)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Success response
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Book added successfully",
// 		"book":    book,
// 	})
// }

func AddBook(c *gin.Context) {
	var request struct {
		ISBN        string `json:"isbn" binding:"required"`
		LibID       uint   `json:"lib_id"` // ✅ Temporarily taking LibID from request
		Title       string `json:"title" binding:"required"`
		Authors     string `json:"authors" binding:"required"`
		Publisher   string `json:"publisher" binding:"required"`
		Version     string `json:"version"`
		TotalCopies int    `json:"total_copies" binding:"required"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	libIDInterface, exists := c.Get("lib_id") // ✅ Use lowercase "lib_id"
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// // Convert interface{} to float64 first, then to uint
	libIDFloat, ok := libIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Library ID format"})
		return
	}
	libID := uint(libIDFloat) // ✅ Convert float64 to uint

	// Create Book using LibID from the request
	book := models.Book{
		ISBN:            request.ISBN,
		LibID:           libID,
		Title:           request.Title,
		Authors:         request.Authors,
		Publisher:       request.Publisher,
		Version:         request.Version,
		TotalCopies:     request.TotalCopies,
		AvailableCopies: request.TotalCopies,
	}

	// Call service to add book
	err := services.AddBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Book added successfully",
		"book":    book,
	})
}

// GetBooksByLibrary retrieves all books for a specific library
func GetBooksByLibrary(c *gin.Context) {
	// libIDstr := c.Param("lib_id")
	// Ensure correct key name (Check how it's stored in AuthMiddleware)
	libIDInterface, exists := c.Get("lib_id") // ✅ Use lowercase "lib_id"
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// // Convert interface{} to float64 first, then to uint
	libIDFloat, ok := libIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Library ID format"})
		return
	}
	libID := uint(libIDFloat) // ✅ Convert float64 to uint

	// Call service to fetch books
	// libIDunit64, err := strconv.ParseUint(libIDstr, 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid libid type"})
	// }

	books, err := services.GetBooksByLibrary(libID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"books": books})
}

// GetBookByISBN retrieves a single book by ISBN
func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	// Call service to fetch book
	book, err := services.GetBookByISBN(isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"book": book})
}

// UpdateBook handles updating book details
func UpdateBook(c *gin.Context) {
	isbn := c.Param("isbn")
	var updatedBook models.Book

	// Bind JSON request
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to update book
	err := services.UpdateBook(isbn, &updatedBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DeleteBook handles removing a book from the library
func DeleteBook(c *gin.Context) {
	isbn := c.Param("isbn")

	// Call service to delete book
	err := services.DeleteBook(isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func SearchByFilter(c *gin.Context) {

}
