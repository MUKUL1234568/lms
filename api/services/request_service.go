package services

import (
	"errors"
	"fmt"
	"library-management-api/config"
	"library-management-api/models"
	"time"

	"gorm.io/gorm"
)

// CreateRequest stores a new book request in the database
func CreateRequest(request *models.RequestEvent) error {
	return config.DB.Create(request).Error
}

// ApproveRequest approves or rejects a book request
func ApproveRequest(request *models.RequestEvent, approve bool) error {
	// If rejecting, delete request and return
	if !approve {
		request.Status = "Rejected"
		config.DB.Save(request)
		return nil
	}

	// Save approved request (ApprovalDate & ApproverID already set in controller)
	config.DB.Save(request)

	// Process request type (Issue/Return)
	if request.RequestType == "Issue" {
		// Check if book is available
		var book models.Book
		fmt.Println("ok")
		if err := config.DB.Where("isbn = ?", request.ISBN).First(&book).Error; err != nil {
			return errors.New("book not found")
		}

		if book.AvailableCopies <= 0 {
			return errors.New("no available copies")
		}

		// Decrease available copies
		book.AvailableCopies -= 1
		config.DB.Save(&book)

		// ✅ Fix: Store `ExpectedReturn` as a pointer
		expectedReturn := request.ApprovalDate.AddDate(0, 0, 14)

		// ✅ Fix: Dereference `ApproverID`
		issue := models.IssueRegistry{
			ISBN:            request.ISBN,
			ReaderID:        request.ReaderID,
			IssueApproverID: request.ApproverID, // ✅ Fix: Use `*request.ApproverID`
			IssueStatus:     "Issued",
			IssueDate:       *request.ApprovalDate,
			ExpectedReturn:  &expectedReturn, // ✅ Fix: Store pointer
		}
		if err := config.DB.Create(&issue).Error; err != nil {
			return err
		}
	} else if request.RequestType == "Return" {
		// Process book return
		var issue models.IssueRegistry
		if err := config.DB.Where("isbn = ? AND reader_id = ? AND issue_status = 'Issued'", request.ISBN, request.ReaderID).First(&issue).Error; err != nil {
			return errors.New("no active issue record found")
		}

		// Mark as returned
		now := time.Now()
		issue.IssueStatus = "Returned"
		issue.ReturnDate = &now

		// ✅ Fix: Dereference `ApproverID`
		issue.ReturnApproverID = request.ApproverID // ✅ Fix: Use `*request.ApproverID`
		config.DB.Save(&issue)

		// Increase available copies
		var book models.Book
		if err := config.DB.Where("isbn = ?", request.ISBN).First(&book).Error; err == nil {
			book.AvailableCopies += 1
			config.DB.Save(&book)
		}
	}

	// Delete request after processing
	return config.DB.Delete(request).Error
}

// GetUserRequests fetches all requests made by a user
func GetUserRequests(readerID uint) ([]models.RequestEvent, error) {
	var requests []models.RequestEvent
	err := config.DB.Preload("Book", func(db *gorm.DB) *gorm.DB {
		return db.Select("isbn", "title", "publisher")
	}).Where("reader_id = ?", readerID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// GetRequestByID fetches a single request by its ID
func GetRequestByID(requestID uint, request *models.RequestEvent) error {
	return config.DB.First(request, requestID).Error
}

// GetAllRequests fetches all requests (Only for LibraryAdmins)
func GetAllRequests(libID uint) ([]models.RequestEvent, error) {
	var requests []models.RequestEvent
	err := config.DB.Preload("Book", func(db *gorm.DB) *gorm.DB {
		return db.Select("isbn", "title", "publisher")
	}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("name", "email", "id")
	}).Where("lib_id=?", libID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
