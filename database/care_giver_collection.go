package database

import (
	"reflect"
	
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/GSC23-HeadHome/HeadHome-Backend/models"
)

var careGiverRef *firestore.CollectionRef

// InitCareGiver initialises the reference to the care_giver 
// Firebase collection
func InitCareGiver(){
	careGiverRef = Client.Collection("care_giver")
}

// CreateCareGiver creates a new document in the care_giver 
// Firebase collection
func CreateCareGiver(c *gin.Context) (error){
	var careGiver models.CareGiver
	if err := c.ShouldBindJSON(&careGiver); err != nil {
		return err
	}

	_, err := careGiverRef.Doc(careGiver.CgId).Set(FBCtx, careGiver)
	if err != nil {
		return err
	}
	return nil
}

// ReadAllCareGivers reads and returns all documents from the 
// care_giver Firebase collection
func ReadAllCareGivers() ([]models.CareGiver, error) {
	var careGivers []models.CareGiver
	iter := careGiverRef.Documents(FBCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var careGiver models.CareGiver
		if err := doc.DataTo(&careGiver); err != nil {
			return nil, err
		}
		careGivers = append(careGivers, careGiver)
	}
	
	return careGivers, nil
}

// ReadCareGiver reads and returns a document with the matching id
// from the care_giver Firebase collection
func ReadCareGiver(id string) (models.CareGiver, error) {
	
	doc, err := careGiverRef.Doc(id).Get(FBCtx)
	if err != nil {
		return models.CareGiver{}, err
	}

	var careGiver models.CareGiver
	if err := doc.DataTo(&careGiver); err != nil {
		return models.CareGiver{}, err
	}
	return careGiver, nil
}

// UpdateCareGiver updates a document with the matching id in the 
// care_giver Firebase collection (use NewCareReceiver and 
// RemoveCareReceiver to modify the care receiver list)
func UpdateCareGiver(c *gin.Context, id string) (error){
	var careGiver models.CareGiver
	if err := c.ShouldBindJSON(&careGiver); err != nil {
		return err
	}

	//remove care receiver modification during normal update
	careGiver.CareReceiver = []models.Relationship{}

	updates := []firestore.Update{}
    v := reflect.ValueOf(careGiver)
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

	_, err := careGiverRef.Doc(id).Update(FBCtx, updates)
	if err != nil {
		return err
	}
	return nil
}

// NewCareReceiver adds a new care receiver to the care receiver list 
// of the document with the matching id in the care_giver Firebase
// collection
func NewCareReceiver(newCareReceiver models.Relationship, id string) (error){
	update := []firestore.Update {
		{
			Path:  "care_receiver",
			Value: firestore.ArrayUnion(newCareReceiver),
		},
	}

	_, err := careGiverRef.Doc(id).Update(FBCtx, update)
	if err != nil {
		return err
	}
	return nil
}

// RemoveCareReceiver removes a new care receiver from the 
// care receiver list of the document with the matching id
// in the care_giver Firebase collection
func RemoveCareReceiver(CgId string, CrId string) (error){

	//ArrayRemove not available in go; Using manual update

	careGiver, err := ReadCareGiver(CgId)
	if err != nil {
		return err
	}

	careReceivers := careGiver.CareReceiver
	for i, cr := range careReceivers {
		if cr.Id == CrId {
			careReceivers = append(careReceivers[:i], careReceivers[i+1:]...)
		}
	}

	update := []firestore.Update {
		firestore.Update{
			Path:  "care_receiver",
			Value: careReceivers,
		},
	}

	
	if _, err := careGiverRef.Doc(CgId).Update(FBCtx, update); err != nil {
		return err
	}
	return nil
}

// Delete the document with the matching id from the 
// care_giver Firebase collection
func DeleteCareGiver(id string) (error) {
	_, err := careGiverRef.Doc(id).Delete(FBCtx)
	if err != nil {
		return err
	}
	return nil
}