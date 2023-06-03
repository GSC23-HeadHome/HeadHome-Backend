package controllers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/GSC23-HeadHome/HeadHome-Backend/database"
)

// AddVolunteer handles the http request to register a new volunteer
func AddVolunteer(c *gin.Context){
	if err := database.CreateVolunteer(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// GetAllVolunteers handles the http request to retrieve a list of information
// of all volunteers
func GetAllVolunteers(c *gin.Context){
	result, err := database.ReadAllVolunteers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// GetVolunteer handles the http request to retrieve information of the specified
// volunteer
func GetVolunteer(c *gin.Context){
	id := c.Param("id")
	result, err := database.ReadVolunteer(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// UpdateVolunteer handles the http request to update information of the specified 
// volunteer
func UpdateVolunteer(c *gin.Context) {
	id := c.Param("id")
	err := database.UpdateVolunteer(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// DeleteVolunteer handles the http request to delete the specified volunteer's 
// information from the system
func DeleteVolunteer(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteVolunteer(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}