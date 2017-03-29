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
	err = bcrypt.CompareHashAndPassword(adminDetails.Info.Password, m.Password)
	if err !=nil{
		log.Println(err)
		return false, adminDetails, companyDetails, adminId

	}
	return true, adminDetails, companyDetails, adminId
}

func(m *Login)CheckSuperAdminLogin(ctx context.Context)(bool,map[string]SuperAdmins){
	/*var superAdminDetails SuperAdmins*/
	superAdmins := map[string]SuperAdmins{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{

		log.Println(err)
		return false,superAdmins
	}
	err = dB.Child("SuperAdmins").OrderBy("Info/Email").EqualTo(m.Email).Value(&superAdmins)
	if err != nil {
		log.Println(err)
		return false,superAdmins
	}
	if len(superAdmins) == 0{
		return false,superAdmins
	}
	dataValue := reflect.ValueOf(superAdmins)
	var tempValueSlice [][]byte
	for _, key := range dataValue.MapKeys() {
		tempValueSlice = append(tempValueSlice, superAdmins[key.String()].Info.Password)
	}

	for i:=0; i< len(tempValueSlice); i++{
		err = bcrypt.CompareHashAndPassword(tempValueSlice[i], m.Password)
		if err !=nil{
			log.Println(err)
			return false,superAdmins
		}

	}
	return true,superAdmins
}

