/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
	"log"
)

type Login struct{
	Email		string
	Password	[]byte
}

func(m *Login)CheckLogin(ctx context.Context){
	company := Company{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connectivity!")
	}
	err = dB.Child("Company").EqualTo(m.Email).Value(&company)
	log.Println("Login user details: ",company)
}

