package handlers

import (
	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAttendances(c *gin.Context) {
	user, _ := c.Get("user")
	userID := user.(models.Auth).ID
	var attendances []models.GetAttendancesResponse
	// result := db.DB.Table("attendances").Where("user_id = ?", userID).Find(&attendances)
	result := db.DB.Table("attendances").Select("attendances.id", "attendances.created_at", "attendances.is_food", "events.name").Joins("inner join events on attendances.event_id = events.id").Where("attendances.user_id = ?", userID).Scan(&attendances)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"attendances": attendances,
	})
	return
}
