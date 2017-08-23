package models
import (
	"log"
	"golang.org/x/net/context"

)
type WorkLocation struct {
	Info 		WorkLocationInfo
	Settings 	WorkLocationSettings
}
type WorkLocationInfo struct {
	CompanyTeamName			string
	WorkLocation       		string
	UsersAndGroupsInWorkLocation	UsersAndGroupsInWork
}

type WorkLocationSettings struct {
	DateOfCreation  	int64
	Status         	 	string
}
type UsersAndGroupsInWork struct {
	User 		map[string]WorkLocationUser
	Group 		map[string]WorkLocationGroup

}
type WorkLocationUser struct {
	FullName	string
	Status		string
}
type WorkLocationGroup struct{
	GroupName	string
	Members	 	map[string]GroupMemberNameInWorkLocation
}
type  GroupMemberNameInWorkLocation struct {
	MemberName	string

}


func(m *WorkLocation) AddWorkLocationToDb(ctx context.Context) (bool){
	log.Println("add group")
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("WorkLocation").Push(m)

	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}


func GetAllWorkLocationDetails(ctx context.Context,CompanyTeamName string) (map[string]WorkLocation,bool){
	log.Println("cp4")
	workLocationValues := map[string]WorkLocation{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("WorkLocation").OrderBy("Info/CompanyTeamName").EqualTo(CompanyTeamName).Value(&workLocationValues)
	//err = db.Child("WorkLocation").Value(&workLocationValues)
	if err != nil {
		log.Println("cp5")
		log.Fatal(err)
		return workLocationValues, false
	}
	return workLocationValues,true
}

func GetAllWorkLocationDetailsByWorkId(ctx context.Context,workLocationId string)(WorkLocation,bool)  {
	workLocationValues := WorkLocation{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/WorkLocation/"+workLocationId).Value(&workLocationValues)
	if err != nil {
		log.Fatal(err)
		return workLocationValues, false
	}
	return workLocationValues,true


}


func(m *WorkLocation)EditWorkLocationToDb(ctx context.Context,workLocationId string,companyTeamName string) (bool){
	workLocationValues := WorkLocation{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/WorkLocation/"+workLocationId).Value(&workLocationValues)
	if err != nil {
		log.Fatal(err)
		return false
	}
	m.Settings.DateOfCreation = workLocationValues.Settings.DateOfCreation
	m.Settings.Status = workLocationValues.Settings.Status
	log.Println("kjjjjh in model",m)
	err = db.Child("/WorkLocation/"+workLocationId).Update(m)

	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}


