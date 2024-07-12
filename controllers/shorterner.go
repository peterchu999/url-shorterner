package controllers

import (
	"net/http"
	. "peterchu999/url-shorterner/dtos"
	services "peterchu999/url-shorterner/services"

	"github.com/gin-gonic/gin"
)

func CreateShort(c *gin.Context) {
	var createShortUrlDto CreateShortUrlDto
	if err := c.ShouldBindJSON(&createShortUrlDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl, errCreate := services.CreateShortUrl(createShortUrlDto)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"short_url": c.Request.Host + "/t/" + shortUrl})
}
