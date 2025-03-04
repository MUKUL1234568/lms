package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"library-management-api/models"
	"library-management-api/services"
	"library-management-api/validator"
	"net/http"
	// "strconv"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var request struct {
		Name          string `json:"name" binding:"required"`
		Email         string `json:"email" binding:"required,email"`
		Password      string `json:"password" binding:"required"`
		ContactNumber string `json:"contact_number"  binding:"required"`
		LibID         uint   `json:"lib_id" binding:"required,numeric"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.Validatephonenumbr(request.ContactNumber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(request.LibID)
	err := services.Getlibbyid(request.LibID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "library not exsits"})
		return
	}

	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Set default role as "Reader"
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

// GetUser fetches user details by ID
func GetUser(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Success Response (Hiding Password)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsersByLibrary(c *gin.Context) {
	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Call service to fetch books
	users, err := services.GetUsersByLibrary(libID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// func MakeAdmin(c *gin.Context) {
// 	useridstr := c.Param("id")

// 	userid64, err := strconv.ParseUint(useridstr, 10, 32)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid ID"})
// 		return
// 	}
// 	userid := uint(userid64)

// 	var request struct {
// 		Role string `json:"role" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if request.Role != "LibraryAdmin" && request.Role != "Reader" {
// 		c.JSON(http.StatusBadRequest, gin.H{"messege": "not valid role"})
// 		return
// 	}
// 	fmt.Println(userid, request.Role)
// 	userc, err := services.GetUserByID(userid)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}
// 	libID, err := GetLibraryID(c)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if userc.LibID != libID {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "your are not registred in this library"})
// 		return
// 	}
// 	user, err := services.MakeAdmin(userid, request.Role)

// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "service is not initiated "})
// 	}
// 	c.JSON(http.StatusOK, gin.H{"admin": user})
// }

// func DeleteUser(c *gin.Context) {
// 	useridstr := c.Param("id")

// 	userid64, err := strconv.ParseUint(useridstr, 10, 32)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid ID"})
// 		return
// 	}
// 	userid := uint(userid64)

// 	userc, err := services.GetUserByID(userid)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	libID, err := GetLibraryID(c)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if userc.LibID != libID {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "your are not registred in this library"})
// 		return
// 	}

// 	errr := services.DeleteUser(userid)
// 	if errr != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found "})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "user deleted succesfully"})

// }
