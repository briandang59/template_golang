package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func DepartmentRoutes(rg *gin.RouterGroup, h *handler.DepartmentHandler) {
	rg.GET("/departments", h.GetAll)
	rg.POST("/departments", h.Create)
	rg.PATCH("/departments/:id", h.Update)
	rg.DELETE("/departments/:id", h.Delete)
}
