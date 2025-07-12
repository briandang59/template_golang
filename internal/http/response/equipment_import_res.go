package response

import (
	"github.com/briandang59/be_scada/internal/model"
)

// EquipmentImportResponse represents the response for CSV import
type EquipmentImportResponse struct {
	Total   int                        `json:"total"`
	Success int                        `json:"success"`
	Failed  int                        `json:"failed"`
	Errors  []string                   `json:"errors"`
	Created []*EquipmentResponseData   `json:"created"`
}

// ToEquipmentImportResponse converts service ImportResult to response
func ToEquipmentImportResponse(total, success, failed int, errors []string, created []*model.Equipment) *EquipmentImportResponse {
	createdResponses := make([]*EquipmentResponseData, len(created))
	for i, equipment := range created {
		createdResponses[i] = ToEquipmentResponse(equipment)
	}

	return &EquipmentImportResponse{
		Total:   total,
		Success: success,
		Failed:  failed,
		Errors:  errors,
		Created: createdResponses,
	}
} 