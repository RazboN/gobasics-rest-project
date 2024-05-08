package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration) //cancel registering to an event
	authenticated.POST("/events", createEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.PUT("/events/:id", updateEvent)

	// server.POST("/events", middlewares.Authenticate, createEvent) //node.js middleware mantığı
}
