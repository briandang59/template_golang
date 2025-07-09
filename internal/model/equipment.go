package model

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"        gorm:"size:100;not null"`
	Location    string         `json:"location"    gorm:"size:50"`
	Status      string         `json:"status"      gorm:"size:20"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"           gorm:"index"`
}
