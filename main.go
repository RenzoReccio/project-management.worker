package main

import (
	"log"
	"os"

	application_createevent "github.com/RenzoReccio/project-management.worker/application/event/command/create-event"
	mongoInfraestructure "github.com/RenzoReccio/project-management.worker/infrastructure/mongo/event"
	"github.com/RenzoReccio/project-management.worker/presentation/controllers"
	"github.com/RenzoReccio/project-management.worker/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

func InitializeMediatR() {
	eventService := mongoInfraestructure.NewEventService()
	createProductCommandHandler := application_createevent.NewCreateEventCommandHandler(eventService)
	mediatr.RegisterRequestHandler[*application_createevent.CreateProductCommand, *string](createProductCommandHandler)

}

func initializeReceiverHttpHandler(router *gin.Engine) {
	// Initialize all layers
	receiverController := controllers.NewEventController()
	// Routers
	router.POST("event", receiverController.InsertEvent)
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
