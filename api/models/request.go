package models

import (
	"time"

	"gorm.io/gorm"
)

type RequestEvent struct {
	ReqID        uint       `gorm:"primaryKey;autoIncrement" json:"req_id"`
	BookID       string     `gorm:"not null" json:"book_id"`
	ReaderID     uint       `gorm:"not null" json:"reader_id"`
	RequestDate  time.Time  `gorm:"not null" json:"request_date"`
	ApprovalDate *time.Time `json:"approval_date"`                // ✅ Nullable (NULL if not approved)
	ApproverID   *uint      `json:"approver_id"`                  // ✅ Nullable (NULL if not approved)
	RequestType  string     `gorm:"not null" json:"request_type"` // "Issue" or "Return"
	Status       string     `gorm:"not null" json:"status"`
}

// Migrate RequestEvent table
func MigrateRequestEvent(db *gorm.DB) {
	db.AutoMigrate(&RequestEvent{})
}
