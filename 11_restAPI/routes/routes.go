package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)

}
