package routes

import (
	"github.com/gin-gonic/gin"
)

func ResgisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.POST("/events",createEvent)
	r.GET("/events/:id", getEvent)
	r.PUT("/events/:id", updateEvent)
	r.DELETE("/events/:id", deleteEvent)
	r.POST("/users", createUser)
}

