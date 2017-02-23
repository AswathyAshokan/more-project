/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
	"log"
	//"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// Struct for Company
type Company struct{
	Admins		map[string]CompanyAdmin
	Info 		CompanyInfo
	Settings 	CompanySettings
	Plan            string
}

type CompanyAdmin struct {
	FirstName	string
	LastName	string
	Status		string
}

type CompanyInfo struct{
	CompanyName	string
	TeamName	string
	Address		string
	State		string
	ZipCode		string
}

type CompanySettings struct{
	Status		string
	DateOfCreation  int64
}

//Struct for Admin
type Admins struct {
	Company		AdminCompany
	Info     	AdminInfo
	Settings 	AdminSettings
}

type AdminInfo struct {
	FirstName	string
	LastName	string
	PhoneNo		string
	Email		string
	Password	[]byte
}

type AdminSettings struct {
	Status		string
	DateOfCreation  int64
}

type AdminCompany struct {
	CompanyName	string
	CompanyId	string
}

//Register new Company Admin
func (m *Admins)CreateAdminAndCompany(ctx context.Context, company Company) bool {

	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
		return false
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(m.Info.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false
	}
	m.Info.Password = hashedPassword
	adminData, err := dB.Child("Admins").Push(m)
	if err != nil {
		log.Println(err)
		return false
	}
	adminDataString := strings.Split(adminData.String(),"/")
	adminUniqueID := adminDataString[len(adminDataString)-2]

	companyAdmin := CompanyAdmin{}
	companyAdmin.FirstName = m.Info.FirstName
	companyAdmin.LastName = m.Info.LastName
	companyAdmin.Status = m.Settings.Status
	adminMap := make(map[string] CompanyAdmin)
	adminMap[adminUniqueID] = companyAdmin
	company.Admins = adminMap

	companyData, err := dB.Child("Company").Push(company)
	if err != nil {
		log.Println(err)
		return false
	}
	companyDataString := strings.Split(companyData.String(),"/")
	companyUniqueID := companyDataString[len(companyDataString)-2]

	adminsCompany := AdminCompany{}
	adminsCompany.CompanyId = companyUniqueID
	adminsCompany.CompanyName = company.Info.CompanyName

	err = dB.Child("Admins/"+adminUniqueID+"/Company").Set(adminsCompany)

	return true
}

func CheckEmailIsUsed(ctx context.Context, emailId string) bool{
	companyAdmins := map[string]Admins{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	if err :=  dB.Child("Admins").OrderBy("Info/Email").EqualTo(emailId).Value(&companyAdmins); err != nil {
		log.Fatal(err)
	}
	if len(companyAdmins)==0{
		log.Println("map null:",companyAdmins)
		return true
	}else{
		log.Println("map not null:",companyAdmins)
		return false
	}
}
