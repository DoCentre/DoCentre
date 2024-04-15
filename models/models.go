package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Identity string `gorm:"not null;default:'user'"` // user, admin
}

type Document struct {
	ID           uint   `gorm:"primaryKey"`
	AuthorsID    []User `gorm:"many2many:document_authors;"`
	Title        string `gorm:"not null"`
	Content      []byte `gorm:"not null"`
	Appendix     []byte
	Status       string `gorm:"not null;default:'EDIT'"` // EDIT, VERIFY, REJECT, APPROVE
	RejectReason string
	LastEditDate time.Time
	ApprovedDate time.Time
	// foreign keys to User
	ApproverID User `gorm:"foreignKey:User.ID"`
	ViewerID   User `gorm:"foreignKey:User.ID"`
}
