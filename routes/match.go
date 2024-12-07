package routes

import (
	"net/http"

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
