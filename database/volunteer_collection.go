package database

import (	
	"reflect"

	"google.golang.org/api/iterator"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"
	
	"github.com/GSC23-HeadHome/HeadHome-Backend/models"
)

var volunteerRef *firestore.CollectionRef

// InitVolunteers initialises the reference to the volunteers 
// Firebase collection
func InitVolunteers(){
	volunteerRef = Client.Collection("volunteers")
}

// CreateVolunteer creates a new document in the volunteers 
// Firebase Collection
func CreateVolunteer(c *gin.Context) (error) {
	var volunteer models.Volunteer
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		return err
	}

	_, err := volunteerRef.Doc(volunteer.VId).Set(FBCtx, volunteer)
	if err != nil {
		return err
	}

	return nil 
}

// ReadAllVolunteers reads and returns all documents from the 
// volunteers Firebase collection
func ReadAllVolunteers() ([]models.Volunteer, error) {
	var volunteers []models.Volunteer
	iter := volunteerRef.Documents(FBCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var volunteer models.Volunteer
		if err := doc.DataTo(&volunteer); err != nil {
			return nil, err
		}
		volunteers = append(volunteers, volunteer)
	}
	
	return volunteers, nil
}

// ReadVolunteer reads and returns the document with matching id 
// from the volunteers Firebase collection
func ReadVolunteer(id string) (models.Volunteer, error) {
	
	doc, err := volunteerRef.Doc(id).Get(FBCtx)
	if err != nil {
		return models.Volunteer{}, err
	}

	var volunteer models.Volunteer
	if err := doc.DataTo(&volunteer); err != nil {
		return models.Volunteer{}, err
	}
	return volunteer, nil
}


// UpdateVolunteer updates the document with matching id from the 
// the volunteers Firebase collection
func UpdateVolunteer(c *gin.Context, id string) (error){
	var volunteer models.Volunteer
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		return err
	}

	updates := []firestore.Update{}
    v := reflect.ValueOf(volunteer)
    for i := 0; i < v.NumField(); i++ {
        field := v.Type().Field(i)
        value := v.Field(i)
        if value.IsZero() {
            continue
        }
        updates = append(updates, firestore.Update{
            Path:  field.Tag.Get("firestore"),
            Value: value.Interface(),
        })
    }


	_, err := volunteerRef.Doc(id).Update(FBCtx, updates)
	if err != nil {
		return err
	}
	return nil
}

// DeleteVolunteer deletes the document with matching id from the 
// volunteer Firebase collection
func DeleteVolunteer(id string) (error) {
	_, err := volunteerRef.Doc(id).Delete(FBCtx)
	if err != nil {
		return err
	}
	return nil
}