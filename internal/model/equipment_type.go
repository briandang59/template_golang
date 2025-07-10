package model

import (
	"time"

	"gorm.io/gorm"
)

type EquipmentType struct {
	ID        uint           `json:"id"`
	NameEn    string         `json:"name_en" `
	NameZh    string         `json:"name_zh" `
	NameVn    string         `json:"name_vn" `
	Code      string         `json:"code"`
	Active    bool           `json:"active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"           gorm:"index"`
}
