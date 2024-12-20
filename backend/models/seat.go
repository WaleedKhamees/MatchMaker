package models

import (
	"github.com/WaleedKhamees/MatchMaker/db"
)

type Seat struct {
	Id         int    `binding:"required"`
	MatchId    int    `binding:"required"`
	SeatRow    int    `binding:"required"`
	SeatColumn int    `binding:"required"`
	Username   string `binding:"required"`
}

type seatOutput struct {
	SeatRow    int    `binding:"required"`
	SeatColumn int    `binding:"required"`
	Username   string `binding:"required"`
}

func GetAllSeats(matchid int) ([]seatOutput, error) {
	rows, err := db.DB.Query("SELECT seatrow, seatcolumn, Username FROM seats WHERE matchId = ?", matchid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []seatOutput

	for rows.Next() {
		var seat seatOutput
		err := rows.Scan(&seat.SeatRow, &seat.SeatColumn, &seat.Username)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}

func GetSeatById(matchid int, seatid int) ([]Seat, error) {
	rows, err := db.DB.Query("SELECT * FROM seats WHERE matchId = ? AND Id = ?", matchid, seatid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var seats []Seat
	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.Id, &seat.MatchId, &seat.SeatRow, &seat.SeatColumn, &seat.Username)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}

func DeleteSeat(matchid int, seatrow int, seatcol int) error {
	_, err := db.DB.Exec("DELETE FROM seats WHERE matchId = ? AND seatRow = ? AND seatColumn = ?", matchid, seatrow, seatcol)
	return err
}

func CheckSeatAvailability(matchId int64, sRow int64, sCol int64) (bool, error) {
	query := "SELECT * FROM seats WHERE seatRow = ? AND seatColumn = ? AND matchId = ?"
	rows, err := db.DB.Query(query, sRow, sCol, matchId)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return !rows.Next(), nil
}

func ReserveSeat(matchId int64, sRow int64, sCol int64, username string) error {
	query := "INSERT INTO seats (matchId, seatRow, seatColumn, username) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, matchId, sRow, sCol, username)
	return err
}
