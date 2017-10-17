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
	var userKey []string
	//var userNameSLice []string

	//expiryDetails := map[string]Expirations{}
	//CompanyDetails := map[string]CompanyData{}
	selectedExpiry := Expirations{}
	var fullName string
	usersInCompany :=map[string] CompanyUsers{}
	/*var userKey []string
	var documentKey []string*/
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil{
		log.Fatal(err)
	}
	err = db.Child("/Company/"+companyTeamname+"/Users/").Value(&usersInCompany)
	log.Println("company usersssss",usersInCompany)
	//err = db.Child("Expirations").Value(&expiryDetails)
	dataValue := reflect.ValueOf(usersInCompany)
	for _, companyKey := range dataValue.MapKeys() {
		log.Println("kety out ", companyKey.String())
		userKey =append(userKey,companyKey.String())
		//userNameSLice =append(userNameSLice,usersInCompany[key].FullName)


	}
	var AllSharedfile [][]string
	for _, key := range userKey {
		eachExpiry :=map[string]Expirations{}

		err = db.Child("/Expirations/" + key).Value(&eachExpiry)
		log.Println("lllllll123")
		if len(eachExpiry) !=0{
			log.Println("insideeeee")
			eachDataValues := reflect.ValueOf(eachExpiry)
			for _, k := range eachDataValues.MapKeys() {
				//err = db.Child("/Expirations/"+key.String()+"/"+k.String()+"/Company").Value(&CompanyDetails)
				/*if CompanyDetails.CompanyName !=""&&CompanyDetails.CompanyStatus !=""{
					userKey = append(userKey,key.String())
					documentKey = append(documentKey,k.String())
				}*/
				/*companyDataValues := reflect.ValueOf(CompanyDetails)
				for _, companykey := range companyDataValues.MapKeys() {

					if CompanyDetails[companykey.String()].CompanyStatus != helpers.UserStatusDeleted{*/
				log.Println("test 1")
				log.Println("user key",key)
				log.Println("expiration key",k.String())
				//if companykey.String() == companyTeamname{
				//log.Println("companyTeamname",companyTeamname,companykey.String())
				err = db.Child("/Expirations/"+key+"/"+k.String()).Value(&selectedExpiry)

				if eachExpiry[k.String()].Info.Mode == "Public"  && len(selectedExpiry.Info.DocumentId) !=0{
					//err = db.Child("Users/" + key + "/Info/FullName").Value(&fullName)

					log.Println("k1111111111111111")
					var tempSlice        []string
					KeySlice = append(KeySlice, k.String())
					tempSlice = append(tempSlice, selectedExpiry.Info.Description)
					expirationDate := strconv.FormatInt(int64(selectedExpiry.Info.ExpirationDate), 10)
					tempSlice = append(tempSlice, expirationDate)
					tempSlice = append(tempSlice, usersInCompany[key].FullName)
					tempSlice = append(tempSlice, selectedExpiry.Info.DocumentId)
					log.Println("tempSlice", tempSlice)
					AllSharedfile = append(AllSharedfile, tempSlice)
					tempSlice = tempSlice[:0]
				}

				/*}*/
				/*}*/
			}
		}

	}

	/*}*/
	/*log.Println("userKey",userKey)*/


	log.Println("AllSharedfile",AllSharedfile)
	return selectedExpiry,true,fullName,KeySlice,AllSharedfile
}


