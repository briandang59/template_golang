package response

import (
	"time"

	"github.com/briandang59/be_scada/internal/model"
)

type EquipmentListResponse struct {
	Code       int               `json:"code" example:"200"`
	Message    string            `json:"message" example:"success"`
	Data       []model.Equipment `json:"data"`
	Pagination *Pagination       `json:"pagination"`
}

type EquipmentResponse struct {
	Code    int             `json:"code" example:"200"`
	Message string          `json:"message" example:"success"`
	Data    model.Equipment `json:"data"`
}

// EquipmentResponseData represents equipment data in response
type EquipmentResponseData struct {
	ID                  uint       `json:"id"`
	NameEn              string     `json:"name_en"`
	NameZh              string     `json:"name_zh"`
	NameVn              string     `json:"name_vn"`
	Code                string     `json:"code"`
	SerialNumber        string     `json:"serial_number"`
	Model               string     `json:"model"`
	Manufacturer        string     `json:"manufacturer"`
	Location            string     `json:"location"`
	PurchaseDate        time.Time  `json:"purchase_date"`
	WarrantyEndDate     *time.Time `json:"warranty_end_date"`
	InstallationDate    *time.Time `json:"installation_date"`
	Status              string     `json:"status"`
	IPAddress           string     `json:"ip_address"`
	MACAddress          string     `json:"mac_address"`
	OperatingSystem     string     `json:"operating_system"`
	Description         string     `json:"description"`
	Notes               string     `json:"notes"`
	LastMaintenanceDate *time.Time `json:"last_maintenance_date"`
	NextMaintenanceDate *time.Time `json:"next_maintenance_date"`
	Active              bool       `json:"active"`
	DepartmentID        uint       `json:"department_id"`
	EquipmentTypeID     uint       `json:"equipment_type_id"`
	ResponsibleUserID   *uint      `json:"responsible_user_id"`
	AssignedUserID      *uint      `json:"assigned_user_id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

// ToEquipmentResponse converts model.Equipment to EquipmentResponseData
func ToEquipmentResponse(equipment *model.Equipment) *EquipmentResponseData {
	return &EquipmentResponseData{
		ID:                  equipment.ID,
		NameEn:              equipment.NameEn,
		NameZh:              equipment.NameZh,
		NameVn:              equipment.NameVn,
		Code:                equipment.Code,
		SerialNumber:        equipment.SerialNumber,
		Model:               equipment.Model,
		Manufacturer:        equipment.Manufacturer,
		Location:            equipment.Location,
		PurchaseDate:        equipment.PurchaseDate,
		WarrantyEndDate:     equipment.WarrantyEndDate,
		InstallationDate:    equipment.InstallationDate,
		Status:              equipment.Status,
		IPAddress:           equipment.IPAddress,
		MACAddress:          equipment.MACAddress,
		OperatingSystem:     equipment.OperatingSystem,
		Description:         equipment.Description,
		Notes:               equipment.Notes,
		LastMaintenanceDate: equipment.LastMaintenanceDate,
		NextMaintenanceDate: equipment.NextMaintenanceDate,
		Active:              equipment.Active,
		DepartmentID:        equipment.DepartmentID,
		EquipmentTypeID:     equipment.EquipmentTypeID,
		ResponsibleUserID:   equipment.ResponsibleUserID,
		AssignedUserID:      equipment.AssignedUserID,
		CreatedAt:           equipment.CreatedAt,
		UpdatedAt:           equipment.UpdatedAt,
	}
}
