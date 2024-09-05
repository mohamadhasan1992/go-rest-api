package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/models"
)

func RegisterUserToEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Unable to extract Event!"})
		return
	}
	event, err := models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found!"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Unabel to register to this event!"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"registration:": event})
	return
}

func DeleteUserRegisteration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found!"})
		return
	}
	event, err := models.GetEventDetail(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found!"})
		return
	}
	fmt.Println("userId", event.UserId)
	fmt.Println("userId", userId)
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Registeration is not yours!"})
		return
	}
	err = event.DeleteRegister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registeration"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registeration canceled!"})
	return
}
