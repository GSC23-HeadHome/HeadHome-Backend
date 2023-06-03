package controllers

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	
	"github.com/gin-gonic/gin"

	"github.com/GSC23-HeadHome/HeadHome-Backend/logic"
	"github.com/GSC23-HeadHome/HeadHome-Backend/models"
	"github.com/GSC23-HeadHome/HeadHome-Backend/database"
)

// AddSOSLog handles the http request to create a new sos log 
// when the care receiver calls for assistance to return home
func AddSOSLog(c *gin.Context) {
	
	//1. Get previous sos log
	var sosLog models.SOSLog
	if err := c.ShouldBindJSON(&sosLog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      	return 
	}

	lastestSOSLog, err := database.ReadLatestSOSLog(sosLog.CrId)
	if err != nil {
	} 
	
	//2. Create incoming request
	jsonData, err := json.Marshal(sosLog)
    if err != nil {
        fmt.Println(err)
        return
    }
	bytesData := []byte(jsonData)

	res, err := database.CreateSOSLog(bytesData); 
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	//3. Update sos log (from 1.) 
	lostMap := map[string] string {
		"Status": "lost",
	}

	lostJson, err := json.Marshal(lostMap)
	if err != nil {
	}

	lostBytes := []byte(lostJson)

	if err := database.UpdateSOSLog(lostBytes, lastestSOSLog.SOSId); err != nil {
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"SOSId": res})
}

// GetAllSOSLogs retrieves all SOS logs related to the specified 
// care receiver
func GetAllSOSLogs(c *gin.Context) {
	result, err := database.ReadAllSOSLogs()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// GetLatestSOSLog reads the specified SOS log from the soslog
func GetLatestSOSLog(c *gin.Context) {
	id := c.Param("id")
	result, err := database.ReadLatestSOSLog(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}


// AcceptSOSRequest handles the http request from the volunteer when they attempt
// to provide assistance to the elderly. It changes the SOS log's status 
// from "lost" to "guided" and adds the volunteer details into the log.
// This is done after verifying that the volunteer has made contact with the 
// care receiver and has a valid certification. 
func AcceptSOSRequest(c *gin.Context) {
	
	type requestBody struct{
		VId string `json:"VId"`
		AuthID string `json:"AuthID"`
		SOSId string `json:"SOSId"`
	}

	var req requestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Retrieve existing SOS log
	sosLog, err := database.FindSOSLog(req.SOSId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "sos record not found"})
		return
	}
	
	// Retrieve care receiver involved
	careReceiver, err := database.ReadCareReceiver(sosLog.CrId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care receiver not found"})
		return
	}

	// Retrieve requesting volunteer
	volunteer, err := database.ReadVolunteer(req.VId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "volunteer not found"})
		return
	}

	//Authenticate volunteer and verify volunteer certification validity
	currentTime := time.Now().Unix()
	if volunteer.CertificationStart >= currentTime || volunteer.CertificationEnd <= currentTime {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "volunteer not certified"})
		return
	} else if req.AuthID != careReceiver.AuthID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "authentication failed"})
		return
	} else if sosLog.Status != "lost" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "care receiver have already received help, thank you!"})
		return
	} else {
		// Update SOS Log with new status and volunteer information
		data := map[string]interface{}{
			"VId": req.VId,
			"Volunteer": volunteer.Name,
			"VolunteerContactNum": volunteer.ContactNum,
			"Status": "guided",
		}
		bytesData, err := json.Marshal(data)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		
		err = database.UpdateSOSLog(bytesData, req.SOSId)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Send instructions for the way home to the care receiver
		careGiver, err := database.ReadCareGiver(careReceiver.CareGiver[0].Id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no care giver found"})
		}

		directions, err := logic.RetrieveDirections(fmt.Sprintf("%f,%f", sosLog.StartLocation.Lat, sosLog.StartLocation.Lng), careGiver.Address)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		//Send response
		resMsg := map[string]interface{} {
			"CrId": careReceiver.CrId,
			"Name": careReceiver.Name,
			"Address": careReceiver.Address,
			"ContactNum": careReceiver.ContactNum,
			"CgName": careGiver.Name,
			"CgContactNum": careGiver.ContactNum,
			"RouteGeom": directions.OverallPolyline,
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message":resMsg})

	} 
}

// UpdateSOSStatus handles the http request to modify the status of a SOS log
func UpdateSOSStatus(c *gin.Context){

	//Extract information for request body
	id := c.Param("id")
	reqBod, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      	return 
    }

	lastestSOSLog, err := database.ReadLatestSOSLog(id)
	if err != nil {
		return 
	}

	//convert io.Reader data type to []byte data type
	bytesData := []byte(reqBod)
	if err = database.UpdateSOSLog(bytesData, lastestSOSLog.SOSId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}