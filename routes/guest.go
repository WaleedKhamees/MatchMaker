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
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Approved = false

	err = utils.ValidateEmail(user.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = utils.ValidateUsername(user.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = utils.ValidatePassword(user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Create()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully wait for approval"})
}
