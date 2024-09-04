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
	}
	context.JSON(http.StatusAccepted, gin.H{"events": events})
}

func createNewEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Fill the form correctly!"})
		return
	}

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save events"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event Id!"})
	}
	event, err := models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event with this Id!"})
	}
	context.JSON(http.StatusCreated, event)
}

func updateOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event Id!"})
	}
	_, err = models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event with this Id!"})
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
}

func deleteOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event Id!"})
	}
	event, err := models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event with this Id!"})
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Coulld not delete!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})

}
