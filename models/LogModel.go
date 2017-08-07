package models

import (
	"golang.org/x/net/context"
	"log"
	//"app/passporte/helpers"
	"reflect"
)

type WorkLog struct {
	Duration 	string
	Latitude	float64
	LogDescription	string
	LogTime		int64
	Longitude	float64
	Type		string
	UserID		string
	UserName 	string
	TaskID		string

}
type GeneralLog struct {
	LogDescription 		string
	Latitude		float64
	LogTime 		int64
	Longitude		float64
	Type 			string
	UserID			string
	UserName  		string
}
func (m *WorkLog)GetLogDetailOfUser(ctx context.Context,companyTeamName string)(bool,map[string]WorkLog) {
	workDetail := map[string]WorkLog{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("WorkLog/"+companyTeamName).Value(&workDetail)
	log.Println("worklog",workDetail)
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



func GetSpecificLogValues(ctx context.Context,userId string)(map[string]GeneralLog) {
	log.Println("id......",userId)
	generalLogData :=map[string]GeneralLog{}
	dB, err := GetFirebaseClient(ctx,"")
	//contactStatus := "Active";
	log.Println("model",userId)
	err = dB.Child("/GeneralLog/"+userId).Value(&generalLogData)
	if err != nil {
		log.Fatal(err)
		return generalLogData
	}
	return generalLogData

}