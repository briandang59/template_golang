package response

import (
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int `json:"page" example:"1"`
	PageSize int `json:"page_size" example:"10"`
	Total    int `json:"total" example:"100"`
}

type Body struct {
	Code       int         `json:"code" example:"200"`
	Message    string      `json:"message" example:"success"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

func Success(c *gin.Context, data interface{}, pg *Pagination) {
	c.JSON(200, Body{
		Code:       200,
		Message:    "success",
		Data:       data,
		Pagination: pg,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, Body{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"error message"`
}
