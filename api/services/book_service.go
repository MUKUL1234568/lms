package services

import (
	"errors"
	"library-management-api/config"
	"library-management-api/models"
)

// AddBook adds a new book to the library
func AddBook(book *models.Book) error {
	return config.DB.Create(book).Error
}

// GetBooksByLibrary retrieves all books in a specific library
func GetBooksByLibrary(libID uint) ([]models.Book, error) {
	var books []models.Book
	err := config.DB.Preload("IssueRecords").Preload("Requests").Where("lib_id = ?", libID).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookByISBN fetches a single book by its ISBN
func GetBookByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	err := config.DB.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return nil, errors.New("book not found")
	}
	return &book, nil
}

// UpdateBook updates an existing book's details
func UpdateBook(isbn string, updatedBook *models.Book) error {
	var book models.Book
	err := config.DB.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return errors.New("book not found")
	}

	// Update book details
	return config.DB.Model(&book).Updates(updatedBook).Error
}

// DeleteBook removes a book from the library
func DeleteBook(isbn string) error {
	var book models.Book

	// Check if the book exists
	if err := config.DB.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		return errors.New("book not found")
	}

	// If available copies > 0, decrement instead of deleting
	if book.AvailableCopies > 0 {
		book.AvailableCopies -= 1
		return config.DB.Save(&book).Error
	}

	// If no copies left, delete the book
	result := config.DB.Where("isbn = ?", isbn).Delete(&models.Book{})
	if result.RowsAffected == 0 {
		return errors.New("no book to delete")
	}

	return nil
}
