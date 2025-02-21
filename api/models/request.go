package models

import "gorm.io/gorm"

type RequestEvent struct {
	ReqID        uint   `gorm:"primaryKey;autoIncrement" json:"req_id"`
	BookID       string `gorm:"not null" json:"book_id"`
	ReaderID     uint   `gorm:"not null" json:"reader_id"`
	RequestDate  string `gorm:"not null" json:"request_date"`
	ApprovalDate string `json:"approval_date"`
	ApproverID   uint   `json:"approver_id"`
	RequestType  string `gorm:"not null" json:"request_type"` // "Issue" or "Return"

}

// Migrate RequestEvent table
func MigrateRequestEvent(db *gorm.DB) {
	db.AutoMigrate(&RequestEvent{})
}
