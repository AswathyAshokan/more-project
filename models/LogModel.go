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
	Longitude		string
	LogTime 		string
	LongitudeVal		string
	Type 			string
	UsedId			string
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


func (m *GeneralLog)GetGeneralLogDataByUserId(ctx context.Context,userId string) {
	log.Println("id......",userId)
	generalLogData :=map[string]GeneralLog{}
	dB, err := GetFirebaseClient(ctx,"")
	//contactStatus := "Active";
	log.Println("model",userId)
	err = dB.Child("GeneralLog/"+userId).Value(&generalLogData)
	log.Println("GeneralLog",generalLogData)
	log.Println("err",err)
	/*log.Println("GeneralLog",generalLogData)
	if err != nil {
		log.Fatal(err)
	}*/

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
	//log.Println("AllTask",AllTask)
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