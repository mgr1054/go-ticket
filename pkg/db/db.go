package db

import (
	"fmt"

	"github.com/mgr1054/go-ticket/pkg/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "database"
	port     = "5432"
	user     = "admin"
	password = "p"
	dbname   = "postgres"
	sslmode  = "disable"
)

var DB *gorm.DB

func Connect() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	log.Info("Using DSN for DB:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Info("Event migrated to DB")

	err = db.AutoMigrate(&models.Ticket{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Info("Ticket migrated to DB")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Info("User migrated to DB")

	DB = db
}
