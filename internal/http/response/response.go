package response

import (
	"github.com/briandang59/be_scada/internal/model"
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int `json:"page" example:"1"`
	PageSize int `json:"page_size" example:"10"`
	Total    int `json:"total" example:"100"`
}

// Body is the generic response struct used in actual code
type Body struct {
	Code       int         `json:"code" example:"200"`
	Message    string      `json:"message" example:"success"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Success sends a successful response
func Success(c *gin.Context, data interface{}, pg *Pagination) {
	c.JSON(200, Body{
		Code:       200,
		Message:    "success",
		Data:       data,
		Pagination: pg,
	})
}

// Error sends an error response
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, Body{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

// ---------- BELOW: Structs for Swagger documentation only ----------

// ErrorResponse is used for Swagger to describe error structure
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"error message"`
}

// FactoryListResponse describes a list of factories (for GetAll)
type FactoryListResponse struct {
	Code       int             `json:"code" example:"200"`
	Message    string          `json:"message" example:"success"`
	Data       []model.Factory `json:"data"`
	Pagination *Pagination     `json:"pagination"`
}

// FactoryResponse describes a single factory (for Create, Update)
type FactoryResponse struct {
	Code    int           `json:"code" example:"200"`
	Message string        `json:"message" example:"success"`
	Data    model.Factory `json:"data"`
}
