package main

import (
	"github.com/WaleedKhamees/MatchMaker/db"
	"github.com/WaleedKhamees/MatchMaker/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
)

// @title			User Service
// @version		1.0
// @description	Testing Swagger APIs.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8000
// @BasePath		/v1
// @schemes		http https
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db.InitDb()
	server := gin.Default()

	server.Use(cors.Default())

	io := socketio.NewServer(nil)

	routes.RegisterRoutes(server, io)

	go io.Serve()

	server.Run(":8000")
}
