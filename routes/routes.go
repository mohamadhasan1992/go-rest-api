package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEventHandler)
	server.POST("/events", createNewEvent)
	server.GET("/events/:id", getOneEvent)
	server.PUT("/events/:id", updateOneEvent)
	server.DELETE("/events/:id", deleteOneEvent)
}
