package models

// A Volunteer records information about a volunteer and 
// records their certification duration. This allows that 
// a central authority to manage all volunteers and will 
// update their certification on a discretionary basis. 
type Volunteer struct {
	VId 				string 	`firestore:"v_id"`
	Name 				string 	`firestore:"name"`
	ContactNum 			string 	`firestore:"contact_num"`
	CertificationStart	int64 	`firestore:"certification_start"`
	CertificationEnd 	int64 	`firestore:"certification_end"`
	ProfilePic 			string 	`firestore:"profile_pic"`
}