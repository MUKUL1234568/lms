package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"library-management-api/models"
	"library-management-api/services"
	"library-management-api/validator"
	"net/http"
)

// CreateLibrary handles the creation of a new library with an owner
func CreateLibrary(c *gin.Context) {
	var request struct {
		LibraryName   string `json:"library_name" binding:"required"`
		OwnerName     string `json:"owner_name" binding:"required"`
		OwnerEmail    string `json:"owner_email" binding:"required,email"`
		OwnerPassword string `json:"owner_password" binding:"required"`
		OwnerContact  string `json:"owner_contact" binding :"required"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.Validatephonenumbr(request.OwnerContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Hash the Password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.OwnerPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create User (Owner)
	owner := models.User{
		Name:          request.OwnerName,
		Email:         request.OwnerEmail,
		Password:      string(hashedPassword), // ✅ Store hashed password
		ContactNumber: request.OwnerContact,
		Role:          "Owner", // ✅ Default role as Owner
	}

	// Create Library
	library := models.Library{
		Name: request.LibraryName,
	}

	// Call Service
	err = services.CreateLibrary(&owner, &library)
	if err != nil {
		if err.Error() == "user with this email already exists" || err.Error() == "choose different library name" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Success Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Library created successfully",
		"library": library,
		"owner":   owner,
	})
}

func GetLibraries(c *gin.Context) {
	libraries, err := services.GetAllLibraries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"libraries": libraries})
}
