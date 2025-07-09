package main

import (
	"github.com/gin-gonic/gin"

	"log"
	"os"

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

	// khởi tạo repository / service / handler
	eqRepo := repository.NewEquipmentRepo()
	eqSvc := service.NewEquipmentService(eqRepo)
	eqHdl := handler.NewEquipmentHandler(eqSvc)

	deps := &handler.Dependencies{
		Equipment: eqHdl,
	}

	r := gin.Default()

	routes.RegisterRoutes(r, deps)

	// WebSocket
	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", ws.ServeWs(hub))

	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
