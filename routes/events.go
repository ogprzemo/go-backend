package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/models"
	"strconv"
)

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event id"})
		return
	}
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(500, gin.H{"error": "Error getting event"})
		return
	}
	context.JSON(200, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"error": "Error getting events"})
		return
	}
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
	}

	context.GetInt64("userId")
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error creating event)"})
		return
	}

	context.JSON(201, event)
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event id"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event.ID = int(eventID)
	err = event.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error updating event"})
		return
	}

	context.JSON(200, event)
	if event.UserID != context.GetInt(strconv.Itoa(int(event.UserID))) {
		context.JSON(403, gin.H{"error": "You are not authorized to update this event"})
		return
	}
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventID)

	err = event.Delete()
	if err != nil {
		context.JSON(500, gin.H{"error": "Error deleting event"})
		return
	}
	context.JSON(204, nil)

	if event.UserID != context.GetInt(strconv.Itoa(int(event.UserID))) {
		context.JSON(403, gin.H{"error": "You are not authorized to delete this event"})
		return
	}
}
