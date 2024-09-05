package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/middleware"
)

func RegisterAuthRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
}

func RegisterEventRoutes(server *gin.Engine) {
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

func RegisterTestRoutes(server *gin.Engine) {

	server.GET("/bindquery", PrintQueryParams)
	server.GET("/binduri/:name/:id", PrintUrlParams)
	server.GET("/multipart-form", PrintMultiPartForm)
	server.POST("/form-query", PrintQuery)

	server.GET("/log-async", LogAsync)
	server.GET("/log-sync", LogSync)
}
