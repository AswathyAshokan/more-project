/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
	"log"
)

type Info struct {
	FirstName	string
	LastName	string
	PhoneNo		string
	Email		string
	Password	[]byte
	CompanyName	string
	Address		string
	State		string
	ZipCode		string
}

type Admin struct {
	FirstName	string
	LastName  	string
	Email     	string
	Status		string
}
type Company struct {
	Info  Info
	CompanyAdmins []Admin
}

func (m *Company)AddUser(ctx context.Context) bool {
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	_, err = dB.Child("Company").Push(m)
	if err != nil {
		log.Println("Company Registration failed!")
		return false
	} else {
		log.Println("Company Registration Successfull!")
		return true
	}
}



