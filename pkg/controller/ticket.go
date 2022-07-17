package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	"github.com/mgr1054/go-ticket/pkg/utils"
)

type TicketRequest struct {
	UserID		uint		`json:"user_id"`
	EventID		uint		`json:"event_id"`
	Price		string		`json:"price"`
	Event		models.Event `json:"event"`	

}


// @Summary 		Create Ticket by EventID
// @Description		Creates Ticket for EventID, also checks if enough capacity is available
// @Description		allowed: user
// @ID				create-ticket
// @Tags 			tickets
// @Produce 		json
// @Success 		200 {object} models.Ticket
// @Failure			401 {string} json "{"error":"Unauthorized for this route"}"
// @Failure			404 {string} json "{"error": "User not found"}"
// @Failure			500 {string} json "{"error": "Could not create Ticket"}"
// @Router 			/secured/tickets/{id} [get]
func CreateTicket (c *gin.Context) {

	if err := utils.CheckUserType(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	event := models.Event{}
	db.DB.First(&event, "id = ?", c.Param("id"))
	eventCapacity := event.Capacity

	usedCapacity := int64(0)
	db.DB.Model(&models.Ticket{}).Where("event_id = ?", c.Param("id")).Count(&usedCapacity)

	if usedCapacity == int64(eventCapacity) {
		c.JSON(http.StatusOK, gin.H{"info": "Unfortunately, this event is fully booked!"})
		return
	}

	var user models.User
	
	if err := db.DB.Where("username = ?", c.GetString("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	NewTicket := models.Ticket {
		UserID: user.ID,
		EventID: event.ID,
		Price: event.Price,
	}

	if err := db.DB.Create(&NewTicket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Ticket"})
		return
	} 

	c.JSON(http.StatusOK, NewTicket)
}

// @Summary 		Get Tickets By EventID
// @Description		Gives back a number of all sold tickets for this event
// @Description		allowed: admin
// @ID				get-tickets-by-event-id
// @Tags 			tickets
// @Produce 		json
// @Success 		200 {object} models.Event
// @Failure			401 {string} json "{"error":"Unauthorized for this route"}"
// @Failure			404 {string} json "{"error": "Tickets not found"}""
// @Router 			/secured/tickets/events/{id} [get]
func GetTicketsByEvent (c* gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}
	
	usedCapacity := int64(0)

	if err := db.DB.Model(&models.Ticket{}).Where("event_id = ?", c.Param("id")).Count(&usedCapacity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tickets not found"})
		return
	}

	c.JSON(http.StatusOK, usedCapacity)
}

// @Summary 		Get Tickets By UserID
// @Description		Gives back all tickets for user
// @Description		allowed: admin
// @ID				get-tickets-by-user-id
// @Tags 			tickets
// @Produce 		json
// @Success 		200 {object} models.Ticket
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Failure			500 {object} string
// @Router 			/secured/tickets/user/{id} [get]
func GetTicketsByID (c* gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	var ticket []models.Ticket

	if err := db.DB.Where("user_id = ?", c.Param("id")).Find(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tickets not found"})
		return
	}

	if len(ticket) < 1 {
		c.JSON(http.StatusOK, gin.H{"info": "There are no tickts for this user"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// @Summary 		Delete Ticket By ID
// @Description		Deletes Ticket by Ticket ID, available up until one week before the event
// @Description		allowed: admin, user
// @ID				delete-tickets-by-user-id
// @Tags 			tickets
// @Produce 		json
// @Success 		200 {string} json "{"message": "Ticket deleted"}"
// @Failure			401 {string} json "{"error":"Unauthorized for this route"}"
// @Failure			404 {string} json "{"error": "Tickets not found"}"
// @Failure			500 {string} json "{"error": "Could not parse time"}"
// @Router 			/secured/tickets/{id} [delete]
func DeleteTicketById (c *gin.Context) {

	if err := utils.CheckUserType(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}
	
	var ticket models.Ticket

	if err := db.DB.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found!"})
        return
    }

	now := time.Now()
	currentDate := now.AddDate(0, 0, 7)
	
	event := models.Event{}
	db.DB.First(&event, "id = ?", ticket.EventID)
	eventDate, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse time"})
	}

	if(currentDate.After(eventDate)) {
		c.JSON(http.StatusOK, gin.H{"info": "Unfortunately, you are too late to cancle your ticket!"})
		return
	}
	

	if err := db.DB.Delete(&ticket).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Ticket"})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}