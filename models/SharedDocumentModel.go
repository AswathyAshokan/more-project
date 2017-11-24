package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"
	//"app/passporte/helpers"
	"strconv"

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
func GetAllSharedDocumentsByCompany(ctx context.Context,companyTeamname string )(Expirations,bool,string,[]string,[][]string){
	//companyData :=map[string]Company{}
	var KeySlice []string
	var AllSharedfile [][]string

	selectedExpiry := Expirations{}
	var fullName string
	usersInCompany :=map[string] CompanyUsers{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil{
		log.Fatal(err)
	}
	err = db.Child("Company/"+companyTeamname+"/Users").Value(&usersInCompany)
	log.Println("company usersssss",usersInCompany)
	dataValue := reflect.ValueOf(usersInCompany)
	for _, key := range dataValue.MapKeys() {
		log.Println("kety out ",key.String())
		eachExpiry :=map[string]Expirations{}
		err = db.Child("Expirations/"+key.String()).Value(&eachExpiry)
		eachDataValues := reflect.ValueOf(eachExpiry)
		for _, k := range eachDataValues.MapKeys() {
			log.Println("haiiiiiii each key",k)
			if eachExpiry[k.String()].Info.Mode =="Public" {
				err = db.Child("Users/"+key.String()+"/Info/FullName").Value(&fullName)
				var tempSlice    []string
				KeySlice = append(KeySlice,k.String())
				tempSlice = append(tempSlice, eachExpiry[k.String()].Info.Description)
				expirationDate := strconv.FormatInt(int64(eachExpiry[k.String()].Info.ExpirationDate), 10)
				tempSlice = append(tempSlice, expirationDate)
				tempSlice = append(tempSlice, fullName)
				tempSlice = append(tempSlice,eachExpiry[k.String()].Info.Type)
				tempSlice = append(tempSlice,eachExpiry[k.String()].Info.TypeNumber)
				tempSlice = append(tempSlice, eachExpiry[k.String()].Info.DocumentId)
				log.Println("tempSlice",tempSlice)
				AllSharedfile = append(AllSharedfile, tempSlice)
				tempSlice = tempSlice[:0]
			}
		}
	}
	log.Println("AllSharedfile",AllSharedfile)
	return selectedExpiry,true,fullName,KeySlice,AllSharedfile
}