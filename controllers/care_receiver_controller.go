package controllers

import (
	"net/http"
	"io/ioutil"
	
	"github.com/gin-gonic/gin"

	"github.com/GSC23-HeadHome/HeadHome-Backend/database"
)


// AddCareReceiver handles the http request to register a new care receiver
func AddCareReceiver(c *gin.Context){
	
	reqBod, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      	return 
    }

	bytesData := []byte(reqBod)

	if err := database.CreateCareReceiver(bytesData); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// GetAllCareReceivers handles the http request to retrieve information of a list of all care 
// receivers
func GetAllCareReceivers(c *gin.Context){
	result, err := database.ReadAllCareReceivers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// GetCareReceiver handles the http request to retrieve information of a specified care receiver
func GetCareReceiver(c *gin.Context){
	id := c.Param("id")
	result, err := database.ReadCareReceiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// ContactCareGiver handles the http request to retrieve the specified care receiver's contact
// information
func ContactCareGiver(c *gin.Context){
	//Process request body
	type requestBody struct {
		CrId string `json:"CrId"`
		CgId string `json:"CgId"`
	}

	CrId := c.Query("CrId")
	CgId := c.Query("CgId")
	if CrId == "" || CgId == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Please enter CrId and CgId"})
		return
	}

	//Retrieve care receiver
	careReceiver, err := database.ReadCareReceiver(CrId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care receiver not found"})
		return
	}

	//Check access permission
	for _, cg := range careReceiver.CareGiver {
		if (cg.Id == CgId){
			//Retrieve care giver infromation
			careGiver, err := database.ReadCareGiver(CgId)
			if err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care giver not found"})
				return
			}

			//Send response message
			resMsg := map[string]interface{} {
				"CgContactNum": careGiver.ContactNum,
			}
			c.IndentedJSON(http.StatusOK, resMsg)
			return
		} 
	}
	//None of the linked care givers match requested care giver
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care giver does not match"})
	return
}

// UpdateCareReceiver handles the http request to update the specified care receiver's information
func UpdateCareReceiver(c *gin.Context) {
	id := c.Param("id")
	err := database.UpdateCareReceiver(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

// DeleteCareReceiver handles the http request to remove the specified care receiver's 
// information from the system
func DeleteCareReceiver(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteCareReceiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}