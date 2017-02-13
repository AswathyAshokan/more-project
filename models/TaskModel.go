/* Author :Aswathy Ashok */
package models

import (
	"log"

	"golang.org/x/net/context"
	//"golang.org/x/crypto/nacl/box"
)

type Task   struct {

	JobName         string
	TaskName        string
	TaskLocation    string
	StartDate       string
	EndDate         string
	LoginType       string
	Status          string
	TaskDescription string
	UserNumber      string
	Log             string
	UserType        string
	ContactId       []string
	FitToWork       string
	CurrentDate     int64

}
type User struct {
	FirstName string
	LastName  string
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
	jobValue := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Job").Value(&jobValue)
	if err != nil {
		log.Fatal(err)
		return false, jobValue
	}
	log.Println(jobValue)
	return true, jobValue


}
func (m *Job)RetrieveJobValueFromDB(ctx context.Context, jobId[] string)([] string) {

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
func (m *ContactUser)RetrieveContactNameFromDB(ctx context.Context, contactId[] string)([] string) {

	c := ContactUser{}
	var contactName []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(contactId) ; i++ {
		err = dB.Child("/Contacts/" + contactId[i]).Value(&c)
		contactName =append(contactName,c.Name)

	}
	if err != nil {
		log.Fatal(err)
		return contactName
	}
	return contactName
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
	taskDetail := Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Task/"+ taskId).Value(&taskDetail)
	if err != nil {
		log.Fatal(err)
		return false, taskDetail
	}
	return true, taskDetail
//log.Println("There are "+v.getChildrenCount());

}
func (m *User ) RetrieveUserFromDB(ctx context.Context)(bool,map[string]User) {
	valueOfUser := map[string]User {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("User").Value(&valueOfUser)
	if err != nil {
		log.Fatal(err)
		return false,valueOfUser
	}

	return true,valueOfUser


}
func (m *User)RetrieveUserNameFromDB(ctx context.Context, userId[] string)([] string) {

	c := User{}
	var allUserNames []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(userId) ; i++ {
		err = dB.Child("/User/" + userId[i]).Value(&c)
		allUserNames = append(allUserNames, (c.FirstName + "" + c.LastName))


	}
	if err != nil {
		log.Fatal(err)
		return allUserNames
	}
	return allUserNames
	//log.Println("There are "+v.getChildrenCount());

}