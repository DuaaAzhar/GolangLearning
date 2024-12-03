package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restAPI/db"
	"github.com/restAPI/routes"
)

func main() {
	db.InitDB()
	server := gin.Default() //GET, POST, PUT, PATCH, DELETE
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080)

}
