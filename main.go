package main

import (
	"github.com/gin-gonic/gin"
	"go-web/db"
	"go-web/routes"
	"os"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.Static("/static", "./static")
	server.LoadHTMLGlob("templates/*")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.RegisterRoutes(server)
	err := server.Run(":" + port)
	if err != nil {
		return
	}
}
