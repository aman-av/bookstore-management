package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// Retrieve database credentials from environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")

	// Construct the DSN (Data Source Name)
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the database
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// Assign the database connection to the global variable
	db = d
}

func GetDB() *gorm.DB {
	return db
}
