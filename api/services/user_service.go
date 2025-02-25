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

func MakeAdmin(userID string, role string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	user.Role = role
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, errors.New("faild to save to make admin ")
	}
	return &user, nil
}

func DeleteUser(userid string) error {
	var user models.User
	errr := config.DB.First(&user, userid).Error
	if errr != nil {
		return errors.New("user not found")
	}

	err := config.DB.Delete(&user).Error
	if err != nil {
		return errors.New("failed to delete ")
	}
	return nil
}
