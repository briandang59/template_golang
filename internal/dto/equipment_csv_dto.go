package dto

import (
	"fmt"
	"time"

	"github.com/briandang59/be_scada/internal/model"
)

// EquipmentCSV represents the CSV structure for equipment import
type EquipmentCSV struct {
	NameEn              string `csv:"name_en"`
	NameZh              string `csv:"name_zh"`
	NameVn              string `csv:"name_vn"`
	Code                string `csv:"code"`
	SerialNumber        string `csv:"serial_number"`
	Model               string `csv:"model"`
	Manufacturer        string `csv:"manufacturer"`
	Location            string `csv:"location"`
	PurchaseDate        string `csv:"purchase_date"`
	WarrantyEndDate     string `csv:"warranty_end_date"`
	InstallationDate    string `csv:"installation_date"`
	Status              string `csv:"status"`
	IPAddress           string `csv:"ip_address"`
	MACAddress          string `csv:"mac_address"`
	OperatingSystem     string `csv:"operating_system"`
	Description         string `csv:"description"`
	Notes               string `csv:"notes"`
	LastMaintenanceDate string `csv:"last_maintenance_date"`
	NextMaintenanceDate string `csv:"next_maintenance_date"`
	DepartmentID        string `csv:"department_id"`
	EquipmentTypeID     string `csv:"equipment_type_id"`
	ResponsibleUserID   string `csv:"responsible_user_id"`
	AssignedUserID      string `csv:"assigned_user_id"`
}

// ToEquipmentModel converts CSV data to Equipment model
func (csv *EquipmentCSV) ToEquipmentModel() (*model.Equipment, error) {
	equipment := &model.Equipment{
		NameEn:          csv.NameEn,
		NameZh:          csv.NameZh,
		NameVn:          csv.NameVn,
		Code:            csv.Code,
		SerialNumber:    csv.SerialNumber,
		Model:           csv.Model,
		Manufacturer:    csv.Manufacturer,
		Location:        csv.Location,
		Status:          csv.Status,
		IPAddress:       csv.IPAddress,
		MACAddress:      csv.MACAddress,
		OperatingSystem: csv.OperatingSystem,
		Description:     csv.Description,
		Notes:           csv.Notes,
		Active:          true,
	}

	// Parse dates
	if csv.PurchaseDate != "" {
		if date, err := time.Parse("2006-01-02T15:04:05Z", csv.PurchaseDate); err == nil {
			equipment.PurchaseDate = date
		}
	}

	if csv.WarrantyEndDate != "" {
		if date, err := time.Parse("2006-01-02T15:04:05Z", csv.WarrantyEndDate); err == nil {
			equipment.WarrantyEndDate = &date
		}
	}

	if csv.InstallationDate != "" {
		if date, err := time.Parse("2006-01-02T15:04:05Z", csv.InstallationDate); err == nil {
			equipment.InstallationDate = &date
		}
	}

	if csv.LastMaintenanceDate != "" {
		if date, err := time.Parse("2006-01-02T15:04:05Z", csv.LastMaintenanceDate); err == nil {
			equipment.LastMaintenanceDate = &date
		}
	}

	if csv.NextMaintenanceDate != "" {
		if date, err := time.Parse("2006-01-02T15:04:05Z", csv.NextMaintenanceDate); err == nil {
			equipment.NextMaintenanceDate = &date
		}
	}

	// Parse IDs
	if csv.DepartmentID != "" {
		if id, err := parseUint(csv.DepartmentID); err == nil {
			equipment.DepartmentID = id
		}
	}

	if csv.EquipmentTypeID != "" {
		if id, err := parseUint(csv.EquipmentTypeID); err == nil {
			equipment.EquipmentTypeID = id
		}
	}

	if csv.ResponsibleUserID != "" {
		if id, err := parseUint(csv.ResponsibleUserID); err == nil {
			equipment.ResponsibleUserID = &id
		}
	}

	if csv.AssignedUserID != "" {
		if id, err := parseUint(csv.AssignedUserID); err == nil {
			equipment.AssignedUserID = &id
		}
	}

	return equipment, nil
}

// parseUint helper function to parse string to uint
func parseUint(s string) (uint, error) {
	var result uint
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
} 