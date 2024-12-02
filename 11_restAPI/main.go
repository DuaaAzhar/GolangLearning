package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restAPI/db"
	"github.com/restAPI/models"
)

func main() {
	db.InitDB()
	server := gin.Default() //GET, POST, PUT, PATCH, DELETE

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080)

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// extracting incoming request data
	var event models.Event

	// BindJSON to bind the received JSON to new event
	// it works like Scan
	err := context.BindJSON(&event)
	fmt.Println(event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse data."})
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}
