package models

type Team struct {
	Id          int    `binding:"required"`
	Name        string `binding:"required"`
	City        string `binding:"required"`
	StadiumId   int    `binding:"required"`
	Coach       string `binding:"required"`
	Description string `binding:"required"`
	Founded     int    `binding:"required"`
	LogoUrl     string `binding:"required"`
}
