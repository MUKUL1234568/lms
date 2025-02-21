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
