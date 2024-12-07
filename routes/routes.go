package routes

import (
	"github.com/WaleedKhamees/MatchMaker/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.PUT("/user/update", UpdateUser)
	server.GET("/users", getAllUsers)
	server.PUT("/user/approve", middlewares.MasterAuth, approveUser)

	server.GET("/matches", getMatches)
	server.POST("/match/create", createMatch)

	server.POST("/stadium/create", createStadium)

	server.POST("/team/create", createTeam)

	server.POST("/register", register)
	server.POST("/login", login)
}
