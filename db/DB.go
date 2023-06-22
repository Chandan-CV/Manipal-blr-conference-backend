package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//connecting to the mysql database
	var dsn = os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		panic("Could not connect to the database")
	}

	// connecting to the sqlite database
	// db, err := gorm.Open(sqlite.Open("MaheblrConference.db"), &gorm.Config{})
	// DB = db
	// if err != nil {
	// 	panic("Could not connect to the database")
	// }

	log.Println("Database connection successful")
}
