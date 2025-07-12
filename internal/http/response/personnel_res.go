package response

import "github.com/briandang59/be_scada/internal/model"

type PersonnelListResponse struct {
	Code       int               `json:"code" example:"200"`
	Message    string            `json:"message" example:"success"`
	Data       []model.Personnel `json:"data"`
	Pagination *Pagination       `json:"pagination"`
}

type PersonnelResponse struct {
	Code    int             `json:"code" example:"200"`
	Message string          `json:"message" example:"success"`
	Data    model.Personnel `json:"data"`
}
