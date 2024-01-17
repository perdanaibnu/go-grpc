package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/gogrpc"))

	if err != nil {
		log.Fatalf("Database connection failed %v", err.Error())
	}

	return db
}
