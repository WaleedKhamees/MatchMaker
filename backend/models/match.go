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
	Username    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	MainReferee string    `binding:"required"`
	Lineman1    string    `binding:"required"`
	Lineman2    string    `binding:"required"`
}

type MatchOutput struct {
	Id          int
	HomeTeam    Team
	AwayTeam    Team
	Stadium     Stadium
	Date        time.Time
	MainReferee string
	Lineman1    string
	Lineman2    string
}

func GetMatches() ([]MatchOutput, error) {
	query := `
		SELECT m.id, m.homeTeamId, m.awayTeamId, m.stadiumId, m.date, m.mainReferee, m.lineman1, m.lineman2,
			   ht.id, ht.name, ht.city, ht.stadiumId,
			   at.id, at.name, at.city, at.stadiumId,
			   s.id, s.name, s.capacity, s.vipRows, s.seatsPerRow
		FROM matches m
		JOIN teams ht ON m.homeTeamId = ht.id
		JOIN teams at ON m.awayTeamId = at.id
		JOIN stadiums s ON m.stadiumId = s.id
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []MatchOutput
	for rows.Next() {
		var match MatchOutput
		var homeTeam, awayTeam Team
		var stadium Stadium

		err := rows.Scan(&match.Id, &match.HomeTeam.Id, &match.AwayTeam.Id, &match.Stadium.Id, &match.Date, &match.MainReferee, &match.Lineman1, &match.Lineman2,
			&homeTeam.Id, &homeTeam.Name, &homeTeam.City, &homeTeam.StadiumId,
			&awayTeam.Id, &awayTeam.Name, &awayTeam.City, &awayTeam.StadiumId,
			&stadium.Id, &stadium.Name, &stadium.Capacity, &stadium.VipRows, &stadium.SeatsPerRow)
		if err != nil {
			return nil, err
		}

		match.HomeTeam = homeTeam
		match.AwayTeam = awayTeam
		match.Stadium = stadium

		matches = append(matches, match)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func (m *Match) Save() error {
	query := `
		INSERT INTO matches (homeTeamId, awayTeamId, stadiumId, date, mainReferee, lineman1, lineman2, username)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(m.HomeTeamId, m.AwayTeamId, m.StadiumId, m.Date,
		m.MainReferee, m.Lineman1, m.Lineman2, m.Username)

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

func GetMatchByID(matchId int) (*MatchOutput, error) {
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

	team1, err := GetTeamByID(match.HomeTeamId)

	if err != nil {
		return nil, err
	}
	team2, err := GetTeamByID(match.AwayTeamId)
	if err != nil {
		return nil, err
	}
	stadium, err := GetStadiumByID(match.StadiumId)
	if err != nil {
		return nil, err
	}
	output := MatchOutput{
		Id:          match.Id,
		HomeTeam:    team1,
		AwayTeam:    team2,
		Stadium:     stadium,
		Date:        match.Date,
		MainReferee: match.MainReferee,
		Lineman1:    match.Lineman1,
		Lineman2:    match.Lineman2,
	}

	return &output, nil
}
