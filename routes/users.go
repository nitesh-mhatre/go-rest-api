package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nitesh-mhatre/go-rest-api/models"
	"net/http"
)

func createUser(c *gin.Context){
	var user models.User

	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)

}

