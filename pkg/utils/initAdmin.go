package utils

import (

	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	log "github.com/sirupsen/logrus"
)

func InitAdmin(){
	var admin models.User
	var exists bool

	db.DB.Model(&admin).Select("count (*) > 0").Where("username = ?", "admin").Find(&exists)

	if exists {
		log.Info("Admin found")
		return
	}

	log.Info("Admin not found, creating new Admin")
	admin = models.User{Name: "admin", Username:"admin", Email:"admin@go-ticket.com", Password: "p", Role: "admin"}

	err1 := db.DB.Create(&admin).Error; if err1 != nil {
		log.Fatalln("Error while creating Admin")
	}

	log.Info("Admin created")
}