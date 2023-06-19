package main

import (
	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/handlers"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	// connecting to the database
	db.Connect()
	// migrating the database
	// db.MigrageDB()

	r := gin.Default()

	r.GET("/ping", handlers.Ping)
	r.POST("/register", handlers.Register)
	r.POST("/signup", handlers.SignUP)
	r.POST("/login", handlers.Login)
	r.GET("/logout", handlers.Logout)
	r.GET("/validate", middlewares.ReqAuth, handlers.Validate)
	r.Static("/docs", "./docs")
	r.Run() // listen and serve on
}
