package dto

import (
	"time"

	"github.com/briandang59/be_scada/internal/model"
)

type EquipmentResponse struct {
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
	WarrantyEndDate     *time.Time `json:"warranty_end_date,omitempty"`
	InstallationDate    *time.Time `json:"installation_date,omitempty"`
	Status              string     `json:"status"`
	IPAddress           string     `json:"ip_address"`
	MACAddress          string     `json:"mac_address"`
	OperatingSystem     string     `json:"operating_system"`
	Description         string     `json:"description"`
	Notes               string     `json:"notes"`
	LastMaintenanceDate *time.Time `json:"last_maintenance_date,omitempty"`
	NextMaintenanceDate *time.Time `json:"next_maintenance_date,omitempty"`

	Department    string `json:"department"`
	EquipmentType string `json:"equipment_type"`

	ResponsibleUser string `json:"responsible_user,omitempty"`
	AssignedUser    string `json:"assigned_user,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToEquipmentResponse(e *model.Equipment) EquipmentResponse {
	resp := EquipmentResponse{
		ID:                  e.ID,
		NameEn:              e.NameEn,
		NameZh:              e.NameZh,
		NameVn:              e.NameVn,
		Code:                e.Code,
		SerialNumber:        e.SerialNumber,
		Model:               e.Model,
		Manufacturer:        e.Manufacturer,
		Location:            e.Location,
		PurchaseDate:        e.PurchaseDate,
		WarrantyEndDate:     e.WarrantyEndDate,
		InstallationDate:    e.InstallationDate,
		Status:              e.Status,
		IPAddress:           e.IPAddress,
		MACAddress:          e.MACAddress,
		OperatingSystem:     e.OperatingSystem,
		Description:         e.Description,
		Notes:               e.Notes,
		LastMaintenanceDate: e.LastMaintenanceDate,
		NextMaintenanceDate: e.NextMaintenanceDate,
		CreatedAt:           e.CreatedAt,
		UpdatedAt:           e.UpdatedAt,
	}

	if e.Department != nil {
		resp.Department = e.Department.NameEn
	}

	if e.EquipmentType != nil {
		resp.EquipmentType = e.EquipmentType.NameEn
	}

	if e.ResponsibleUser != nil {
		resp.ResponsibleUser = e.ResponsibleUser.FullName
	}

	if e.AssignedUser != nil {
		resp.AssignedUser = e.AssignedUser.FullName
	}

	return resp
}
func ToEquipmentResponseList(equipments []*model.Equipment) []EquipmentResponse {
	res := make([]EquipmentResponse, 0, len(equipments))
	for _, e := range equipments {
		res = append(res, ToEquipmentResponse(e))
	}
	return res
}
