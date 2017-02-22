/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
	"log"
	"reflect"
	"golang.org/x/crypto/bcrypt"
)

type Login struct{
	Email		string
	Password	[]byte
}

func(m *Login)CheckLogin(ctx context.Context)(bool, Admins){
	companyAdmins := map[string]Admins{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("No DB Connectivity!")
	}
	log.Println("Email: ", m.Email)
	log.Println("Password: ", m.Password)
	if err := dB.Child("CompanyAdmins").OrderBy("Info/Email").EqualTo(m.Email).Value(&companyAdmins); err != nil {
	    log.Println(err)
	}
	log.Println("Login user details: ",companyAdmins)

	var adminDetails Admins
	dataValue := reflect.ValueOf(companyAdmins)
	for _, key := range dataValue.MapKeys() {
		adminDetails = companyAdmins[key.String()]
	}
	log.Println(adminDetails.Info.Password)
	err = bcrypt.CompareHashAndPassword(adminDetails.Info.Password, m.Password)
	if err !=nil{
		log.Println(err)
		return false, adminDetails

	}
	return true,adminDetails
}

