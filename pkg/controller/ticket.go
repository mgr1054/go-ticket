package controller

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	"github.com/mgr1054/go-ticket/pkg/utils"
)

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

	NewTicket := models.Ticket {
		EventID: event.ID,
		Price: event.Price,
		Event: event,
	}

	if err := db.DB.Create(&NewTicket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Ticket"})
		return
	} 

	c.JSON(http.StatusOK, NewTicket)
}

func GetTicketsByEvent (c* gin.Context) {

	if err := utils.CheckUserType(c, "user"); err != nil {
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
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}