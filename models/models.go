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
	ID           uint `gorm:"primaryKey"`
	AuthorID     uint `gorm:"foreignKey:AuthorID"`
	Title        string
	Content      string
	Appendix     string
	Status       string    `gorm:"not null;check:status IN ('EDIT', 'VERIFY', 'REJECT', 'APPROVE');default:'EDIT'"` // EDIT, VERIFY, REJECT, APPROVE
	Comment      string    `gorm:"default:''"`
	CreatedAt    time.Time // 建立時間（由GORM自動管理）
	UpdatedAt    time.Time // 最後一次更新時間（由GORM自動管理）
	ApprovedDate time.Time
	// foreign keys to User
	ApproverID uint `gorm:"foreignKey:ApproverID"`
	Author     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Approver   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DocumentViewer struct {
	ID uint `gorm:"primaryKey"`
	// foreign keys to Document and User
	DocumentID uint     `gorm:"foreignKey:DocumentID"`
	ViewerID   uint     `gorm:"foreignKey:ViewerID"`
	Document   Document `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Viewer     User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
