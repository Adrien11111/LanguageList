package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
	"os"
)

var DB *gorm.DB


func Connect() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || dbName == "" {
		log.Fatal("Database connection details are not set in environment variables")
	}

	var err error

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}
