package controllers

import (
	"library-management-api/models"
	"library-management-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var request struct {
		Name          string `json:"name" binding:"required"`
		Email         string `json:"email" binding:"required,email"`
		Password      string `json:"password" binding:"required"`
		ContactNumber string `json:"contact_number"`
		Role          string `json:"role"` // Optional, defaults to "Reader"
		LibID         uint   `json:"lib_id" binding:"required"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Set default role as "Reader" if no role is provided
	userRole := "Reader"

	// Create User
	user := models.User{
		Name:          request.Name,
		Email:         request.Email,
		Password:      string(hashedPassword),
		ContactNumber: request.ContactNumber,
		Role:          userRole,
		LibID:         request.LibID,
	}

	// Call Service to Register User
	err = services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success Response (Hiding Password for Security)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":             user.ID,
			"name":           user.Name,
			"email":          user.Email,
			"contact_number": user.ContactNumber,
			"role":           user.Role,
			"lib_id":         user.LibID,
		},
	})
}

// LoginUser handles user authentication
// func LoginUser(c *gin.Context) {
// 	var request struct {
// 		Email    string `json:"email" binding:"required,email"`
// 		Password string `json:"password" binding:"required"`
// 	}

// 	// Bind JSON request
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Authenticate User
// 	user, err := services.AuthenticateUser(request.Email, request.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	// Success Response (Hiding Password)
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Login successful",
// 		"user": gin.H{
// 			"id":             user.ID,
// 			"name":           user.Name,
// 			"email":          user.Email,
// 			"contact_number": user.ContactNumber,
// 			"role":           user.Role,
// 			"lib_id":         user.LibID,
// 		},
// 	})
// }

// GetUser fetches user details by ID
func GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Success Response (Hiding Password)
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":             user.ID,
			"name":           user.Name,
			"email":          user.Email,
			"contact_number": user.ContactNumber,
			"role":           user.Role,
			"lib_id":         user.LibID,
		},
	})
}

func GetUsersByLibrary(c *gin.Context) {
	libIDInterface, exists := c.Get("lib_id") // ✅ Use lowercase "lib_id"
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert interface{} to float64 first, then to uint
	libIDFloat, ok := libIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Library ID format"})
		return
	}
	libID := uint(libIDFloat) // ✅ Convert float64 to uint

	// Call service to fetch books
	books, err := services.GetUsersByLibrary(libID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"books": books})
}
