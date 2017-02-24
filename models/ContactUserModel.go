
/* Author :Aswathy Ashok */
package models

import (
	"golang.org/x/net/context"
	"log"
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

func GetAllContact(ctx context.Context)(bool,map[string]ContactUser) {
	contactDetail := map[string]ContactUser{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Contacts").Value(&contactDetail)
	if err != nil {
		log.Fatal(err)
		return false, contactDetail
	}
	log.Println(contactDetail)
	return true, contactDetail

}

/*Function for delete contact from DB*/

func (m *ContactUser) DeleteContactFromDB(ctx context.Context, contactId string)(bool)  {

	log.Println(contactId)

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Contacts/"+ contactId).Remove()
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

	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Contacts/"+ contactId).Update(&m)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false
	}
	return true
}