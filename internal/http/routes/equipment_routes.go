package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func EquipmentRoutes(rg *gin.RouterGroup, h *handler.EquipementHandler) {
	rg.GET("/equipments", h.GetAll)
	rg.POST("/equipments", h.Create)
	rg.PATCH("/equipments/:id", h.Update)
	rg.DELETE("/equipments/:id", h.Delete)
}
