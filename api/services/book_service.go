package services

import (
	"errors"
	"gorm.io/gorm"
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
	if updatedBook.TotalCopies != 0 {
		if updatedBook.AvailableCopies > updatedBook.TotalCopies {
			return errors.New("available copies can not be greater than total copies")
		}
	} else {
		if updatedBook.AvailableCopies > book.TotalCopies {
			return errors.New("available copies can not be greater than total copies")
		}
	}

	return config.DB.Model(&book).Updates(updatedBook).Error
}
func UpdateBookCopies(isbn string, totalcopies int) error {
	result := config.DB.Model(&models.Book{}).
		Where("isbn = ?", isbn).
		Updates(map[string]interface{}{
			"total_copies":     gorm.Expr("total_copies + ?", totalcopies),
			"available_copies": gorm.Expr("available_copies + ?", totalcopies),
		})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteBook removes a book from the library
func DeleteBook(isbn string) error {

	if err := config.DB.Where("isbn = ?", isbn).Delete(&models.Book{}).Error; err != nil {
		return err
	}

	return nil
}
