package response

import "github.com/briandang59/be_scada/internal/model"

type EquipmentTypeListResponse struct {
	Code       int                   `json:"code" example:"200"`
	Message    string                `json:"message" example:"success"`
	Data       []model.EquipmentType `json:"data"`
	Pagination *Pagination           `json:"pagination"`
}

type EquipmentTypeResponse struct {
	Code    int                 `json:"code" example:"200"`
	Message string              `json:"message" example:"success"`
	Data    model.EquipmentType `json:"data"`
}
