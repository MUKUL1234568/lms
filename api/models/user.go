package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID            uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string          `gorm:"not null" binding:"required" json:"name"`
	Email         string          `gorm:"unique;not null" binding:"required"  json:"email"`
	Password      string          `gorm:"not null" json:"-"`
	ContactNumber string          `gorm:"unique;not null" binding:"required"  json:"contact_number"`
	Role          string          `gorm:"not null" json:"role"` // Owner, LibraryAdmin, Reader
	LibID         uint            `gorm:"not null" json:"lib_id"`
	IssueRecords  []IssueRegistry `gorm:"foreignKey:ReaderID" json:"issue_records"`
	Requests      []RequestEvent  `gorm:"foreignKey:ReaderID" json:"requests"`
}

// Migrate User table
func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
