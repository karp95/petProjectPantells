package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitDb() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=my_pass dbname=postgres port=5432 sslmode=disable"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
	}
	return Db, nil
}
