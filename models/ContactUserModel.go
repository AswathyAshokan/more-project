
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


func (m *ContactUser) AddToDB(ctx context.Context)  {
	log.Println("values in m:",m)
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	_, err = dB.Child("Contacts").Push(m)
	if err!=nil{
		log.Println("Insertion error:",err)
	}

 }


func (m *ContactUser) RetrieveFromDB(ctx context.Context)(bool,map[string]ContactUser) {
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
func (m *ContactUser) DeleteFromDB(ctx context.Context,key string)(bool)  {

	log.Println(key)

	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Contacts/"+key).Remove()
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
 	return true
}

func (m *ContactUser) RetrieveFromDBId(ctx context.Context,key string)(bool, ContactUser) {
	log.Println( "keyyy in model",key)
	c := ContactUser{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Contacts/"+key).Value(&c)
	if err != nil {
		log.Fatal(err)
		return false,c
	}
	return true,c
	//log.Println("There are "+v.getChildrenCount());

}