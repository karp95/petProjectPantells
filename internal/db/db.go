package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDb() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=my_pass dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
	}

	return db, nil
}
