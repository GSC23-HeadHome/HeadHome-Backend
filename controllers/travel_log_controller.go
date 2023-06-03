package controllers

import (
	"net/http"
	"io/ioutil"
	
	"github.com/gin-gonic/gin"

	"github.com/GSC23-HeadHome/HeadHome-Backend/database"
)

// AddTravelLog handles the http request to create a new travel log 
// when the care receiver's device uploads their location periodically
func AddTravelLog(c *gin.Context) {
	//Extract request body 
	reqBod, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	//convert to bytes
	data := []byte(reqBod)

	//Create 
	
	lastHome, err := database.CreateTravelLog(data)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"LastHome": lastHome})

}

// GetTravelLog retrives all travel logs of the specified care receiver
func GetTravelLog(c *gin.Context) {
	id := c.Param("id")
	result, err := database.ReadTravelLog(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// GetLatestTravelLog handles the http request to retrieve the lastest travel 
// log of the specified care receiver
func GetLatestTravelLog(c *gin.Context) {
	id := c.Param("id")
	result, err := database.ReadLatestTravelLog(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

