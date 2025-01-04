package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", getWelcomePage)

	server.GET("/api/v1/others/health-check", testBackend)
	server.GET("/api/v1/others/get-time", getTime)

	server.POST("/api/v1/registration/signup", signUp)
	server.POST("/api/v1/registration/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/api/v1/events/get-all-events", getEvents)

}
