
/* Author :Aswathy Ashok */
package models

import (
	"golang.org/x/net/context"
	"log"
	"app/passporte/helpers"
)
type ContactInfo struct {
	Name        		string
	Address    	 	string
	State      	 	string
	ZipCode    	 	string
	Email       		string
	PhoneNumber 		string
	CompanyTeamName 	string
}
type ContactSettings struct {
	DateOfCreation 		int64
	Status         		string
}
type ContactUser   struct {
	Info     	ContactInfo
	Settings 	ContactSettings
	Tasks		map[string] TasksContact

}
type TasksContact struct {

	TaskContactStatus		string
}

/*Function for add Contact to DB*/

func (m *ContactUser) AddContactToDB(ctx context.Context) (bool) {
	log.Println("values in m:",m)
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	_, err = dB.Child("Contacts").Push(m)

	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}

/*Function for get all contact details*/

func GetAllContact(ctx context.Context,companyTeamName string)(bool,map[string]ContactUser) {
	contactDetail := map[string]ContactUser{}
	dB, err := GetFirebaseClient(ctx,"")
	//contactStatus := "Active";
	log.Println("model",companyTeamName)
	err = dB.Child("Contacts").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&contactDetail)
	if err != nil {
		log.Fatal(err)
		return false, contactDetail
	}
	return true, contactDetail

}

/*Function for delete contact from DB*/

func (m *ContactUser) DeleteContactFromDB(ctx context.Context, contactId string)(bool)  {

	contactUpdate :=ContactSettings{}
	contactDelete := ContactSettings{}

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Contacts/"+ contactId+"/Settings").Value(&contactUpdate)
	contactDelete.Status = helpers.StatusInActive
	contactDelete.DateOfCreation = contactUpdate.DateOfCreation
	err = dB.Child("/Contacts/"+ contactId+"/Settings").Update(&contactDelete)
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}

/* Get contact detail of specific id*/
func (m *ContactUser) RetrieveContactIdFromDB(ctx context.Context, contactId string)(bool, ContactUser) {
	contactDetail := ContactUser{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Contacts/"+ contactId).Value(&contactDetail)
	if err != nil {
		log.Fatal(err)
		return false, contactDetail
	}
	return true, contactDetail
}

/*Function for Update contact detail*/
func (m *ContactUser) UpdateContactToDB( ctx context.Context, contactId string)(bool)  {
	contactDetail := ContactUser{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Contacts/"+ contactId).Value(&contactDetail)
	m.Settings.DateOfCreation =contactDetail.Settings.DateOfCreation
	err = dB.Child("/Contacts/"+ contactId).Update(&m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}
func (m *TasksContact) IsContactUsedForTask( ctx context.Context, contactId string)(bool,map[string]TasksContact)  {
	contactDetail := map[string]TasksContact{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Contacts/"+ contactId+"/Tasks/").Value(&contactDetail)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false,contactDetail
	}

	return true,contactDetail
}
//func (m *TasksContact) RemoveContactFromTaskForDelete(ctx context.Context, contactId string)(bool, ContactUser) {
//	contactDetail := ContactUser{}
//	dB, err := GetFirebaseClient(ctx,"")
//	err = dB.Child("/Contacts/"+ contactId+"/Tasks/").Value(&contactDetail)
//	if err != nil {
//		log.Fatal(err)
//		return false, contactDetail
//	}
//	return true, contactDetail
//}
func (m *TasksContact) DeleteContactFromTask(ctx context.Context,contactId string,TaskSlice []string)(bool) {


	contactDetailForUpdate :=TasksContact{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}

	contactDetailForUpdate.TaskContactStatus =helpers.StatusInActive
	for i:=0;i<len(TaskSlice);i++{
		err = dB.Child("/Contacts/"+ contactId+"/Tasks/"+TaskSlice[i]).Set(contactDetailForUpdate)

	}
	taskContactDetail :=TaskContact{}
	taskContactForUpdate :=TaskContact{}
	for i:=0;i<len(TaskSlice);i++ {
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Value(&taskContactDetail)
		taskContactForUpdate.ContactName =taskContactDetail.ContactName
		taskContactForUpdate.EmailId =taskContactDetail.EmailId
		taskContactForUpdate.PhoneNumber =taskContactDetail.PhoneNumber
		taskContactForUpdate.ContactStatus =helpers.StatusInActive
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Set(taskContactForUpdate)

	}
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true

}