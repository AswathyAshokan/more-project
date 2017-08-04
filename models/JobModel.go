/* Author :Aswathy Ashok */
package models

import (
	"log"
	"golang.org/x/net/context"
	"app/passporte/helpers"
	"reflect"
)
type JobInfo struct {
	JobName		string
	JobNumber	string
	NumberOfTask	int64
	CompanyTeamName	string
	OrderNumber	string
	OrderDate	int64
}
type JobSettings struct {
	Status         string
	DateOfCreation int64
}
type JobCustomer struct {
	CustomerId		string
	CustomerName		string
	CustomerStatus 		string
}
type Job   struct {
	Info 		JobInfo
	Settings 	JobSettings
	Customer	JobCustomer
	Tasks		map[string] TasksJob
}
type TasksJob struct {
	TasksJobStatus	string
}
/*Function for add job details to DB*/
func (m *Job) AddJobToDB( ctx context.Context)(bool)  {
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
		return false
	}
	_, err = dB.Child("Jobs").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}

/*Function for get all job details*/
func (m *Job ) GetAllJobs(ctx context.Context,companyTeamName string)(bool,map[string]Job) {
	jobDetail := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	//jobStatus := "Active"

	err = dB.Child("Jobs").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&jobDetail)

	if err != nil {
		log.Fatal(err)
		return false, jobDetail
	}
	return true, jobDetail

}

/*delete job detail from DB*/
//func (m *Job) DeleteJobFromDB(ctx context.Context, jobId string)(bool)  {
//
//	jobUpdate :=JobSettings{}
//	jobDeletion :=JobSettings{}
//	dB, err := GetFirebaseClient(ctx,"")
//	err = dB.Child("/Jobs/"+ jobId+"/Settings").Value(&jobUpdate)
//	jobDeletion.DateOfCreation =jobUpdate.DateOfCreation
//	jobDeletion.Status =helpers.StatusInActive
//
//	if err!=nil{
//		log.Println("Connection error:",err)
//	}
//	err = dB.Child("/Jobs/"+ jobId+"/Settings").Update(&jobDeletion)
//	if err!=nil{
//		log.Println("Deletion error:",err)
//		return false
//	}
//	return true
//}

/*get job details of specific id*/
func (m *Job) GetJobDetailById(ctx context.Context, jobId string)(bool, Job) {
	job := Job{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Jobs/"+ jobId).Value(&job)
	if err != nil {
		log.Fatal(err)
		return false, job
	}
	return true, job

}

/* Update job details to DB*/
func (m *Job) UpdateJobToDB( ctx context.Context,jobId string)(bool)  {
	job :=Job{}
	jobDetail :=Job{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	job.Info.JobName =m.Info.JobName
	job.Info.CompanyTeamName = m.Info.CompanyTeamName
	job.Info.JobNumber = m.Info.JobNumber

	job.Customer.CustomerName = m.Customer.CustomerName
	job.Customer.CustomerId = m.Customer.CustomerId
	err = dB.Child("/Jobs/"+ jobId).Value(&jobDetail)
	job.Info.NumberOfTask = jobDetail.Info.NumberOfTask
	job.Settings.Status =jobDetail.Settings.Status
	job.Settings.DateOfCreation =jobDetail.Settings.DateOfCreation
	job.Customer.CustomerStatus =jobDetail.Customer.CustomerStatus
	job.Info.NumberOfTask =jobDetail.Info.NumberOfTask
	job.Info.OrderNumber = m.Info.OrderNumber
	err = dB.Child("/Jobs/"+ jobId).Update(&job)

	//....updation in task
	jobDetailForUpdation := map[string]Tasks{}
	taskJobForUpdate :=TaskJob{}
	taskJobDetail :=TaskJob{}

	err = dB.Child("/Tasks/").Value(&jobDetailForUpdation)
	dataValue := reflect.ValueOf(jobDetailForUpdation)
	for _, key := range dataValue.MapKeys() {

		if jobDetailForUpdation[key.String()].Job.JobId  ==jobId{

			err = dB.Child("Tasks/" + key.String()+"/Job/").Value(&taskJobDetail)
			taskJobForUpdate.JobId =taskJobDetail.JobId
			taskJobForUpdate.JobStatus =taskJobDetail.JobStatus
			taskJobForUpdate .JobName =m.Info.JobName
			err = dB.Child("Tasks/" + key.String()+"/Job/").Update(&taskJobForUpdate)

		}
	}
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}

	return true

}

func CheckJobNameIsUsed(ctx context.Context, jobName string)bool{
	job := map[string]Job{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("Jobs").OrderBy("JobName").EqualTo(jobName).Value(&job)
	if err!=nil{
		log.Println("Error:",err)
	}
	if len(job)==0{
		return true
	}else{
		return false
	}
}

func CheckJobNumberIsUsed(ctx context.Context, jobNumber string)bool{
	job := map[string]Job{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("Jobs").OrderBy("JobNumber").EqualTo(jobNumber).Value(&job)
	if err!=nil{
		log.Println("Error:",err)
	}
	if len(job)==0{
		return true
	}else{
		return false
	}
}
func (m *Job) DeleteJobFromDB(ctx context.Context, jobId string,TaskSlice []string,companyTeamName string)(bool)  {

	jobDetailForUpdate :=TasksJob{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	jobDetailForUpdate.TasksJobStatus =helpers.StatusInActive
	for i:=0;i<len(TaskSlice);i++{
		err = dB.Child("/Jobs/"+ jobId+"/Tasks/"+TaskSlice[i]).Update(&jobDetailForUpdate)

	}

	jobDetail := JobSettings{}
	updatedJobDetail :=JobSettings{}
	log.Println("gggg")
	err = dB.Child("/Jobs/"+ jobId).Value(&jobDetail)
	updatedJobDetail.DateOfCreation =jobDetail.DateOfCreation
	updatedJobDetail.Status =helpers.StatusInActive
	err = dB.Child("/Jobs/"+ jobId+"/Settings").Update(&updatedJobDetail)
	if err != nil {
		log.Fatal(err)
		return false
	}
	taskJobDetail :=TaskJob{}
	taskJobForUpdate :=TaskJob{}
	taskDetailForUser :=Tasks{}
	for i:=0;i<len(TaskSlice);i++ {
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Job/").Value(&taskJobDetail)
		log.Println("details from task job",)
		taskJobForUpdate.JobId =taskJobDetail.JobId
		taskJobForUpdate.JobName =taskJobDetail.JobName
		taskJobForUpdate.JobStatus =helpers.StatusInActive

		log.Println("fhsgjs",taskJobForUpdate)
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Job/").Update(&taskJobForUpdate)
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
			userTaskDetail.CompanyId = companyTeamName
			userKey := key.String()
			err = dB.Child("/Users/" + userKey + "/Tasks/" + TaskSlice[i]).Update(&userTaskDetail)
			if err!=nil{
				log.Println("Deletion error:",err)
			}
		}
		log.Println("deleted successfully")
	}
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}
func (m *TasksJob) IsJobUsedForTask( ctx context.Context, jobId string)(bool,map[string]TasksJob)  {
	jobDetail := map[string]TasksJob{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Jobs/"+ jobId+"/Tasks/").Value(&jobDetail)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false,jobDetail
	}
	log.Println(jobDetail)
	log.Println("job inside task",jobDetail)

	return true,jobDetail
}
func (m *Job) DeleteJobFromDBForNonTask(ctx context.Context, jobId string)(bool) {
	jobDetail := Job{}
	updatedJobDetail :=Job{}
	log.Println("gggg")

	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Jobs/"+ jobId).Value(&jobDetail)
	updatedJobDetail.Settings.DateOfCreation =jobDetail.Settings.DateOfCreation
	updatedJobDetail.Settings.Status =helpers.StatusInActive
	updatedJobDetail.Info.JobName =jobDetail.Info.JobName
	updatedJobDetail.Info.CompanyTeamName =jobDetail.Info.CompanyTeamName
	updatedJobDetail.Info.JobNumber =jobDetail.Info.JobNumber
	updatedJobDetail.Info.NumberOfTask =jobDetail.Info.NumberOfTask
	updatedJobDetail.Customer.CustomerId =jobDetail.Customer.CustomerId
	updatedJobDetail.Customer.CustomerName =jobDetail.Customer.CustomerName
	err = dB.Child("/Jobs/"+ jobId).Update(&updatedJobDetail)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
//func (m *TasksContact) DeleteContactFromTask(ctx context.Context,contactId string,TaskSlice []string)(bool) {
//
//
//	contactDetailForUpdate :=TasksContact{}
//	dB, err := GetFirebaseClient(ctx,"")
//	if err!=nil{
//		log.Println("Connection error:",err)
//	}
//	contactDetailForUpdate.TaskContactStatus =helpers.StatusInActive
//	for i:=0;i<len(TaskSlice);i++{
//		log.Println(TaskSlice[i])
//		err = dB.Child("/Contacts/"+ contactId+"/Tasks/"+TaskSlice[i]).Update(&contactDetailForUpdate)
//
//	}
//	taskContactDetail :=TaskContact{}
//	taskContactForUpdate :=TaskContact{}
//	for i:=0;i<len(TaskSlice);i++ {
//		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Value(&taskContactDetail)
//		taskContactForUpdate.ContactName =taskContactDetail.ContactName
//		taskContactForUpdate.EmailId =taskContactDetail.EmailId
//		taskContactForUpdate.PhoneNumber =taskContactDetail.PhoneNumber
//		taskContactForUpdate.ContactStatus =helpers.StatusInActive
//		log.Println("fhsgjs",taskContactForUpdate)
//		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Update(&taskContactForUpdate)
//
//	}
//	if err!=nil{
//		log.Println("Insertion error:",err)
//		return false
//	}
//	return true
//}