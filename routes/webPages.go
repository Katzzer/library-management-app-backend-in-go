package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getWelcomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
