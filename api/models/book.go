package models

import "gorm.io/gorm"

type Book struct {
	ISBN            string `gorm:"primaryKey" json:"isbn"`
	LibID           uint   `gorm:"not null" json:"lib_id"`
	Title           string `gorm:"not null" json:"title"`
	Authors         string `gorm:"not null" json:"authors"`
	Publisher       string `gorm:"not null" json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     int    `gorm:"not null" json:"total_copies"`
	AvailableCopies int    `gorm:"not null" json:"available_copies"`
}

// Migrate Book table
func MigrateBook(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
