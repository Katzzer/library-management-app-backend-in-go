package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func HealthCheckHandler(context *gin.Context) {
	uptime := time.Since(startTime)

	// Convert uptime to a human-readable format (e.g., hours, minutes, seconds)
	uptimeReadable := fmt.Sprintf("%02dh:%02dm:%02ds",
		int(uptime.Hours()),
		int(uptime.Minutes())%60,
		int(uptime.Seconds())%60,
	)

	context.JSON(200, gin.H{
		"status":     "healthy",
		"timestamp":  time.Now().UTC().Format(time.RFC3339),
		"start_time": startTime.Format(time.RFC3339),
		"uptime":     uptimeReadable,
	})
}
