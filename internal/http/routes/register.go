package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, deps *handler.Dependencies) {
	api := r.Group("/api")

	FactoryRoutes(api, deps.Factory)
	DepartmentRoutes(api, deps.Department)
	EquipmentTypeRoutes(api, deps.EquipmentType)
	EquipmentRoutes(api, deps.Equipment)
}
