/* Author :Aswathy Ashok */
package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"
	"strings"
	"app/passporte/helpers"
)

type Tasks   struct {

	Info           TaskInfo
	Contacts       map[string]TaskContact
	Customer       TaskCustomer
	Job            TaskJob
	UsersAndGroups UsersAndGroups
	Settings       TaskSetting

}
type TaskInfo struct {

	TaskName        string
	TaskLocation    string
	StartDate       int64
	EndDate         int64
	LoginType       string
	TaskDescription string
	UserNumber      string
	Log             string
	FitToWork      	string
	CompanyTeamName	string
	Latitude	string
	Longitude	string
	StartTime	string
	EndTime		string
}
type TaskContact struct {
	ContactName	string
	PhoneNumber	string
	EmailId		string
}
type TaskCustomer struct{
	CustomerId	string
	CustomerName	string
}
type TaskJob struct {
	JobId		string
	JobName		string
}
type TaskUser struct {
	FullName	string
}
type TaskGroup struct{
	GroupName	string
	Members	 	map[string]GroupMemberName
}
type  GroupMemberName struct {
	MemberName	string
}
type TaskSetting struct {
	Status		string
	DateOfCreation	int64

}
type UsersAndGroups struct {
	User 		map[string]TaskUser
	Group 		map[string]TaskGroup

}

/*add task details to DB*/
func (m *Tasks) AddTaskToDB(ctx context.Context  ,companyId string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}

	taskData, err := dB.Child("Tasks").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}

	//For inserting task details to User
	taskDataString := strings.Split(taskData.String(),"/")
	taskUniqueID := taskDataString[len(taskDataString)-2]
	userData := reflect.ValueOf(m.UsersAndGroups.User)
	for _, key := range userData.MapKeys() {
		userTaskDetail := UserTasks{}
		userTaskDetail.DateOfCreation = m.Settings.DateOfCreation
		userTaskDetail.TaskName = m.Info.TaskName
		userTaskDetail.CustomerName = m.Customer.CustomerName
		userTaskDetail.EndDate = m.Info.EndDate
		userTaskDetail.StartDate =m.Info.StartDate
		userTaskDetail.JobName = m.Job.JobName
		userTaskDetail.Status = helpers.StatusPending
		userTaskDetail.CustomerName=m.Info.CompanyTeamName
		userTaskDetail.CompanyId = companyId

		userKey :=key.String()
		err = dB.Child("/Users/"+userKey+"/Tasks/"+taskUniqueID).Set(userTaskDetail)
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}

	}
	return true
}

/*get all task details from DB*/
func (m *Tasks) RetrieveTaskFromDB(ctx context.Context,companyTeamName string)(bool,map[string]Tasks) {
	taskValue := map[string]Tasks{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Tasks").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&taskValue)
	if err != nil {
		log.Fatal(err)
		return false, taskValue
	}
	log.Println(taskValue)
	return true, taskValue
}

/*delete  task details from DB*/
func (m *Tasks) DeleteTaskFromDB(ctx context.Context, taskId string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Tasks/"+ taskId).Remove()
	log.Println("deleted successfully")
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}

/*get all job details from DB*/
func GetAllJobs(ctx context.Context)(bool,map[string]Job) {
	jobValue := map[string]Job{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Jobs").Value(&jobValue)
	if err != nil {
		log.Fatal(err)
		return false, jobValue
	}
	return true, jobValue
}

/*get all contact details from DB*/
func (m *Tasks) GetAllContact(ctx context.Context)(bool,map[string]Tasks) {
	contactValue := map[string]Tasks{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Contacts").Value(&contactValue)
	if err != nil {
		log.Fatal(err)
		return false, contactValue
	}
	return true, contactValue


}

/* Function for update task on DB*/
func (m *Tasks) UpdateTaskToDB( ctx context.Context, taskId string)(bool)  {


	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Tasks/"+ taskId).Update(&m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true

}

/*get a specific task detail by id*/
func (m *Tasks) GetTaskDetailById(ctx context.Context, taskId string)(bool, Tasks) {
	taskDetail := Tasks{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Tasks/"+ taskId).Value(&taskDetail)
	if err != nil {
		log.Fatal(err)
		return false, taskDetail
	}
	return true, taskDetail

}

/*get user details from DB*/
func (m *UsersAndGroups ) GetAllUsers(ctx context.Context)(bool,map[string]UsersAndGroups) {
	valueOfUser := map[string]UsersAndGroups{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("User").Value(&valueOfUser)
	if err != nil {
		log.Fatal(err)
		return false,valueOfUser
	}

	return true,valueOfUser
}

func(m *Tasks) GetContactDetailById(ctx context.Context, contactId string) (Tasks,bool){
	contactDetails := Tasks{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Tasks/"+ contactId).Value(&contactDetails)
	if err != nil {
		log.Fatal(err)
		return contactDetails, false
	}
	return contactDetails,true
}
