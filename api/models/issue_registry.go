package models

import "gorm.io/gorm"

type IssueRegistry struct {
	IssueID          uint   `gorm:"primaryKey;autoIncrement" json:"issue_id"`
	ISBN             string `gorm:"not null" json:"isbn"`
	ReaderID         uint   `gorm:"not null" json:"reader_id"`
	IssueApproverID  uint   `json:"issue_approver_id"`
	IssueStatus      string `gorm:"not null" json:"issue_status"`
	IssueDate        string `gorm:"not null" json:"issue_date"`
	ExpectedReturn   string `json:"expected_return_date"`
	ReturnDate       string `json:"return_date"`
	ReturnApproverID uint   `json:"return_approver_id"`
}

// Migrate IssueRegistry table
func MigrateIssueRegistry(db *gorm.DB) {
	db.AutoMigrate(&IssueRegistry{})
}
