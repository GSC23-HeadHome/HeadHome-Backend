// Package firebase_app manages the connection between this server and Firebase, providing initialization and access to the Firebase app instance.
package firebase_app

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var App *firebase.App

func init() {
	var err error
	
	_, exists := os.LookupEnv(("FIREBASE_ADMIN_PRIVATE_KEY"))
	if !exists {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	conf := &firebase.Config{ProjectID: "gsc23-12e94"}
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_ADMIN_PRIVATE_KEY")))
	App, err = firebase.NewApp(context.Background(), conf ,opt)
	if err != nil {
	  	log.Fatalln(err)
	}
}
