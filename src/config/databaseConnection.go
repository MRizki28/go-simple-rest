package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnetion() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file!")
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE")

	database, err := gorm.Open(mysql.Open(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName))

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
