package models

import (
	"log"
	"golang.org/x/net/context"

	"golang.org/x/crypto/bcrypt"
)

func GetAllSuperAdminsDetails(ctx context.Context)(bool,map[string]SuperAdmins) {
	superAdmin := map[string]SuperAdmins{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("SuperAdmins").Value(&superAdmin)
	if err != nil{
		log.Fatal(err)
		return false, superAdmin
	}
	return true, superAdmin


}


func(m *SuperAdmins) EditSuperAdminDetails(ctx context.Context ,superAdminId string) (bool){
	superAdminsSettings := SuperAdminSettings{}
	superAdminInfo := SuperAdmins{}
	dB,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = dB.Child("SuperAdmins/"+superAdminId+"/Settings").Value(&superAdminsSettings)
	if err != nil{
		log.Fatal(err)
		return false
	}
	err = dB.Child("SuperAdmins/"+superAdminId).Value(&superAdminInfo)
	if err != nil{
		log.Fatal(err)
		return false
	}
	m.Settings.DateOfCreation = superAdminsSettings.DateOfCreation
	m.Settings.Status = superAdminsSettings.Status
	m.Info.LastName = superAdminInfo.Info.LastName
	m.Info.Password = superAdminInfo.Info.Password

	err = dB.Child("/SuperAdmins/"+superAdminId).Update(&m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}



func(m *SuperAdmins) EditSuperAdminPassword(ctx context.Context ,superAdminId string) (bool){
	superAdminsSettings := SuperAdminSettings{}
	superAdminInfo := SuperAdmins{}
	dB,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}

	err = dB.Child("SuperAdmins/"+superAdminId+"/Settings").Value(&superAdminsSettings)
	if err != nil{
		log.Fatal(err)
		return false
	}


	err = dB.Child("SuperAdmins/"+superAdminId).Value(&superAdminInfo)
	if err != nil{
		log.Fatal(err)
		return false
	}
	m.Settings.DateOfCreation = superAdminsSettings.DateOfCreation
	m.Settings.Status = superAdminsSettings.Status
	m.Info.LastName = superAdminInfo.Info.LastName
	m.Info.FirstName = superAdminInfo.Info.FirstName
	m.Info.PhoneNo = superAdminInfo.Info.PhoneNo
	m.Info.Email = superAdminInfo.Info.Email
	hashedPassword, err := bcrypt.GenerateFromPassword(m.Info.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false
	}
	m.Info.Password = hashedPassword
	err = dB.Child("/SuperAdmins/"+superAdminId).Update(&m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}


func IsEnteredPasswordCorrect(ctx context.Context ,superAdminId string,enteredOldPassword []byte) (bool){
	superAdminInfo := SuperAdmins{}
	dB,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = dB.Child("SuperAdmins/"+superAdminId).Value(&superAdminInfo)
	if err != nil{
		log.Fatal(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword(superAdminInfo.Info.Password, enteredOldPassword)
	if err !=nil{
		log.Println(err)
		return false
	}
	return true

}
