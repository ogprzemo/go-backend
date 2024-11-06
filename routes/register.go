package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/models"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Error registering for event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Error registering for event"})
		return
	}
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Error cancelling registration"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Error cancelling registration"})
		return
	}
}
