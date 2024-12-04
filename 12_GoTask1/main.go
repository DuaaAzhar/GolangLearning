package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gotask1/db"
	"github.com/gotask1/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
