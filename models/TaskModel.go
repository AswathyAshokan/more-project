/* Author :Aswathy Ashok */
package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"
	//"strings"
	"app/passporte/helpers"
	"time"
	"github.com/kjk/betterguid"
	"math/rand"


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


	log.Println("information",m)
	var UserInsertionCount []string
	var UserNotificationCount []string


	//For inserting task details to User
	//taskDataString := strings.Split(taskData.String(),"/")
	taskUniqueID := betterguid.New()
	//for adding fit to work to database
	 if len(m.Info.CompanyTeamName )!=0 {
		 err = dB.Child("Tasks/"+taskUniqueID).Set(m)

	 }
	//setting notification  task in user
	if len(m.Info.CompanyTeamName )!=0 {
		userDataDetails := reflect.ValueOf(m.UsersAndGroups.User)
		notifyId := betterguid.New()
		for _, key := range userDataDetails.MapKeys() {
			log.Println("inside  notificationnnnn")
			userNotificationDetail := UserNotification{}
			userNotificationDetail.Date = time.Now().Unix()
			userNotificationDetail.IsRead = false
			userNotificationDetail.IsViewed = false
			userNotificationDetail.TaskId = taskUniqueID
			userNotificationDetail.TaskName = m.Info.TaskName
			userNotificationDetail.Category = "Tasks"
			userNotificationDetail.Status = "New"
			userNotificationDetail.IsDeleted = false
			err = dB.Child("/Users/" + key.String() + "/Settings/Notifications/Tasks/" + notifyId).Set(userNotificationDetail)
			UserNotificationCount = append(UserNotificationCount, "true")
			if err != nil {
				log.Println("Insertion error:", err)
				return false
			}
		}
	}


	//...........................................................
	if len(m.Info.CompanyTeamName )!=0 {
		userData := reflect.ValueOf(m.UsersAndGroups.User)
		for _, key := range userData.MapKeys() {
			log.Println("inside task in user")
			log.Println("user key", key.String())
			userTaskDetail := UserTasks{}
			userTaskDetail.DateOfCreation = m.Settings.DateOfCreation
			userTaskDetail.TaskName = m.Info.TaskName
			userTaskDetail.CustomerName = m.Customer.CustomerName
			userTaskDetail.EndDate = m.Info.EndDate
			userTaskDetail.StartDate = m.Info.StartDate
			userTaskDetail.JobName = m.Job.JobName
			userTaskDetail.Status = helpers.StatusPending
			userTaskDetail.CompanyId = companyId
			userKey := key.String()
			err = dB.Child("/Users/" + userKey + "/Tasks/" + taskUniqueID).Set(userTaskDetail)
			UserInsertionCount = append(UserInsertionCount, "true")
			log.Println("indideeeeeeeee our testttttttttttttttttttttttttttttttttttttttttt")

			if err != nil {
				log.Println("Insertion error:", err)
				return false
			}

		}
	}
	//if len(m.UsersAndGroups.User)==len(UserInsertionCount) && len(m.Info.CompanyTeamName )!=0{
	//	err = dB.Child("Tasks/"+taskUniqueID).Set(m)
	//}else{
	//	return false
	//}

	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	//log.Println("last inserted data",taskData)


	//setting number of task in job
	jobDetail := map[string]Job {}

	updatedInfo :=JobInfo{}
	updatedcustomer :=JobCustomer{}
	updatedSettings :=JobSettings{}
	JobTask :=TasksJob{}
	JobTask.TasksJobStatus =helpers.StatusActive
	err = dB.Child("Jobs").OrderBy("Info/CompanyTeamName").EqualTo(companyId).Value(&jobDetail)
	jobData := reflect.ValueOf(jobDetail)
	for _, key := range jobData.MapKeys() {
		if key.String() ==m.Job.JobId  {
			log.Println("inside job addddddddddd")
			NumberOfTask :=jobDetail[key.String()].Info.NumberOfTask
			NumberOfTaskNew :=NumberOfTask+1
			updatedInfo.JobName =jobDetail[key.String()].Info.JobName
			updatedInfo.JobNumber = jobDetail[key.String()].Info.JobNumber
			updatedInfo.NumberOfTask = NumberOfTaskNew
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
	if len(WorkBreakSlice) !=0{
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

	}
	log.Println("l1")

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
	log.Println("k1")
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


	log.Println("k5")
	//setting task id to Job



	err = dB.Child("/Jobs/"+ JobId+"/Tasks/"+taskUniqueID).Set(JobTask)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	if len(m.UsersAndGroups.User) !=len(UserInsertionCount) && len(m.Info.CompanyTeamName )!=0{
		log.Println("danger111111")
		userDataForTest := reflect.ValueOf(m.UsersAndGroups.User)
		userTaskDetailForTest := UserTasks{}
		userDetailsDupilcate :=UserTasks{}
		userNotificationDetail :=UserNotification{}
		notifyId := betterguid.New()
		userNotificationDetail.Date =time.Now().Unix()
		userNotificationDetail.IsRead =false
		userNotificationDetail.IsViewed =false
		userNotificationDetail.TaskId =taskUniqueID
		userNotificationDetail.TaskName =m.Info.TaskName
		userNotificationDetail.Category ="Tasks"
		userNotificationDetail.Status ="New"
		userNotificationDetail.IsDeleted =false
		for _, userKeyForTest := range userDataForTest.MapKeys() {
			err = dB.Child("/Users/"+userKeyForTest.String()+"/Tasks/"+taskUniqueID).Value(&userTaskDetailForTest)
			if len(userTaskDetailForTest.TaskName) ==0{
				userDetailsDupilcate.DateOfCreation = m.Settings.DateOfCreation
				userDetailsDupilcate.TaskName = m.Info.TaskName
				userDetailsDupilcate.CustomerName = m.Customer.CustomerName
				userDetailsDupilcate.EndDate = m.Info.EndDate
				userDetailsDupilcate.StartDate =m.Info.StartDate
				userDetailsDupilcate.JobName = m.Job.JobName
				userDetailsDupilcate.Status = helpers.StatusPending
				userDetailsDupilcate.CompanyId = companyId
				err = dB.Child("/Users/"+userKeyForTest.String()+"/Tasks/"+taskUniqueID).Set(userDetailsDupilcate)
				err = dB.Child("/Users/"+userKeyForTest.String()+"/Settings/Notifications/Tasks/"+notifyId).Set(userNotificationDetail)

			}
		}
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
	taskValues :=Tasks{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	log.Println("taskkkkk id in model",taskId)
	err = dB.Child("/Tasks/"+ taskId).Value(&taskValues)
	log.Println("task details",taskValues)
	taskDeletion.Status =helpers.StatusInActive
	err = dB.Child("/Tasks/"+ taskId+"/Settings").Value(&taskUpdate)
	//function to decrement the number of task  when deleting job
	jobForTask :=map[string]Job{}
	updatedInfo :=JobInfo{}
	err = dB.Child("/Jobs/").Value(&jobForTask)
	jobData := reflect.ValueOf(jobForTask)
	for _, key := range jobData.MapKeys() {
		log.Println("job details",len(taskValues.Job.JobId) )
		log.Println("task details",taskValues.Job.JobId)
		if len(taskValues.Job.JobName) !=0{
			if key.String() == taskValues.Job.JobId {
				log.Println("inside delet=ction and updation of jobbb")
				NumberOfTask := jobForTask[key.String()].Info.NumberOfTask
				if NumberOfTask > 0 {
					NewNumberOfTask := NumberOfTask - 1
					updatedInfo.NumberOfTask = NewNumberOfTask
				}

				updatedInfo.JobName = jobForTask[key.String()].Info.JobName
				updatedInfo.JobNumber = jobForTask[key.String()].Info.JobNumber

				updatedInfo.CompanyTeamName = companyId
				updatedInfo.OrderDate = jobForTask[key.String()].Info.OrderDate
				updatedInfo.OrderNumber = jobForTask[key.String()].Info.OrderNumber
				err = dB.Child("/Jobs/" + key.String() + "/Info").Update(&updatedInfo)

			}

		}


	}

	taskDeletion.DateOfCreation =taskUpdate.DateOfCreation
	err = dB.Child("/Tasks/"+ taskId+"/Settings").Update(&taskDeletion)
	err = dB.Child("/Tasks/"+ taskId).Value(&taskDetailForUser)
	userData := reflect.ValueOf(taskDetailForUser.UsersAndGroups.User)
	for _, key := range userData.MapKeys() {
		log.Println("username ",key.String())
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

	//setting task id to contact
	ContactTask :=TasksContact{}
	ContactTask.TaskContactStatus =helpers.StatusInActive
	userDatas := reflect.ValueOf(taskDetailForUser.Contacts)
	for _, key := range userDatas.MapKeys() {
		err = dB.Child("/Contacts/"+ key.String()+"/Tasks/"+taskId).Set(ContactTask)
	}
	JobTask :=TasksJob{}
	JobTask.TasksJobStatus =helpers.StatusInActive
	JobIdInTask :=taskDetailForUser.Job.JobId
	err = dB.Child("/Jobs/"+ JobIdInTask+"/Tasks/"+taskId).Set(JobTask)
	//updated on user task notification
	notifyDeleteId := betterguid.New()
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 1)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	generatedString :=string(result)
	log.Println("genertedstring",generatedString)
	newGeneratedKey:=notifyDeleteId[0:len(notifyDeleteId)-1]+generatedString
	log.Println("newly gener",newGeneratedKey)
	userDataDetails := reflect.ValueOf(taskDetailForUser.UsersAndGroups.User)

	log.Println("iddddddddddd",notifyDeleteId)
	for _, key := range userDataDetails.MapKeys() {
		log.Println("inside  notificationnnnn")
		log.Println("deleted user",key.String())
		userNotificationDetail :=UserNotification{}
		userNotificationDetail.Date =time.Now().Unix()
		userNotificationDetail.IsRead =false
		userNotificationDetail.IsViewed =false
		userNotificationDetail.TaskId =taskId
		userNotificationDetail.TaskName =taskDetailForUser.Info.TaskName
		userNotificationDetail.Category ="Tasks"
		userNotificationDetail.Status ="Deleted"
		userNotificationDetail.IsDeleted= false
		err = dB.Child("/Users/"+key.String()+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
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

func Inslice(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}

func InArray(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}
//func Remove(s []string, r string) []string {
//	for i, v := range s {
//		if v == r {
//			return append(s[:i], s[i+1:]...)
//		}
//	}
//	return s
//}

/* Function for update task on DB*/
func (m *Tasks) UpdateTaskToDB( ctx context.Context, taskId string , companyId string,WorkBreakSlice []string,TaskWorkTimeSlice []string,fitToWorkName string,ContactId []string,GroupId []string,JobId string,CustomerId string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}


	//generating unique id
	//userDataDetails := reflect.ValueOf(m.UsersAndGroups.User)
	notifyUpdatedId := betterguid.New()
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 2)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	generatedString :=string(result)
	newGeneratedKey:=notifyUpdatedId[0:len(notifyUpdatedId)-2]+generatedString
	log.Println("newly gener",newGeneratedKey)

	taskValues :=Tasks{}
	var tempUserKeySlice []string
	var tempUserKeyForOldTask []string
	var uniqueUserKey  []string
	var tempContactKeySlice []string
	//var oldElementNotInNew []string
	userName := TaskUser{}
	err = dB.Child("/Tasks/"+ taskId).Value(&taskValues)
	userStatusInTask := reflect.ValueOf(taskValues.UsersAndGroups.User)
	userStatusForTaskFromForm :=reflect.ValueOf(m.UsersAndGroups.User)
	for _, userKeyForTask := range userStatusForTaskFromForm.MapKeys() {
		tempUserKeySlice = append(tempUserKeySlice, userKeyForTask.String())
	}
	for _, userKeyForOldTask := range userStatusInTask.MapKeys() {
		if taskValues.UsersAndGroups.User[userKeyForOldTask.String()].Status !=helpers.UserStatusDeleted{
			tempUserKeyForOldTask = append(tempUserKeyForOldTask, userKeyForOldTask.String())

		}
	}
	var EleminatedArray []string


	for _, s := range tempUserKeyForOldTask {
		if !Inslice(s, tempUserKeySlice) {
			EleminatedArray = append(EleminatedArray, s)
		}
	}


	log.Println("the actuall new formed array",tempUserKeySlice)
	log.Println("the old user array",tempUserKeyForOldTask)
	log.Println("the array i got from here that is removed user ayyar",EleminatedArray)



	//for i:=0;i<len(tempUserKeySlice);i++{
	//	tempUserKeyForOldTask = append(tempUserKeyForOldTask,tempUserKeySlice[i])
	//}

	for i := 0; i < 2; i++ {
		for _, s1 := range tempUserKeySlice {
			found := false
			for _, s2 := range tempUserKeyForOldTask {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				uniqueUserKey = append(uniqueUserKey, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			tempUserKeySlice, tempUserKeyForOldTask = tempUserKeyForOldTask, tempUserKeySlice
		}
	}
	log.Println("the elemets of the old tasks",tempUserKeySlice)


	log.Println("the atlast thingsss",uniqueUserKey)
	for i:=0;i<len(uniqueUserKey);i++{
		userNotificationDetail :=UserNotification{}
		userNotificationDetail.Date =time.Now().Unix()
		userNotificationDetail.IsRead =false
		userNotificationDetail.IsViewed =false
		userNotificationDetail.TaskId =taskId
		userNotificationDetail.TaskName =m.Info.TaskName
		userNotificationDetail.Category ="Tasks"
		userNotificationDetail.Status ="New"
		userNotificationDetail.IsDeleted =false
		if taskValues.UsersAndGroups.User[uniqueUserKey[i]].UserTaskStatus !=helpers.StatusCompleted {
			err = dB.Child("/Users/"+uniqueUserKey[i]+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)

		}
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
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
	//updating the user
	userNames :=TaskUser{}
	usersInTask :=reflect.ValueOf(m.UsersAndGroups.User)
	for _, usersInTaskKey := range usersInTask.MapKeys() {
		userInOldTask :=reflect.ValueOf(taskValues.UsersAndGroups.User)
		for _, usersInOldTaskKey := range userInOldTask.MapKeys() {
			if usersInTaskKey.String() ==usersInOldTaskKey.String(){
				userTaskStatus:=  taskValues.UsersAndGroups.User[usersInTaskKey.String()].UserTaskStatus
				userNames.UserTaskStatus=userTaskStatus
				userNames.FullName =m.UsersAndGroups.User[usersInTaskKey.String()].FullName
				userNames.Status =m.UsersAndGroups.User[usersInTaskKey.String()].Status

				m.UsersAndGroups.User[usersInTaskKey.String()]=userNames
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
	//setting task id to contact
	ContactTask :=TasksContact{}
	ContactTask.TaskContactStatus =helpers.StatusActive
	for i:=0;i<len(ContactId);i++{
		err = dB.Child("/Contacts/"+ ContactId[i]+"/Tasks/"+taskId).Set(ContactTask)

	}
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}


	JobTask :=TasksJob{}
	//JobTaskFOrUpdate :=TasksJob{}
	//TotalJobTask :=map[string]TasksJob {}
	JobTask.TasksJobStatus =helpers.StatusActive
	err = dB.Child("/Jobs/"+ JobId+"/Tasks/"+taskId).Set(JobTask)

	//setting task id to Group
	GroupTask :=TasksGroup{}
	GroupTask.TasksGroupStatus =helpers.StatusActive
	for i:=0;i<len(GroupId);i++{
		err = dB.Child("/Group/"+ GroupId[i] +"/Tasks/"+taskId).Set(GroupTask)

	}






	log.Println("under removeeeeeeeeeeeeeeee")
	//remove elements from array
	var newUserArray []string
	if len(EleminatedArray)!=0{
		for _, s := range tempUserKeySlice {
			if !InArray(s, EleminatedArray) {
				newUserArray = append(newUserArray, s)
			}
		}
		log.Println("if any error",tempUserKeySlice)
		for _, key := range userStatusInTask.MapKeys() {
			for i:=0;i<len(newUserArray);i++{
				if newUserArray[i]==key.String()&& taskValues.UsersAndGroups.User[key.String()].UserTaskStatus !=helpers.StatusCompleted {
					log.Println("key in old task",key.String())
					userName.UserTaskStatus =taskValues.UsersAndGroups.User[key.String()].UserTaskStatus
					userName.FullName = taskValues.UsersAndGroups.User[key.String()].FullName
					userName.Status =helpers.StatusActive
					m.UsersAndGroups.User[key.String()] =userName
					userNotificationDetail :=UserNotification{}
					userNotificationDetail.Date =time.Now().Unix()
					userNotificationDetail.IsRead =false
					userNotificationDetail.IsViewed =false
					userNotificationDetail.TaskId =taskId
					userNotificationDetail.TaskName =m.Info.TaskName
					userNotificationDetail.Category ="Tasks"
					userNotificationDetail.Status ="Updated"
					userNotificationDetail.IsDeleted =false
					err = dB.Child("/Users/"+key.String()+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)
					if err!=nil{
						log.Println("Insertion error:",err)
						return false
					}
					break

				}else {


					//log.Println("key not in old task",tempUserKeySlice[i])
					//userNotificationDetail :=UserNotification{}
					//userNotificationDetail.Date =m.Settings.DateOfCreation
					//userNotificationDetail.IsRead =false
					//userNotificationDetail.IsViewed =false
					//userNotificationDetail.TaskId =taskId
					//userNotificationDetail.TaskName =m.Info.TaskName
					//userNotificationDetail.Category ="Tasks"
					//userNotificationDetail.Status ="New"
					//userNotificationDetail.IsDeleted =false
					//err = dB.Child("/Users/"+tempUserKeySlice[i]+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)
					//if err!=nil{
					//	log.Println("Insertion error:",err)
					//	return false
					//}
				}
			}
		}
	}else{
		log.Println("if any error",tempUserKeySlice)
		for _, key := range userStatusInTask.MapKeys() {
			for i:=0;i<len(tempUserKeySlice);i++{
				if tempUserKeySlice[i]==key.String()&& taskValues.UsersAndGroups.User[key.String()].UserTaskStatus !=helpers.StatusCompleted {
					log.Println("key in old task",key.String())
					userName.UserTaskStatus =taskValues.UsersAndGroups.User[key.String()].UserTaskStatus
					userName.FullName = taskValues.UsersAndGroups.User[key.String()].FullName
					userName.Status =helpers.StatusActive
					m.UsersAndGroups.User[key.String()] =userName
					userNotificationDetail :=UserNotification{}
					userNotificationDetail.Date =time.Now().Unix()
					userNotificationDetail.IsRead =false
					userNotificationDetail.IsViewed =false
					userNotificationDetail.TaskId =taskId
					userNotificationDetail.TaskName =m.Info.TaskName
					userNotificationDetail.Category ="Tasks"
					userNotificationDetail.Status ="Updated"
					userNotificationDetail.IsDeleted =false
					err = dB.Child("/Users/"+key.String()+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)
					if err!=nil{
						log.Println("Insertion error:",err)
						return false
					}
					break

				}else {


					//log.Println("key not in old task",tempUserKeySlice[i])
					//userNotificationDetail :=UserNotification{}
					//userNotificationDetail.Date =m.Settings.DateOfCreation
					//userNotificationDetail.IsRead =false
					//userNotificationDetail.IsViewed =false
					//userNotificationDetail.TaskId =taskId
					//userNotificationDetail.TaskName =m.Info.TaskName
					//userNotificationDetail.Category ="Tasks"
					//userNotificationDetail.Status ="New"
					//userNotificationDetail.IsDeleted =false
					//err = dB.Child("/Users/"+tempUserKeySlice[i]+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)
					//if err!=nil{
					//	log.Println("Insertion error:",err)
					//	return false
					//}
				}
			}
		}
	}




	for i:=0;i<len(EleminatedArray);i++{

		userNotificationDetail :=UserNotification{}
		userNotificationDetail.Date =time.Now().Unix()
		userNotificationDetail.IsRead =false
		userNotificationDetail.IsViewed =false
		userNotificationDetail.TaskId =taskId
		userNotificationDetail.TaskName =taskValues.Info.TaskName
		userNotificationDetail.Category ="Tasks"
		userNotificationDetail.Status ="Removed"
		userNotificationDetail.IsDeleted =false
		if taskValues.UsersAndGroups.User[EleminatedArray[i]].UserTaskStatus !=helpers.StatusCompleted {
			err = dB.Child("/Users/"+EleminatedArray[i]+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)

		}

		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}

	}



	if err!=nil{
		log.Println("updation error:",err)
		return false
	}

	//updated task notification
	//updated on user task notification




	//for _, key := range userDataDetails.MapKeys() {
	//	log.Println("inside  notificationnnnn")
	//	userNotificationDetail :=UserNotification{}
	//	userNotificationDetail.Date =m.Settings.DateOfCreation
	//	userNotificationDetail.IsRead =false
	//	userNotificationDetail.IsViewed =false
	//	userNotificationDetail.TaskId =taskId
	//	userNotificationDetail.TaskName =m.Info.TaskName
	//	userNotificationDetail.Category ="Tasks"
	//	userNotificationDetail.Status ="Updated"
	//	err = dB.Child("/Users/"+key.String()+"/Settings/Notifications/Tasks/"+newGeneratedKey).Set(userNotificationDetail)
	//	if err!=nil{
	//		log.Println("Insertion error:",err)
	//		return false
	//	}
	//
	//
	//}



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
	userTaskDetailOfOriginal := UserTasks{}
	userTaskDetailDeleted := UserTasks{}
	userTaskDetailOfDeleted := UserTasks{}


	for _, key := range userData.MapKeys() {

		log.Println("updation of task",newUserArray)
		userKey :=key.String()
		err = dB.Child("/Users/"+userKey+"/Tasks/"+taskId).Value(&userTaskDetailOfOriginal)
		log.Println("task id",taskId)
		log.Println("user id",userKey)

		log.Println("task value",userTaskDetailOfOriginal)
		log.Println("task user value",userTaskDetailOfOriginal.TaskName)
		log.Println("task user value",userTaskDetailOfOriginal.Status)
		if len(userTaskDetailOfOriginal.Status) !=0{
			if(userTaskDetailOfOriginal.Status =="Open"||userTaskDetailOfOriginal.Status=="Completed"){
				log.Println("user key",userKey)
				log.Println("inside 1")
				userTaskDetail.Status =userTaskDetailOfOriginal.Status
				log.Println("st1",userTaskDetail.Status)
			}else{
				log.Println("inside2")
				log.Println("user key",userKey)
				userTaskDetail.Status =helpers.StatusPending
				log.Println("st2",userTaskDetail.Status)



			}
		}else{
			log.Println("inside 3")
			log.Println("user key",userKey)
			userTaskDetail.Status =helpers.StatusPending
			log.Println("st3",userTaskDetail.Status)
		}
		userTaskDetail.CompanyId = companyId
		userTaskDetail.CustomerName = m.Customer.CustomerName
		userTaskDetail.EndDate = m.Info.EndDate
		userTaskDetail.JobName = m.Job.JobName
		userTaskDetail.TaskName = m.Info.TaskName
		userTaskDetail.StartDate = m.Info.StartDate
		userTaskDetail.DateOfCreation =taskValues.Settings.DateOfCreation
		userTaskDetail.Id=taskId
		if taskValues.UsersAndGroups.User[userKey].UserTaskStatus !=helpers.StatusCompleted{
			log.Println("kkkkkkkk")
			err = dB.Child("/Users/"+userKey+"/Tasks/"+taskId).Update(&userTaskDetail)
		}

	}



	//insertion on user for new users
	for i :=0;i<len(uniqueUserKey);i++{
		log.Println("updation of task")
		userTaskDetail.Status =helpers.StatusPending
		userTaskDetail.CompanyId = companyId
		userTaskDetail.CustomerName = m.Customer.CustomerName
		userTaskDetail.EndDate = m.Info.EndDate
		userTaskDetail.JobName = m.Job.JobName
		userTaskDetail.TaskName = m.Info.TaskName
		userTaskDetail.StartDate = m.Info.StartDate
		userTaskDetail.DateOfCreation =taskValues.Settings.DateOfCreation
		userTaskDetail.Id=taskId
		err = dB.Child("/Users/"+uniqueUserKey[i]+"/Tasks/"+taskId).Update(&userTaskDetail)


	}














	//deleted user status
	for i :=0;i<len(EleminatedArray);i++{
		err = dB.Child("/Users/"+EleminatedArray[i]+"/Tasks/"+taskId).Value(&userTaskDetailOfDeleted)
		userTaskDetailDeleted.CompanyId = companyId
		userTaskDetailDeleted.CustomerName = userTaskDetailOfDeleted.CustomerName
		userTaskDetailDeleted.EndDate = userTaskDetailOfDeleted.EndDate
		userTaskDetailDeleted.JobName = userTaskDetailOfDeleted.JobName
		userTaskDetailDeleted.TaskName = userTaskDetailOfDeleted.TaskName
		userTaskDetailDeleted.StartDate = userTaskDetailOfDeleted.StartDate
		userTaskDetailDeleted.DateOfCreation =taskValues.Settings.DateOfCreation
		userTaskDetailDeleted.Status =helpers.StatusInActive
		userTaskDetailDeleted.Id =taskId
		if taskValues.UsersAndGroups.User[EleminatedArray[i]].UserTaskStatus !=helpers.StatusCompleted{
			err = dB.Child("/Users/"+EleminatedArray[i]+"/Tasks/"+taskId).Update(&userTaskDetailDeleted)

		}

	}
	CustomerTask :=TasksCustomer{}
	CustomerTask.TasksCustomerStatus =helpers.StatusActive
	job := Job{}
	//JobId := m.Job.JobId
	err = dB.Child("/Jobs/"+ JobId).Value(&job)
	CustomerIdForTask :=job.Customer.CustomerId
	customerInTask :=TaskCustomer{}
	customerInTask.CustomerId =CustomerIdForTask
	log.Println("customer id",CustomerIdForTask)
	customerInTask.CustomerName =m.Customer.CustomerName
	customerInTask.CustomerStatus =m.Customer.CustomerStatus
	err = dB.Child("/Tasks/"+ taskId+"/Customer/").Set(customerInTask)


	//change in job when the task assigned to another job
	jobDetail := map[string]Job {}
	//updatedJob :=Job{}
	updatedInfo :=JobInfo{}
	updatedcustomer :=JobCustomer{}
	updatedSettings :=JobSettings{}

	if taskValues.Job.JobId == m.Job.JobId{
		log.Println("no change for job")
	}else{
		err = dB.Child("Jobs").OrderBy("Info/CompanyTeamName").EqualTo(companyId).Value(&jobDetail)
		jobData := reflect.ValueOf(jobDetail)
		for _, key := range jobData.MapKeys() {
			if key.String() ==m.Job.JobId && jobDetail[key.String()].Settings.Status==helpers.StatusActive {
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

			if key.String() ==taskValues.Job.JobId && jobDetail[key.String()].Settings.Status==helpers.StatusActive {
				NumberOfTask :=jobDetail[key.String()].Info.NumberOfTask
				log.Println("number of task",NumberOfTask)
				if NumberOfTask >0{
					log.Println("inside job decrement")
					NewNumberOfTask :=NumberOfTask-1
					updatedInfo.NumberOfTask = NewNumberOfTask
				}

				updatedInfo.JobName =jobDetail[key.String()].Info.JobName
				updatedInfo.JobNumber = jobDetail[key.String()].Info.JobNumber
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
	log.Println("w8")
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

func (m *Tasks) TaskDeleteStatusChck(ctx context.Context, taskId string,companyId string)(bool) {
	taskDetail := Tasks{}
	var condition =""
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Tasks/"+ taskId).Value(&taskDetail)
	if err != nil {
		log.Fatal(err)
		return false
	}
	userData := reflect.ValueOf(taskDetail.UsersAndGroups.User)
	for _, key := range userData.MapKeys() {
		if taskDetail.UsersAndGroups.User[key.String()].UserTaskStatus=="Open" &&taskDetail.UsersAndGroups.User[key.String()].Status==helpers.StatusActive{
			condition="true"
			break
		}else{
			condition="false"
		}

	}
	if condition == "true" {
		return true

	}

	return false
}