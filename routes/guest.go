package routes

import (
	"net/http"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/WaleedKhamees/MatchMaker/utils"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var email, username, password string
	email_err := context.BindJSON(&email)
	username_err := context.BindJSON(&username)
	password_err := context.BindJSON(&password)

	if email_err != nil && username_err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email or username is required"})
		return
	}
	if password_err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	user, err := models.GetUser(username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	generatedToken, err := utils.GenerateToken(user.Username, user.Email, user.Role)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": generatedToken})

}
func register(context *gin.Context) {

}
