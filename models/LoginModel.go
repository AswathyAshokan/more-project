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
	nfcDetail :=  map[string]NFC{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connectivity!")
	}
	log.Println("Email: ", m.Email)
	//err = dB.Child("CompanyAdmins").OrderBy("Test").EqualTo("one").Value(&companyAdmins)
	err =dB.Child("CompanyAdmins").OrderBy("Test").EqualTo("one").Value(&companyAdmins)
	err =dB.Child("NFCTag").Value(&nfcDetail)

	log.Println("Login user details: ",companyAdmins)
	log.Println("NFC details: ",nfcDetail)
}

