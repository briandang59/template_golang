package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes gom mọi group API vào 1 chỗ
func RegisterRoutes(r *gin.Engine, deps *handler.Dependencies) {
	// tạo nhóm /api
	api := r.Group("/api")

	// Equipment
	FactoryRoutes(api, deps.Factory)
}
