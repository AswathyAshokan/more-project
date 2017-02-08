/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
)

type Info struct{
	FirstName	string
	LastName	string
	PhoneNo		string
	Email 		string
	Password	[]byte
	CompanyName	string
	Address 	string
	State 		string
	ZipCode 	string
}

type Admin struct {
	FirstName	string
	LastName	string
	Email 		string
}
type Company struct{
	Info	Info
	Admin 	Admin
}

func (m *Company)AddUser(ctx context.Context){
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {

	}
	_,err = dB.Child("Company").Push(m)
	if err != nil{

	}
}



