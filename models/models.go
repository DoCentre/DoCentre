package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"unique;not null"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	Identity  string     `gorm:"not null;default:'user'"` // user, admin
	Documents []Document `gorm:"foreignKey:AuthorID"`     // foreign key only ; not a real field
}

type Document struct {
	ID           uint `gorm:"primaryKey"`
	AuthorID     uint
	Title        string `gorm:"default:'Untitled'"`
	Content      []byte
	Appendix     []byte
	Status       string `gorm:"not null;default:'EDIT'"` // EDIT, VERIFY, REJECT, APPROVE
	RejectReason string
	LastEditDate time.Time
	// ApprovedDate time.Time
	// foreign keys to User
	// ApproverID User `gorm:"foreignKey:User.ID"`
	// ViewerID   User `gorm:"foreignKey:User.ID"`
}
