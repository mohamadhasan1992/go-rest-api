package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/models"
)

func getEventHandler(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"events": events})
	return
}

func createNewEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Fill the form correctly!"})
		return
	}
	userId := context.GetInt64("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save events"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
	return
}

func getOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event Id!"})
		return
	}
	event, err := models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event with this Id!"})
		return
	}
	context.JSON(http.StatusCreated, event)
	return
}

func updateOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event Id!"})
		return
	}
	event, err := models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event with this Id!"})
		return
	}
	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Event is not yours!"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Fill the form correctly!"})
		return
	}
	updatedEvent.Id = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Coulld not update!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "updated successfully!"})
	return
}

func deleteOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event Id!"})
		return
	}
	event, err := models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event with this Id!"})
		return
	}
	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Event is not yours!"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Coulld not delete!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})
	return
}
