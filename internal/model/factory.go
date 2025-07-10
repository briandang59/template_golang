package model

import (
	"time"

	"gorm.io/gorm"
)

type Factory struct {
	ID        uint           `json:"id"`
	Name_En   string         `json:"name_en" gorm:"size:100"`
	Name_Zh   string         `json:"name_zh" gorm:"size:100"`
	Name_Vn   string         `json:"name_vn" gorm:"size:100"`
	Code      string         `json:"code"`
	Address   string         `json:"address"`
	Location  string         `json:"location"`
	Active    bool           `json:"active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"           gorm:"index"`
}
