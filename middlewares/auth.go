package middlewares

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-web/db"
	"go-web/utils"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Check if the token matches the latest token stored in the database
	var latestToken string
	query := `SELECT latest_jwt_token FROM users WHERE id = ?`
	err = db.DB.QueryRow(query, userId).Scan(&latestToken)
	if err != nil {
		if err == sql.ErrNoRows {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			return
		}
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	// Compare the provided token with the latest token in the database
	if latestToken != token {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is invalid or outdated"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
