package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Health(c *gin.Context) {
	log.Info("API Health is OK")
	c.JSON(http.StatusOK, gin.H{"message": "Alive"})
}
