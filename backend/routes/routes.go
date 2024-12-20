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
	server.DELETE("/match/seats", cancelReservation)

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

	// io.OnDisconnect("/", func(s socketio.Conn, reason string) {
	// 	log.Print("disconnected")
	// })

	// io.OnEvent("/", "reserve", func(s socketio.Conn, msg string) {
	// 	var seatStruct struct {
	// 		Token      string `json:"token" binding:"required"`
	// 		MatchID    int    `json:"matchid" binding:"required"`
	// 		SeatRow    int    `json:"seatrow" binding:"required"`
	// 		SeatColumn int    `json:"seatcol" binding:"required"`
	// 	}

	// 	err := json.Unmarshal([]byte(msg), &seatStruct)
	// 	if err != nil {
	// 		log.Println("error parsing", err)
	// 		return
	// 	}
	// 	fmt.Printf("%+v\n", seatStruct)
	// 	username, _, _, err := utils.VerifyToken(seatStruct.Token)

	// 	if err != nil {
	// 		log.Println("error verifying token", err)
	// 		return
	// 	}

	// 	match, err := models.GetMatchByID(seatStruct.MatchID)
	// 	if err != nil {
	// 		log.Println("error getting match", err)
	// 		return
	// 	}

	// 	if time.Until(match.Date) <= 0 {
	// 		log.Println("match has already started")
	// 		return
	// 	}
	// 	available, err := models.CheckSeatAvailability(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn))
	// 	if err != nil {
	// 		log.Println("error checking seat availability", err)
	// 		return
	// 	}
	// 	if !available {
	// 		log.Println("seat is not available")
	// 		return
	// 	}

	// 	err = models.ReserveSeat(int64(seatStruct.MatchID), int64(seatStruct.SeatRow), int64(seatStruct.SeatColumn), username)
	// 	if err != nil {
	// 		log.Println("error reserving seat", err)
	// 		return
	// 	}

	// 	var outputStruct struct {
	// 		typeofReq  string
	// 		SeatRow    int `binding:"required"`
	// 		SeatColumn int `binding:"required"`
	// 		Username   string
	// 	}
	// 	outputStruct.SeatRow = seatStruct.SeatRow
	// 	outputStruct.SeatColumn = seatStruct.SeatColumn
	// 	outputStruct.Username = username
	// 	outputStruct.typeofReq = "reserve"

	// 	outputMsg, err := json.Marshal(outputStruct)
	// 	if err != nil {
	// 		log.Println("error marshaling output struct", err)
	// 		return
	// 	}

	// 	fmt.Print("success")

	// 	io.BroadcastToRoom("/", "match:"+fmt.Sprint(seatStruct.MatchID), "reserve", string(outputMsg))
	// })

	// io.OnEvent("/", "cancel", func(s socketio.Conn, msg string) {
	// })

	// io.OnError("/", func(s socketio.Conn, e error) {
	// 	log.Println("meet error:", e)
	// })
	// io.OnDisconnect("/", func(s socketio.Conn, reason string) {
	// 	log.Println("closed", reason)
	// })

	// server.GET("/socket.io/*any", gin.WrapH(io))
	// server.POST("/socket.io/*any", gin.WrapH(io))

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

	})

	server.GET("/socket.io/*any", gin.WrapH(socketio.ServeHandler(c)))
	server.POST("/socket.io/*any", gin.WrapH(socketio.ServeHandler(c)))
}
