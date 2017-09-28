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
	FitToWork	FitToWorkForkWorkLocation
}
type WorkLocationInfo struct {
	LoginType        		string
	CompanyTeamName			string
	WorkLocation       		string
	Latitude			string
	Longitude			string
	StartDate			int64
	EndDate				int64
	DailyStartDate                  int64
	DailyEndDate			int64
	UsersAndGroupsInWorkLocation	UsersAndGroupsInWork
	LogTimeInMinutes 		int64
	WorkLocationExposure		map[string]WorkExposure
}
type WorkExposure struct {
	BreakDurationInMinutes  string
	BreakStartTimeInMinutes string
	Status                  string
	DateOfCreation          int64

}

type FitToWorkForkWorkLocation struct {
	FitToWorkInstruction    map[string]WorkLocationFitToWork
	Settings		WorkFitToWorkSettings
	Info			WOrkFitToWorkInfo

}

type WOrkFitToWorkInfo struct {
	FitToWorkName  		string
	FitToWorkId 		string
}
type  WorkFitToWorkSettings struct {

	Status			string

}
type WorkLocationFitToWork struct {
	Description    string
	Status         string
	DateOfCreation int64


}

type WorkLocationSettings struct {
	FitToWorkDisplayStatus	string
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


func(m *WorkLocation) AddWorkLocationToDb(ctx context.Context,companyTeamName string,fitToWorksName string,WorkBreakSlice []string,TaskWorkTimeSlice []string) (bool){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	workLocationUniqueID := betterguid.New()
	userData := reflect.ValueOf(m.Info.UsersAndGroupsInWorkLocation.User)
	for _, key := range userData.MapKeys() {
		if m.Info.UsersAndGroupsInWorkLocation.User[key.String()].Status!=helpers.UserStatusDeleted {

			workLocationData := WorkLocationInUser{}
			workLocationData.CompanyId = companyTeamName
			workLocationData.DateOfCreation = m.Settings.DateOfCreation
			workLocationData.WorkLocationForTask = m.Info.WorkLocation
			workLocationData.StartDate = m.Info.StartDate
			workLocationData.EndDate = m.Info.EndDate
			workLocationData.DailyStartDate = m.Info.DailyStartDate
			workLocationData.DailyEndDate = m.Info.DailyEndDate
			workLocationData.Latitude = m.Info.Latitude
			workLocationData.Longitude = m.Info.Longitude
			workLocationData.Status = helpers.StatusPending
			userKey := key.String()
			err = db.Child("/Users/" + userKey + "/WorkLocation/" + workLocationUniqueID).Set(workLocationData)
			if err != nil {
				log.Println("w16")
				log.Println("Insertion error:", err)
				return false
			}
		}
	}
	if len(fitToWorksName) !=0{
		FitToWorkForSetting :=WorkFitToWorkSettings{}
		FitToWorkForInfo  :=WOrkFitToWorkInfo{}
		var tempKeySlice []string
		var fitToWOrkKey string
		instructionOfFitWork :=map[string]WorkLocationFitToWork{}
		fitToWork :=map[string]FitToWork{}
		db,err :=GetFirebaseClient(ctx,"")
		if err!=nil{
			log.Println("Connection error:",err)
		}
		err = db.Child("FitToWork/"+ companyTeamName).Value(&fitToWork)
		fitToWorkDataValues := reflect.ValueOf(fitToWork)
		for _, fitToWorkKey := range fitToWorkDataValues.MapKeys() {
			tempKeySlice = append(tempKeySlice, fitToWorkKey.String())
		}
		FitToWorkForSetting.Status =helpers.StatusActive
		err = db.Child("WorkLocation/"+workLocationUniqueID+"/FitToWork/Settings").Set(FitToWorkForSetting)
		for _, eachKey := range tempKeySlice {
			log.Println(reflect.TypeOf(fitToWork[eachKey].FitToWorkName))
			log.Println(reflect.TypeOf(fitToWorksName))
			string1 :=fitToWork[eachKey].FitToWorkName
			string2 :=fitToWorksName
			if Compare(string1,string2) ==0 {
				fitToWOrkKey =eachKey
				err = db.Child("FitToWork/"+companyTeamName+"/"+eachKey+"/Instructions").Value(&instructionOfFitWork)
				err = db.Child("WorkLocation/"+workLocationUniqueID+"/FitToWork/FitToWorkInstruction").Set(instructionOfFitWork)

			}

		}
		FitToWorkForSetting.Status =helpers.StatusActive
		err = db.Child("WorkLocation/"+workLocationUniqueID+"/FitToWork/Settings").Set(FitToWorkForSetting)
		FitToWorkForInfo.FitToWorkName =fitToWorksName
		FitToWorkForInfo.FitToWorkId =fitToWOrkKey
		err = db.Child("WorkLocation/"+workLocationUniqueID+"/FitToWork/Info").Set(FitToWorkForInfo)
	}


	// for adding work break to database

	ExposureMap := make(map[string]WorkExposure)
	ExposureTask :=WorkExposure{}
	if WorkBreakSlice[0] !=""{

		for i := 0; i < len(WorkBreakSlice); i++ {

			ExposureTask.BreakDurationInMinutes =WorkBreakSlice[i]
			ExposureTask.BreakStartTimeInMinutes =TaskWorkTimeSlice[i]
			ExposureTask.DateOfCreation =time.Now().Unix()
			ExposureTask.Status = helpers.StatusActive
			id := betterguid.New()
			ExposureMap[id] = ExposureTask
			err = db.Child("WorkLocation/"+workLocationUniqueID+"/WorkExposure/").Set(ExposureMap)

		}
	}


	WorkLocationInfoData := WorkLocationInfo{}
	WorkLocationInfoData.LogTimeInMinutes = m.Info.LogTimeInMinutes
	WorkLocationInfoData.LoginType = m.Info.LoginType
	WorkLocationInfoData.CompanyTeamName = m.Info.CompanyTeamName
	WorkLocationInfoData.DailyEndDate = m.Info.DailyEndDate
	WorkLocationInfoData.DailyStartDate = m.Info.DailyEndDate
	WorkLocationInfoData.EndDate = m.Info.EndDate
	WorkLocationInfoData.Latitude = m.Info.Latitude
	WorkLocationInfoData.Longitude = m.Info.Longitude
	WorkLocationInfoData.Latitude = m.Info.Latitude
	WorkLocationInfoData.StartDate = m.Info.StartDate
	WorkLocationInfoData.UsersAndGroupsInWorkLocation = m.Info.UsersAndGroupsInWorkLocation
	WorkLocationInfoData.WorkLocation = m.Info.WorkLocation
	err = db.Child("WorkLocation/"+workLocationUniqueID+"/Info").Set(WorkLocationInfoData)

	if err != nil {
		log.Println("w14")
		log.Println(err)
		return false
	}
	WorkLocationSettingsData := WorkLocationSettings{}
	WorkLocationSettingsData.DateOfCreation = m.Settings.DateOfCreation
	WorkLocationSettingsData.FitToWorkDisplayStatus = m.Settings.FitToWorkDisplayStatus
	WorkLocationSettingsData.Status = m.Settings.Status
	err = db.Child("WorkLocation/"+workLocationUniqueID+"/Settings").Set(WorkLocationSettingsData)

	if err != nil {
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
		log.Println("error connection")
		//.log.Fatal(err)
		return workLocationValues
	}
	log.Println("workLocationValues",workLocationValues)
	return workLocationValues
}


func(m *WorkLocation)EditWorkLocationToDb(ctx context.Context,workLocationId string,companyTeamName string) (bool){
	workLocationValues := WorkLocationSettings{}
	//workLocationInfoData := WorkLocationInfo{}
	oldUserData  := UsersAndGroupsInWork{}
	var keySlice []string
	workLocationData := WorkLocationInUser{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}

	err = db.Child("/WorkLocation/"+workLocationId+"/Settings").Value(&workLocationValues)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("/WorkLocation/"+workLocationId+"/Info/UsersAndGroupsInWorkLocation").Value(&oldUserData)
	if err != nil {
		log.Fatal(err)
		return false
	}
	userValues :=reflect.ValueOf(oldUserData.User)
	for _,eachUserKey := range userValues.MapKeys() {
		keySlice = append(keySlice,eachUserKey.String())

	}
	var UserKeySlice []string
	OldUserWorkLocation :=  WorkLocationInUser{}
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
	m.Settings.DateOfCreation = workLocationValues.DateOfCreation
	m.Settings.Status = workLocationValues.Status
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




