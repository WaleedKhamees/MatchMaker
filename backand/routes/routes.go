package routes

import (
	"github.com/WaleedKhamees/MatchMaker/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.PUT("/user/update", UpdateUser)
	server.GET("/users", getAllUsers)
	server.PUT("/user/approve", middlewares.MasterAuth, approveUser)
	server.DELETE("/user/delete", middlewares.MasterAuth, deleteUser)
	server.POST("/user/reserve", middlewares.CustomerAuth, makeReservation)

	server.GET("/matches", getMatches)
	server.POST("/match/create", createMatch)
	server.PUT("/match/update", middlewares.EFAAuth, updateMatch)
	server.GET("/match/seats", getSeats)
	server.DELETE("/match/seats", cancelReservation)

	server.POST("/stadium/create", createStadium)

	server.POST("/team/create", createTeam)

	server.POST("/register", register)
	server.POST("/login", login)
}
