package main

import (
	_ "net/http"

	"github.com/WaleedKhamees/MatchMaker/db"
	"github.com/WaleedKhamees/MatchMaker/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8000")
}
