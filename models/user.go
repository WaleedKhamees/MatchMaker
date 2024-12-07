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
