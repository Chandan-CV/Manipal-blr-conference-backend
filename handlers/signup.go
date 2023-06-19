package handlers

import (
	"github.com/Chandan-CV/Manipal-blr-conference-backend/db"
	"github.com/Chandan-CV/Manipal-blr-conference-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUP(c *gin.Context) {
	var body models.Auth
	c.BindJSON(&body)

	var user models.User

	r := db.DB.Table("users").Where("email = ?", body.Email).First(&user)

	if r.Error != nil {
		c.JSON(500, gin.H{
			"message": "User not registerd!!",
		})
		return
	}

	var count int64
	db.DB.Table("auth").Where("email = ?", body.Email).Count(&count)

	if count != 0 {
		c.JSON(400, gin.H{
			"message": "User already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error, could not hash password",
			"error":   err,
		})
		return
	}

	result := db.DB.Table("auths").Create(&models.Auth{Email: body.Email, Password: string(hashedPassword), EmailVerified: false})

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"data":    user,
	})

}
