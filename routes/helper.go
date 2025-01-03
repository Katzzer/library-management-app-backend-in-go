package routes

import (
	"github.com/gin-gonic/gin"
	"time"
)

func getTime(context *gin.Context) {
	currentTime := time.Now()
	// Correct formats for time and date
	formattedTime := currentTime.Format("15:04:05.000")
	formattedDate := currentTime.Format("02.01.2006")

	context.JSON(200, gin.H{
		"time": formattedTime,
		"date": formattedDate,
	})
}

func testBackend(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Backend is working",
	})
}
