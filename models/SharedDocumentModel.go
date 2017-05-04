package models

import (
	"log"
	"golang.org/x/net/context"

	"reflect"
)


//Fetch all the details of invite user from database

func GetAllInvitationDetail(ctx context.Context,userId string) (CompanyInvitations,bool) {
	//user := User{}
	companyData := map[string]Company{}
	var keySlice []string
	invitationData := CompanyInvitations{}
	db,err :=GetFirebaseClient(ctx,"")
	/*err = db.Child("/Invitation/"+userId+"/Info").Value(&invitationDetails)
	if err != nil {
		log.Fatal(err)
		return invitationDetails,false
	}*/
	err = db.Child("Company").Value(&companyData)
	if err != nil {
		log.Println("t3")
		return invitationData,false
	}
	dataValue := reflect.ValueOf(companyData)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, key := range keySlice {
		err = db.Child("Company/" + key + "/Invitation/" + userId).Value(&invitationData)
		log.Println("haii",invitationData)
		if err != nil {
			log.Println("t4")
			return invitationData,false
		}
	}
	return invitationData,true
}



func GetAllUserDetail(ctx context.Context,tempEmailId string ) (map[string]Users,bool) {
	usersDetails := map[string]Users{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil{
		log.Fatal(err)
		return usersDetails, false
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(tempEmailId).Value(&usersDetails)
	if err != nil{
		log.Println(err)
		log.Println(usersDetails)
		return usersDetails, false
	}
	return usersDetails,true
}


func GetExpireDetailsOfUser(ctx context.Context,specifiedUserId string ) (map[string]Expirations,bool) {
	expiryDetails := map[string]Expirations{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Fatal(err)
		return expiryDetails, false
	}
	err = db.Child("/Expirations/"+specifiedUserId).Value(&expiryDetails)
	if err != nil{
		log.Fatal(err)
		return expiryDetails, false
	}
	return expiryDetails,true


}


