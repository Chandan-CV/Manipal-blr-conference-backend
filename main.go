package main

import (
	"log"
	"os"

	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/handlers"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// connecting to the database
	db.Connect()
	// migrating the database
	db.MigrageDB()

	r := gin.Default()

	r.GET("/ping", handlers.Ping)
	r.POST("/register", handlers.Register)
	r.POST("/signup", handlers.SignUP)
	r.POST("/login", handlers.Login)
	r.GET("/logout", handlers.Logout)
	r.GET("/validate", middlewares.ReqAuth, handlers.Validate)
	r.Static("/docs", "./docs")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port) // listen and serve on
}
