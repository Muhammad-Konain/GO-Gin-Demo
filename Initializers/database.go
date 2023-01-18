package initializers

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "sqlserver://go-user:go-user123@localhost:1433?database=go-demo"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB.")
	}
}
