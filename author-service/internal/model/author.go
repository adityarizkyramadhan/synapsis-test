package model

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint32         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" validate:"required,min=3,max=100" json:"name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" validate:"required,email" json:"email"`
	Bio       string         `gorm:"type:text" json:"bio"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
