package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/matches", getMatches)
	server.PUT("/user/update", UpdateUser)

	server.POST("/register", register)
	server.POST("/login", login)
}
