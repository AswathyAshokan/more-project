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
	Exposure		map[string]TaskExposure

}
type TaskFitToWork struct {
	Description    string
	Status         string
	DateOfCreation int64

}
type TaskExposure struct {
	BreakDurationInMinutes  string
	BreakStartTimeInMinutes string
	Status                  string
	DateOfCreation          int64

}
type TaskInfo struct {

	TaskName         string
	StartDate        int64
	EndDate          int64
	LoginType        string
	TaskDescription  string
	UserNumber       string
	LogTimeInMinutes int64
	TaskLocation     string
	CompanyTeamName  string
	CompanyName	string
	NFCTagID	string

}
type TaskContact struct {
	ContactName	string
	PhoneNumber	string
	EmailId		string
	ContactStatus	string
}
type TaskLocation struct{
	Latitude	string
	Longitude	string
}
type TaskCustomer struct{
	CustomerId	string
	CustomerName	string
	CustomerStatus	string
}
type TaskJob struct {
	JobId		string
	JobName		string
	JobStatus	string
}
type TaskUser struct {
	FullName	string
	Status		string
	UserTaskStatus	string
}
type TaskGroup struct{
	GroupName	string
	GroupStatus	string
	Members	 	map[string]GroupMemberName
}
type  GroupMemberName struct {
	MemberName	string

}
type TaskSetting struct {
	Status			string
	DateOfCreation		int64
	FitToWorkDisplayStatus	string
	TaskStatus		string
	CompletedPercentage	float32
	PendingPercentage	float32

}
type UsersAndGroups struct {
	User 		map[string]TaskUser
	Group 		map[string]TaskGroup

}

/*add task details to DB*/
func (m *Tasks) AddTaskToDB(ctx context.Context  ,companyId string ,FitToWorkSlice []string,WorkBreakSlice []string,TaskWorkTimeSlice []string, ContactId []string,GroupId []string,JobId string,CustomerId string)(bool)  {

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
	if FitToWorkSlice[0] !=""{

		for i := 0; i < len(FitToWorkSlice); i++ {

			fitToWorkForTask.Description =FitToWorkSlice[i]
			fitToWorkForTask.DateOfCreation =time.Now().Unix()
			fitToWorkForTask.Status = helpers.StatusActive
			id := betterguid.New()
			fitToWorkMap[id] = fitToWorkForTask
			err = dB.Child("/Tasks/"+taskUniqueID+"/FitToWork/").Set(fitToWorkMap)

		}
	}
	// for adding work break to database

	ExposureMap := make(map[string]TaskExposure)
	ExposureTask :=TaskExposure{}
	if WorkBreakSlice[0] !=""{

		for i := 0; i < len(WorkBreakSlice); i++ {

			ExposureTask.BreakDurationInMinutes =WorkBreakSlice[i]
			ExposureTask.BreakStartTimeInMinutes =TaskWorkTimeSlice[i]
			ExposureTask.DateOfCreation =time.Now().Unix()
			ExposureTask.Status = helpers.StatusActive
			id := betterguid.New()
			ExposureMap[id] = ExposureTask
			err = dB.Child("/Tasks/"+taskUniqueID+"/WorkExposure/").Set(ExposureMap)

		}
	}

	//...........................................................
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
	//setting task Id to company
	TaskIdForCompany :=TaskIdInfo{}

	TaskIdForCompany.DateOfCreation =m.Settings.DateOfCreation
	TaskIdForCompany.FitToWorkDisplayStatus =m.Settings.FitToWorkDisplayStatus
	TaskIdForCompany.Status =m.Settings.Status
	TaskIdForCompany.TaskStatus =m.Settings.TaskStatus
	err = dB.Child("/Company/"+companyId+"/Tasks/"+taskUniqueID).Set(TaskIdForCompany)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}

	//setting task id to contact
	ContactTask :=TasksContact{}
	ContactTask.TaskContactStatus =helpers.StatusActive
	 for i:=0;i<len(ContactId);i++{
		 err = dB.Child("/Contacts/"+ ContactId[i]+"/Tasks/"+taskUniqueID).Set(ContactTask)

	 }
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	//setting task id to customer
	CustomerTask :=TasksCustomer{}
	CustomerTask.TasksCustomerStatus =helpers.StatusActive
	job := Job{}
	err = dB.Child("/Jobs/"+ JobId).Value(&job)
	CustomerIdForTask :=job.Customer.CustomerId
	customerInTask :=TaskCustomer{}
	customerInTask.CustomerId =CustomerIdForTask
	log.Println("customer id",CustomerIdForTask)
	customerInTask.CustomerName =m.Customer.CustomerName
	customerInTask.CustomerStatus =m.Customer.CustomerStatus
	err = dB.Child("/Tasks/"+ taskUniqueID+"/Customer/").Set(customerInTask)
	//log.Println(customerInTask)
	//err = dB.Child("/Customers/"+ CustomerIdForTask+"/Tasks/"+taskUniqueID).Set(CustomerTask)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	//setting task id to Job
	JobTask :=TasksJob{}
	JobTask.TasksJobStatus =helpers.StatusActive
	err = dB.Child("/Jobs/"+ JobId+"/Tasks/"+taskUniqueID).Set(JobTask)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	//setting task id to Group
	GroupTask :=TasksGroup{}
	GroupTask.TasksGroupStatus =helpers.StatusActive
	log.Println("dsfsjdfh",GroupId)
	for i:=0;i<len(GroupId);i++{
		log.Println("inside group add")
		err = dB.Child("/Group/"+ GroupId[i] +"/Tasks/"+taskUniqueID).Set(GroupTask)

	}


	//setting number of task in job
	jobDetail := map[string]Job {}
	updatedJob :=Job{}
	err = dB.Child("Jobs").OrderBy("Info/CompanyTeamName").EqualTo(companyId).Value(&jobDetail)
	jobData := reflect.ValueOf(jobDetail)
	for _, key := range jobData.MapKeys() {
		if key.String() ==m.Job.JobId  {
			NumberOfTask :=jobDetail[key.String()].Info.NumberOfTask
			NumberOfTask =NumberOfTask+1
			updatedJob.Info.JobName =jobDetail[key.String()].Info.JobName
			updatedJob.Info.JobNumber = jobDetail[key.String()].Info.JobNumber
			updatedJob.Info.NumberOfTask = NumberOfTask
			updatedJob.Info.CompanyTeamName = companyId
			updatedJob.Customer.CustomerId= jobDetail[key.String()].Customer.CustomerId
			updatedJob.Customer.CustomerName= jobDetail[key.String()].Customer.CustomerName
			updatedJob.Customer.CustomerStatus= jobDetail[key.String()].Customer.CustomerStatus
			updatedJob.Settings.Status = jobDetail[key.String()].Settings.Status
			updatedJob.Settings.DateOfCreation = jobDetail[key.String()].Settings.DateOfCreation
			err = dB.Child("/Jobs/"+ key.String()).Update(&updatedJob)

		}

	}

	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}

	return true
}

/*get all task details from DB*/
func (m *Tasks) RetrieveTaskFromDB(ctx context.Context,companyTeamName string)(bool,map[string]Tasks) {
	taskValue := map[string]Tasks{}
	dB, err := GetFirebaseClient(ctx,"")
	//taskStatus := "Active"
	err = dB.Child("Tasks").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&taskValue)
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
func GetAllJobs(ctx context.Context,companyTeamName string)(bool,map[string]Job) {
	jobValue := map[string]Job{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Jobs").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&jobValue)
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
func (m *Tasks) UpdateTaskToDB( ctx context.Context, taskId string , companyId string,FitToWorkSlice []string,WorkBreakSlice []string,TaskWorkTimeSlice []string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	taskValues :=Tasks{}
	err = dB.Child("/Tasks/"+ taskId).Value(&taskValues)
	m.Settings.TaskStatus=taskValues.Settings.TaskStatus
	m.Settings.DateOfCreation =taskValues.Settings.DateOfCreation
	m.Settings.FitToWorkDisplayStatus =taskValues.Settings.FitToWorkDisplayStatus
	m.Settings.Status =taskValues.Settings.Status
	m.Settings.CompletedPercentage =taskValues.Settings.CompletedPercentage
	m.Settings.PendingPercentage =taskValues.Settings.PendingPercentage
	m.Settings.Status =taskValues.Settings.Status
	m.Customer.CustomerStatus =taskValues.Customer.CustomerStatus
	m.Job.JobStatus = taskValues.Job.JobStatus

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
	ExposureMap := make(map[string]TaskExposure)
	ExposureTask :=TaskExposure{}
	if WorkBreakSlice[0] !=""{

		for i := 0; i < len(WorkBreakSlice); i++ {

			ExposureTask.BreakDurationInMinutes =WorkBreakSlice[i]
			ExposureTask.BreakStartTimeInMinutes = TaskWorkTimeSlice[i]
			ExposureTask.DateOfCreation =time.Now().Unix()
			ExposureTask.Status = helpers.StatusActive
			id := betterguid.New()
			ExposureMap[id] = ExposureTask
			err = dB.Child("/Tasks/"+taskId+"/WorkExposure/").Set(ExposureMap)

		}
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
	CustomerTask :=TasksCustomer{}
	CustomerTask.TasksCustomerStatus =helpers.StatusActive
	job := Job{}
	JobId := m.Job.JobId
	err = dB.Child("/Jobs/"+ JobId).Value(&job)
	CustomerIdForTask :=job.Customer.CustomerId
	customerInTask :=TaskCustomer{}
	customerInTask.CustomerId =CustomerIdForTask
	log.Println("customer id",CustomerIdForTask)
	customerInTask.CustomerName =m.Customer.CustomerName
	customerInTask.CustomerStatus =m.Customer.CustomerStatus
	err = dB.Child("/Tasks/"+ taskId+"/Customer/").Set(customerInTask)

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
	err = dB.Child("Tasks/"+ taskId).Value(&taskDetail)
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
	err = dB.Child("Users").Value(&valueOfUser)
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
func (m *Users) GetActiveUsersEmailForDropDown(ctx context.Context,userKeys string,email string,companyTeamName string)bool {

	var keySlice []string

	invitationData := map[string]CompanyInvitations{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Company/" + companyTeamName + "/Invitation" ).Value(&invitationData)
	dataValue := reflect.ValueOf(invitationData)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, key := range keySlice {
		if invitationData[key].Email ==email && invitationData[key].Status =="Active" &&invitationData[key].UserResponse=="Accepted"{

			return true
			break
		}
	}

	if err != nil {

		log.Fatal(err)
		return false
	}
	return false
}
//function to get  break time deatil
func (m *TaskExposure) GetTaskWorkBreakDetailById(ctx context.Context, taskId string)(bool, map[string] TaskExposure) {
	taskBreakDetail :=map[string] TaskExposure{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Tasks/"+ taskId+"/WorkExposure/").Value(&taskBreakDetail)
	if err != nil {
		log.Fatal(err)
		return false, taskBreakDetail
	}
	return true, taskBreakDetail

}