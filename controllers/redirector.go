package controllers

import (
	"fmt"
	"log"
	"net/http"
	"peterchu999/url-shorterner/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectShort(c *gin.Context) {

	log.Println(c.ClientIP())
	shortUrl := c.Param("short")
	longUrl, err := services.GetRedirectUrl(shortUrl)
	if err != nil {
		log.Println(longUrl, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !(strings.Contains(longUrl, "http:") || strings.Contains(longUrl, "https:")) {
		longUrl = fmt.Sprintf("https://%s", longUrl)
	}

	c.Redirect(http.StatusPermanentRedirect, longUrl)
}
