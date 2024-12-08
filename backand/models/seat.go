package models

import "github.com/WaleedKhamees/MatchMaker/db"

type Seat struct {
	Id         int `binding:"required"`
	MatchId    int `binding:"required"`
	SeatRow    int `binding:"required"`
	SeatColumn int `binding:"required"`
	UserId     int `binding:"required"`
}

func GetSeats(mId int64) ([]Seat, error) {
	rows, err := db.DB.Query("SELECT * FROM seats WHERE matchId = ?", mId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var seats []Seat
	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.Id, &seat.MatchId, &seat.SeatRow, &seat.SeatColumn, &seat.UserId)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}

func DeleteSeat(seatId int64) error {
	_, err := db.DB.Exec("DELETE FROM seats WHERE id = ?", seatId)
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
