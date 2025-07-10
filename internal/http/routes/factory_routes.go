package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func FactoryRoutes(rg *gin.RouterGroup, h *handler.FactoryHandler) {
	rg.GET("/factories", h.GetAll)
	rg.POST("/factories", h.Create)
	rg.PATCH("/factories/:id", h.Update)

}
