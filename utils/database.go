package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBBas struct {
	Orm *gorm.DB
}

func NewDatabase() *DBBas {
	dsn := "host=146.190.84.201 user=training_7 password=12345678 dbname=training_7 port=3002 sslmode=disable TimeZone=Asia/Jakarta" //manggil db postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})                                                                         //package postgres
	if err != nil {
		log.Fatal(err)
	}

	dbObject := &DBBas{}
	dbObject.Orm = db

	return dbObject
}
