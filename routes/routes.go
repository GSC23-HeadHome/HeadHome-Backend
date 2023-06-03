// Package routes provides the implementation of the API routes for the HeadHome-Backend server. 
//
// It defines the RESTful endpoints and associates them with the corresponding controller functions.
package routes

import (
	"github.com/gin-gonic/gin"
	
	"github.com/GSC23-HeadHome/HeadHome-Backend/controllers"
)

// InitRoutes declares router groups and initialises different 
// API endpoints for them. 
func InitRoutes(router *gin.Engine){
	// API health 
	router.HEAD("/", func(c *gin.Context){c.Status(200)})
	router.GET("/", func(c *gin.Context){c.String(200, "API HEALTHY")})

	
	// Volunteers
	volunteerR := router.Group("/volunteers")
	volunteerR.GET("", controllers.GetAllVolunteers)
	volunteerR.GET("/:id", controllers.GetVolunteer)
	volunteerR.POST("", controllers.AddVolunteer)
	volunteerR.PUT("/:id", controllers.UpdateVolunteer)
	volunteerR.DELETE("/:id", controllers.DeleteVolunteer)

	// Caregivers
	careGiverR := router.Group("/caregiver")
	careGiverR.GET("", controllers.GetAllCareGivers)
	careGiverR.GET("/:id", controllers.GetCareGiver)
	careGiverR.POST("", controllers.AddCareGiver)
	careGiverR.PUT("/:id", controllers.UpdateCareGiver)
	careGiverR.PUT("/:id/newcr", controllers.NewCareReceiver)
	careGiverR.PUT("/:id/rmcr", controllers.RemoveCareReceiver)
	careGiverR.DELETE("/:id", controllers.DeleteCareGiver)
	

	// Care Receiver
	careReceiverR := router.Group("/carereceiver")
	careReceiverR.GET("", controllers.GetAllCareReceivers)
	careReceiverR.GET("/:id", controllers.GetCareReceiver)
	careReceiverR.GET("/contactcg", controllers.ContactCareGiver)
	careReceiverR.POST("/route", controllers.PlanRoute)
	careReceiverR.POST("", controllers.AddCareReceiver)
	careReceiverR.POST("/:id/help", controllers.Help)
	careReceiverR.PUT("/:id", controllers.UpdateCareReceiver)
	careReceiverR.DELETE("/:id", controllers.DeleteCareReceiver)

	// SOS calls
	sosR := router.Group("/sos")
	sosR.GET("", controllers.GetAllSOSLogs)
	sosR.GET("/:id", controllers.GetLatestSOSLog)
	sosR.POST("/", controllers.AddSOSLog)
	sosR.PUT("/accept", controllers.AcceptSOSRequest)
	sosR.PUT("/:id", controllers.UpdateSOSStatus) 

	// Travel logs
	travelLogR := router.Group("/travellog")
	travelLogR.GET("/:id", controllers.GetLatestTravelLog)
	travelLogR.GET("/:id/all", controllers.GetTravelLog)
	travelLogR.POST("/:id", controllers.AddTravelLog)
}