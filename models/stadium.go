package models

import "github.com/WaleedKhamees/MatchMaker/db"

type Stadium struct {
	id          int    `binding:"required"`
	name        string `binding:"required"`
	capacity    int    `binding:"required"`
	vipRows     int    `binding:"required"`
	SeatsPerRow int    `binding:"required"`
}

func (s *Stadium) Save() error {
	query := `
		INSERT INTO staduims (name, capacity, vipRows, seatsPerRow)
		VALUES (?, ?, ?, ?)
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(s.name, s.capacity, s.vipRows, s.SeatsPerRow)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	s.id = int(id)
	return err
}
