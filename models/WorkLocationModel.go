package models
import (
	"log"
	"golang.org/x/net/context"

	"app/passporte/helpers"
	"strings"
	"reflect"
)
type WorkLocation struct {
	Info 		WorkLocationInfo
	Settings 	WorkLocationSettings
}
type WorkLocationInfo struct {

	CompanyTeamName			string
	WorkLocation       		string
	Latitude			string
	Longitude			string
	StartDate			int64
	EndDate				int64
	DailyStartDate                   int64
	DailyEndDate			int64
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


func(m *WorkLocation) AddWorkLocationToDb(ctx context.Context,companyTeamName string) (bool){
	log.Println("add group")
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	workData,err := db.Child("WorkLocation").Push(m)

	if err != nil {
		log.Println(err)
		return false
	}
	workDataString := strings.Split(workData.String(),"/")
	workLocationUniqueID := workDataString[len(workDataString)-2]
	userData := reflect.ValueOf(m.Info.UsersAndGroupsInWorkLocation.User)
	for _, key := range userData.MapKeys() {
		workLocationData := WorkLocationInUser{}
		workLocationData.CompanyId = companyTeamName
		workLocationData.DateOfCreation = m.Settings.DateOfCreation
		workLocationData.WorkLocationForTask = m.Info.WorkLocation
		workLocationData.StartDate =1503878400
		workLocationData.DailyStartDate = 1503907200
		workLocationData.DailyEndDate = 1503939600
		workLocationData.EndDate =1503964800
		workLocationData.Latitude ="9.7321201"
		workLocationData.Longitude ="76.35365479999996"
		workLocationData.Status = helpers.StatusPending
		userKey := key.String()
		err = db.Child("/Users/"+userKey+"/WorkLocation/"+workLocationUniqueID).Set(workLocationData)
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}
	}
	return  true
}


func GetAllWorkLocationDetails(ctx context.Context,CompanyTeamName string) (map[string]WorkLocation,bool){
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
func DeleteWorkLog(ctx context.Context,workLocationId string) (bool) {
	workLocationValues := WorkLocationSettings{}
	deleteWorkLocation := WorkLocationSettings{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/WorkLocation/"+workLocationId+"/Settings").Value(&workLocationValues)
	if err != nil {
		log.Fatal(err)
		return false
	}
	deleteWorkLocation.Status = helpers.UserStatusDeleted
	deleteWorkLocation.DateOfCreation = workLocationValues.DateOfCreation
	err = db.Child("/WorkLocation/"+workLocationId+"/Settings").Update(&deleteWorkLocation)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true




}


