package services

import (
	"errors"
	"library-management-api/config"
	"library-management-api/models"
)

// CreateLibrary creates a new library and assigns the user as the Owner
func CreateLibrary(owner *models.User, library *models.Library) error {
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
