package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	"github.com/mgr1054/go-ticket/pkg/utils"
)

type NewEvent struct {
	Band_Name	string		`json:"band_name" binding:"required" example:"Deichkind"`
	Location	string		`json:"location" binding:"required" example:"Olympiastadion"`
	Price		string		`json:"price" binding:"required" example:"55"`
	Capacity	int			`json:"capacity" binding:"required" example:"35000"`
	Date 		string		`json:"date" binding:"required" example:"2022-10-11"`
}

type EventUpdate struct {
	Band_Name	string		`json:"band_name"`
	Location	string		`json:"location"`
	Price		string		`json:"price"`
	Capacity	int			`json:"capacity"`
	Date 		string		`json:"date"`
}


// @Summary 		Get All Events
// @Description		Sends Array Of Events 
// @Description		allowed: user, admin
// @ID				get-events
// @Tags 			events
// @Produce 		json
// @Success 		200 {object} []models.Event
// @Failure			404 {object} string
// @Router 			/secured/events [get]
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

// @Summary 		Create Event
// @Description		Creates a new Event
// @Description		allowed: admin
// @ID				create-event
// @Tags 			events
// @Accept			json
// @Produce 		json
// @Param			event body NewEvent true "Create Event"
// @Success 		201 {object} models.Event
// @Failure			400 {object} string
// @Failure			401 {object} string
// @Failure			500 {object} string
// @Router 			/secured/events [post]
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
		Capacity: event.Capacity, 
		Date: event.Date,
	}

	if err := db.DB.Create(&newEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Event"})
		return
	} 

	c.JSON(http.StatusCreated, newEvent)
}

// @Summary 		Get Event By ID
// @Description		Sends a Event with ID
// @Description		allowed: user, admin
// @ID				get-event-by-id
// @Tags 			events
// @Produce 		json
// @Success 		200 {object} models.Event
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Router 			/secured/events/{id} [get]
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

// @Summary 		Get Event By Location
// @Description		Sends Events for a Location
// @Description		allowed: user, admin
// @ID				get-event-by-location
// @Tags 			events
// @Produce 		json
// @Success 		200 {object} []models.Event
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Router 			/secured/events/{location} [get]
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

// @Summary 		Get Event By Date
// @Description		Sends Events for a Date
// @Description		allowed: user, admin
// @ID				get-event-by-date
// @Tags 			events
// @Produce 		json
// @Success 		200 {object} []models.Event
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Router 			/secured/events/{date} [get]
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

// @Summary 		Update Event By ID
// @Description		Updates Event with given ID
// @Description		allowed: admin
// @ID				update-event-by-id
// @Tags 			events
// @Produce 		json
// @Success 		200 {object} models.Event
// @Failure			400 {object} string
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Failure			500 {object} string
// @Router 			/secured/events/{id} [put]
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

// @Summary 		Delete Event By ID
// @Description		Deletes Event with given ID
// @Description		allowed: admin
// @ID				delete-event-by-id
// @Tags 			events
// @Produce 		plain
// @Success 		200 {object} string
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Failure			500 {object} string
// @Router 			/secured/events/{id} [delete]
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