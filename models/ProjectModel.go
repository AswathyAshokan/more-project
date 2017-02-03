/* Author :Aswathy Ashok */
package models

import (
	"log"

	"golang.org/x/net/context"
)

type Project   struct {

	CustomerName	string
	ProjectName	string
	ProjectNumber	string
	NumberOfTask	string
	Status		string
	CurrentDate	int64
}
func (m *Project) AddProjectToDB( ctx context.Context)(bool)  {


	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	_, err = dB.Child("Project").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true


}
func (m *Project ) RetrieveProjectFromDB(ctx context.Context)(bool,map[string]Project) {
	v := map[string]Project {}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Project").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v


}
func (m *Project) DeleteProjectFromDB(ctx context.Context, projectId string)(bool)  {

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Project/"+ projectId).Remove()
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}