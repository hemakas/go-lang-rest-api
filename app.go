package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rest-api/models"
)

func main() {
	r := gin.Default()

	// routes
	r.GET("/events", getEvents)
	r.POST("/events", createEvent)

	// listen to port 8080
	r.Run()
}

// view all events
func getEvents(context *gin.Context) {
	events := models.GetAllEvents
	context.JSON(http.StatusOK, events)
}

// create event
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	// check for errors
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request",
		})
		return
	}

	// save event
	event.Save()

	// response
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})

}
