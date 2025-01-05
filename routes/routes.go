package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	// Public Routes
	server.GET("/", getWelcomePage)
	server.GET("/api/v1/others/health-check", testBackend)
	server.GET("/api/v1/others/get-time", getTime)

	server.POST("/api/v1/registration/signup", signUp)
	server.POST("/api/v1/registration/login", login)

	// Authenticated Routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// Book Routes
	authenticated.GET("/api/v1/books", getAllBooks)
	authenticated.GET("/api/v1/borrowed-books", borrowBook)
	authenticated.GET("/api/v1/books/:id", getBook)
	authenticated.POST("/api/v1/books/:id/borrow", borrowBook)
	authenticated.POST("/api/v1/books/:id/return", returnBook)
}
