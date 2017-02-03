/* Author :Aswathy Ashok */
package models

import (
	"log"

	"golang.org/x/net/context"
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
func (m *Task) AddToDB(ctx context.Context )  {

	log.Println("values in m:",m)
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	_, err = dB.Child("Task").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
	}


}
func (m *Task) RetrieveFromDB(ctx context.Context)(bool,map[string]Task) {
	v := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Task").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v
	//log.Println("There are "+v.getChildrenCount());

}
func (m *Task) DeleteFromDB(ctx context.Context,key string)(bool)  {

	log.Println(key)
	log.Println("deleteDb")

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Task/"+key).Remove()
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}
func (m *Task) RetrieveFromUserDB(ctx context.Context)(bool,map[string]Task) {
	v := map[string]Task{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Users").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v
	//log.Println("There are "+v.getChildrenCount());

}