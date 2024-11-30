package model

import "time"

type Book struct {
	ID          uint32    `json:"id" gorm:"primaryKey,autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	AuthorID    uint32    `json:"author_id" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Year        uint32    `json:"year" gorm:"not null"`
	Stock       uint32    `json:"stock" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
