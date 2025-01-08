package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/middlewares"
	"time"
)

var startTime time.Time

func RegisterRoutes(server *gin.Engine) {
	startTime = time.Now()

	//server.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"}, // Allow all origins
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

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
