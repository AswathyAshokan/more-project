/* Author :Aswathy Ashok */
package models

import (
	"log"

	"golang.org/x/net/context"
	//"golang.org/x/crypto/nacl/box"
)

type Task   struct {

	ProjectName	string
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
func (m *Task) RetrieveProjectFromDB(ctx context.Context)(bool,map[string]Task) {
	v := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Project").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


}
func (m *Project)RetrieveProjectValueFromDB(ctx context.Context, projectId[] string)([] string) {
	log.Println( "keyyy in model", projectId)
	c := Project{}
	var s []string
	dB, err := GetFirebaseClient(ctx,"")
	for i := 0; i <len(projectId) ; i++ {
		err = dB.Child("/Project/" + projectId[i]).Value(&c)
		 s =append(s,c.ProjectName)

	}
	if err != nil {
		log.Fatal(err)
		return s
	}
	return s
	//log.Println("There are "+v.getChildrenCount());

}