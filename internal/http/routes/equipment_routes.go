package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

// EquipmentRoutes gắn đường dẫn liên quan thiết bị
func EquipmentRoutes(rg *gin.RouterGroup, h *handler.EquipmentHandler) {
	rg.GET("/equipment", h.GetAll)
	rg.GET("/equipment/:id", h.Get)
	rg.POST("/equipment", h.Create)
	rg.PUT("/equipment/:id", h.Update)
	rg.DELETE("/equipment/:id", h.Delete)
}
