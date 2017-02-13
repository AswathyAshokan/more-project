/* Author :Aswathy Ashok */
package models

import (
	"log"
	"golang.org/x/net/context"
)

type Job   struct {

	CustomerName	string
	JobName		string
	JobNumber	string
	NumberOfTask	string
	Status		string
	CurrentDate	int64
}


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


func (m *Job ) RetrieveJobFromDB(ctx context.Context)(bool,map[string]Job) {
	jobDetail := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Job").Value(&jobDetail)
	if err != nil {
		log.Fatal(err)
		return false, jobDetail
	}
	log.Println(jobDetail)
	return true, jobDetail

}


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


func (m *Job ) RetrieveCustomerFromDB(ctx context.Context)(bool,map[string]Job) {
	customerDetail := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Customer").Value(&customerDetail)
	if err != nil {
		log.Fatal(err)
		return false, customerDetail
	}
	log.Println(customerDetail)
	return true, customerDetail
}


func (m *Job)RetrieveCustomerNameFromDB(ctx context.Context, customerId[] string)([] string) {

	job := Job{}
	var customerName []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(customerId) ; i++ {
		err = dB.Child("/Customer/" + customerId[i]).Value(&job)
		customerName =append(customerName, job.CustomerName)

	}
	if err != nil {
		log.Fatal(err)
		return customerName
	}
	return customerName
	//log.Println("There are "+v.getChildrenCount());

}


func (m *Job) RetrieveJobDetailFromDB(ctx context.Context, jobId string)(bool, Job) {
	job := Job{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Job/"+ jobId).Value(&job)
	if err != nil {
		log.Fatal(err)
		return false, job
	}
	return true, job
	//log.Println("There are "+v.getChildrenCount());

}


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