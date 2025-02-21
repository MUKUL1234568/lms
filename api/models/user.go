package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"not null" json:"name"`
	Email         string `gorm:"unique;not null" json:"email"`
	Password      string `gorm:"not null" json:"password"` // ✅ Added Password Field
	ContactNumber string `json:"contact_number"`
	Role          string `gorm:"not null" json:"role"` // Owner, LibraryAdmin, Reader
	LibID         uint   `gorm:"not null" json:"lib_id"`
}

// Migrate User table
func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
