package controllers

import (
	"github.com/gin-gonic/gin"
    "net/http"
    "BeerBook/internal/models"
    "strconv"
)

func GetBeers(c *gin.Context) {

	// DATA HARDCODED FOR LEARNING PURPOSES
	beers := []models.Beer {
		{ID: 1, Name: "IPA", Price: 30.5},
		{ID: 2, Name: "Weizen", Price: 5.5},
		{ID: 3, Name: "Porter", Price: 50.2},
		{ID: 4, Name: "RIS", Price: 100.5},
	}
	c.JSON(http.StatusOK, beers)
}

func GetBeer(c *gin.Context) {
		
	// DATA HARDCODED FOR LEARNING PURPOSES
	beers := []models.Beer {
		{ID: 1, Name: "IPA", Price: 30.5},
		{ID: 2, Name: "Weizen", Price: 5.5},
		{ID: 3, Name: "Porter", Price: 50.2},
		{ID: 4, Name: "RIS", Price: 100.5},
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
      		"error": "Error Bad Request",
    	})
    	return
	}
	
	for i := 0; i < len(beers); i++ {
		if beers[i].ID == int(id) {
			c.JSON(http.StatusOK, beers[i])
			return
		} 
	}

	c.JSON(http.StatusNotFound, gin.H{
      "error": "Not Found",
    })
	
}
