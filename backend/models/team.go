package models

import "github.com/WaleedKhamees/MatchMaker/db"

type Team struct {
	Id          int
	Name        string `binding:"required"`
	City        string `binding:"required"`
	StadiumId   int    `binding:"required"`
	Coach       string `binding:"required"`
	Description string `binding:"required"`
	Founded     int    `binding:"required"`
	LogoUrl     string `binding:"required"`
}

func (t *Team) Save() error {
	query := `
		INSERT INTO teams (name, city, stadiumId, coach, description, founded, logoUrl)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.Name, t.City, t.StadiumId, t.Coach, t.Description, t.Founded, t.LogoUrl)

	id, err := result.LastInsertId()

	t.Id = int(id)

	return err
}

func GetAllTeams() ([]Team, error) {
	teams := []Team{}
	query := `
		SELECT id, name, city, stadiumId, coach, description, founded, logoUrl
		FROM teams
		`
	rows, err := db.DB.Query(query)
	if err != nil {
		return teams, err
	}
	defer rows.Close()
	for rows.Next() {
		team := Team{}
		err := rows.Scan(&team.Id, &team.Name, &team.City, &team.StadiumId, &team.Coach, &team.Description, &team.Founded, &team.LogoUrl)
		if err != nil {
			return teams, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func GetTeamByID(teamid int) (Team, error) {
	team := Team{}
	query := `
		SELECT id, name, city, stadiumId, coach, description, founded, logoUrl
		FROM teams
		WHERE id = ?
		`
	err := db.DB.QueryRow(query, teamid).Scan(&team.Id, &team.Name, &team.City, &team.StadiumId, &team.Coach, &team.Description, &team.Founded, &team.LogoUrl)
	return team, err
}

func (t *Team) Update() error {
	query := `
		UPDATE teams
		SET name = ?, city = ?, stadiumId = ?, coach = ?, description = ?, founded = ?, logoUrl = ?
		WHERE id = ?
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(t.Name, t.City, t.StadiumId, t.Coach, t.Description, t.Founded, t.LogoUrl, t.Id)
	return err
}
func DeleteTeam(teamid int) error {
	query := `
		DELETE FROM teams
		WHERE id = ?
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(teamid)
	return err
}
