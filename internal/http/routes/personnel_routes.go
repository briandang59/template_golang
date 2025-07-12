package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func PersonnelRoutes(rg *gin.RouterGroup, h *handler.PersonnelHandler) {
	rg.GET("/personnels", h.GetAll)
	rg.POST("/personnels", h.Create)
	rg.PATCH("/personnels/:id", h.Update)
}
