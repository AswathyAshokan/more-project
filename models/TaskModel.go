/* Author :Aswathy Ashok */
package models

import (
	"log"
	"golang.org/x/net/context"
)

type Task   struct {

	Info           TaskInfo
	Contact        []TaskContact
	Customer       TaskCustomer
	Job            TaskJob
	UsersAndGroups UsersAndGroups
	Settings       TaskSetting

}
type TaskInfo struct {

	TaskName        string
	TaskLocation    string
	StartDate       string
	EndDate         string
	LoginType       string
	TaskDescription string
	UserNumber      string
	Log             string
	FitToWork       string
}
type TaskContact struct {
	ContactId 	string
	ContactName	string
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
func (m *Task) AddTaskToDB(ctx context.Context )(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	_, err = dB.Child("Task").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}

/*get all task details from DB*/
func (m *Task) RetrieveTaskFromDB(ctx context.Context)(bool,map[string]Task) {
	taskValue := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Task").Value(&taskValue)
	if err != nil {
		log.Fatal(err)
		return false, taskValue
	}
	log.Println(taskValue)
	return true, taskValue
}

/*delete  task details from DB*/
func (m *Task) DeleteTaskFromDB(ctx context.Context, taskId string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Task/"+ taskId).Remove()
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
	err = dB.Child("Job").Value(&jobValue)
	if err != nil {
		log.Fatal(err)
		return false, jobValue
	}
	log.Println("job value",jobValue)
	return true, jobValue
}

/*get all contact details from DB*/
func (m *Task) GetAllContact(ctx context.Context)(bool,map[string]Task) {
	contactValue := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Contacts").Value(&contactValue)
	if err != nil {
		log.Fatal(err)
		return false, contactValue
	}
	log.Println(contactValue)
	return true, contactValue


}

/* Function for update task on DB*/
func (m *Task) UpdateTaskToDB( ctx context.Context, taskId string)(bool)  {


	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Task/"+ taskId).Update(&m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true

}

/*get a specific task detail by id*/
func (m *Task) GetTaskDetailById(ctx context.Context, taskId string)(bool, Task) {
	taskDetail := Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Task/"+ taskId).Value(&taskDetail)
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

func(m *Task) GetContactDetailById(ctx context.Context, contactId string) (Task,bool){
	contactDetails := Task{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Task/"+ contactId).Value(&contactDetails)
	if err != nil {
		log.Fatal(err)
		return contactDetails, false
	}
	return contactDetails,true
}
