package services

import (
	"library-management-api/config"
	"library-management-api/models"

	"gorm.io/gorm"
)

// GetAllIssuedBooks fetches all issued books (LibraryAdmin Only)
func GetAllIssuedBooks() ([]models.IssueRegistry, error) {
	var issuedBooks []models.IssueRegistry
	err := config.DB.
		Preload("Book", func(db *gorm.DB) *gorm.DB {
			return db.Select("isbn, title, authors")
		}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, email")
		}).
		Find(&issuedBooks).Error
	if err != nil {
		return nil, err
	}
	return issuedBooks, nil
}

// GetIssuedBooksByReader fetches issued books for a specific reader
func GetIssuedBooksByReader(readerID uint) ([]models.IssueRegistry, error) {
	var issuedBooks []models.IssueRegistry
	err := config.DB.Where("reader_id = ?", readerID).Find(&issuedBooks).Error
	if err != nil {
		return nil, err
	}
	return issuedBooks, nil
}
