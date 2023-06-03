// Package database provides functionality for managing the connection and interactions with the Firestore database. 
//
// It includes initialization of the Firestore client and CRUD operations for all Firestore collections. 
package database

import (
	"log"
	"context"
	
	"cloud.google.com/go/firestore"

	"github.com/GSC23-HeadHome/HeadHome-Backend/firebase_app"
)

var FBCtx context.Context
var Client *firestore.Client

// init automatically initialises the Firebase application context 
// and references to all Firebase collections when the database 
// package is first referenced 
func init(){
	var err error
	FBCtx = context.Background()
	Client, err = firebase_app.App.Firestore(FBCtx)
	if err != nil {
	  log.Fatalln(err)
	}

	//Init collections 
	InitVolunteers()
	InitCareGiver()
	InitCareReceiver()
	InitSosLog()
	InitTravelLog()
}

// CloseDB destructs the closes the Firebase App Client instance
func CloseDB(){
	Client.Close()
}