package services

import (
	"library-management-api/config"
	"library-management-api/models"

	"gorm.io/gorm"
)

// GetAllIssuedBooks fetches all issued books (LibraryAdmin Only)
func GetAllIssuedBooks(libId uint) ([]models.IssueRegistry, error) {
	var issuedBooks []models.IssueRegistry
	err := config.DB.
		Joins("JOIN books ON books.isbn = issue_registries.isbn").
		Where("books.lib_id = ?", libId). // Filter by book's lib_id
		Preload("Book", func(db *gorm.DB) *gorm.DB {
			return db.Select("isbn, title, authors, lib_id")
		}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, email, lib_id")
		}).
		Find(&issuedBooks).Error

	if err != nil {
		return nil, err
	}
	return issuedBooks, nil
}

// GetIssuedBooksByReader fetches issued books for a specific reader
// func GetIssuedBooksByReader(readerID uint) ([]models.IssueRegistry, error) {
// 	var issuedBooks []models.IssueRegistry
// 	err := config.DB.Where("reader_id = ?", readerID).Find(&issuedBooks).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return issuedBooks, nil
// }
