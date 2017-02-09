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
	companyAdmins := CompanyAdmins{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connectivity!")
	}
	log.Println("Email: ", m.Email)
	err = dB.Child("CompanyAdmin").OrderBy("Email").EqualTo(m.Email).ChildAdded()

	//err = dB.Child("Company").EqualTo(m.Email).Value(&companyAdmins)
	log.Println("Login user details: ",companyAdmins)
}

