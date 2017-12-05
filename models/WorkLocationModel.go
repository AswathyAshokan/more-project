package models
import (
	"log"
	"golang.org/x/net/context"

	"app/passporte/helpers"
	//"strings"
	"reflect"
	"github.com/kjk/betterguid"
	"time"
	"math/rand"
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
	NFCTagID			string
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
	Users			map[string]WorkLocationFitToWorkResponse

}
type WorkLocationFitToWorkResponse struct{
	ResponseTime		int64
	UserName		string
	UserResponse		string
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


func(m *WorkLocation) AddWorkLocationToDb(ctx context.Context,companyTeamName string,fitToWorksName string,WorkBreakSlice []string,TaskWorkTimeSlice []string,companyName string) (bool){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	workLocationUniqueID := betterguid.New()
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 3)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	generatedString :=string(result)
	log.Println("genertedstring",generatedString)
	newGeneratedKey:=workLocationUniqueID[0:len(workLocationUniqueID)-1]+generatedString
	log.Println("newly gener",newGeneratedKey)
	workLocationUniqueID =newGeneratedKey
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
			workLocationData.CompanyName = companyName
			workLocationData.Status = helpers.StatusPending
			userKey := key.String()
			if workLocationData.WorkLocationForTask !="" && workLocationData.CompanyName !=""{
				err = db.Child("/Users/" + userKey + "/WorkLocation/" + workLocationUniqueID).Set(workLocationData)
				if err != nil {
					log.Println("w16")
					log.Println("Insertion error:", err)
					return false
				}
			}
			notifyId := betterguid.New()
			userNotificationDetail :=WorkLocationNotification{}
			userNotificationDetail.Date =time.Now().Unix()
			userNotificationDetail.IsRead =false
			userNotificationDetail.IsViewed =false
			userNotificationDetail.WorkLocationId=workLocationUniqueID
			userNotificationDetail.Category ="WorkLocation"
			userNotificationDetail.Status ="New"
			userNotificationDetail.CompanyName = companyName
			userNotificationDetail.WorkLocation = m.Info.WorkLocation
			userNotificationDetail.IsDeleted =false
			if userNotificationDetail.WorkLocation!="" &&userNotificationDetail.CompanyName  !=""{
				err = db.Child("/Users/"+key.String()+"/Settings/Notifications/WorkLocationNotification/"+notifyId).Set(userNotificationDetail)


			}

			if err!=nil{
				log.Println("Insertion error:",err)
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
	WorkLocationInfoData.DailyStartDate = m.Info.DailyStartDate
	WorkLocationInfoData.EndDate = m.Info.EndDate
	WorkLocationInfoData.Latitude = m.Info.Latitude
	WorkLocationInfoData.Longitude = m.Info.Longitude
	WorkLocationInfoData.Latitude = m.Info.Latitude
	WorkLocationInfoData.StartDate = m.Info.StartDate
	WorkLocationInfoData.NFCTagID = m.Info.NFCTagID
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

func GetWorkLocationDataFromUsers(ctx context.Context,workLocationId string,usersId string)(WorkLocationInUser)  {
	usersWorkLocation := WorkLocationInUser{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("workLoaction id",workLocationId)
	err = db.Child("/Users/"+usersId+"/WorkLocation/"+workLocationId).Value(&usersWorkLocation)
	if err != nil {
		log.Fatal(err)
		return usersWorkLocation
	}
	log.Println("workLocation User data form modal",usersWorkLocation)
	return usersWorkLocation

}






func IsWorkAssignedToUser(ctx context.Context ,companyTeamName string )( map[string]WorkLocation)  {
	workLocationValues := map[string]WorkLocation{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("WorkLocation").Value(&workLocationValues)
	if err != nil {
		log.Println("error connection")
		//.log.Fatal(err)
		return workLocationValues
	}
	log.Println("workLocationValues",workLocationValues)
	return workLocationValues
}


func(m *WorkLocation)EditWorkLocationToDb(ctx context.Context,workLocationId string,companyTeamName string,fitToWorksName string,WorkBreakSlice []string,TaskWorkTimeSlice []string,companyName string) (bool){



	notifyUpdatedId := betterguid.New()
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 4)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	generatedString :=string(result)
	newGeneratedKey:=notifyUpdatedId[0:len(notifyUpdatedId)-2]+generatedString

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


		for j:=0;j<len(UserKeySlice);j++{
			log.Println("keySlice[i]",keySlice[i])
			log.Println("UserKeySlice[j]",UserKeySlice[j])
			if keySlice[i]==UserKeySlice[j] {
				err = db.Child("/Users/" + keySlice[i] + "/WorkLocation/" + workLocationId).Value(&OldUserWorkLocation)
				log.Println("kkkkkkkkkkk inside if")
				workLocationData.CompanyId = companyTeamName
				workLocationData.WorkLocationForTask = m.Info.WorkLocation
				workLocationData.StartDate = m.Info.StartDate
				workLocationData.EndDate = m.Info.EndDate
				workLocationData.DailyStartDate = m.Info.DailyStartDate
				workLocationData.DailyEndDate = m.Info.DailyEndDate
				workLocationData.Latitude = m.Info.Latitude
				workLocationData.Longitude = m.Info.Longitude
				workLocationData.DateOfCreation = m.Settings.DateOfCreation
				workLocationData.Status = OldUserWorkLocation.Status
				workLocationData.DateOfCreation = OldUserWorkLocation.DateOfCreation
				workLocationData.CompanyName = companyName
				if workLocationData.WorkLocationForTask !="" &&workLocationData.CompanyName !=""{
					err = db.Child("/Users/" + UserKeySlice[j] + "/WorkLocation/" + workLocationId).Set(workLocationData)

				}
				if err != nil {
					log.Println("Insertion error:", err)
					return false
				}

				userNotificationDetail := WorkLocationNotification{}
				userNotificationDetail.Date = time.Now().Unix()
				userNotificationDetail.IsRead = false
				userNotificationDetail.IsViewed = false
				userNotificationDetail.WorkLocationId = workLocationId
				userNotificationDetail.Category = "WorkLocation"
				userNotificationDetail.Status = "Updated"
				userNotificationDetail.WorkLocation = m.Info.WorkLocation
				userNotificationDetail.CompanyName = companyName
				userNotificationDetail.IsDeleted = false
				if userNotificationDetail.WorkLocation !="" &&userNotificationDetail.CompanyName !=""{
					err = db.Child("/Users/" + keySlice[i] + "/Settings/Notifications/WorkLocationNotification/" + newGeneratedKey).Set(userNotificationDetail)

				}
				if err != nil {
					log.Println("Insertion error:", err)
					return false
				}
				/*} else {
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
				workLocationData.CompanyName = companyName

			}
			err = db.Child("/Users/"+UserKeySlice[j]+"/WorkLocation/"+workLocationId).Set(workLocationData)
			if err!=nil{
				log.Println("Insertion error:",err)
				return false
			}*/
			}

		}

	}

	m.Settings.DateOfCreation = workLocationValues.DateOfCreation
	m.Settings.Status = workLocationValues.Status

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
		fitToWorkDetail := map[string]WorkLocationFitToWorkResponse{}
		userMap := make(map[string]WorkLocationFitToWorkResponse)
		fitStruct :=WorkLocationFitToWorkResponse{}
		fitWorkStructForResponse :=FitToWorkForkWorkLocation{}

		err = db.Child("WorkLocation/"+workLocationId+"/FitToWork/Settings").Set(FitToWorkForSetting)
		for _, eachKey := range tempKeySlice {
			log.Println(reflect.TypeOf(fitToWork[eachKey].FitToWorkName))
			log.Println(reflect.TypeOf(fitToWorksName))
			string1 :=fitToWork[eachKey].FitToWorkName
			string2 :=fitToWorksName
			if Compare(string1,string2) ==0 {
				fitToWOrkKey =eachKey
				err = db.Child("FitToWork/"+companyTeamName+"/"+eachKey+"/Instructions").Value(&instructionOfFitWork)
				err = db.Child("WorkLocation/"+workLocationId+"/FitToWork/FitToWorkInstruction").Set(instructionOfFitWork)
				err = db.Child("WorkLocation/"+workLocationId+"FitToWork/Users").Value(&fitToWorkDetail)
				fitToWorkResponse := reflect.ValueOf(fitToWorkDetail)
				for _, fitToWorkResponseKey := range fitToWorkResponse.MapKeys() {

					fitStruct.ResponseTime=fitToWorkDetail[fitToWorkResponseKey.String()].ResponseTime
					fitStruct.UserName =fitToWorkDetail[fitToWorkResponseKey.String()].UserName
					fitStruct.UserResponse =fitToWorkDetail[fitToWorkResponseKey.String()].UserResponse
					userMap[fitToWorkResponseKey.String()] =fitStruct
				}

				fitWorkStructForResponse.Users=userMap
				err = db.Child("WorkLocation/"+workLocationId+"FitToWork/Users/").Set(fitWorkStructForResponse.Users)
			}
		}
		FitToWorkForSetting.Status =helpers.StatusActive
		err = db.Child("WorkLocation/"+workLocationId+"/FitToWork/Settings").Set(FitToWorkForSetting)
		FitToWorkForInfo.FitToWorkName =fitToWorksName
		FitToWorkForInfo.FitToWorkId =fitToWOrkKey
		err = db.Child("WorkLocation/"+workLocationId+"/FitToWork/Info").Set(FitToWorkForInfo)
	}


	// for adding work break to database
	log.Println("op8")
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
			err = db.Child("WorkLocation/"+workLocationId+"/WorkExposure/").Set(ExposureMap)

		}
	}

	log.Println("op8")
	WorkLocationInfoData := WorkLocationInfo{}
	WorkLocationInfoData.LogTimeInMinutes = m.Info.LogTimeInMinutes
	WorkLocationInfoData.LoginType = m.Info.LoginType
	WorkLocationInfoData.CompanyTeamName = m.Info.CompanyTeamName
	WorkLocationInfoData.DailyEndDate = m.Info.DailyEndDate
	WorkLocationInfoData.DailyStartDate = m.Info.DailyStartDate
	WorkLocationInfoData.EndDate = m.Info.EndDate
	WorkLocationInfoData.Latitude = m.Info.Latitude
	WorkLocationInfoData.Longitude = m.Info.Longitude
	WorkLocationInfoData.Latitude = m.Info.Latitude
	WorkLocationInfoData.StartDate = m.Info.StartDate
	WorkLocationInfoData.NFCTagID = m.Info.NFCTagID
	WorkLocationInfoData.UsersAndGroupsInWorkLocation = m.Info.UsersAndGroupsInWorkLocation
	WorkLocationInfoData.WorkLocation = m.Info.WorkLocation
	err = db.Child("WorkLocation/"+workLocationId+"/Info").Set(WorkLocationInfoData)

	if err != nil {
		log.Println("w14")
		log.Println(err)
		return false
	}
	log.Println("op9")
	WorkLocationSettingsData := WorkLocationSettings{}
	WorkLocationSettingsData.DateOfCreation = m.Settings.DateOfCreation
	WorkLocationSettingsData.FitToWorkDisplayStatus = m.Settings.FitToWorkDisplayStatus
	WorkLocationSettingsData.Status = m.Settings.Status
	err = db.Child("WorkLocation/"+workLocationId+"/Settings").Set(WorkLocationSettingsData)
	log.Println("op10")
	if err != nil {
		log.Println(err)
		return false
	}
	//get newly selected users
	newUser := make([]string, 0)
	s_one :=  UserKeySlice
	s_two :=  keySlice
	for _, s := range s_one {
		if !inslice(s, s_two) {
			newUser = append(newUser, s)
		}
	}

	//get removed users

	var removedUsers  []string
	for _, s := range keySlice {
		if !inslice(s, UserKeySlice) {
			removedUsers = append(removedUsers, s)
		}
	}
	for i:=0; i<len(removedUsers);i++{
		log.Println("iam in removed section")
		userNotificationDetail :=WorkLocationNotification{}
		userNotificationDetail.Date =time.Now().Unix()
		userNotificationDetail.IsRead =false
		userNotificationDetail.IsViewed =false
		userNotificationDetail.WorkLocationId=workLocationId
		userNotificationDetail.Category ="WorkLocation"
		userNotificationDetail.Status ="Removed"
		userNotificationDetail.IsDeleted =false
		userNotificationDetail.WorkLocation =  m.Info.WorkLocation
		userNotificationDetail.CompanyName = companyName
		if userNotificationDetail.WorkLocation !="" &&userNotificationDetail.CompanyName !=""{
			err = db.Child("/Users/"+removedUsers[i]+"/Settings/Notifications/WorkLocationNotification/"+newGeneratedKey).Set(userNotificationDetail)

		}
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}



		workLocationData.CompanyId = companyTeamName
		workLocationData.WorkLocationForTask = m.Info.WorkLocation
		workLocationData.StartDate =m.Info.StartDate
		workLocationData.EndDate =m.Info.EndDate
		workLocationData.DailyStartDate = m.Info.DailyStartDate
		workLocationData.DailyEndDate = m.Info.DailyEndDate
		workLocationData.Latitude =m.Info.Latitude
		workLocationData.Longitude =m.Info.Longitude
		workLocationData.Status =helpers.StatusInActive
		workLocationData.DateOfCreation = m.Settings.DateOfCreation
		workLocationData.CompanyName = companyName
		if workLocationData.CompanyName !="" && workLocationData.WorkLocationForTask !=""{
			err = db.Child("/Users/"+removedUsers[i]+"/WorkLocation/"+workLocationId).Set(workLocationData)

		}

		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}
	}

	for i:=0; i<len(newUser);i++{
		log.Println("iam in new dded section")
		userNotificationDetail :=WorkLocationNotification{}
		userNotificationDetail.Date =time.Now().Unix()
		userNotificationDetail.IsRead =false
		userNotificationDetail.IsViewed =false
		userNotificationDetail.WorkLocationId=workLocationId
		userNotificationDetail.Category ="WorkLocation"
		userNotificationDetail.Status ="New"
		userNotificationDetail.WorkLocation =  m.Info.WorkLocation
		userNotificationDetail.CompanyName = companyName
		userNotificationDetail.IsDeleted =false
		if userNotificationDetail.WorkLocation !=""&&userNotificationDetail.CompanyName !=""{

			err = db.Child("/Users/"+newUser[i]+"/Settings/Notifications/WorkLocationNotification/"+newGeneratedKey).Set(userNotificationDetail)

		}
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}


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
		workLocationData.CompanyName = companyName
		if workLocationData.CompanyName !="" &&workLocationData.WorkLocationForTask !=""{
			err = db.Child("/Users/"+newUser[i]+"/WorkLocation/"+workLocationId).Set(workLocationData)

		}
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}
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


func GetWorkLocationBreakDetailById(ctx context.Context, workLocationId string)( bool,map[string]WorkExposure) {
	breakDetail := map[string]WorkExposure{}
	dB, err := GetFirebaseClient(ctx, "")
	err = dB.Child("WorkLocation/" + workLocationId + "/WorkExposure/").Value(&breakDetail)
	if err != nil {
		log.Fatal(err)
		return false, breakDetail
	}
	return true, breakDetail
}

//get the fit to work details

func GetFitToWorkDetailsResponseById(ctx context.Context, workLocationId string)( bool,map[string]WorkLocationFitToWorkResponse) {
	fitToWorkDetail := map[string]WorkLocationFitToWorkResponse{}
	log.Println("hiii1")
	dB, err := GetFirebaseClient(ctx, "")
	err = dB.Child("WorkLocation/" + workLocationId + "/FitToWork/Users").Value(&fitToWorkDetail)
	if err != nil {
		log.Fatal(err)
		return false, fitToWorkDetail
	}
	log.Println("llll",fitToWorkDetail)
	return true, fitToWorkDetail
}
func  WorkLocationDeleteStatusCheck(ctx context.Context, workId string,companyId string)(bool,int64) {
	usersOfWorkLocation := UsersAndGroupsInWork{}
	var condition =""
	var keySlice []string
	workDataFromUsers := WorkLocationInUser{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("WorkLocation/"+ workId+"/Info/UsersAndGroupsInWorkLocation").Value(&usersOfWorkLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("jjjjj",usersOfWorkLocation.User)
	userData := reflect.ValueOf(usersOfWorkLocation.User)
	for _, key := range userData.MapKeys() {
		if usersOfWorkLocation.User[key.String()].Status != helpers.UserStatusDeleted{
			keySlice= append(keySlice,key.String())
		}
	}
	for i := 0;i<len(keySlice);i++{
		err = dB.Child("/Users/"+keySlice[i]+"/WorkLocation/"+workId).Value(&workDataFromUsers)
		if workDataFromUsers.Status == "Open"{

			condition="true"
			break
		}else {
			condition="false"
		}
	}
	log.Println("conditoion array",condition)
	if condition == "true" {
		return true,workDataFromUsers.EndDate

	}
	return false,workDataFromUsers.EndDate
}