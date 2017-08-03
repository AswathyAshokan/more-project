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
func (m *WorkLog)GetLogDetailOfUser(ctx context.Context,companyTeamName string)(bool,map[string]WorkLog) {
	workDetail := map[string]WorkLog{}
	dB, err := GetFirebaseClient(ctx,"")
	//contactStatus := "Active";
	log.Println("model",companyTeamName)
	err = dB.Child("WorkLog/"+companyTeamName).Value(&workDetail)
	if err != nil {
		log.Fatal(err)
		return false, workDetail
	}
	log.Println("work",workDetail)
	return true, workDetail

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
	log.Println("AllTask",AllTask)
	dataValue := reflect.ValueOf(AllTask)
	for _, key := range dataValue.MapKeys() {
		if key.String() == taskId{
			taskName = AllTask[key.String()].Info.TaskName
			jobName = AllTask[key.String()].Job.JobName
			return taskName,jobName

		}


	}
	return taskName,jobName
}