package controllers

import (
	"library-management-api/models"
	"library-management-api/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateRequest handles a reader's book request (Issue/Return)
// CreateRequest handles a reader's book request (Issue/Return)
func CreateRequest(c *gin.Context) {
	var request struct {
		BookID      string `json:"book_id" binding:"required"`
		RequestType string `json:"request_type" binding:"required"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request type
	if request.RequestType != "Issue" && request.RequestType != "Return" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type"})
		return
	}

	// Get `ReaderID` from the session (logged-in user)
	readerIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert `interface{}` to `uint`
	readerIDFloat, ok := readerIDInterface.(float64) // ✅ Fix: First convert to float64
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	readerID := uint(readerIDFloat) // ✅ Fix: Convert float64 to uint

	// Create RequestEvent with RequestDate
	reqEvent := models.RequestEvent{
		BookID:      request.BookID,
		ReaderID:    readerID, // ✅ Fixed conversion issue
		RequestType: request.RequestType,
		RequestDate: time.Now(),
		Status:      "pending",
	}

	// Call service to create request
	err := services.CreateRequest(&reqEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Request submitted successfully",
		"request": reqEvent,
	})
}

// ApproveRequest handles approval or rejection of a book request
// ApproveRequest handles approval or rejection of a book request
func ApproveRequest(c *gin.Context) {
	// Parse request ID from URL
	requestID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	// Parse JSON body
	var request struct {
		Approve bool `json:"approve" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readerIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert `interface{}` to `uint`
	readerIDFloat, ok := readerIDInterface.(float64) // ✅ Fix: Convert to float64 first
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	readerID := uint(readerIDFloat) // ✅ Fix: Convert float64 to uint

	// Fetch the request from DB
	var reqEvent models.RequestEvent
	if err := services.GetRequestByID(uint(requestID), &reqEvent); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	// If approving, set approval date and approver ID
	if request.Approve {
		now := time.Now()
		reqEvent.ApprovalDate = &now
		reqEvent.ApproverID = &readerID // ✅ Fix: Use a pointer
	}

	// Call service to approve/reject request
	err = services.ApproveRequest(&reqEvent, request.Approve)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Request processed successfully",
		"request": reqEvent,
	})
}

// GetUserRequests fetches all requests made by the logged-in user
func GetUserRequests(c *gin.Context) {
	// Get `ReaderID` from session (logged-in user)
	readerIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert `interface{}` to `uint`
	readerIDFloat, ok := readerIDInterface.(float64) // ✅ Fix: Convert to float64 first
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	readerID := uint(readerIDFloat) // ✅ Fix: Convert float64 to uint

	// Call service to fetch user requests
	requests, err := services.GetUserRequests(readerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func GetAllRequestsForAdmin(c *gin.Context) {
	// // Get `Role` from session (logged-in user)
	// role, exists := c.Get("role")
	// if !exists || role != "LibraryAdmin" {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
	// 	return
	// }

	// Call service to fetch all requests

	requests, err := services.GetAllRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"requests": requests})
}
