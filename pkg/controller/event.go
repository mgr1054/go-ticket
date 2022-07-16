package controller

import (
	"github.com/mgr1054/go-ticket/pkg/models"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

var events = []event{
	
}

func getEvents (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}