package routes

import (
	"net/http"
	"strconv"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/gin-gonic/gin"
)

func createTeam(context *gin.Context) {
	var team models.Team
	if err := context.ShouldBindJSON(&team); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := team.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, team)
}
func getTeamById(context *gin.Context) {
	teamIdStr := context.Param("id")
	teamId, err := strconv.Atoi(teamIdStr)

	team, err := models.GetTeamByID(teamId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, team)
}
func getAllTeams(context *gin.Context) {
	teams, err := models.GetAllTeams()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, teams)
}
func updateTeam(context *gin.Context) {
	var team models.Team
	if err := context.ShouldBindJSON(&team); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := team.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, team)
}
func deleteTeam(context *gin.Context) {
	teamIdStr := context.Param("id")
	teamId, err := strconv.Atoi(teamIdStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DeleteTeam(teamId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "team deleted"})
}
