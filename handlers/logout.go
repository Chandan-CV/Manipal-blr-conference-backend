package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization", "logged out :)", 10, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "user logged in! and token sent as a cookie"})
}
