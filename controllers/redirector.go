package controllers

import (
	"log"
	"net/http"
	"peterchu999/url-shorterner/services"

	"github.com/gin-gonic/gin"
)

func RedirectShort(c *gin.Context) {

	log.Println(c.ClientIP())
	shortUrl := c.Param("short")
	longUrl, err := services.GetRedirectUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, longUrl)
}
