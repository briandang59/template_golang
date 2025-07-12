package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/briandang59/be_scada/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, deps *handler.Dependencies) {
	auth := r.Group("/api/auth")
	api := r.Group("/api", middleware.AuthMiddleware())

	AuthRoutes(auth, deps.Account)

	FactoryRoutes(api, deps.Factory)
	DepartmentRoutes(api, deps.Department)
	EquipmentTypeRoutes(api, deps.EquipmentType)
	EquipmentRoutes(api, deps.Equipment)
	PersonnelRoutes(api, deps.Personnel)
}
