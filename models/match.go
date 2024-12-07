package models

import (
	"time"
)

type Match struct {
	id                int       `binding:"required"`
	homeTeamId        int       `binding:"required"`
	awayTeamId        int       `binding:"required"`
	stadiumId         int       `binding:"required"`
	date              time.Time `binding:"required"`
	mainReferee       string    `binding:"required"`
	assistantReferee1 string    `binding:"required"`
	assistantReferee2 string    `binding:"required"`
}
