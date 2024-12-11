package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restAPI/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)

}
