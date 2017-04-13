/* Author :Aswathy Ashok */
package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"
	"strings"
	"app/passporte/helpers"
	"time"
	"github.com/kjk/betterguid"

)

type Tasks   struct {

	Info           		TaskInfo
	Location		TaskLocation
	Contacts      	 	map[string]TaskContact
	Customer       		TaskCustomer
	Job           	 	TaskJob
	UsersAndGroups 		UsersAndGroups
	Settings       		TaskSetting
	FitToWork		map[string]TaskFitToWork

}
type TaskFitToWork struct {
	Description    string
	Status         string
	DateOfCreation int64

}
type TaskInfo struct {

	TaskName        string
	StartDate       int64
	EndDate         int64
	LoginType       string
	TaskDescription string
	UserNumber      string
	Log             string
	TaskLocation	string
	CompanyTeamName	string

}
type TaskContact struct {
	ContactName	string
	PhoneNumber	string
	EmailId		string
}
type TaskLocation struct{
	Latitude	string
	Longitude	string
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
func (m *Tasks) AddTaskToDB(ctx context.Context  ,companyId string ,FitToWorkSlice []string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}

	/*for i :=0; i<len(m.FitToWork.Info); i++ {


	}*/
	taskData, err := dB.Child("Tasks").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	//For inserting task details to User
	taskDataString := strings.Split(taskData.String(),"/")
	taskUniqueID := taskDataString[len(taskDataString)-2]
	//for adding fit to work to database
	fitToWorkMap := make(map[string]TaskFitToWork)
	fitToWorkForTask :=TaskFitToWork{}
	for i := 0; i < len(FitToWorkSlice); i++ {

		fitToWorkForTask.Description =FitToWorkSlice[i]
		fitToWorkForTask.DateOfCreation =time.Now().Unix()
		fitToWorkForTask.Status = helpers.StatusActive
		id := betterguid.New()
		fitToWorkMap[id] = fitToWorkForTask
		err = dB.Child("/Tasks/"+taskUniqueID+"/FitToWork/").Set(fitToWorkMap)

	}

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
	taskStatus := "Active"
	err = dB.Child("Tasks").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).OrderBy("Settings/Status").EqualTo(taskStatus).Value(&taskValue)
	if err != nil {
		log.Fatal(err)
		return false, taskValue
	}
	return true, taskValue
}

/*delete  task details from DB*/
func (m *Tasks) DeleteTaskFromDB(ctx context.Context, taskId string,companyId string)(bool)  {

	 taskUpdate := TaskSetting{}
	 taskDeletion :=TaskSetting{}
	taskDetailForUser :=Tasks{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	taskDeletion.Status =helpers.StatusInActive
	err = dB.Child("/Tasks/"+ taskId+"/Settings").Value(&taskUpdate)
	taskDeletion.DateOfCreation =taskUpdate.DateOfCreation
	err = dB.Child("/Tasks/"+ taskId+"/Settings").Update(&taskDeletion)
	err = dB.Child("/Tasks/"+ taskId).Value(&taskDetailForUser)
	userData := reflect.ValueOf(taskDetailForUser.UsersAndGroups.User)
	for _, key := range userData.MapKeys() {
		userTaskDetail := UserTasks{}
		userTaskDetail.DateOfCreation = taskDetailForUser.Settings.DateOfCreation
		userTaskDetail.TaskName = taskDetailForUser.Info.TaskName
		userTaskDetail.CustomerName = taskDetailForUser.Customer.CustomerName
		userTaskDetail.EndDate = taskDetailForUser.Info.EndDate
		userTaskDetail.StartDate = taskDetailForUser.Info.StartDate
		userTaskDetail.JobName = taskDetailForUser.Job.JobName
		userTaskDetail.Status = helpers.StatusInActive
		userTaskDetail.CompanyId = companyId
		userKey := key.String()
		err = dB.Child("/Users/" + userKey + "/Tasks/" + taskId).Update(&userTaskDetail)
		if err!=nil{
			log.Println("Deletion error:",err)
		}
	}
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
func (m *Tasks) UpdateTaskToDB( ctx context.Context, taskId string , companyId string,FitToWorkSlice []string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Tasks/"+ taskId).Update(&m)
	//for adding fit to work to database
	fitToWorkMap := make(map[string]TaskFitToWork)
	fitToWorkForTask :=TaskFitToWork{}
	for i := 0; i < len(FitToWorkSlice); i++ {

		fitToWorkForTask.Description =FitToWorkSlice[i]
		fitToWorkForTask.DateOfCreation =time.Now().Unix()
		fitToWorkForTask.Status = helpers.StatusActive
		id := betterguid.New()
		fitToWorkMap[id] = fitToWorkForTask
		err = dB.Child("/Tasks/"+taskId+"/FitToWork/").Set(fitToWorkMap)

	}
	userData := reflect.ValueOf(m.UsersAndGroups.User)
	userTaskDetail := UserTasks{}
	for _, key := range userData.MapKeys() {
		userKey :=key.String()
		userTaskDetail.CompanyId = companyId
		userTaskDetail.CustomerName = m.Customer.CustomerName
		userTaskDetail.EndDate = m.Info.EndDate
		userTaskDetail.JobName = m.Job.JobName
		userTaskDetail.TaskName = m.Info.TaskName
		userTaskDetail.StartDate = m.Info.StartDate
		userTaskDetail.Status =helpers.StatusPending
		err = dB.Child("/Users/"+userKey+"/Tasks/"+taskId).Update(&userTaskDetail)
	}

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
//getting users from company
func (m *Company) GetUsersForDropdownFromCompany(ctx context.Context,companyTeamName string)(bool,map[string]Company) {
	companyUsers := map[string]Company{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Company").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&companyUsers)
	if err != nil {
		log.Fatal(err)
		return false, companyUsers
	}
	return true, companyUsers
}