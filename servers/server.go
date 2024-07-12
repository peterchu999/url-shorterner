package servers

import (
	"peterchu999/url-shorterner/controllers"
	. "peterchu999/url-shorterner/database"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	server := gin.Default()

	server.GET("/t/:short", controllers.RedirectShort)
	server.POST("/short", controllers.CreateShort)

	ConnectMonggoDB()
	return server
}
