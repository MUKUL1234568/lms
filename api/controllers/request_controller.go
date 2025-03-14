package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "library-management-api/config"
	"library-management-api/models"
	"library-management-api/services"
	"net/http"
	"strconv"
	"time"
)

// CreateRequest handles a reader's book request (Issue/Return)
// CreateRequest handles a reader's book request (Issue/Return)

func HasPendingRequest(user *models.User, isbn string) bool {
	for _, req := range user.Requests {
		if req.ISBN == isbn && req.Status == "pending" {
			return true
		}
	}
	return false
}

func HasIssuedBook(user *models.User, isbn string) bool {
	for _, req := range user.IssueRecords {
		if req.ISBN == isbn && req.IssueStatus == "Issued" {
			return true
		}
	}
	return false
}

func CreateRequest(c *gin.Context) {
	var request struct {
		ISBN        string `json:"isbn" binding:"required"`
		RequestType string `json:"request_type" binding:"required"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readerID, err := GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, errr := services.GetUserByID(readerID)
	if errr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if HasPendingRequest(user, request.ISBN) {
		c.JSON(http.StatusAlreadyReported, gin.H{"error": "you have made already request for this book wait for admin action"})
		return
	}
	if request.RequestType != "Return" {
		if HasIssuedBook(user, request.ISBN) {
			c.JSON(http.StatusAlreadyReported, gin.H{"error": "you have  already  this book"})
			return
		}
	}

	if request.RequestType == "Return" {
		if !HasIssuedBook(user, request.ISBN) {
			c.JSON(http.StatusAlreadyReported, gin.H{"error": "you have not this book to return "})
			return
		}
	}
	book, err := services.GetBookByISBN(request.ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if libID != book.LibID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not registred with same library"})
		return
	}
	// Create RequestEvent with RequestDate
	reqEvent := models.RequestEvent{
		ISBN:        request.ISBN,
		ReaderID:    readerID, // ✅ Fixed conversion issue
		RequestType: request.RequestType,
		RequestDate: time.Now(),
		Status:      "pending",
	}

	// Call service to create request
	if err := services.CreateRequest(&reqEvent); err != nil {
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
		Approve bool `json:"approve" `
	}
	fmt.Println("ok")
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ApproverID, err := GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Fetch the request from DB
	var reqEvent models.RequestEvent
	if err := services.GetRequestByID(uint(requestID), &reqEvent); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}
	if reqEvent.Status != "pending" {
		c.JSON(http.StatusAlreadyReported, gin.H{"error": "Request already processed"})
		return
	}
	fmt.Println(libID, reqEvent.Book.LibID)
	if libID != reqEvent.Book.LibID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to approve this request"})
		return
	}
	// If approving, set approval date and approver ID
	if request.Approve {
		now := time.Now()
		reqEvent.ApprovalDate = &now
		reqEvent.ApproverID = &ApproverID // ✅ Fix: Use a pointer
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
// func GetUserRequests(c *gin.Context) {
// 	// Get `ReaderID` from session (logged-in user)

// 	readerID, err := GetUserID(c)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Call service to fetch user requests
// 	requests, err := services.GetUserRequests(readerID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Success response
// 	c.JSON(http.StatusOK, gin.H{"requests": requests})
// }

func GetAllRequestsForAdmin(c *gin.Context) {
	// // Get `Role` from session (logged-in user)
	// role, exists := c.Get("role")
	// if !exists || role != "LibraryAdmin" {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
	// 	return
	// }

	// Call service to fetch all requests

	libID, err := GetLibraryID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	requests, err := services.GetAllRequests(libID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

// func GetRequestByID(c *gin.Context) {
// 	var reqEvent models.RequestEvent
// 	if err := services.GetRequestByID(uint(requestID), &reqEvent); err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
// 		return
// 	}
// }
