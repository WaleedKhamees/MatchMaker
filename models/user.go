package models

import (
	"time"
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
