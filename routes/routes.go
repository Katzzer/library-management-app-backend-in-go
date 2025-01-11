package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-web/middlewares"
	"time"
)

var startTime time.Time

func RegisterRoutes(server *gin.Engine) {
	startTime = time.Now()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://localhost:3000", "https://library-management-app-frontend.pavelkostal.com", "https://library-management-app-frontend-angular.pavelkostal.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow cookies to be sent
		MaxAge:           12 * time.Hour,
	}))

	// Public Routes
	server.GET("/", getWelcomePage)
	server.GET("/api/health-check", HealthCheckHandler)

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
