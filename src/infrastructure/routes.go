package infrastructure

import (
	"sensor/src/application"
	"sensor/src/infrastructure/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, createEventUseCase *application.CreateEventUseCase) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	eventController := controllers.NewEventController(createEventUseCase)

	router.POST("/events", eventController.CreateEvent)
}
