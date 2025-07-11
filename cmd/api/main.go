// @title SCADA API
// @version 1.0
// @description This is a SCADA system backend.
// @host localhost:8080/api
// @BasePath /
// @schemes http
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
	deps := &handler.Dependencies{
		Factory:       factoryHdl,
		Department:    dpmHdl,
		EquipmentType: eqtHdl,
		Equipment:     eqHdl,
	}

	r := gin.Default()

	routes.RegisterRoutes(r, deps)

	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", ws.ServeWs(hub))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
