package routes

import (
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, h *handler.AuthHandler) {
	rg.POST("/login", h.Login)
	rg.POST("/register", h.Register)

}
