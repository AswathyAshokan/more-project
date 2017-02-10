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
	companyAdmins := map[string]CompanyAdmins{}
	//nfcDetail := map[string]NFC{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connectivity!")
	}
	log.Println("Email: ", m.Email)
	//err = dB.Child("CompanyAdmins").OrderBy("Test").EqualTo("one").Value(&companyAdmins)
	err =dB.Child("CompanyAdmins").OrderBy("Info/Email").EqualTo("john@gmail.com").Value(&companyAdmins)

	log.Println("Login user details: ",companyAdmins)
}

