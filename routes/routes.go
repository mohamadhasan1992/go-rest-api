package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// AUTH
	server.POST("/signup", signup)
	server.POST("/login", login)
	// EVENTS
	server.GET("/events", getEventHandler)
	server.GET("/events/:id", getOneEvent)
	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middleware.Authenticate)
	authenticatedRoutes.POST("/events", createNewEvent)
	authenticatedRoutes.PUT("/events/:id", updateOneEvent)
	authenticatedRoutes.DELETE("/events/:id", deleteOneEvent)
	authenticatedRoutes.POST("/events/:id/register", RegisterUserToEvent)
	authenticatedRoutes.DELETE("/events/:id/register", DeleteUserRegisteration)
}
