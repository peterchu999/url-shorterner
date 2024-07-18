package controllers

import (
	"net/http"
	"peterchu999/url-shorterner/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractPath(path string, n int) string {
	splitted := strings.Split(path, "/")
	return splitted[n]
}

func RedirectShort(c *gin.Context) {
	path := c.FullPath()
	redirectPath := extractPath(path, 1)
	shortUrl := c.Param("short")

	var getRedirectUrl func(string) (string, error)
	switch redirectPath {
	case "t":
		getRedirectUrl = services.GetRedirectUrl
	case "f":
		getRedirectUrl = services.GetRedirectUrlFast
	case "c":
		getRedirectUrl = services.GetCustomRedirectUrl
	default:
		getRedirectUrl = services.GetRedirectUrl
	}
	longUrl, err := getRedirectUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, longUrl)
}
