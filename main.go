package main

import (
	"github.com/gin-gonic/gin"
	"go-web/routes"
	"os"
)

func main() {
	server := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.RegisterRoutes(server)
	server.Run(":" + port)
}
