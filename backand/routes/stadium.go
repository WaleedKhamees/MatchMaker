package routes

import (
	"net/http"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/gin-gonic/gin"
)

func createStadium(context *gin.Context) {
	var stadium models.Stadium
	if err := context.ShouldBindJSON(&stadium); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := stadium.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, stadium)
}

func getStadiumById(context *gin.Context) {
	var stadiumStruct struct {
		Id int `binding:"required"`
	}

	if err := context.ShouldBindJSON(&stadiumStruct); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stadium, err := models.GetStadiumByID(int64(stadiumStruct.Id))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"stadium": stadium})
}
