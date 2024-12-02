package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restAPI/models"
)

func main() {
	server := gin.Default() //GET, POST, PUT, PATCH, DELETE
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// extracting incoming request data
	var event models.Event

	// it works like Scan
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "couldnot parse data"})
		return
	}
	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}
