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
		INSERT INTO stadiums (name, capacity, vipRows, seatsPerRow)
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

func GetStadiumByID(stadiumid int) (Stadium, error) {
	stadium := Stadium{}
	query := `
		SELECT id, name, capacity, vipRows, seatsPerRow
		FROM stadiums
		WHERE id = ?
		`
	err := db.DB.QueryRow(query, stadiumid).Scan(&stadium.Id, &stadium.Name, &stadium.Capacity, &stadium.VipRows, &stadium.SeatsPerRow)
	return stadium, err
}

func GetAllStadiums() ([]Stadium, error) {
	stadiums := []Stadium{}
	query := `
		SELECT id, name, capacity, vipRows, seatsPerRow
		FROM stadiums
		`
	rows, err := db.DB.Query(query)
	if err != nil {
		return stadiums, err
	}
	defer rows.Close()
	for rows.Next() {
		stadium := Stadium{}
		err := rows.Scan(&stadium.Id, &stadium.Name, &stadium.Capacity, &stadium.VipRows, &stadium.SeatsPerRow)
		if err != nil {
			return stadiums, err
		}
		stadiums = append(stadiums, stadium)
	}
	return stadiums, nil
}

func (s *Stadium) Update() error {
	query := `
		UPDATE stadiums
		SET name = ?, capacity = ?, vipRows = ?, seatsPerRow = ?
		WHERE id = ?
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.Name, s.Capacity, s.VipRows, s.SeatsPerRow, s.Id)
	return err
}

func DeleteStadium(stadiumid int) error {
	query := `
		DELETE FROM stadiums
		WHERE id = ?
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(stadiumid)
	return err
}
