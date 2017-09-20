package models
import (
	"log"
	"golang.org/x/net/context"

	"app/passporte/helpers"
	//"strings"
	"reflect"
	"github.com/kjk/betterguid"
	"time"

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
	DailyStartDate                  int64
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
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	log.Println("db",db)
	startDateInt := m.Info.StartDate
	endDateInt := m.Info.EndDate
	starta := time.Unix(startDateInt, 0).UTC()
	log.Println("utc time in golang",starta)
	enda := time.Unix(endDateInt, 0).UTC()
	log.Println("utc enda",enda)
	workLocationUniqueID := betterguid.New()
	userData := reflect.ValueOf(m.Info.UsersAndGroupsInWorkLocation.User)
	for _, key := range userData.MapKeys() {
		log.Println("w15")
		if m.Info.UsersAndGroupsInWorkLocation.User[key.String()].Status!=helpers.UserStatusDeleted{

		}

		workLocationData := WorkLocationInUser{}
		workLocationData.CompanyId = companyTeamName
		workLocationData.DateOfCreation = m.Settings.DateOfCreation
		workLocationData.WorkLocationForTask = m.Info.WorkLocation
		workLocationData.StartDate =m.Info.StartDate
		workLocationData.EndDate =m.Info.EndDate
		workLocationData.DailyStartDate = m.Info.DailyStartDate
		workLocationData.DailyEndDate = m.Info.DailyEndDate
		workLocationData.Latitude =m.Info.Latitude
		workLocationData.Longitude =m.Info.Longitude
		workLocationData.Status = helpers.StatusPending
		userKey := key.String()
		err = db.Child("/Users/"+userKey+"/WorkLocation/"+workLocationUniqueID).Set(workLocationData)
		if err!=nil{
			log.Println("w16")
			log.Println("Insertion error:",err)
			return false
		}
	}
	err = db.Child("WorkLocation/"+workLocationUniqueID).Set(m)

	if err != nil {
		log.Println("w14")
		log.Println(err)
		return false
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


func IsWorkAssignedToUser(ctx context.Context ,companyTeamName string )( map[string]WorkLocation)  {
	workLocationValues := map[string]WorkLocation{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("WorkLocation").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&workLocationValues)
	if err != nil {
		//.log.Fatal(err)
		return workLocationValues
	}
	return workLocationValues
}


func(m *WorkLocation)EditWorkLocationToDb(ctx context.Context,workLocationId string,companyTeamName string) (bool){
	workLocationValues := WorkLocationSettings{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}

	err = db.Child("/WorkLocation/"+workLocationId+"/Settings").Value(&workLocationValues)
	if err != nil {
		log.Fatal(err)
		return false
	}
	var keySlice []string
	var UserKeySlice []string
	OldUserWorkLocation :=  WorkLocationInUser{}
	workLocationForUpdation :=map[string]Users{}
	err = db.Child("Users").Value(&workLocationForUpdation)
	workLocationData := WorkLocationInUser{}
	userValues :=reflect.ValueOf(workLocationForUpdation)
	for _,eachUserKey := range userValues.MapKeys() {
		keySlice = append(keySlice,eachUserKey.String())

	}
	userDataForEditing := reflect.ValueOf(m.Info.UsersAndGroupsInWorkLocation.User)
	for _, key := range userDataForEditing.MapKeys() {
		UserKeySlice = append(UserKeySlice, key.String())
	}
	for i:=0;i<len(keySlice);i++{
		err = db.Child("/Users/" + keySlice[i] + "/WorkLocation/" + workLocationId).Value(&OldUserWorkLocation)
		for j:=0;j<len(UserKeySlice);j++{
			if keySlice[i]==UserKeySlice[j]{
				workLocationData.CompanyId = companyTeamName
				workLocationData.WorkLocationForTask = m.Info.WorkLocation
				workLocationData.StartDate =m.Info.StartDate
				workLocationData.EndDate =m.Info.EndDate
				workLocationData.DailyStartDate = m.Info.DailyStartDate
				workLocationData.DailyEndDate = m.Info.DailyEndDate
				workLocationData.Latitude =m.Info.Latitude
				workLocationData.Longitude =m.Info.Longitude
				workLocationData.DateOfCreation = m.Settings.DateOfCreation
				workLocationData.Status = OldUserWorkLocation.Status
				workLocationData.DateOfCreation = OldUserWorkLocation.DateOfCreation


			} else {
				workLocationData.CompanyId = companyTeamName
				workLocationData.WorkLocationForTask = m.Info.WorkLocation
				workLocationData.StartDate =m.Info.StartDate
				workLocationData.EndDate =m.Info.EndDate
				workLocationData.DailyStartDate = m.Info.DailyStartDate
				workLocationData.DailyEndDate = m.Info.DailyEndDate
				workLocationData.Latitude =m.Info.Latitude
				workLocationData.Longitude =m.Info.Longitude
				workLocationData.Status = helpers.StatusPending
				workLocationData.DateOfCreation = m.Settings.DateOfCreation

			}
			err = db.Child("/Users/"+UserKeySlice[j]+"/WorkLocation/"+workLocationId).Set(workLocationData)
			if err!=nil{
				log.Println("Insertion error:",err)
				return false
			}

		}

	}

	log.Println("worklocation values in model oid users",OldUserWorkLocation)
	m.Settings.DateOfCreation = workLocationValues.DateOfCreation
	m.Settings.Status = workLocationValues.Status

	userData := reflect.ValueOf(m.Info.UsersAndGroupsInWorkLocation.User)
	log.Println("edited values",m)
	for _, key := range userData.MapKeys() {
		log.Println("idddd",key)
	}
	err = db.Child("/WorkLocation/"+workLocationId).Set(m)

	if err != nil {
		log.Println(err)
		return false
	}


	return  true
}
func DeleteWorkLog(ctx context.Context,workLocationId string) (bool) {
	workLocationValues := WorkLocationSettings{}
	deleteWorkLocation := WorkLocationSettings{}
	userDetails := map[string]Users{}
	workLocationFromUser := map[string]WorkLocationInUser{}
	WorkLocationOfSpecifiedUser :=WorkLocationInUser{}
	UpdatewOrKLocationFromUser := WorkLocationInUser{}

	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/WorkLocation/"+workLocationId+"/Settings").Value(&workLocationValues)
	if err != nil {
		log.Fatal(err)
		return false
	}
	deleteWorkLocation.Status = helpers.StatusInActive
	deleteWorkLocation.DateOfCreation = workLocationValues.DateOfCreation
	err = db.Child("/WorkLocation/"+workLocationId+"/Settings").Update(&deleteWorkLocation)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("Users").Value(&userDetails)
	dataValue := reflect.ValueOf(userDetails)
	for _, key := range dataValue.MapKeys() {
		log.Println("key",key.String())
		err = db.Child("Users/"+key.String()+"/WorkLocation/").Value(&workLocationFromUser)
		//log.Println("workLocationFromUser",workLocationFromUser)
		dataValueOfUserWorkLocation := reflect.ValueOf(workLocationFromUser)
		for _, k := range dataValueOfUserWorkLocation.MapKeys() {
			log.Println("keys of work log",k)
			if k.String() == workLocationId{
				err = db.Child("Users/"+key.String()+"/WorkLocation/"+k.String()).Value(&WorkLocationOfSpecifiedUser)
				UpdatewOrKLocationFromUser.EndDate =WorkLocationOfSpecifiedUser.EndDate
				UpdatewOrKLocationFromUser.StartDate=WorkLocationOfSpecifiedUser.StartDate
				UpdatewOrKLocationFromUser.Status =helpers.StatusInActive
				UpdatewOrKLocationFromUser.CompanyId =WorkLocationOfSpecifiedUser.CompanyId
				UpdatewOrKLocationFromUser.DailyEndDate =WorkLocationOfSpecifiedUser.DailyEndDate
				UpdatewOrKLocationFromUser.DailyStartDate =WorkLocationOfSpecifiedUser.DailyStartDate
				UpdatewOrKLocationFromUser.Latitude =WorkLocationOfSpecifiedUser.Latitude
				UpdatewOrKLocationFromUser.Longitude =WorkLocationOfSpecifiedUser.Longitude
				UpdatewOrKLocationFromUser.WorkLocationForTask =WorkLocationOfSpecifiedUser.WorkLocationForTask
				err = db.Child("Users/"+key.String()+"/WorkLocation/"+k.String()).Update(&UpdatewOrKLocationFromUser)

			}
			log.Println("WorkLocationOfSpecifiedUser",WorkLocationOfSpecifiedUser)

		}

	}

	return true

}




