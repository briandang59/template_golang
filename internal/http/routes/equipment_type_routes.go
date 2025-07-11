package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func EquipmentTypeRoutes(rg *gin.RouterGroup, h *handler.EquipementTypeHandler) {
	rg.GET("/equipment-types", h.GetAll)
	rg.POST("/equipment-types", h.Create)
	rg.PATCH("/equipment-types/:id", h.Update)
	rg.DELETE("/equipment-types/:id", h.Delete)
}
