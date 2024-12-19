package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/WaleedKhamees/MatchMaker/db"
	"github.com/WaleedKhamees/MatchMaker/utils"
)

type User struct {
	Username   string    `binding:"required"`
	Firstname  string    `binding:"required"`
	Lastname   string    `binding:"required"`
	Gender     string    `binding:"required"`
	Email      string    `binding:"required"`
	Password   string    `binding:"required"`
	Role       string    `binding:"required"`
	Birthdate  time.Time `binding:"required"`
	City       string    `binding:"required"`
	Address    string    `binding:"required"`
	Creditcard string
	Creditpin  string
	Approved   bool
}

func (u *User) Create() error {
	query := `INSERT INTO users
	(userName, firstName, lastName, email, gender, password, role, birthdate, city, approved, address)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("error preparing statement")
	}
	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.New("error hashing password")
	}

	_, err = stmt.Exec(u.Username, u.Firstname, u.Lastname, u.Email, u.Gender, hashPassword, u.Role, u.Birthdate, u.City, u.Approved, u.Address)
	if err != nil {
		return errors.New("username or email already exists")
	}

	return nil
}

func (u *User) Update() error {
	query := `
		UPDATE users 
		SET  firstName = ?, lastName = ? , 
		gender = ?, password = ?, role = ?,
		birthdate = ?, address = ?, city = ?, creditCard = ?,
		pin = ?, approved = ?
		where userName = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	_, err = stmt.Exec(
		u.Firstname, u.Lastname,
		u.Gender, hashPassword, u.Role,
		u.Birthdate, u.Address, u.City, u.Creditcard,
		u.Creditpin, u.Approved, u.Username)

	return err
}

func GetUser(username, email *string) (*User, error) {
	var query string

	if username != nil && *username != "" {
		query = `SELECT 
			userName, firstName, lastName, gender, 
			email, password, role, birthdate, city,
			address, approved 
		FROM users WHERE userName = ?`
	} else {
		query = `SELECT
			userName, firstName, lastName, gender,
			email, password, role, birthdate, city,
			address, approved
		FROM users WHERE email = ?`
	}

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user User
	var row *sql.Row

	if username != nil && *username != "" {
		row = stmt.QueryRow(*username)
	} else {
		row = stmt.QueryRow(*email)
	}

	err = row.Scan(&user.Username, &user.Firstname, &user.Lastname, &user.Gender,
		&user.Email, &user.Password, &user.Role, &user.Birthdate, &user.City,
		&user.Address, &user.Approved)

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

func (u *User) Approve(username string, approved bool) error {
	query := `SELECT approved FROM users WHERE userName = ?`
	stmt, err := db.DB.Prepare(query)

	fmt.Printf("username: %s\n", username)

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

func (u *User) Validate(password string) error {
	query := "SELECT password FROM users WHERE username = ? OR email = ?"
	row := db.DB.QueryRow(query, u.Username, u.Email)
	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		return err
	}
	valid := utils.ComparePassword(password, hashedPassword)
	if !valid {
		return errors.New("invalid password")
	}

	return nil

}

func GetAllUsers() ([]User, error) {
	query := `SELECT userName, password, firstName, lastName, email, gender, city, address, birthdate, role, approved FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var birthdate string

		err := rows.Scan(
			&user.Username, &user.Password, &user.Firstname, &user.Lastname,
			&user.Email, &user.Gender, &user.City, &user.Address, &birthdate,
			&user.Role, &user.Approved,
		)
		if err != nil {
			return nil, err
		}

		// Parse the birthdate string into a time.Time object
		user.Birthdate, err = time.Parse("2006-01-02T15:04:05Z07:00", birthdate) // Adjust format to match your DB format
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func GetUnapprovedUsers() ([]User, error) {
	query := `SELECT userName, firstName, lastName, gender, email, role, birthdate, city, address, approved FROM users WHERE approved = false`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var birthdate string

		err := rows.Scan(
			&user.Username, &user.Firstname, &user.Lastname, &user.Gender,
			&user.Email, &user.Role, &birthdate,
			&user.City, &user.Address, &user.Approved,
		)
		user.Password = ""

		if err != nil {
			return nil, err
		}

		// Parse the birthdate string into a time.Time object
		user.Birthdate, err = time.Parse("2006-01-02T15:04:05Z07:00", birthdate) // Adjust format to match your DB format
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
