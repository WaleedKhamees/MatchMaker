package routes

import (
	"net/http"
	"time"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/gin-gonic/gin"
)

func getSeats(context *gin.Context) {
	var matchStruct struct {
		MatchId int `binding:"required"`
	}

	err := context.ShouldBindJSON(&matchStruct)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seats, err := models.GetSeats(int64(matchStruct.MatchId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, seats)
}

func cancelReservation(context *gin.Context) {
	var seatStruct struct {
		SeatID  int `binding:"required"`
		MatchID int `binding:"required"`
	}

	err := context.ShouldBindJSON(&seatStruct)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match, err := models.GetMatchByID(seatStruct.MatchID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if time.Until(match.Date) < 72*time.Hour {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot cancel reservation within 3 days of the match"})
		return
	}

	err = models.DeleteSeat(int64(seatStruct.SeatID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}

func makeReservation(context *gin.Context) {
	var seatStruct struct {
		MatchID    int `binding:"required"`
		SeatRow    int `binding:"required"`
		SeatColumn int `binding:"required"`
		UserID     int `binding:"required"`
	}

	err := context.ShouldBindJSON(&seatStruct)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	match, err := models.GetMatchByID(seatStruct.MatchID)
	stadium, err := models.GetStadiumByID(match.StadiumID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if time.Until(match.Date) <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot make a reservation for a past match"})
		return
	}
	if stadium.Capacity <= 

}
