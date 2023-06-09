// Package models provides data models that represent documents in the Firestore collection.
// 
// These models define the structure and fields of each document in the collection, 
// enabling convenient manipulation and interaction with the Firestore data.
package models

// A Relationship is used by CareGiver and CareReceiver types to show the relationship betweeen
// instances of these objects. 
type Relationship struct{
	Id 				string `firestore:"id"`
	Relationship 	string `firestore:"relationship"`
}

// A CareGiver contains information about a care giver instance and the care receivers they care for.
type CareGiver struct {
	CgId 			string 			`firestore:"cg_id"`
	Name 			string 			`firestore:"name"`
	Address 		string 			`firestore:"address"`
	ContactNum 		string 			`firestore:"contact_num"`
	CareReceiver 	[]Relationship  `firestore:"care_receiver"`
	ProfilePic 		string 			`firestore:"profile_pic"`
}