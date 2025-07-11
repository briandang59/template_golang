package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	NameEn      string         `json:"name_en"`
	NameZh      string         `json:"name_zh"`
	NameVn      string         `json:"name_vn"`
	Tag         string         `json:"tag"`
	Description string         `json:"description"`
	Permissions []Permission   `json:"permissions" gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"           gorm:"index"`
}
