package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/restAPI/models"
)

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

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse data."})
		return
	}
	userID := context.GetInt64("userId")

	event.UserID = userID

	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data. Try again later"})
		return
	}

	updatedEvent.ID = eventId

	userID := context.GetInt64("userId")
	if userID != updatedEvent.UserID {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not Authorized to Update event"})
		return
	}

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event. Try again later"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "EVENT updated successfully!!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch event"})
		return
	}
	userID := context.GetInt64("userId")
	if userID != event.UserID {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not Authorized to Delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		fmt.Println("Error===>>>", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not Delete event. Try again later"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "EVENT Deleted successfully!!"})
}
