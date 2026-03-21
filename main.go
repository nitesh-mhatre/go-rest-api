package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nitesh-mhatre/go-rest-api/models"
	"net/http"

)

func main() {
	r := gin.Default()
	r.GET("/events", getEvents)
	r.POST("/events",createEvent)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getEvents(c *gin.Context){
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)

}

func createEvent(c *gin.Context){
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		// fallback to query binding when json/form is absent
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.Save()

	c.JSON(http.StatusCreated, event)

}