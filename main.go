package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/db"
	"github.com/mohamadhasan1992/go-rest-api.git/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
