package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/WaleedKhamees/MatchMaker/models"
	"github.com/WaleedKhamees/MatchMaker/utils"
	"github.com/zishang520/socket.io/v2/socket"
)

func handleReserve(socketio *socket.Server, client *socket.Socket, args ...interface{}) {
	// Define a struct to hold the seat data
	var seatStruct struct {
		Token      string `json:"token" binding:"required"`
		MatchID    int    `json:"matchid" binding:"required"`
		SeatRow    int    `json:"seatrow" binding:"required"`
		SeatColumn int    `json:"seatcol" binding:"required"`
	}

	var errorstruct struct {
		Typeofreq string `json:"typeofreq"`
		Error     string `json:"error"`
	}

	msg, ok := args[0].(string)
	if !ok {
		log.Println("Invalid message format, expected string")

		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Invalid message format, expected string"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))

		return
	}

	err := json.Unmarshal([]byte(msg), &seatStruct)
	if err != nil {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Error parsing JSON"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))

		return
	}

	fmt.Printf("%+v\n", seatStruct)

	username, _, _, err := utils.VerifyToken(seatStruct.Token)
	if err != nil {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Error verifying token"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))

		return
	}

	match, err := models.GetMatchByID(seatStruct.MatchID)
	if err != nil {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Error getting match"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))

		return
	}

	// Check if the match has already started
	if time.Until(match.Date) <= 0 {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Match has already started"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))

		return
	}

	available, err := models.CheckSeatAvailability(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn))

	if err != nil {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Error checking seat availability"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))

		return
	}

	if !available {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Seat is already reserved"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))
		return
	}

	canReserve, err := models.CanReserveSeat(seatStruct.MatchID, username)

	if err != nil {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Error checking if seat can be reserved"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))
		return
	}

	if !canReserve {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "cannot reserve more than one seat"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))
		return
	}

	err = models.ReserveSeat(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn), username)

	if err != nil {
		errorstruct.Typeofreq = "error"
		errorstruct.Error = "Error reserving seat"
		outputMsg, _ := json.Marshal(errorstruct)
		client.Emit("onreserve", string(outputMsg))

		return
	}

	var outputStruct struct {
		TypeOfReq  string `json:"typeofreq" binding:"required"`
		SeatRow    int    `json:"seatrow" binding:"required"`
		SeatColumn int    `json:"seatcol" binding:"required"`
		Username   string `json:"username" binding:"required"`
	}

	outputStruct.SeatRow = seatStruct.SeatRow
	outputStruct.SeatColumn = seatStruct.SeatColumn
	outputStruct.TypeOfReq = "reserve"
	outputStruct.Username = username

	// Marshal the output struct into JSON
	outputMsg, err := json.Marshal(outputStruct)
	if err != nil {
		log.Println("Error marshaling output struct:", err)
		return
	}

	fmt.Print("outputMsg: ", string(outputMsg))

	room := socket.Room("match-" + fmt.Sprint(seatStruct.MatchID))

	client.To(room).Emit("onreserve", string(outputMsg))
	client.Emit("onreserve", string(outputMsg))
}

func handleSubscribeToMatch(_ *socket.Server, client *socket.Socket, args ...interface{}) {
	msg, _ := args[0].(string)
	// if !ok {
	// 	log.Println("Invalid message format, expected string")
	// 	return
	// }

	var matchStruct struct {
		MatchID int `json:"matchid"`
	}

	err := json.Unmarshal([]byte(msg), &matchStruct)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	client.Join(socket.Room("match-" + fmt.Sprint(matchStruct.MatchID)))
}

func handleCancel(_ *socket.Server, client *socket.Socket, args ...interface{}) {
	var seatStruct struct {
		Token      string `json:"token" binding:"required"`
		SeatRow    int    `json:"seatrow" binding:"required"`
		SeatColumn int    `json:"seatcol" binding:"required"`
		MatchID    int    `json:"matchid" binding:"required"`
	}
	msg, ok := args[0].(string)

	if !ok {
		log.Println("Invalid message format, expected string")
		return
	}

	err := json.Unmarshal([]byte(msg), &seatStruct)

	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	username, _, _, err := utils.VerifyToken(seatStruct.Token)

	if err != nil {
		log.Println("Error verifying token:", err)
		return
	}

	match, err := models.GetMatchByID(seatStruct.MatchID)
	if err != nil {
		log.Println("Error getting match:", err)
		return
	}

	if time.Until(match.Date) < 72*time.Hour {
		log.Println("Cannot cancel reservation within 3 days of the match")
		return
	}

	err = models.DeleteSeat(match.Id, seatStruct.SeatRow, seatStruct.SeatColumn)
	if err != nil {

		log.Println("Error deleting seat:", err)
		return
	}

	var outputStruct struct {
		TypeOfReq  string `json:"typeofreq" binding:"required"`
		SeatRow    int    `json:"seatrow" binding:"required"`
		SeatColumn int    `json:"seatcol" binding:"required"`
		Username   string `json:"username" binding:"required"`
	}

	outputStruct.SeatRow = seatStruct.SeatRow
	outputStruct.SeatColumn = seatStruct.SeatColumn
	outputStruct.TypeOfReq = "cancel"
	outputStruct.Username = username

	outputMsg, err := json.Marshal(outputStruct)
	if err != nil {
		log.Println("Error marshaling output struct:", err)
		return
	}

	room := socket.Room("match-" + fmt.Sprint(seatStruct.MatchID))

	client.To(room).Emit("onreserve", string(outputMsg))
	client.Emit("onreserve", string(outputMsg))
}
