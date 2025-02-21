package services

import (
	"errors"
	"library-management-api/config"
	"library-management-api/models"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser adds a new user to the database
func RegisterUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// AuthenticateUser checks if the user exists and verifies the password
func AuthenticateUser(email string, password string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Compare stored hashed password with provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

// GetUserByID retrieves user details by ID
func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
