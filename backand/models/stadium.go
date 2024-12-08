package models

import "github.com/WaleedKhamees/MatchMaker/db"

type Stadium struct {
	Id          int
	Name        string `binding:"required"`
	Capacity    int    `binding:"required"`
	VipRows     int    `binding:"required"`
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
	result, err := stmt.Exec(s.Name, s.Capacity, s.VipRows, s.SeatsPerRow)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	s.Id = int(id)
	return err
}

func GetStadiumByID(stadiumid int64) (Stadium, error) {
	stadium := Stadium{}
	query := `
		SELECT id, name, capacity, vipRows, seatsPerRow
		FROM staduims
		WHERE id = ?
		`
	err := db.DB.QueryRow(query, stadiumid).Scan(&stadium.Id, &stadium.Name, &stadium.Capacity, &stadium.VipRows, &stadium.SeatsPerRow)
	return stadium, err
}
