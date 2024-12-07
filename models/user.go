package models

import (
	"time"

	"github.com/WaleedKhamees/MatchMaker/db"
)

type User struct {
	Username   string    `binding:"required"`
	Password   string    `binding:"required"`
	Firstname  string    `binding:"required"`
	Lastname   string    `binding:"required"`
	Email      string    `binding:"required"`
	Gender     string    `binding:"required"`
	City       string    `binding:"required"`
	Birthdate  time.Time `binding:"required"`
	Role       string    `binding:"required"`
	Creditcard string
	Creditpin  string
	Approved   bool
}

func (u *User) Update() error {
	query := `
		UPDATE users 
		SET userName = ?, firstName = ?, lastName = ? , 
		gender = ?, email = ?, password = ?, role = ?,
		birthdate = ?, address = ?, city = ?, creditCard = ?,
		pin = ?, approved = ?
		where userName = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Username, u.Firstname, u.Lastname,
		u.Gender, u.Email, u.Password, u.Role, u.Birthdate,
		u.City, u.Creditcard, u.Creditpin, u.Approved, u.Username)

	return err
}

func GetUser(username string) (*User, error) {
	query := `SELECT * FROM users WHERE userName = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	user := User{}
	err = stmt.QueryRow(username).Scan(&user.Username, &user.Firstname, &user.Lastname,
		&user.Gender, &user.Email, &user.Password, &user.Role, &user.Birthdate,
		&user.City, &user.Creditcard, &user.Creditpin, &user.Approved)

	return &user, err
}

func DeleteUser(username string) error {
	query := `DELETE FROM users WHERE userName = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	return err
}

func (u *User) approved(username string, approved bool) error {
	query := `SELECT approved FROM users WHERE userName = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	err = row.Scan(&u.Approved)

	if err != nil {
		return err
	}

	if u.Approved == approved == false {
		err = DeleteUser(username)
		if err != nil {
			return err
		}
	} else if approved == true {
		u.Approved = approved
		err = u.Update()
		if err != nil {
			return err
		}
	}
	return nil
}
