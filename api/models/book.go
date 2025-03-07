package models

import "gorm.io/gorm"

type Book struct {
	ISBN            string          `gorm:"primaryKey" json:"isbn"`
	LibID           uint            `gorm:"not null" json:"lib_id"`
	Title           string          `gorm:"not null" json:"title"`
	Authors         string          `gorm:"not null" json:"authors"`
	Publisher       string          `gorm:"not null" json:"publisher"`
	Version         string          `json:"version,omitempty"`
	TotalCopies     int             `gorm:"not null" json:"total_copies"`
	AvailableCopies int             ` json:"available_copies"`
	IssueRecords    []IssueRegistry `gorm:"foreignKey:ISBN" json:"issue_records"`
	Requests        []RequestEvent  `gorm:"foreignKey:ISBN" json:"requests"`
}

// Migrate Book table
func MigrateBook(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
