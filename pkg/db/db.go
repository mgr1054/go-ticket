package db

import (

	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = "5432"
	user = "admin"
	password = "p"
	dbname = "database"
)

func connect() (*gorm.DB, error){
	var err error
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	log.Info("Using DSN for DB:", dsn) 

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	return db, nil
}
