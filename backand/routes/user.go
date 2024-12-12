package routes

import (
	"fmt"
	"net/http"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/gin-gonic/gin"
)

// getAllUsers returns all users
//
//	@Summary		gets all users
//	@Description	gets all users from the database for the admin
//	@Tags			users
//	@Produce		json
//	@Success		200		{object}	[]models.User
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/user	[get]
func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

// UpdateUser updates the user
//
//	@Summary		updates the user
//	@Description	updates the user in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user body models.User true "User object that needs to be updated"
//	@Success		200		{object}  models.User
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/user	[put]
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

	context.JSON(http.StatusOK, gin.H{"user": user})
}

// getUserByUsername returns the user by username
//
// @Summary		gets the user by username
// @Description	gets the user by username from the database
// @Tags			users
// @Produce		json
// @Param			username path string true "Username of the user"
// @Success		200		{object}	models.User
// @Failure		500		{object}	models.ErrorResponse
// @Router			/user/{username}	[get]
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
