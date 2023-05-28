// Package main is the entry point for the backend server of HeadHome-Backend, 
// a project focused on providing a robust and scalable solution to help alleviate the problem of dementia wandering.
//
// It initializes the server, configures the routes, and manages the database connection.
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/GSC23-HeadHome/HeadHome-Backend/database"
	"github.com/GSC23-HeadHome/HeadHome-Backend/routes"
)

// main is the application entry point and manages API endpoints and routing. 
func main(){
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Initialises all API endpoints
	routes.InitRoutes(router)	
	
	router.Run("0.0.0.0:8080")
	defer database.CloseDB()
}