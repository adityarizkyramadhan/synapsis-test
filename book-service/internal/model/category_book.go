package model

type CategoryBook struct {
	ID         uint32 `json:"id" gorm:"primaryKey,autoIncrement"`
	CategoryID uint32 `json:"category_id" gorm:"not null"`
	BookID     uint32 `json:"book_id" gorm:"not null"`
}
