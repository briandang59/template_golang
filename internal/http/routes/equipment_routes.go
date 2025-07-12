package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func EquipmentRoutes(rg *gin.RouterGroup, h *handler.EquipementHandler) {
	rg.GET("/equipments", h.GetAll)
	rg.GET("/equipments/template", h.DownloadCSVTemplate)
	rg.POST("/equipments", h.Create)
	rg.POST("/equipments/import", h.ImportFromCSV)
	rg.PATCH("/equipments/:id", h.Update)
	rg.DELETE("/equipments/:id", h.Delete)
}
