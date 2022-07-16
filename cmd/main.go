package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mgr1054/go-ticket/pkg/controller"
)

func main() {
	log.Info("Starting API server")
	router := gin.Defautl()
	router.GET("/events", getEvents)
	router.Run("localhost:8080")
}