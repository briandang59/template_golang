package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Tag         string         `json:"tag" gorm:"uniqueIndex;not null"`
	NameEn      string         `json:"name_en" `
	NameZh      string         `json:"name_zh" `
	NameVn      string         `json:"name_vn" `
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"           gorm:"index"`
}
