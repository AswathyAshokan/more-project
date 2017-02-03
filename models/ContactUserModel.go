
/* Author :Aswathy Ashok */
package models

import (
	"golang.org/x/net/context"
	"log"



	//"app/go_appengine/goroot/src/go/doc/testdata"
)

type ContactUser   struct {

	Name       	string
	Address		string
	State		string
	Zipcode		string
	Email		string
	PhoneNumber	string
	CurrentDate	int64
	Status		string


}


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


func (m *ContactUser) RetrieveContactFromDB(ctx context.Context)(bool,map[string]ContactUser) {
	v := map[string]ContactUser{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Contacts").Value(&v)
	if err != nil {
		log.Fatal(err)
		return false,v
	}
	log.Println( v)
	return true,v
	//log.Println("There are "+v.getChildrenCount());

}
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

func (m *ContactUser) RetrieveContactIdFromDB(ctx context.Context, contactId string)(bool, ContactUser) {
	log.Println( "keyyy in model", contactId)
	c := ContactUser{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Contacts/"+ contactId).Value(&c)
	if err != nil {
		log.Fatal(err)
		return false,c
	}
	return true,c
	//log.Println("There are "+v.getChildrenCount());

}