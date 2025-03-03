package controllers

import (
	"errors"
	"library-management-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles user authentication
func Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user
	token, err := services.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func GetUserID(c *gin.Context) (uint, error) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user not authenticated")
	}

	// Convert interface{} to float64, then to uint
	userIDFloat, ok := userIDInterface.(float64)
	if !ok {
		return 0, errors.New("invalid user ID format")
	}

	return uint(userIDFloat), nil
}
func GetLibraryID(c *gin.Context) (uint, error) {
	libIDInterface, exists := c.Get("lib_id")
	if !exists {
		return 0, errors.New("library not found in session")
	}

	// Convert interface{} to float64, then to uint
	libIDFloat, ok := libIDInterface.(float64)
	if !ok {
		return 0, errors.New("invalid library ID format")
	}

	return uint(libIDFloat), nil
}
