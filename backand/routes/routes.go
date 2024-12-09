package routes

import (
	"github.com/WaleedKhamees/MatchMaker/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/user", getAllUsers)
	server.GET("/user/:username", getUserByUsername)
	server.PUT("/user", UpdateUser)
	server.PUT("/user/approve", middlewares.MasterAuth, approveUser)
	server.POST("/user/reserve", middlewares.CustomerAuth, makeReservation)
	server.DELETE("/user/delete", middlewares.MasterAuth, deleteUser)

	server.GET("/match", getMatches)
	server.GET("/match/:id", getMatchById)
	server.POST("/match/create", createMatch)
	server.PUT("/match/update", middlewares.EFAAuth, updateMatch)
	server.GET("/match/seats", getSeats)
	server.DELETE("/match/seats", cancelReservation)

	server.POST("/stadium/create", createStadium)
	server.GET("/stadium", getStadiumById)

	server.GET("/team", getAllTeams)
	server.GET("/team/:id", getTeamById)
	server.POST("/team/create", createTeam)

	server.POST("/register", register)
	server.POST("/login", login)
}
