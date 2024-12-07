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
