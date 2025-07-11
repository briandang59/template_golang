package response

import "github.com/briandang59/be_scada/internal/model"

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
