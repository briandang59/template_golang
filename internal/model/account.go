package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Roles     []Role         `json:"roles" gorm:"many2many:account_roles;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"           gorm:"index"`
}
