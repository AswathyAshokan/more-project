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



func GetAllUserDetail(ctx context.Context,tempEmailId string ) (map[string]Users) {
	usersDetails := map[string]Users{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil{
		log.Fatal(err)
		return usersDetails
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(tempEmailId).Value(&usersDetails)
	log.Println("gggg",usersDetails)
	if err != nil{
		log.Println(err)

		return usersDetails
	}
	return usersDetails
}


func GetExpireDetailsOfUser(ctx context.Context,specifiedUserId string ) (map[string]Expirations,bool,string) {
	expiryDetails := map[string]Expirations{}
	var fullName string
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Fatal(err)
		return expiryDetails, false,fullName
	}
	err = db.Child("Users/"+specifiedUserId+"/Info/FullName").Value(&fullName)
	err = db.Child("/Expirations/"+specifiedUserId).Value(&expiryDetails)
	if err != nil{
		log.Fatal(err)
		return expiryDetails, false,fullName
	}
	return expiryDetails,true,fullName


}
func GetAllSharedDocumentsByCompany(ctx context.Context,companyTeamname string )(map[string]Expirations,bool,string){
	companyData :=map[string]Company{}
	userDetails := map[string]Users{}
	expiryDetails := map[string]Expirations{}
	var name string
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil{
		log.Fatal(err)
	}
	err = db.Child("Users").Value(&userDetails)
	dataValue := reflect.ValueOf(userDetails)
	for _, key := range dataValue.MapKeys() {
		err = db.Child("/Users/"+key.String()+"/Company/").Value(&companyData)
		companyDataValues := reflect.ValueOf(companyData)
		for _, k := range companyDataValues.MapKeys() {
			log.Println("k",k,"companyTeamname",companyTeamname)
			if k.String() == companyTeamname{
				err = db.Child("/Expirations/"+key.String()).Value(&expiryDetails)
				err = db.Child("/Users/"+key.String()+"/Info/FullName").Value(&name)
				return expiryDetails,true,name

			}
		}


	}

	return expiryDetails,true,name


}


