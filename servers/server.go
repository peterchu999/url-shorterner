package servers

import (
	controllers "peterchu999/url-shorterner/controllers"
	. "peterchu999/url-shorterner/database"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {

	server := gin.New()

	server.GET("/t/:short", controllers.RedirectShort)
	server.GET("/c/:short", controllers.RedirectShort)
	server.GET("/f/:short", controllers.RedirectShort)
	server.POST("/short", controllers.CreateShort)

	ConnectMonggoDB()
	ConnectRedis()

	return server
}
