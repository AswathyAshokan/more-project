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

func(m *Login)CheckLogin(ctx context.Context)bool{
	companyAdmins := CompanyAdmins{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connectivity!")
	}
	log.Println("Email: ", m.Email)
	var result map[string]interface{}

	if err := dB.Child("CompanyAdmins/Info").OrderBy("Email").Value(&result); err != nil {
	    log.Println(err)
	}

	log.Println("result:", result)
	log.Println("Login user details: ",companyAdmins)
	return true
}

