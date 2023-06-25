package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body models.Auth
	var user models.Auth
	c.BindJSON(&body)

	result := db.DB.Table("auths").Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect password",
		})
		return
	}

	//generate a token
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(os.Getenv("SECRET"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   user.ID,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	)
	s, err = t.SignedString(key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	//send the token to the user as a cookie
	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization", s, 3600*24*30, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "user logged in! and token sent as a cookie"})
}
