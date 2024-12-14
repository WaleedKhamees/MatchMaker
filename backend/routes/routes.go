package routes

import (
	"github.com/WaleedKhamees/MatchMaker/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/user", getAllUsers)
	server.GET("/user/:username", getUserByUsername)
	server.PUT("/user", UpdateUser)
	server.PUT("/user/approve", middlewares.MasterAuth, approveUser)
	server.POST("/user/reserve", middlewares.CustomerAuth, makeReservation)
	server.DELETE("/user/:username", middlewares.MasterAuth, deleteUser)

	server.POST("/register", register)
	server.POST("/login", login)

	server.GET("/match", getMatches)
	server.GET("/match/:id", getMatchById)
	server.POST("/match/create", createMatch)
	server.PUT("/match/update", middlewares.EFAAuth, updateMatch)
	server.GET("/match/:id/seat", getMatchSeats)
	server.GET("/match/:id/seat/:seatid", getSeatById)

	server.DELETE("/match/seats", cancelReservation)

	server.GET("/stadium", getAllStadiums)
	server.GET("/stadium/:id", getStadiumById)
	server.POST("/stadium/create", createStadium)

	server.GET("/team", getAllTeams)
	server.GET("/team/:id", getTeamById)
	server.POST("/team/create", createTeam)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
