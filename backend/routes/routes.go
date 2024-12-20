package routes

import (
	"time"

	"github.com/WaleedKhamees/MatchMaker/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io/v2/socket"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/user", middlewares.MasterAuth, getAllUsers)
	server.GET("/user/:username", getUserByUsername)
	server.PUT("/user", UpdateUser)

	server.GET("/user/approve", middlewares.MasterAuth, getUnapprovedUsers)
	server.PUT("/user/approve", middlewares.MasterAuth, approveUser)
	server.POST("/user/reserve", middlewares.CustomerAuth, makeReservation)
	server.DELETE("/user/:username", middlewares.MasterAuth, deleteUser)

	server.POST("/register", register)
	server.POST("/login", login)

	server.GET("/match", getMatches)
	server.GET("/match/:id", getMatchById)
	server.POST("/match", middlewares.EFAAuth, createMatch)
	server.PUT("/match", middlewares.EFAAuth, updateMatch)
	server.DELETE("/match/:id", middlewares.EFAAuth, deleteMatch)

	server.GET("/match/:id/seat", getMatchSeats)
	server.GET("/match/:id/seat/:seatid", getSeatById)
	// server.DELETE("/match/seats", cancelReservation)

	server.GET("/stadium", getAllStadiums)
	server.GET("/stadium/:id", getStadiumById)
	server.POST("/stadium", middlewares.EFAAuth, createStadium)
	server.PUT("/stadium", middlewares.EFAAuth, updateStadium)
	server.DELETE("/stadium/:id", middlewares.EFAAuth, deleteStadium)

	server.GET("/team", getAllTeams)
	server.GET("/team/:id", getTeamById)
	server.POST("/team", middlewares.EFAAuth, createTeam)
	server.PUT("/team", middlewares.EFAAuth, updateTeam)
	server.DELETE("/team/:id", middlewares.EFAAuth, deleteTeam)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	initSocket(server)

}
func initSocket(server *gin.Engine) {
	c := socket.DefaultServerOptions()
	c.SetServeClient(true)
	c.SetConnectionStateRecovery(&socket.ConnectionStateRecovery{})
	c.SetAllowEIO3(true)
	c.SetPingInterval(300 * time.Millisecond)
	c.SetPingTimeout(200 * time.Millisecond)
	c.SetMaxHttpBufferSize(1000000)
	c.SetConnectTimeout(1000 * time.Millisecond)
	c.SetCors(&types.Cors{
		Origin:      "*",
		Credentials: true,
	})

	socketio := socket.NewServer(nil, nil)

	socketio.On("connection", func(clients ...interface{}) {
		client := clients[0].(*socket.Socket)

		client.On("message", func(args ...interface{}) {
			client.Emit("message-back", args...)
		})
		client.Emit("auth", client.Handshake().Auth)

		client.On("message-with-ack", func(args ...interface{}) {
			ack := args[len(args)-1].(socket.Ack)
			ack(args[:len(args)-1], nil)
		})
	})

	socketio.Of("/socket/match", nil).On("connection", func(clients ...interface{}) {
		client := clients[0].(*socket.Socket)

		client.On("subscribe", func(args ...interface{}) {
			handleSubscribeToMatch(socketio, client, args...)
		})

		client.On("reserve", func(args ...interface{}) {
			handleReserve(socketio, client, args...)
		})

		client.On("cancel", func(args ...interface{}) {
			handleCancel(socketio, client, args...)
		})

	})

	server.GET("/socket.io/*any", gin.WrapH(socketio.ServeHandler(c)))
	server.POST("/socket.io/*any", gin.WrapH(socketio.ServeHandler(c)))
}
