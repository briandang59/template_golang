package model

import (
	"time"

	"gorm.io/gorm"
)

type Personnel struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FullName     string         `json:"full_name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	AccountID    uint           `json:"account_id"`
	DepartmentID uint           `json:"department_id"`
	Workplace    []Factory      `json:"work_place,omitempty" gorm:"many2many:personnel_workplaces;"`
	Account      *Account       `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	Department   *Department    `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
