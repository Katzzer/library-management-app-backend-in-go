package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", getWelcomePage)

	server.GET("/api/v1/test", testBackend)
	server.GET("/api/v1/get-time", getTime)

}
