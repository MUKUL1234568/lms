package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID            uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string          `gorm:"not null"  json:"name"`
	Email         string          `gorm:"unique;not null"  json:"email"`
	Password      string          `gorm:"not null" json:"-"`
	ContactNumber string          `gorm:"not null" json:"contact_number"`
	Role          string          `gorm:"not null" json:"role"` // Owner, LibraryAdmin, Reader
	LibID         uint            `gorm:"not null" json:"lib_id"`
	IssueRecords  []IssueRegistry `gorm:"foreignKey:ReaderID" json:"issue_records"`
	Requests      []RequestEvent  `gorm:"foreignKey:ReaderID" json:"requests"`
}

// Migrate User table
func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
