// @title SCADA API
// @version 1.0
// @description This is a SCADA system backend for industrial equipment management.
// @host localhost:5000
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	"log"
	"os"

	_ "github.com/briandang59/be_scada/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/joho/godotenv"

	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/http/handler"
	"github.com/briandang59/be_scada/internal/http/routes"
	"github.com/briandang59/be_scada/internal/repository"
	"github.com/briandang59/be_scada/internal/service"
	ws "github.com/briandang59/be_scada/internal/websocket"
)

func main() {
	_ = godotenv.Load()
	config.Init()

	factoryRepo := repository.NewFactoryRepo()
	factorySvc := service.NewFactoryService(factoryRepo)
	factoryHdl := handler.NewFactoryHandler(factorySvc)

	dpmRepo := repository.NewDepartmentRepo()
	dpmySvc := service.NewDepartmentService(dpmRepo)
	dpmHdl := handler.NewDepartmentHandler(dpmySvc)

	eqtRepo := repository.NewEquipmentTypeRepo()
	eqtySvc := service.NewEquipmentTypeService(eqtRepo)
	eqtHdl := handler.NewEquipmentTypeHandler(eqtySvc)

	eqRepo := repository.NewEquipmentRepo()
	eqySvc := service.NewEquipmentService(eqRepo)
	eqHdl := handler.NewEquipmentHandler(eqySvc)

	accountRepo := repository.NewAccountRepository()
	authSvc := service.NewAuthService(accountRepo)
	authHdl := handler.NewAuthHandler(authSvc)

	psnRepo := repository.NewPersonnelRepository()
	psnSvc := service.NewPersonnelService(psnRepo)
	psnHdl := handler.NewPersonnelHandler(psnSvc)

	deps := &handler.Dependencies{
		Factory:       factoryHdl,
		Department:    dpmHdl,
		EquipmentType: eqtHdl,
		Equipment:     eqHdl,
		Account:       authHdl,
		Personnel:     psnHdl,
	}

	r := gin.Default()

	// Add CORS middleware for Swagger UI
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	routes.RegisterRoutes(r, deps)

	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", ws.ServeWs(hub))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
