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
	FitToWork		 FitToWorkForTask
	Exposure		map[string]TaskExposure

}
type FitToWorkForTask struct {
	FitToWorkInstruction  map[string]TaskFitToWork
	Settings	TaskFitToWorkSettings
	Info		TaskFitToWorkInfo

}


type TaskFitToWork struct {
	Description    string
	Status         string
	DateOfCreation int64


}
type  TaskFitToWorkSettings struct {

	Status			string

}
type TaskFitToWorkInfo struct {
	TaskFitToWorkName  	string
	FitToWorkId 		string
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
func (m *Tasks) AddTaskToDB(ctx context.Context  ,companyId string ,WorkBreakSlice []string,TaskWorkTimeSlice []string, ContactId []string,GroupId []string,JobId string,CustomerId string,fitToWorksName string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}

	/*for i :=0; i<len(m.FitToWork.Info); i++ {


	}*/
	log.Println("information",m)
	taskData, err := dB.Child("Tasks").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	//For inserting task details to User
	taskDataString := strings.Split(taskData.String(),"/")
	taskUniqueID := taskDataString[len(taskDataString)-2]
	//for adding fit to work to database

	//setting notification  task in user
	userDataDetails := reflect.ValueOf(m.UsersAndGroups.User)
	for _, key := range userDataDetails.MapKeys() {
		log.Println("inside  notificationnnnn")
		userNotificationDetail :=UserNotification{}
		userNotificationDetail.Date =m.Settings.DateOfCreation
		userNotificationDetail.IsRead =false
		userNotificationDetail.IsViewed =false
		userNotificationDetail.TaskId =taskUniqueID
		userNotificationDetail.TaskName =m.Info.TaskName
		userNotificationDetail.Category ="Tasks"
		err = dB.Child("/Users/"+key.String()+"/Settings/Notifications/Tasks/"+taskUniqueID).Set(userNotificationDetail)
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}


	}

	if len(fitToWorksName) !=0{
		FitToWorkForSetting :=TaskFitToWorkSettings{}
		FitToWorkForInfo  :=TaskFitToWorkInfo{}
		var tempKeySlice []string
		var fitToWOrkKey =""
		instructionOfFitWork :=map[string]TaskFitToWorks{}
		fitToWork :=map[string]FitToWork{}
		db,err :=GetFirebaseClient(ctx,"")
		if err!=nil{
			log.Println("Connection error:",err)
		}
		err = db.Child("FitToWork/"+ companyId).Value(&fitToWork)
		fitToWorkDataValues := reflect.ValueOf(fitToWork)
		for _, fitToWorkKey := range fitToWorkDataValues.MapKeys() {
			tempKeySlice = append(tempKeySlice, fitToWorkKey.String())
		}
		log.Println("value in tempslice",tempKeySlice)
		log.Println("insideeeedfgdgd")
		for _, eachKey := range tempKeySlice {
			log.Println(reflect.TypeOf(fitToWork[eachKey].FitToWorkName))
			log.Println(reflect.TypeOf(fitToWorksName))
			string1 :=fitToWork[eachKey].FitToWorkName
			string2 :=fitToWorksName
			if Compare(string1,string2) ==0 {
				log.Println("insideeee")
				fitToWOrkKey =eachKey
				err = db.Child("FitToWork/"+companyId+"/"+eachKey+"/Instructions").Value(&instructionOfFitWork)
				log.Println("instructions .....",instructionOfFitWork)
				err = dB.Child("/Tasks/"+taskUniqueID+"/FitToWork/FitToWorkInstruction").Set(instructionOfFitWork)

			}

		}
		FitToWorkForSetting.Status =helpers.StatusActive
		err = dB.Child("/Tasks/"+taskUniqueID+"/FitToWork/Settings").Set(FitToWorkForSetting)
		FitToWorkForInfo.TaskFitToWorkName =fitToWorksName
		FitToWorkForInfo.FitToWorkId =fitToWOrkKey
		err = dB.Child("/Tasks/"+taskUniqueID+"/FitToWork/Info").Set(FitToWorkForInfo)
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
	customerInTask.CustomerName =m.Customer.CustomerName
	customerInTask.CustomerStatus =m.Customer.CustomerStatus
	err = dB.Child("/Tasks/"+ taskUniqueID+"/Customer/").Set(customerInTask)
	//log.Println(customerInTask)
	//err = dB.Child("/Customers/"+ CustomerIdForTask+"/Tasks/"+taskUniqueID).Set(CustomerTask)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}

	//setting task id to Group
	GroupTask :=TasksGroup{}
	GroupTask.TasksGroupStatus =helpers.StatusActive
	for i:=0;i<len(GroupId);i++{
		err = dB.Child("/Group/"+ GroupId[i] +"/Tasks/"+taskUniqueID).Set(GroupTask)

	}


	//setting number of task in job
	jobDetail := map[string]Job {}
	//updatedJob :=Job{}
	updatedInfo :=JobInfo{}
	updatedcustomer :=JobCustomer{}
	updatedSettings :=JobSettings{}
	JobTask :=TasksJob{}
	//JobTaskFOrUpdate :=TasksJob{}
	//TotalJobTask :=map[string]TasksJob {}
	JobTask.TasksJobStatus =helpers.StatusActive
	//var StatusArray []string
	err = dB.Child("Jobs").OrderBy("Info/CompanyTeamName").EqualTo(companyId).Value(&jobDetail)
	jobData := reflect.ValueOf(jobDetail)
	for _, key := range jobData.MapKeys() {
		if key.String() ==m.Job.JobId  {
			NumberOfTask :=jobDetail[key.String()].Info.NumberOfTask
			NumberOfTask =NumberOfTask+1
			updatedInfo.JobName =jobDetail[key.String()].Info.JobName
			updatedInfo.JobNumber = jobDetail[key.String()].Info.JobNumber
			updatedInfo.NumberOfTask = NumberOfTask
			updatedInfo.CompanyTeamName = companyId
			updatedInfo.OrderDate =jobDetail[key.String()].Info.OrderDate
			updatedInfo.OrderNumber =jobDetail[key.String()].Info.OrderNumber
			err = dB.Child("/Jobs/"+ key.String()+"/Info").Update(&updatedInfo)
			updatedcustomer.CustomerId= jobDetail[key.String()].Customer.CustomerId
			updatedcustomer.CustomerName= jobDetail[key.String()].Customer.CustomerName
			updatedcustomer.CustomerStatus= jobDetail[key.String()].Customer.CustomerStatus
			err = dB.Child("/Jobs/"+ key.String()+"/Customer").Update(&updatedcustomer)
			updatedSettings.Status = jobDetail[key.String()].Settings.Status
			updatedSettings.DateOfCreation = jobDetail[key.String()].Settings.DateOfCreation
			err = dB.Child("/Jobs/"+ key.String()+"/Settings").Update(&updatedSettings)
			//totalJobDataStatus := reflect.ValueOf(jobDetail[key.String()].Tasks)
			//for _, key := range totalJobDataStatus.MapKeys() {
			//	StatusArray =append(StatusArray,jobDetail[key.String()].Tasks[key.String()].TasksJobStatus)
			//}
			//err = dB.Child("/Jobs/"+ key.String()).Update(&updatedJob)
			//
			////err = dB.Child("/Jobs/"+ JobId+"/Tasks/").Value(&TotalJobTask)
			//totalJobData := reflect.ValueOf(jobDetail[key.String()].Tasks)
			//for _, key := range totalJobData.MapKeys() {
			//	for i:=0;i<len(StatusArray);i++ {
			//		JobTaskFOrUpdate.TasksJobStatus =StatusArray[i]
			//		err = dB.Child("/Jobs/" + JobId + "/Tasks/" + key.String()).Set(JobTaskFOrUpdate)
			//	}
			//}

		}

	}
	//setting task id to Job



	err = dB.Child("/Jobs/"+ JobId+"/Tasks/"+taskUniqueID).Set(JobTask)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}

	//JobTaskMap := make(map[string]TasksJob)
	//JobTask :=TasksJob{}
	//JobTask.TasksJobStatus =helpers.StatusActive
	//JobTaskMap[taskUniqueID] = JobTask
	//err = dB.Child("/Jobs/"+ JobId+"/Tasks/").Set(JobTaskMap)


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

	//function to decrement the number of task  when deleting job
	jobForTask :=map[string]Job{}
	updatedInfo :=JobInfo{}
	err = dB.Child("/Jobs/").Value(&jobForTask)
	jobData := reflect.ValueOf(jobForTask)
	for _, key := range jobData.MapKeys() {
		jobTaskData := reflect.ValueOf(jobForTask[key.String()].Tasks)
		for _, taskKey := range jobTaskData.MapKeys() {
			if taskKey.String() == taskId{
				NumberOfTask :=jobForTask[key.String()].Info.NumberOfTask
				NumberOfTask =NumberOfTask-1
				updatedInfo.JobName =jobForTask[key.String()].Info.JobName
				updatedInfo.JobNumber = jobForTask[key.String()].Info.JobNumber
				updatedInfo.NumberOfTask = NumberOfTask
				updatedInfo.CompanyTeamName = companyId
				updatedInfo.OrderDate =jobForTask[key.String()].Info.OrderDate
				updatedInfo.OrderNumber =jobForTask[key.String()].Info.OrderNumber
				err = dB.Child("/Jobs/"+ key.String()+"/Info").Update(&updatedInfo)
			}


		}

	}
	//delete task from company
	companyTaskDetail := map[string]TaskIdInfo{}
	updatedcompanyTaskDetail :=TaskIdInfo{}
	err = dB.Child("/Company/"+companyId+"/Tasks/").Value(&companyTaskDetail)
	taskDataInCompany := reflect.ValueOf(companyTaskDetail)
	for _, key := range taskDataInCompany.MapKeys() {
		if key.String() == taskId{
			log.Println("inside deletion from company",taskId,key.String())
			updatedcompanyTaskDetail.Status =helpers.StatusInActive
			updatedcompanyTaskDetail.DateOfCreation =companyTaskDetail[key.String()].DateOfCreation
			updatedcompanyTaskDetail.FitToWorkDisplayStatus =companyTaskDetail[key.String()].FitToWorkDisplayStatus
			updatedcompanyTaskDetail.TaskStatus =companyTaskDetail[key.String()].TaskStatus

			err = dB.Child("/Company/"+companyId+"/Tasks/"+taskId).Set(updatedcompanyTaskDetail)
			if err != nil {
				log.Println(err)
				return false
			}

		}
	}
	if err != nil {
		log.Fatal(err)
		return false
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
func (m *Tasks) UpdateTaskToDB( ctx context.Context, taskId string , companyId string,WorkBreakSlice []string,TaskWorkTimeSlice []string,fitToWorkName string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	taskValues :=Tasks{}
	var tempUserKeySlice []string
	var tempContactKeySlice []string
	userName := TaskUser{}
	err = dB.Child("/Tasks/"+ taskId).Value(&taskValues)
	userStatusInTask := reflect.ValueOf(taskValues.UsersAndGroups.User)
	userStatusForTaskFromForm :=reflect.ValueOf(m.UsersAndGroups.User)
	for _, userKeyForTask := range userStatusForTaskFromForm.MapKeys() {
		tempUserKeySlice = append(tempUserKeySlice, userKeyForTask.String())
	}

	for _, key := range userStatusInTask.MapKeys() {
		for i:=0;i<len(tempUserKeySlice);i++{
			if tempUserKeySlice[i]==key.String() {
				userName.UserTaskStatus =taskValues.UsersAndGroups.User[key.String()].UserTaskStatus
				userName.FullName = taskValues.UsersAndGroups.User[key.String()].FullName
				userName.Status =taskValues.UsersAndGroups.User[key.String()].Status
				m.UsersAndGroups.User[key.String()] =userName
			}
		}
	}
	//update contactstatus in task while updating
	taskContactDetail := TaskContact{}
	contactStatusInTask := reflect.ValueOf(taskValues.Contacts)
	contactStatusForTaskFromForm :=reflect.ValueOf(m.Contacts)
	for _, contactKeyForTask := range contactStatusForTaskFromForm.MapKeys() {
		tempContactKeySlice = append(tempContactKeySlice, contactKeyForTask.String())
	}
	for _, key := range contactStatusInTask.MapKeys() {
		for i := 0; i < len(tempContactKeySlice); i++ {
			if tempContactKeySlice[i] == key.String() {
				taskContactDetail.ContactName =taskValues.Contacts[key.String()].ContactName
				taskContactDetail.EmailId =taskValues.Contacts[key.String()].EmailId
				taskContactDetail.PhoneNumber =taskValues.Contacts[key.String()].PhoneNumber
				taskContactDetail.ContactStatus =taskValues.Contacts[key.String()].ContactStatus
				m.Contacts[key.String()] =taskContactDetail

			}
		}
	}

	m.Settings.TaskStatus=taskValues.Settings.TaskStatus
	m.Settings.DateOfCreation =taskValues.Settings.DateOfCreation
	m.Settings.FitToWorkDisplayStatus =taskValues.Settings.FitToWorkDisplayStatus
	m.Settings.Status =taskValues.Settings.Status
	m.Settings.CompletedPercentage =taskValues.Settings.CompletedPercentage
	m.Settings.PendingPercentage =taskValues.Settings.PendingPercentage
	m.Settings.Status =taskValues.Settings.Status
	m.Customer.CustomerStatus =taskValues.Customer.CustomerStatus
	m.Job.JobStatus = taskValues.Job.JobStatus
	log.Println("fdgdfgdfg",m)

	err = dB.Child("/Tasks/"+ taskId).Update(&m)
	if err!=nil{
		log.Println("updation error:",err)
		return false
	}



	//for adding fit to work to database
	log.Println("adsfdfdgdfgdfgdfgdfgdfgdfgfdgdfgdgdfgfdgfdgdf",fitToWorkName)
	var fitToWorkKey string
	FitToWorkForSetting :=TaskFitToWorkSettings{}
	FitToWorkForInfo  :=TaskFitToWorkInfo{}
	var tempKeySlice []string
	instructionOfFitWork :=map[string]TaskFitToWorks{}
	fitToWork :=map[string]FitToWork{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("FitToWork/"+ companyId).Value(&fitToWork)
	fitToWorkDataValues := reflect.ValueOf(fitToWork)
	for _, fitToWorkKey := range fitToWorkDataValues.MapKeys() {
		tempKeySlice = append(tempKeySlice, fitToWorkKey.String())
	}
	if len(fitToWorkName) !=0{
		log.Println("value in tempslice",tempKeySlice)
		for _, eachKey := range tempKeySlice {
			log.Println(reflect.TypeOf(fitToWork[eachKey].FitToWorkName))
			log.Println(reflect.TypeOf(fitToWorkName))
			fitToWorkKey =eachKey
			string1 :=fitToWork[eachKey].FitToWorkName
			string2 :=fitToWorkName
			if Compare(string1,string2) ==0 {
				log.Println("insideeee")
				err = db.Child("FitToWork/"+companyId+"/"+eachKey+"/Instructions").Value(&instructionOfFitWork)
				log.Println("instructions .....",instructionOfFitWork)
				err = dB.Child("/Tasks/"+taskId+"/FitToWork/FitToWorkInstruction").Set(instructionOfFitWork)

			}

		}
		FitToWorkForSetting.Status =helpers.StatusActive
		err = dB.Child("/Tasks/"+taskId+"/FitToWork/Settings").Set(FitToWorkForSetting)
		FitToWorkForInfo.TaskFitToWorkName =fitToWorkName
		FitToWorkForInfo.FitToWorkId = fitToWorkKey
		err = dB.Child("/Tasks/"+taskId+"/FitToWork/Info").Set(FitToWorkForInfo)
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


func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func (m *Tasks) TaskStatusChck(ctx context.Context, taskId string,companyId string)(bool) {
	taskDetail := Tasks{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Tasks/"+ taskId).Value(&taskDetail)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Println("status",taskDetail.Settings.TaskStatus)
	if taskDetail.Settings.TaskStatus == helpers.StatusCompleted {
		return true

	}

	return false
}