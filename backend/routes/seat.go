package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/gin-gonic/gin"
)

func getMatchSeats(context *gin.Context) {
	matchid, err := strconv.Atoi(context.Param("matchid"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seats, err := models.GetAllSeats(matchid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"seats": seats})
}

func getSeatById(context *gin.Context) {
	matchid, err := strconv.Atoi(context.Param("matchid"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seatid, err := strconv.Atoi(context.Param("seatid"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat, err := models.GetSeatById(matchid, seatid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"seat": seat})
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
	}

	username := context.GetString("username")

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

	if time.Until(match.Date) <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot make a reservation for a past match"})
		return
	}
	available, err := models.CheckSeatAvailability(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !available {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Seat is already reserved"})
		return
	}

	err = models.ReserveSeat(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn), username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Reservation made successfully", "match": match})
}
