package models

import (
	"time"

	"github.com/WaleedKhamees/MatchMaker/db"
)

type User struct {
	username   string    `binding:"required"`
	password   string    `binding:"required"`
	firstname  string    `binding:"required"`
	lastname   string    `binding:"required"`
	email      string    `binding:"required"`
	gender     string    `binding:"required"`
	city       string    `binding:"required"`
	birthdate  time.Time `binding:"required"`
	role       string    `binding:"required"`
	creditcard string
	creditpin  string
	approved   bool
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

	_, err = stmt.Exec(u.username, u.firstname, u.lastname,
		u.gender, u.email, u.password, u.role, u.birthdate,
		u.city, u.creditcard, u.creditpin, u.approved, u.username)

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
	err = stmt.QueryRow(username).Scan(&user.username, &user.firstname, &user.lastname,
		&user.gender, &user.email, &user.password, &user.role, &user.birthdate,
		&user.city, &user.creditcard, &user.creditpin, &user.approved)

	return &user, err
}
