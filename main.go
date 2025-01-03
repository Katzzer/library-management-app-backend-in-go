package main

import (
	"github.com/gin-gonic/gin"
	"go-web/routes"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	routes.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
