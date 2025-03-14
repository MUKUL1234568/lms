package services

import (
	"errors"
	"library-management-api/config"
	"library-management-api/models"
)

// CreateLibrary creates a new library and assigns the user as the Owner
func CreateLibrary(owner *models.User, library *models.Library) error {

	var existingUser models.User
	if err := config.DB.Where("email = ?", owner.Email).First(&existingUser).Error; err == nil {
		return errors.New("user with this email already exists")
	}
	var existingUserr models.Library
	if err := config.DB.Where("name = ?", library.Name).First(&existingUserr).Error; err == nil {
		return errors.New("choose different library name")
	}
	tx := config.DB.Begin()

	// Step 1: Create Library FIRST
	if err := tx.Create(library).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to create library")
	}

	// Step 2: Assign Library ID to Owner
	owner.LibID = library.ID // ✅ Correctly assigning Library ID to Owner
	owner.Role = "Owner"     // ✅ Set Role as Owner

	// Step 3: Create User (Owner)
	if err := tx.Create(owner).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to create library owner")
	}

	// Commit transaction
	tx.Commit()
	return nil
}

func GetAllLibraries() ([]models.Library, error) {
	var libraries []models.Library

	if err := config.DB.Preload("Users").Preload("Books").Find(&libraries).Error; err != nil {
		return nil, errors.New("failed to fetch libraries")
	}

	return libraries, nil
}
func Getlibbyid(id uint) error {
	var library *models.Library
	err := config.DB.Where("id=?", id).First(&library).Error
	return err
}

// Stats holds the count of users, libraries, and total book copies
type Stats struct {
	TotalLibraries int `json:"total_libraries"`
	TotalUsers     int `json:"total_users"`

	TotalBooks int `json:"total_books"`
}

// GetStats retrieves the total number of users, libraries, and total book copies
func GetStats() (*Stats, error) {
	var stats Stats
	var userCount int64
	var libraryCount int64
	var bookCopies int64

	// Count total users
	if err := config.DB.Model(&models.User{}).Count(&userCount).Error; err != nil {
		return nil, err
	}

	// Count total libraries
	if err := config.DB.Model(&models.Library{}).Count(&libraryCount).Error; err != nil {
		return nil, err
	}

	// Sum total book copies
	if err := config.DB.Model(&models.Book{}).Select("COALESCE(SUM(total_copies), 0)").Scan(&bookCopies).Error; err != nil {
		return nil, err
	}

	// Convert int64 to int
	stats.TotalUsers = int(userCount)
	stats.TotalLibraries = int(libraryCount)
	stats.TotalBooks = int(bookCopies)

	return &stats, nil
}

type Statsl struct {
	TotalIssuedBook int `json:"total_issued_book"`
	TotalUsers      int `json:"total_users"`

	TotalBooks int `json:"total_books"`
}

func GetStatsBylib(libId uint) (*Statsl, error) {
	var stats Statsl
	var userCount int64
	var totalIssuedBook int64
	var bookCopies int64

	// Count total users for the given libId
	if err := config.DB.Model(&models.User{}).Where("lib_id = ?", libId).Count(&userCount).Error; err != nil {
		return nil, err
	}

	// Sum total book copies for books with the given libId
	if err := config.DB.Model(&models.Book{}).Where("lib_id = ?", libId).Select("COALESCE(SUM(total_copies), 0)").Scan(&bookCopies).Error; err != nil {
		return nil, err
	}

	// Count total issued books for books belonging to the given libId
	if err := config.DB.Model(&models.IssueRegistry{}).
		Joins("JOIN books ON books.isbn = issue_registries.isbn").
		Where("books.lib_id = ? AND issue_registries.issue_status = ?", libId, "Issued").
		Count(&totalIssuedBook).Error; err != nil {
		return nil, err
	}

	// Convert int64 to int
	stats.TotalUsers = int(userCount)
	stats.TotalIssuedBook = int(totalIssuedBook)
	stats.TotalBooks = int(bookCopies)

	return &stats, nil
}
