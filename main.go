package main

import (
	"log"
	"os"

	config_application "github.com/RenzoReccio/project-management.worker/config/application"
	config_service "github.com/RenzoReccio/project-management.worker/config/service"
	"github.com/RenzoReccio/project-management.worker/presentation/controllers"
	"github.com/RenzoReccio/project-management.worker/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initializeReceiverHttpHandler(router *gin.Engine) {
	// Initialize all layers
	receiverController := controllers.NewEventController()
	// Routers
	router.POST("event", receiverController.InsertEvent)
}

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	configService := config_service.NewConfigService()
	config_application.InitApplication(configService)
	initializeReceiverHttpHandler(router)
	router.Run(":" + port)
}
