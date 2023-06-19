package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("here is the db:" + os.Getenv("DATABASE_URL"))
	fmt.Println("here is the secret:" + os.Getenv("SECRET"))
	var dsn = os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		panic("Could not connect to the database")
	}

	log.Println("Database connection successful")
}
