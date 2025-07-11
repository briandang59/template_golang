package model

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID              uint           `json:"id"`
	NameEn          string         `json:"name_en" `
	NameZh          string         `json:"name_zh" `
	NameVn          string         `json:"name_vn" `
	Code            string         `json:"code"`
	Active          bool           `json:"active" gorm:"default:true"`
	DepartmentID    uint           `json:"department_id"`
	EquipmentTypeID uint           `json:"equipment_type_id"`
	Department      *Department    `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	EquipmentType   *EquipmentType `json:"equipment_type,omitempty" gorm:"foreignKey:EquipmentTypeID"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-"           gorm:"index"`
}
