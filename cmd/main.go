package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/controller"
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/middleware"
	"github.com/mgr1054/go-ticket/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	db.Connect()
	utils.InitAdmin()
}

func main() {
	log.Info("Starting API server")
	router := gin.Default()

	api := router.Group("/api") 
	{
		api.GET("/", controller.Health)
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/events", controller.GetEvents)
			secured.GET("/events/:id", controller.GetEventByID)
			secured.GET("/events/location/:location", controller.GetEventByLocation)
			secured.GET("/events/date/:date", controller.GetEventByDate)
			secured.POST("/events", controller.CreateEvent)
			secured.PUT("/events/:id", controller.UpdateEventById)
			secured.DELETE("/events/:id", controller.DeleteEventById)
			secured.GET("/ticket/:id", controller.CreateTicket)
			secured.GET("/tickets/event/:id", controller.GetTicketsByEvent)
			secured.DELETE("/ticket/:id", controller.DeleteTicketById)
			secured.GET("/tickets/user/:id", controller.GetTicketsByID)
			secured.GET("/user/:id", controller.GetUserById)
			secured.PUT("/user/:id", controller.UpdateUserById)
			secured.DELETE("/user/:id", controller.DelteUserById)
		}
	}

	router.Run("0.0.0.0:8080")
}
