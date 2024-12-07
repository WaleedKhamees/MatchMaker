package utils

import (
	"errors"
	"os"
	"time"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
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
			return nil, errors.New("Unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return "", "", "", errors.New("Error parsing token")
	}
	if !parsedToken.Valid {
		return "", "", "", errors.New("Invalid token")
	}
	username := parsedToken.Claims.(jwt.MapClaims)["username"].(string)
	email := parsedToken.Claims.(jwt.MapClaims)["email"].(string)
	role := parsedToken.Claims.(jwt.MapClaims)["role"].(string)

	return username, email, role, nil
}
