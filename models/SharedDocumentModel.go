package models

import (
	"log"
	"golang.org/x/net/context"

)


//Fetch all the details of invite user from database

func GetAllInvitationDetail(ctx context.Context,userId string ) (inviteUser,bool) {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	invitationDetails := inviteUser{}
	err = db.Child("/Invitation/"+userId+"/Info").Value(&invitationDetails)
	if err != nil {
		log.Fatal(err)
		return invitationDetails,false
	}



	return invitationDetails,true
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
		log.Println("13")
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


