package routes

import (
	"fmt"
	"net/http"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/gin-gonic/gin"
)

func UpdateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	username := context.GetString("username")

	if username != user.Username {
		context.JSON(http.StatusBadRequest, gin.H{"error": "You can only update your own user"})
		return
	}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

func getUserByUsername(context *gin.Context) {
	username := context.Param("username")
	user, err := models.GetUser(&username, nil)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func approveUser(context *gin.Context) {
	var approveStruct struct {
		Username string `binding:"required"`
		Approved bool   `binding:"required"`
	}

	err := context.ShouldBindJSON(&approveStruct)

	fmt.Printf("approveStruct: %v\n", approveStruct)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := context.GetString("username")
	user, err := models.GetUser(&username, nil)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = user.Approve(approveStruct.Username, approveStruct.Approved)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User approved successfully"})
}

func deleteUser(context *gin.Context) {
	var user struct {
		Username string `binding:"required"`
	}
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteUser(user.Username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
