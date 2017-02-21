/* Author :Aswathy Ashok */
package models

import (
	"log"
	"golang.org/x/net/context"
)
type JobInfo struct {
	JobName		string
	JobNumber	string
	NumberOfTask	string
}
type JobSettings struct {
	Status         string
	DateOfCreation int64
}
type JobCustomer struct {
	CustomerId	string
	CustomerName	string
}
type Job   struct {

	Info 		JobInfo
	Settings 	JobSettings
	Customer	JobCustomer
}

/*Function for add job details to DB*/
func (m *Job) AddJobToDB( ctx context.Context)(bool)  {
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
		return false
	}
	_, err = dB.Child("Job").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}

/*Function for get all job details*/
func (m *Job ) GetAllJobs(ctx context.Context)(bool,map[string]Job) {
	jobDetail := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Job").Value(&jobDetail)
	if err != nil {
		log.Fatal(err)
		return false, jobDetail
	}
	return true, jobDetail

}

/*delete job detail from DB*/
func (m *Job) DeleteJobFromDB(ctx context.Context, jobId string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Job/"+ jobId).Remove()
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}

/*get job details of specific id*/
func (m *Job) GetJobDetailById(ctx context.Context, jobId string)(bool, Job) {
	job := Job{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Job/"+ jobId).Value(&job)
	if err != nil {
		log.Fatal(err)
		return false, job
	}
	return true, job

}

/* Update job details to DB*/
func (m *Job) UpdateJobToDB( ctx context.Context,jobId string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Job/"+ jobId).Update(&m)
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
	err = dB.Child("Job").OrderBy("JobName").EqualTo(jobName).Value(&job)
	if err!=nil{
		log.Println("Error:",err)
	}
	if len(job)==0{
		log.Println("map null:",job)
		return true
	}else{
		log.Println("map not null:",job)
		return false
	}
}

func CheckJobNumberIsUsed(ctx context.Context, jobNumber string)bool{
	job := map[string]Job{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	log.Println("JOB NUMBER:",jobNumber)
	err = dB.Child("Job").OrderBy("JobNumber").EqualTo(jobNumber).Value(&job)
	if err!=nil{
		log.Println("Error:",err)
	}
	if len(job)==0{
		log.Println("map null:",job)
		return true
	}else{
		log.Println("map not null:",job)
		return false
	}
}
