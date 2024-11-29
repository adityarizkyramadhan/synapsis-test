package model

import "time"

type Category struct {
	ID        uint32    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" validate:"required,min=3,max=100" json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
