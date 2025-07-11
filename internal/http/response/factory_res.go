package response

import "github.com/briandang59/be_scada/internal/model"

type FactoryListResponse struct {
	Code       int             `json:"code" example:"200"`
	Message    string          `json:"message" example:"success"`
	Data       []model.Factory `json:"data"`
	Pagination *Pagination     `json:"pagination"`
}

type FactoryResponse struct {
	Code    int           `json:"code" example:"200"`
	Message string        `json:"message" example:"success"`
	Data    model.Factory `json:"data"`
}
