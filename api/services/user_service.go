package services

import (
	"errors"
	"library-management-api/config"
	"library-management-api/models"
)

// RegisterUser adds a new user to the database
func RegisterUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// GetUserByID retrieves user details by ID
func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetUsersByLibrary(libID uint) ([]models.User, error) {
	var users []models.User
	err := config.DB.Where("lib_id = ?", libID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
