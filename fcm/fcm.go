// Package fcm implements push notifications for users by leveraging the Firebase Cloud Messaging (FCM) service.
//
// It provides initialization and access to the FCM client for sending push notification messages to users subscribed to a specific FCM topic.
package fcm

import (
	"fmt"
	"log"
	"strings"
	"context"

    "firebase.google.com/go/messaging"

	"github.com/GSC23-HeadHome/HeadHome-Backend/firebase_app"
)

var FCMContext context.Context
var FCMClient *messaging.Client

func init(){
	var err error

	FCMContext = context.Background()
	FCMClient, err = firebase_app.App.Messaging(FCMContext)
	if err != nil {
		log.Fatalln(err)
	}
}

func TopicSend(body map[string]string, topic string) (error){

	domainStartIndex := strings.Index(topic, "@")
	if (domainStartIndex > -1){
		topic = topic[:domainStartIndex]
	}
	message := &messaging.Message{
		Notification: &messaging.Notification{
            Title: "HeadHome",
            Body: fmt.Sprintf("%s requires your assistance!", topic),
        },
        Topic: topic,
	}
	  
	// Send a message to the devices subscribed to the provided topic.
	_, err := FCMClient.Send(FCMContext, message)
	if err != nil {
		return err
	}

	return nil
}