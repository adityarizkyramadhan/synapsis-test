package model

import "time"

type Borrowing struct {
	ID         uint32    `json:"id" gorm:"primaryKey,autoIncrement"`
	BookID     uint32    `json:"book_id" gorm:"not null"`
	UserID     uint32    `json:"user_id" gorm:"not null"`
	BorrowedAt time.Time `json:"borrowed_at" gorm:"not null"`
	ReturnedAt time.Time `json:"returned_at" gorm:"default:null"`
	Amount     uint32    `json:"amount" gorm:"not null"`
	Book       *Book     `json:"book" gorm:"foreignKey:BookID"`
}
