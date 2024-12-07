package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24 * 2).Unix(),
	})
	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString string) (string, string, string, error) {
	secret := os.Getenv("JWT_SECRET")

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return "", "", "", errors.New("error parsing token")
	}
	if !parsedToken.Valid {
		return "", "", "", errors.New("invalid token")
	}
	username := parsedToken.Claims.(jwt.MapClaims)["username"].(string)
	email := parsedToken.Claims.(jwt.MapClaims)["email"].(string)
	role := parsedToken.Claims.(jwt.MapClaims)["role"].(string)

	return username, email, role, nil
}
