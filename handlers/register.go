package handlers

import (
	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var Body models.User
	c.BindJSON(&Body)
	result := db.DB.Table("users").Create(&Body)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
}
