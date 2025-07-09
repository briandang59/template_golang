package response

import "github.com/gin-gonic/gin"

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

type Body struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
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
