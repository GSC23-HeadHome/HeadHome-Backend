// Package controllers provides controller or handler functions for handling API endpoints related to data collections. 
//
// It links the [route] and [database] packages to provide functionalities for managing and processing data collections, 
// including CRUD operations and other operations specific to the collections in the application.

package controllers

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/GSC23-HeadHome/HeadHome-Backend/models"
	"github.com/GSC23-HeadHome/HeadHome-Backend/database"
)

// AddCareGiver handles the http request to register a new care giver
func AddCareGiver(c *gin.Context){
	if err := database.CreateCareGiver(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// GetAllCareGivers handles the http request to retrieve information of a list of all care givers 
func GetAllCareGivers(c *gin.Context){
	result, err := database.ReadAllCareGivers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// GetCareGiver handles the http request to retrieve information of a specified care giver
func GetCareGiver(c *gin.Context){
	id := c.Param("id")
	result, err := database.ReadCareGiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// UpdateCareGiver handles the http request to update specified care giver details
// To update the list of care receivers under the care giver's care, use NewCareReceiver and DeleteCareReceiver
func UpdateCareGiver(c *gin.Context) {
	id := c.Param("id")
	err := database.UpdateCareGiver(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// NewCareReceiver handles the http request to add a new care receiver to a specified care giver's care receiver list
func NewCareReceiver(c *gin.Context) {
	cgId := c.Param("id")

	type reqBod struct {
		CrId			string	`json:"CrId"`
		AuthId			string	`json:"AuthId"`
		Relationship	string	`json:"Relationship"`
	}
	
	var req reqBod
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authentication: Check if care receiver has entered the correct care receiver AuthId
	careReceiver, err := database.ReadCareReceiver(req.CrId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if careReceiver.AuthID != req.AuthId {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errors.New("authentication failed")})
		return
	}	

	// Add care receiver to care giver document 
	newCareReceiver := models.Relationship{
		Id: req.CrId,
		Relationship: req.Relationship,
	}
	
	if err := database.NewCareReceiver(newCareReceiver, cgId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Mofify care give in care receiver document
	newCareGiver := []models.Relationship {
		{
			Id: cgId,
			Relationship: req.Relationship,
		},
	}

	if err := database.ChangeCareGiver(newCareGiver, req.CrId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// RemoveCareReceiver handles the http request to remove a care receiver from a specified care giver's care receiver list
func RemoveCareReceiver(c *gin.Context) {
	cgId := c.Param("id")

	type reqBod struct {
		CrId	string `json:"CrId"`
	}
	var req reqBod
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//R emove care receiver from care giver docs
	if err := database.RemoveCareReceiver(cgId, req.CrId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Remove care giver from care receiver docs
	if err := database.ChangeCareGiver([]models.Relationship{}, req.CrId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// DeleteCareGiver handles the http request to permanently remove a care giver from the database
func DeleteCareGiver(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteCareGiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}