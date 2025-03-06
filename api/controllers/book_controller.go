package controllers

import (
	// "fmt"
	"fmt"
	"library-management-api/models"
	"library-management-api/services"
	"library-management-api/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

import "strings"

func IsBook(books []models.Book, isbn string) *models.Book {
	isbn = strings.TrimSpace(strings.ToLower(isbn)) // Normalize input
	for i := range books {
		fmt.Println(books[i].ISBN)
		if strings.TrimSpace(strings.ToLower(books[i].ISBN)) == isbn {
			return &books[i]
		}
	}
	return nil
}

func AddBook(c *gin.Context) {
	var request struct {
		ISBN        string `json:"isbn" binding:"required"`
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
	if err := validator.Validateisbn(request.ISBN); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} // ✅ Convert float64 to uint

	// Create Book using LibID from the request

	books, err := services.GetBooksByLibrary(libID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(books)
	fmt.Println(request.ISBN)
	bookk := IsBook(books, request.ISBN)

	fmt.Println(bookk)
	if bookk != nil {
		// ✅ Update total_copies and available_copies in the database
		err := services.UpdateBookCopies(request.ISBN, request.TotalCopies)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book copies"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book copies updated successfully", "Updated book": books})
		return
	}

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
	if err := services.AddBook(&book); err != nil {
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

	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

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

	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if book.LibID != libID {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not registred in this library"})
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
	book, err := services.GetBookByISBN(isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if book.LibID != libID {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not registred in this library"})
		return
	}
	// Call service to update book
	if err := services.UpdateBook(isbn, &updatedBook); err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "available copies can not be greater than total copies" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DeleteBook handles removing a book from the library
func DeleteBook(c *gin.Context) {
	isbn := c.Param("isbn")

	// Call service to delete book
	book, err := services.GetBookByISBN(isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if book.LibID != libID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not registred in this library"})
		return
	}

	if err := services.DeleteBook(isbn); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "  book deleted"})
}
