package middlewares

import (
	"net/http"

	"github.com/WaleedKhamees/MatchMaker/utils"
	"github.com/gin-gonic/gin"
)

func MasterAuth(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")

	username, email, role, err := utils.VerifyToken(tokenString)

	if err != nil || role != "master" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	context.Set("username", username)
	context.Set("email", email)
	context.Set("role", role)
	context.Next()
}
