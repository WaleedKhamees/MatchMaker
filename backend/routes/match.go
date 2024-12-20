package routes

import (
	"net/http"
	"strconv"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/gin-gonic/gin"
)

func getMatches(context *gin.Context) {
	matches, err := models.GetMatches()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, matches)
}

func getMatchById(context *gin.Context) {
	matchIdStr := context.Param("id")
	matchId, err := strconv.Atoi(matchIdStr)

	match, err := models.GetMatchByID(matchId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, match)
}

func createMatch(context *gin.Context) {
	var match models.Match
	if err := context.ShouldBindJSON(&match); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := match.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	matchOutput, err := models.GetMatchByID(match.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, matchOutput)
}

func updateMatch(context *gin.Context) {
	var match models.Match
	if err := context.ShouldBindJSON(&match); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := match.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, match)
}
