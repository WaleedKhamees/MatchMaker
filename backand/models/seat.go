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

