package models

import (
	"time"

	"github.com/WaleedKhamees/MatchMaker/db"
)

type Match struct {
	Id          int
	HomeTeamId  int       `binding:"required"`
	AwayTeamId  int       `binding:"required"`
	StadiumId   int       `binding:"required"`
	Date        time.Time `binding:"required"`
	MainReferee string    `binding:"required"`
	Lineman1    string    `binding:"required"`
	Lineman2    string    `binding:"required"`
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
		err := rows.Scan(&match.Id, &match.HomeTeamId, &match.AwayTeamId, &match.StadiumId, &match.Date, &match.MainReferee,
			&match.Lineman1, &match.Lineman2)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}
	return matches, nil
}

func (m *Match) Save() error {
	query := `
		INSERT INTO matches (homeTeamId, awayTeamId, stadiumId, date, mainReferee, lineman1, lineman2)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(m.HomeTeamId, m.AwayTeamId, m.StadiumId, m.Date,
		m.MainReferee, m.Lineman1, m.Lineman2)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.Id = int(id)

	return err

}

func (m *Match) Update() error {
	query := `
		UPDATE matches 
		SET homeTeamId = ?, awayTeamId = ?, stadiumId = ?, date = ?, mainReferee = ?, lineman1 = ?, lineman2 = ?
		where id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.HomeTeamId, m.AwayTeamId, m.StadiumId, m.Date, m.MainReferee, m.Lineman1, m.Lineman2, m.Id)
	return err
}

func GetMatchByID(matchId int) (*Match, error) {
	query := `
		SELECT id, homeTeamId, awayTeamId, stadiumId, date, mainReferee, lineman1, lineman2
		FROM matches
		WHERE id = ?
	`

	row := db.DB.QueryRow(query, matchId)
	var match Match
	err := row.Scan(&match.Id, &match.HomeTeamId, &match.AwayTeamId, &match.StadiumId, &match.Date, &match.MainReferee, &match.Lineman1, &match.Lineman2)
	if err != nil {
		return nil, err
	}

	return &match, nil
}

func Get 