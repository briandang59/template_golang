package model

import (
	"time"

	"gorm.io/gorm"
)

type Personnel struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	FullName    string         `json:"full_name"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	WorkplaceID []uint         `json:"work_place_id"`
	Workplace   []Factory      `json:"work_place,omitempty" gorm:"foreignKey:WorkplaceID"`
	AccountID   uint           `json:"account_id"`
	Account     Account        `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
