package models

import "gorm.io/gorm"

type Library struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`

	Users []User `gorm:"foreignKey:LibID" json:"users"` // One library has many users
	Books []Book `gorm:"foreignKey:LibID" json:"books"` // One library has many books
}

// Migrate Library table
func MigrateLibrary(db *gorm.DB) {
	db.AutoMigrate(&Library{})
}
