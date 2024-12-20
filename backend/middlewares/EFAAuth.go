package middlewares

import (
	"net/http"

	"github.com/WaleedKhamees/MatchMaker/utils"
	"github.com/gin-gonic/gin"
)

func EFAAuth(context *gin.Context) {
	token := context.GetHeader("Authorization")

	username, email, role, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if role != "efa" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	context.Set("username", username)
	context.Set("email", email)
	context.Set("role", role)
	context.Next()
}
