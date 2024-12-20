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

	// Expect args[0] to be a string
	msg, ok := args[0].(string)
	if !ok {
		log.Println("Invalid message format, expected string")
		return
	}

	// Unmarshal the message into seatStruct
	err := json.Unmarshal([]byte(msg), &seatStruct)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	// Print out the parsed struct
	fmt.Printf("%+v\n", seatStruct)

	// Verify the token
	username, _, _, err := utils.VerifyToken(seatStruct.Token)
	if err != nil {
		log.Println("Error verifying token:", err)
		return
	}

	// Get the match from the database
	match, err := models.GetMatchByID(seatStruct.MatchID)
	if err != nil {
		log.Println("Error getting match:", err)
		return
	}

	// Check if the match has already started
	if time.Until(match.Date) <= 0 {
		log.Println("Match has already started")
		return
	}

	// Check if the seat is available
	available, err := models.CheckSeatAvailability(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn))
	if err != nil {
		log.Println("Error checking seat availability:", err)
		return
	}

	if !available {
		log.Println("Seat is not available")
		return
	}

	// Reserve the seat
	err = models.ReserveSeat(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn), username)
	if err != nil {
		log.Println("Error reserving seat:", err)
		return
	}

	// Prepare the output struct
	var outputStruct struct {
		typeofReq  string
		SeatRow    int `binding:"required"`
		SeatColumn int `binding:"required"`
	}

	outputStruct.SeatRow = seatStruct.SeatRow
	outputStruct.SeatColumn = seatStruct.SeatColumn
	outputStruct.typeofReq = "reserve"

	// Marshal the output struct into JSON
	outputMsg, err := json.Marshal(outputStruct)
	if err != nil {
		log.Println("Error marshaling output struct:", err)
		return
	}

	// Emit the result to the socket
	fmt.Println("Success")
	socketio.Emit("match:"+fmt.Sprint(seatStruct.MatchID), "reserve", string(outputMsg))
}

func handleSubscribeToMatch(socketio *socket.Server, client *socket.Socket, args ...interface{}) {
	msg, ok := args[0].(string)
	if !ok {
		log.Println("Invalid message format, expected string")
		return
	}

	var matchStruct struct {
		MatchID int `json:"matchid" binding:"required"`
	}

	err := json.Unmarshal([]byte(msg), &matchStruct)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	client.Join(socket.Room("match:" + fmt.Sprint(matchStruct.MatchID)))
}
