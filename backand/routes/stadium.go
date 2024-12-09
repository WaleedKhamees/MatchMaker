package routes

import (
	"net/http"
	"strconv"

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

func getAllStadiums(context *gin.Context) {
	stadiums, err := models.GetAllStadiums()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, stadiums)
}

func getStadiumById(context *gin.Context) {
	stadiumIDstr := context.Param("id")
	stadiumID, err := strconv.Atoi(stadiumIDstr)

	stadium, err := models.GetStadiumByID(stadiumID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"stadium": stadium})
}
