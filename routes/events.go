package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nitesh-mhatre/go-rest-api/models"
	"github.com/nitesh-mhatre/go-rest-api/utils"
	"net/http"
	"strconv"
	"strings"
)	

func getEvent(c *gin.Context){
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)

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

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	// Strip "Bearer " prefix if present
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	username, err := utils.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user, err := models.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		// fallback to query binding when json/form is absent
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = user.ID
	if err = event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)

}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = int(id)
	err = event.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})

}

func deleteEvent(c *gin.Context){
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	err = models.DeleteEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})	
}


