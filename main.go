package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nitesh-mhatre/go-rest-api/models"
	"net/http"
	"github.com/nitesh-mhatre/go-rest-api/db"

)

func main() {
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/events", getEvents)
	r.POST("/events",createEvent)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getEvents(c *gin.Context){
	events, err := models.GetAllEvents()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)

}

func createEvent(c *gin.Context){
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		// fallback to query binding when json/form is absent
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)

}



