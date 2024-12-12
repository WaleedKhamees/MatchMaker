package routes

import (
	"net/http"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/WaleedKhamees/MatchMaker/utils"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var loginStruct struct {
		Username string
		Email    string
		Password string `binding:"required"`
	}
	err := context.ShouldBindJSON(&loginStruct)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if loginStruct.Username == "" && loginStruct.Email == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email is required"})
		return
	}

	user, err := models.GetUser(&loginStruct.Username, &loginStruct.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !utils.ComparePassword(loginStruct.Password, user.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	if !user.Approved {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "User not approved yet"})
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
