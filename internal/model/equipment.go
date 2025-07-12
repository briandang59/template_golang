package model

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID                  uint       `json:"id" gorm:"primaryKey"`
	NameEn              string     `json:"name_en"`
	NameZh              string     `json:"name_zh"`
	NameVn              string     `json:"name_vn"`
	Code                string     `json:"code"`
	SerialNumber        string     `json:"serial_number" `
	Model               string     `json:"model"`
	Manufacturer        string     `json:"manufacturer"`
	Location            string     `json:"location"`
	PurchaseDate        time.Time  `json:"purchase_date"`
	WarrantyEndDate     *time.Time `json:"warranty_end_date"`
	InstallationDate    *time.Time `json:"installation_date"`
	Status              string     `json:"status" gorm:"default:'active';not null"`
	IPAddress           string     `json:"ip_address"`
	MACAddress          string     `json:"mac_address"`
	OperatingSystem     string     `json:"operating_system"`
	Description         string     `json:"description"`
	Notes               string     `json:"notes"`
	LastMaintenanceDate *time.Time `json:"last_maintenance_date"`
	NextMaintenanceDate *time.Time `json:"next_maintenance_date"`
	Active              bool       `json:"active" gorm:"default:true;not null"`

	DepartmentID      uint  `json:"department_id"`
	EquipmentTypeID   uint  `json:"equipment_type_id"`
	ResponsibleUserID *uint `json:"responsible_user_id"`
	AssignedUserID    *uint `json:"assigned_user_id"`

	Department      *Department    `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	EquipmentType   *EquipmentType `json:"equipment_type,omitempty" gorm:"foreignKey:EquipmentTypeID"`
	ResponsibleUser *Personnel     `json:"responsible_user,omitempty" gorm:"foreignKey:ResponsibleUserID"`
	AssignedUser    *Personnel     `json:"assigned_user,omitempty" gorm:"foreignKey:AssignedUserID"`

	// MaintenanceLogs []MaintenanceLog `json:"maintenance_logs,omitempty" gorm:"foreignKey:EquipmentID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
