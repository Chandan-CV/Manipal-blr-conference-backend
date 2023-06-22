package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetEventJWTs(c *gin.Context) {
	var eventJWTs []models.EventJWT
	var events []models.Event
	db.DB.Find(&events)
	fmt.Println(events[0])
	for _, event := range events {
		var (
			key []byte
			t   *jwt.Token
			s   string
			err error
		)

		key = []byte(os.Getenv("SECRET"))
		t = jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"id":     event.ID,
				"isfood": event.IsFood,
				"name":   event.Name,
				"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
			},
		)
		s, err = t.SignedString(key)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		eventJWTs = append(eventJWTs, models.EventJWT{ID: event.ID, JWT: s, IsFood: event.IsFood, Name: event.Name})
	}

	c.JSON(200, gin.H{
		"eventJWTs": eventJWTs,
	})
	return
}
