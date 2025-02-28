package models

import (
	"time"

	"gorm.io/gorm"
)

type RequestEvent struct {
	ReqID        uint       `gorm:"primaryKey;autoIncrement" json:"req_id"`
	ISBN         string     `gorm:"not null" json:"isbn"` // Changed from BookID to ISBN
	ReaderID     uint       `gorm:"not null" json:"reader_id"`
	RequestDate  time.Time  `gorm:"not null" json:"request_date"`
	ApprovalDate *time.Time `json:"approval_date,omitempty"`      // Nullable
	ApproverID   *uint      `json:"approver_id,omitempty"`        // Nullable
	RequestType  string     `gorm:"not null" json:"request_type"` // "Issue" or "Return"
	Status       string     `gorm:"not null" json:"status"`

	// Relationships
	Book Book `gorm:"foreignKey:ISBN;references:ISBN" json:"book"`
	User User `gorm:"foreignKey:ReaderID;references:ID" json:"user"`
}

// Migrate RequestEvent table
func MigrateRequestEvent(db *gorm.DB) {
	db.AutoMigrate(&RequestEvent{})
}
