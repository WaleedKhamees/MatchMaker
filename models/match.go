package models

import (
	"time"

	"github.com/WaleedKhamees/MatchMaker/db"
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

func GetMatches() ([]Match, error) {
	rows, err := db.DB.Query("SELECT * FROM matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var matches []Match
	for rows.Next() {
		var match Match
		err := rows.Scan(&match.id, &match.homeTeamId, &match.awayTeamId, &match.stadiumId, &match.date, &match.mainReferee, &match.assistantReferee1, &match.assistantReferee2)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}
	return matches, nil
}

func (m *Match) Save() error {
	query := `
		INSERT INTO matches (homeTeamId, awayTeamId, stadiumId, date, mainReferee, assistantReferee1, assistantReferee2)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.homeTeamId, m.awayTeamId, m.stadiumId, m.date, m.mainReferee, m.assistantReferee1, m.assistantReferee2)
	return err
}

func (m *Match) Update() error {
	query := `
		UPDATE matches 
		SET homeTeamId = ?, awayTeamId = ?, stadiumId = ?, date = ?, mainReferee = ?, assistantReferee1 = ?, assistantReferee2 = ?
		where id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.homeTeamId, m.awayTeamId, m.stadiumId, m.date, m.mainReferee, m.assistantReferee1, m.assistantReferee2, m.id)
	return err
}
