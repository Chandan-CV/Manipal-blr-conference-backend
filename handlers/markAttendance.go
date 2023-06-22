package handlers

import (
	"os"

	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MarkAttendance(c *gin.Context) {
	var body models.AttendanceBody
	c.BindJSON(&body)
	token := body.JWT
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{
			"message": "User not found, please login again",
		})
		return
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	// ... error handling
	if err != nil {
		c.JSON(500, gin.H{
			"message": "wrong qr code, please reach out to the admin",
		})
		return
	}

	// check if the user has already marked attendance

	var countAttendance int64
	result := db.DB.Table("attendances").Where("user_id = ? AND event_id = ?", user.(models.Auth).ID, claims["id"]).Count(&countAttendance)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if countAttendance > 0 {
		c.JSON(402, gin.H{
			"message": "You have already marked your attendance",
		})
		return
	}
	var Attendance models.Attendance
	Attendance.UserID = user.(models.Auth).ID
	Attendance.EventID = uint(claims["id"].(float64))
	Attendance.IsFood = claims["isfood"].(bool)
	result = db.DB.Table("attendances").Create(&Attendance)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error, could not mark attendance, please try again",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Attendance marked successfully",
	})
	return
}
