package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary 		Get Health
// @Description		Sends alive
// @ID				get-health
// @Tags 			health
// @Produce 		json
// @Success 		200 {object} map[string]interface{}
// @Router 			/ [get]
func Health(c *gin.Context) {
	log.Info("API Health is OK")
	c.JSON(http.StatusOK, gin.H{"message": "Alive"})
}
