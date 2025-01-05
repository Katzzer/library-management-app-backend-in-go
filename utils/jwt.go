package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-web/db"
	"time"
)

const secretKey = "supersecretkey" // TODO: change this secretKey, use Environmental Variables

func GenerateToken(email string, userId int64) (string, error) {
	// Create a new JWT token with the given claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // Token valid for 2 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	// Save the generated token to the database as the latest JWT for the user
	query := `UPDATE users SET latest_jwt_token = ? WHERE id = ?`
	_, err = db.DB.Exec(query, tokenString, userId)
	if err != nil {
		return "", fmt.Errorf("failed to save token to database: %v", err)
	}

	return tokenString, nil
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	tokenIdValid := parsedToken.Valid

	if !tokenIdValid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
