package db

import (
	"fmt"

	"github.com/Chandan-CV/Manipal-blr-conference-backend/models"
)

func MigrageDB() {
	fmt.Println("Migrating the database")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Auth{})
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.Attendance{})
	fmt.Println("Migrated the database")

}
