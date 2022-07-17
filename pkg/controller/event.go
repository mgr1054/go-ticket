package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	"github.com/mgr1054/go-ticket/pkg/utils"
)

type NewEvent struct {
	Band_Name	string		`json:"band_name" binding:"required"`
	Location	string		`json:"location" binding:"required"`
	Price		string		`json:"price" binding:"required"`
	Capacity	int			`json:"capacity" binding:"required"`
	Date 		string		`json:"date" binding:"required"`
}

type EventUpdate struct {
	Band_Name	string		`json:"band_name"`
	Location	string		`json:"location"`
	Price		string		`json:"price"`
	Capacity	int			`json:"capacity"`
	Date 		string		`json:"date"`
}

func GetEvents (c *gin.Context) {

	if err := utils.CheckUserType(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	var events []models.Event
	if err := db.DB.Find(&events).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not get events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

func CreateEvent (c *gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	var event NewEvent
	
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create Event"})
		return
	}
	
	newEvent := models.Event{
		Band_Name: event.Band_Name, 
		Location: event.Location, 
		Price: event.Price, 
		Capacity: 
		event.Capacity, 
		Date: event.Date,
	}

	if err := db.DB.Create(&newEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Event"})
		return
	} 

	c.JSON(http.StatusOK, newEvent)
}

func GetEventByID (c *gin.Context) {

	if err := utils.CheckUserType(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	var event models.Event

	if err := db.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func GetEventByLocation (c *gin.Context) {

	if err := utils.CheckUserType(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	var event []models.Event

	if err := db.DB.Where("location = ?", c.Param("location")).Find(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if len(event) < 1 {
		c.JSON(http.StatusOK, gin.H{"info": "There are no events in this location"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func GetEventByDate (c *gin.Context) {

	if err := utils.CheckUserType(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	var event []models.Event

	if err := db.DB.Where("date = ?", c.Param("date")).Find(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if len(event) < 1 {
		c.JSON(http.StatusOK, gin.H{"info": "There are no events at this date"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func UpdateEventById (c *gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}
	
	var event models.Event

	if err := db.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var updateEvent EventUpdate

	if err:= c.ShouldBindJSON(&updateEvent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Event could not be updated with provided data"})
        return
    }

	if err:= db.DB.Model(&event).Updates(models.Event{
		Band_Name: updateEvent.Band_Name, 
		Location: updateEvent.Location, 
		Price: updateEvent.Price, 
		Capacity: updateEvent.Capacity, 
		Date: updateEvent.Date,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Event"})
        return
	}

	c.JSON(http.StatusOK, event)
}

func DeleteEventById (c *gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}
	
	var event models.Event

	if err := db.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found!"})
        return
    }

	if err := db.DB.Delete(&event).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}