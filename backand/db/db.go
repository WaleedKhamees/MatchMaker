package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		userName TEXT PRIMARY KEY,
		firstName TEXT NOT NULL,
		lastName TEXT NOT NULL,
		gender TEXT NOT NULL,
		email TEXT NOT NULL unique,
		password TEXT NOT NULL,
		role TEXT NOT NULL,
		birthdate datetime NOT NULL,
		address TEXT,
		city TEXT NOT NULL,
		creditCard TEXT,
		pin TEXT,
		approved INTEGER NOT NULL	
	);`

	staduimsTable := `
	CREATE TABLE IF NOT EXISTS staduims (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		capacity INTEGER NOT NULL,
		vipRows INTEGER NOT NULL,
		seatsPerRow INTEGER NOT NULL
	);`

	matchesTable := `
	CREATE TABLE IF NOT EXISTS matches (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		homeTeamId INTEGER NOT NULL REFERENCES teams(id),
		awayTeamId INTEGER NOT NULL REFERENCES teams(id),
		stadiumId INTEGER NOT NULL REFERENCES staduims(id),
		date datetime NOT NULL,
		mainReferee TEXT NOT NULL,
		lineman1 TEXT NOT NULL,
		lineman2 TEXT NOT NULL
	);`

	teamsTable := `
	CREATE TABLE IF NOT EXISTS teams (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		city TEXT NOT NULL,
		stadiumId INTEGER NOT NULL REFERENCES staduims(id),
		founded TEXT NOT NULL,
		coach TEXT NOT NULL,
		description TEXT NOT NULL,
		logoUrl TEXT NOT NULL
	);`

	seatsTable := `
	CREATE TABLE IF NOT EXISTS seats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		matchId INTEGER NOT NULL REFERENCES matches(id),
		seatRow INTEGER NOT NULL,
		seatColumn INTEGER NOT NULL,
		userId INTEGER REFERENCES users(id)
	);`

	_, err := DB.Exec(usersTable)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(staduimsTable)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(matchesTable)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(teamsTable)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(seatsTable)
	if err != nil {
		panic(err)
	}

}
