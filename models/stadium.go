package models

type Stadium struct {
	id          int    `binding:"required"`
	name        string `binding:"required"`
	capacity    int    `binding:"required"`
	vipRows     int    `binding:"required"`
	SeatsPerRow int    `binding:"required"`
}
