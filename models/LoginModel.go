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

func(m *Login)CheckLogin(ctx context.Context)(bool, Admins, Company, string){
	var adminDetails Admins
	var adminId string
	companyDetails := Company{}
	admins := map[string]Admins{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println(err)
		return false, adminDetails, companyDetails, adminId
	}
	log.Println("Email: ", m.Email)
	log.Println("Password: ", m.Password)
	if err := dB.Child("Admins").OrderBy("Info/Email").EqualTo(m.Email).Value(&admins); err != nil {
	    	log.Println(err)
		return false, adminDetails, companyDetails, adminId
	}


	dataValue := reflect.ValueOf(admins)
	for _, key := range dataValue.MapKeys() {
		adminDetails = admins[key.String()]
		adminId = key.String()
	}

	if err := dB.Child("/Company/"+adminDetails.Company.CompanyId).Value(&companyDetails); err !=nil{
		log.Println(err)
		return false, adminDetails, companyDetails, adminId
	}

	log.Println(adminDetails.Info.Password)
	err = bcrypt.CompareHashAndPassword(adminDetails.Info.Password, m.Password)
	if err !=nil{
		log.Println(err)
		return false, adminDetails, companyDetails, adminId

	}
	return true, adminDetails, companyDetails, adminId
}

func(m *Login)CheckSuperAdminLogin(ctx context.Context)(bool,map[string]Admins){
	var superAdminDetails SuperAdmins
	admins := map[string]Admins{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{

		log.Println(err)
		return false,admins
	}
	log.Println("enter email",m.Email)
	if err := dB.Child("SuperAdmins").OrderBy("Info/Email").EqualTo(m.Email).Value(&admins); err != nil {
		log.Println(err)
		log.Println("cp4")
		return false,admins
	}
	if len(admins) == 0{
		return false,admins
	}
	err = bcrypt.CompareHashAndPassword(superAdminDetails.Info.Password, m.Password)
	if err !=nil{
		log.Println(err)
		return false, admins

	}
	return true,admins
}

