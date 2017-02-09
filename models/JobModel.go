/* Author :Aswathy Ashok */
package models

import (
	"log"

	"golang.org/x/net/context"
)

type Job   struct {

	CustomerName	string
	JobName	string
	JobNumber	string
	NumberOfTask	string
	Status		string
	CurrentDate	int64
}
func (m *Job) AddJobToDB( ctx context.Context)(bool)  {


	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	_, err = dB.Child("Job").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true


}
func (m *Job ) RetrieveJobFromDB(ctx context.Context)(bool,map[string]Job) {
	v := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Job").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


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
	v := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Customer").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


}
func (m *Job)RetrieveCustomerNameFromDB(ctx context.Context, customerId[] string)([] string) {

	c := Job{}
	var s []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(customerId) ; i++ {
		err = dB.Child("/Customer/" + customerId[i]).Value(&c)
		s =append(s,c.CustomerName)

	}
	if err != nil {
		log.Fatal(err)
		return s
	}
	return s
	//log.Println("There are "+v.getChildrenCount());

}
func (m *Job) RetrieveJobDetailFromDB(ctx context.Context, jobId string)(bool, Job) {
	c := Job{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Job/"+ jobId).Value(&c)
	if err != nil {
		log.Fatal(err)
		return false,c
	}
	return true,c
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