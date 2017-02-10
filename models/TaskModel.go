/* Author :Aswathy Ashok */
package models

import (
	"log"

	"golang.org/x/net/context"
	//"golang.org/x/crypto/nacl/box"
)

type Task   struct {

	JobName	string
	TaskName	string
	TaskLocation	string
	StartDate	string
	EndDate		string
	LoginType	string
	Status		string
	TaskDescription	string
	UserNumber	string
	Log		string
	UserType	string
	Contact		string
	FitToWork	string
	CurrentDate	int64

}
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
func (m *Task) RetrieveTaskFromDB(ctx context.Context)(bool,map[string]Task) {
	v := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Task").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


}
func (m *Task) DeleteTaskFromDB(ctx context.Context, taskId string)(bool)  {



	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Task/"+ taskId).Remove()
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}
func (m *Task) RetrieveJobFromDB(ctx context.Context)(bool,map[string]Task) {
	v := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Job").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


}
func (m *Job)RetrieveJobValueFromDB(ctx context.Context, jobId[] string)([] string) {
	log.Println( "keyyy in model", jobId)
	c := Job{}
	var s []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(jobId) ; i++ {
		err = dB.Child("/Job/" + jobId[i]).Value(&c)
		 s =append(s,c.JobName)

	}
	if err != nil {
		log.Fatal(err)
		return s
	}
	return s
	//log.Println("There are "+v.getChildrenCount());

}
func (m *Task) RetrieveContactFromDB(ctx context.Context)(bool,map[string]Task) {
	v := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Contacts").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


}
func (m *ContactUser)RetrieveContactNameFromDB(ctx context.Context, contactId[] string)([] string) {
	log.Println( "keyyy contact model", contactId)
	c := ContactUser{}
	var s []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(contactId) ; i++ {
		err = dB.Child("/Contacts/" + contactId[i]).Value(&c)
		s =append(s,c.Name)

	}
	if err != nil {
		log.Fatal(err)
		return s
	}
	return s
	//log.Println("There are "+v.getChildrenCount());

}
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
func (m *Task) RetrieveTaskDetailFromDB(ctx context.Context, taskId string)(bool, Task) {
	c := Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Task/"+ taskId).Value(&c)
	if err != nil {
		log.Fatal(err)
		return false,c
	}
	return true,c
//log.Println("There are "+v.getChildrenCount());

}
func (m *Task ) RetrieveGroupFromDB(ctx context.Context)(bool,map[string]Job) {
	v := map[string]Job {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Group").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


}
func (m *Task)RetrieveGroupNameFromDB(ctx context.Context, groupId[] string)([] string) {

	c := Task{}
	var s []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(groupId) ; i++ {
		err = dB.Child("/Group/" + groupId[i]).Value(&c)
		s =append(s,c.UserType)

	}
	if err != nil {
		log.Fatal(err)
		return s
	}
	return s
	//log.Println("There are "+v.getChildrenCount());

}