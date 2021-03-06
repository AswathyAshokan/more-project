package models

import (
	"golang.org/x/net/context"
	"log"
	//"app/passporte/helpers"
	"reflect"
)

type WorkLog struct {

	Category    	string
	Date        	string
	WorkId        	string
	Duration     	string
	Latitude    	float64
	LogDescription  string
	LogTime        	int64
	Longitude    	float64
	Type        	string
	UserID       	string
	UserName     	string
	TaskID        	string

}
type GeneralLog struct {
	LastLogin               LastLoginInfo
	LastLogout         	LastLogoutInfo
	LoginLogs               PreviousLoginInfo
	LogoutLogs        	PreviousLogOutInfo
	UserTrackLogs		UserTrackInfo

}

type UserTrackInfo struct {
	Latitude		float64
	LogDescription		string
	LogTime			int64
	Longitude		float64
	Type			string
	UserID			string
	UserName		string
}


type LastLoginInfo struct {
	LogDescription          string
	Latitude        	float64
	Duration         	string
	LogTime         	int64
	Longitude       	float64
	Type             	string
	UserID            	string
	UserName          	string

}
type LastLogoutInfo struct {
	LogDescription          string
	Latitude        	float64
	LogTime         	int64
	Duration         	string
	Longitude        	float64
	Type             	string
	UserID            	string
	UserName          	string

}
type PreviousLoginInfo struct {
	LogDescription          string
	Latitude        	float64
	LogTime         	int64
	Longitude        	float64
	Type             	string
	UserID            	string
	UserName         	string

}
type PreviousLogOutInfo struct {
	LogDescription          string
	Latitude        	float64
	LogTime         	int64
	Longitude        	float64
	Type             	string
	UserID           	string
	UserName          	string

}

func (m *WorkLog)GetLogDetailOfUser(ctx context.Context,companyTeamName string)(bool,map[string]WorkLog) {
	workDetail := map[string]WorkLog{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("WorkLog/"+companyTeamName).Value(&workDetail)
	log.Println("worklog deatils",workDetail)
	if err != nil {
		log.Fatal(err)
		return false, workDetail
	}
	return true, workDetail

}


func GetGeneralLogDataByUserId(ctx context.Context)(bool,map[string]GeneralLog) {
	generalLogData :=map[string]GeneralLog{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("GeneralLog").Value(&generalLogData)
	if err != nil {
		log.Fatal(err)
		return false,generalLogData
	}
	return true,generalLogData

}

func GetUserTrackLogValues(ctx context.Context,userId string)(map[string]UserTrackInfo){

	userTrackMap := map[string]UserTrackInfo{}
	dB, err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
	}
	err = dB.Child("/GeneralLog/"+userId+"/UserTrackLogs").Value(&userTrackMap)
	if err != nil {
		log.Fatal(err)
		return userTrackMap
	}
	return userTrackMap

}

func GetSpecificUserTrackDetails(ctx context.Context,userId string,trackKey string) (UserTrackInfo) {
	userTrackSrtuct := UserTrackInfo{}
	dB, err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
	}
	err = dB.Child("/GeneralLog/"+userId+"/UserTrackLogs/"+trackKey).Value(&userTrackSrtuct)
	if err != nil {
		log.Fatal(err)
		return userTrackSrtuct
	}
	return userTrackSrtuct
}

func GetCompanyUsers(ctx context.Context,companyId string) []string {
	CompanyUsers :=map[string]CompanyUsers{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Company/"+companyId+"/Users/").Value(&CompanyUsers)
	if err != nil {
		log.Fatal(err)
	}
	var keySlice []string
	dataValue := reflect.ValueOf(CompanyUsers)
	//to store the keys of slice
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	return keySlice

}


func GetTaskDataById(ctx context.Context,taskId string)(string,string) {
	AllTask := map[string]Tasks{}
	var taskName string
	var jobName string
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println(err)
	}
	err = dB.Child("Tasks").Value(&AllTask)
	if err != nil {
		log.Fatal(err)
		//return false
	}
	dataValue := reflect.ValueOf(AllTask)
	for _, key := range dataValue.MapKeys() {
		//log.Println("key.String()",key.String())
		log.Println("taskId",taskId)
		if key.String() == taskId{
			log.Println("iam here")
			taskName = AllTask[key.String()].Info.TaskName
			jobName = AllTask[key.String()].Job.JobName
			return taskName,jobName

		}


	}
	return taskName,jobName
}




func GetSpecificLogValues(ctx context.Context,userId string)(LastLoginInfo) {
	log.Println("id......",userId)
	lastLoginData :=LastLoginInfo{}
	dB, err := GetFirebaseClient(ctx,"")
	//contactStatus := "Active";
	log.Println("model",userId)
	err = dB.Child("/GeneralLog/"+userId+"/LastLogin").Value(&lastLoginData)
	if err != nil {
		log.Fatal(err)
		return lastLoginData
	}
	return lastLoginData

}






func GetSpecificLogoutValues(ctx context.Context,userId string)(LastLogoutInfo) {
	log.Println("id......",userId)
	lastLogoutData :=LastLogoutInfo{}
	dB, err := GetFirebaseClient(ctx,"")
	//contactStatus := "Active";
	log.Println("model",userId)
	err = dB.Child("/GeneralLog/"+userId+"/LastLogout").Value(&lastLogoutData)
	if err != nil {
		log.Fatal(err)
		return lastLogoutData
	}
	return lastLogoutData

}