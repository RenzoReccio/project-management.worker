package main

import (
	"log"
	"os"

	application "github.com/RenzoReccio/project-management.worker/application/receiver/insertReceiver"
	mongoInfraestructure "github.com/RenzoReccio/project-management.worker/infrastructure/mongo/receiver"
	"github.com/RenzoReccio/project-management.worker/presentation/controllers"
	"github.com/RenzoReccio/project-management.worker/presentation/middleware"
	"github.com/gin-gonic/gin"
)

func initializeReceiverHttpHandler(router *gin.Engine) {
	// Initialize all layers
	receiverService := mongoInfraestructure.NewReceiverService()

	receiverUsecase := application.NewInsertReceiverUseCase(
		receiverService,
	)

	receiverController := controllers.NewReceiverController(receiverUsecase)
	// Routers
	router.POST("receiver", receiverController.InsertReceiver)
}

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	initializeReceiverHttpHandler(router)
	router.Run(":" + port)
}
