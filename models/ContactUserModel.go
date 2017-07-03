
/* Author :Aswathy Ashok */
package models

import (
	"golang.org/x/net/context"
	"log"
	"app/passporte/helpers"
	"reflect"
)
type ContactInfo struct {
	Name        		string
	Address    	 	string
	State      	 	string
	ZipCode    	 	string
	Email       		string
	PhoneNumber 		string
	CompanyTeamName 	string
	Country			string
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

func (m *ContactUser) DeleteContactFromDB(ctx context.Context, contactId string,TaskSlice []string)(bool)  {

	contactDetailForUpdate :=TasksContact{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	contactDetailForUpdate.TaskContactStatus =helpers.StatusInActive
	for i:=0;i<len(TaskSlice);i++{
		log.Println(TaskSlice[i])
		err = dB.Child("/Contacts/"+ contactId+"/Tasks/"+TaskSlice[i]).Update(&contactDetailForUpdate)

	}
	taskContactDetail :=TaskContact{}
	taskContactForUpdate :=TaskContact{}
	for i:=0;i<len(TaskSlice);i++ {
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Value(&taskContactDetail)
		taskContactForUpdate.ContactName =taskContactDetail.ContactName
		taskContactForUpdate.EmailId =taskContactDetail.EmailId
		taskContactForUpdate.PhoneNumber =taskContactDetail.PhoneNumber
		taskContactForUpdate.ContactStatus =helpers.StatusInActive
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Update(&taskContactForUpdate)

	}
	contactDetail := ContactUser{}
	updatedContactDetail :=ContactUser{}
	err = dB.Child("/Contacts/"+ contactId).Value(&contactDetail)
	updatedContactDetail.Settings.DateOfCreation =contactDetail.Settings.DateOfCreation
	updatedContactDetail.Settings.Status =helpers.StatusInActive
	updatedContactDetail.Info.Address =contactDetail.Info.Address
	updatedContactDetail.Info.CompanyTeamName =contactDetail.Info.CompanyTeamName
	updatedContactDetail.Info.Email =contactDetail.Info.Email
	updatedContactDetail.Info.Name =contactDetail.Info.Name
	updatedContactDetail.Info.PhoneNumber =contactDetail.Info.PhoneNumber
	updatedContactDetail.Info.State =contactDetail.Info.State
	updatedContactDetail.Info.ZipCode =contactDetail.Info.ZipCode
	updatedContactDetail.Info.Country =contactDetail.Info.Country
	log.Println("dfkfj",)

	err = dB.Child("/Contacts/"+ contactId).Update(&updatedContactDetail)
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


	//.....update in task


	contactDetailForUpdation := map[string]Tasks{}
	taskContactForUpdate :=TaskContact{}
	taskContactDetail :=TaskContact{}

	err = dB.Child("/Tasks/").Value(&contactDetailForUpdation)
	dataValue := reflect.ValueOf(contactDetailForUpdation)
	for _, key := range dataValue.MapKeys() {
		log.Println("hhhh")
		dataValueContact := reflect.ValueOf(contactDetailForUpdation[key.String()].Contacts)
		for _, contactkey := range dataValueContact.MapKeys() {
			if  contactkey.String()== contactId {
				log.Println("task id",key.String())
				err = dB.Child("Tasks/" + key.String() + "/Contacts/"+contactId).Value(&taskContactDetail)
				log.Println("contact inside task",taskContactDetail)
				taskContactForUpdate.ContactName = m.Info.Name
				taskContactForUpdate.EmailId = m.Info.Email
				taskContactForUpdate.ContactStatus = taskContactDetail.ContactStatus
				taskContactForUpdate.PhoneNumber =m.Info.PhoneNumber
				err = dB.Child("Tasks/" + key.String() + "/Contacts/"+contactId).Update(&taskContactForUpdate)

			}
		}
	}
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
func (m *ContactUser) DeleteContactFromDBForNonTask(ctx context.Context, contactId string)(bool) {
	contactDetail := ContactUser{}
	updatedContactDetail :=ContactUser{}
	log.Println("gggg")

	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Contacts/"+ contactId).Value(&contactDetail)
	updatedContactDetail.Settings.DateOfCreation =contactDetail.Settings.DateOfCreation
	updatedContactDetail.Settings.Status =helpers.StatusInActive
	updatedContactDetail.Info.Address =contactDetail.Info.Address
	updatedContactDetail.Info.CompanyTeamName =contactDetail.Info.CompanyTeamName
	updatedContactDetail.Info.Email =contactDetail.Info.Email
	updatedContactDetail.Info.Name =contactDetail.Info.Name
	updatedContactDetail.Info.PhoneNumber =contactDetail.Info.PhoneNumber
	updatedContactDetail.Info.State =contactDetail.Info.State
	updatedContactDetail.Info.ZipCode =contactDetail.Info.ZipCode
	updatedContactDetail.Info.Country = contactDetail.Info.Country

	err = dB.Child("/Contacts/"+ contactId).Update(&updatedContactDetail)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
func (m *TasksContact) DeleteContactFromTask(ctx context.Context,contactId string,TaskSlice []string)(bool) {


	contactDetailForUpdate :=TasksContact{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	contactDetailForUpdate.TaskContactStatus =helpers.StatusInActive
	for i:=0;i<len(TaskSlice);i++{
		log.Println(TaskSlice[i])
		err = dB.Child("/Contacts/"+ contactId+"/Tasks/"+TaskSlice[i]).Update(&contactDetailForUpdate)

	}
	taskContactDetail :=TaskContact{}
	taskContactForUpdate :=TaskContact{}
	for i:=0;i<len(TaskSlice);i++ {
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Value(&taskContactDetail)
		taskContactForUpdate.ContactName =taskContactDetail.ContactName
		taskContactForUpdate.EmailId =taskContactDetail.EmailId
		taskContactForUpdate.PhoneNumber =taskContactDetail.PhoneNumber
		taskContactForUpdate.ContactStatus =helpers.StatusInActive
		log.Println("fhsgjs",taskContactForUpdate)
		err = dB.Child("Tasks/" + TaskSlice[i]+"/Contacts/"+contactId).Update(&taskContactForUpdate)

	}
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}