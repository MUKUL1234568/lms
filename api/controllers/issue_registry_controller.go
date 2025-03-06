package controllers

import (
	"library-management-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllIssuedBooks retrieves all issued books (Only for LibraryAdmin)
func GetAllIssuedBooks(c *gin.Context) {
	// Check if the user is a LibraryAdmin
	// libID, err := GetLibraryID(c)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }

	issuedBooks, err := services.GetAllIssuedBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"issued_books": issuedBooks})
}

// GetIssuedBooksByReader retrieves issued books for a specific reader
// func GetIssuedBooksByReader(c *gin.Context) {
// 	// Get user_id from session
// 	readerID, err := GetUserID(c)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Fetch issued books for this reader
// 	issuedBooks, err := services.GetIssuedBooksByReader(readerID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Success response
// 	c.JSON(http.StatusOK, gin.H{"issued_books": issuedBooks})
// }
