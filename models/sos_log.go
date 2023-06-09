package models

// A SOSLog records a SOS call that has been made, its status,
// and the volunteer who responded to that call. 
type SOSLog struct {
	CrId 				string 	`json:"CrId" firestore:"cr_id"`
	Datetime 			int64 	`json:"Datetime" firestore:"datetime"`
	SOSId				string 	`json:"SOSId" firestore:"sos_id"`
	StartLocation struct{
		Lat float64 `json:"Lat" firestore:"lat"`
		Lng float64 `json:"Lng" firestore:"lng"`
	} `json:"StartLocation" firestore:"start_location"`
	Status 				string 	`json:"Status" firestore:"status"`
	VId 				string 	`json:"VId" firestore:"v_id"`
	Volunteer 			string 	`json:"Volunteer" firestore:"volunteer"`
	VolunteerContactNum	string 	`json:"VolunteerContactNum" firestore:"volunteer_contact_num"`
}