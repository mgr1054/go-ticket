package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/controller"
	"github.com/mgr1054/go-ticket/pkg/db"
	log "github.com/sirupsen/logrus"
)

func init() {
	db.Connect()
}

func main() {
	log.Info("Starting API server")
	router := gin.Default()
	router.GET("/events", controller.GetEvents)
	router.GET("/events/:id", controller.GetEventByID)
	router.GET("/events/location/:location", controller.GetEventByLocation)
	router.GET("/events/date/:date", controller.GetEventByDate)
	router.POST("/events", controller.CreateEvent)
	router.PUT("/events/:id", controller.UpdateEventById)
	router.DELETE("events/:id", controller.DeleteEventById)
	router.Run("localhost:8080")
}
