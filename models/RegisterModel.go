/*Author: Sarath
Date:01/02/2017*/
package models

import "golang.org/x/net/context"

type User struct{
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

func (m *User)AddUser(ctx context.Context){
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {

	}
	_,err = dB.Child("Users").Push(m)
	if err != nil{

	}
}



