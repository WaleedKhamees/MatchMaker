package main

import (
	_ "net/http"

	"github.com/WaleedKhamees/MatchMaker/db"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.Run(":8000")
}
