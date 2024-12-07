package models

type Seat struct {
	Id         int `binding:"required"`
	MatchId    int `binding:"required"`
	SeatRow    int `binding:"required"`
	SeatColumn int `binding:"required"`
	UserId     int `binding:"required"`
}
