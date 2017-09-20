/*Author: Sarath
Date:01/02/2017*/
package models

import (
	"golang.org/x/net/context"
	"log"
	//"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"app/passporte/helpers"
)



// Struct for Company
type Company struct{
	Admins			map[string]CompanyAdmin
	Info 			CompanyInfo
	Settings 		CompanySettings
	Plan            	string
	Users			map[string]CompanyUsers
	Invitation       	map[string]CompanyInvitations
	Tasks			map[string] TaskIdInfo
}
type TaskIdInfo struct {
	Status			string
	DateOfCreation		int64
	FitToWorkDisplayStatus	string
	TaskStatus		string
}
type CompanyInvitations struct {
	FirstName	string
	LastName	string
	UserType	string
	Status 		string
	Email 		string
	UserResponse    string
}

type CompanyAdmin struct {
	FirstName	string
	LastName	string
	Status		string
}

type CompanyInfo struct{
	CompanyName	string
	CompanyTeamName	string
	Country		string
	Number 		string
	State		string
	City 		string
	Street 		string
	ZipCode		string
}

type CompanySettings struct{
	Status		 string
	DateOfCreation   int64
	PaymentStatus    string
	LimitedUsers     string
	UserID		string
}

type CompanyUsers struct{
	DateOfJoin	int64
	Status		string
	FullName	string
	Email           string

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
	Status			string
	DateOfCreation  	int64
	ProfilePicture		string
	ThumbProfilePicture	string
}

type AdminCompany struct {
	CompanyName	string
	CompanyId	string
	CompanyStatus	string
}
type Storage struct {
	Bucket       string
	Token        string
	RefreshToken string
	APIKey       string
}


//Register new Company Admin
func (m *Admins)CreateAdminAndCompany(ctx context.Context, company Company) (bool,Company) {

	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
		return false,company
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(m.Info.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false,company
	}
	m.Info.Password = hashedPassword
	adminData, err := dB.Child("Admins").Push(m)
	if err != nil {
		log.Println(err)
		return false,company
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
		return false,company
	}
	companyDataString := strings.Split(companyData.String(),"/")
	companyUniqueID := companyDataString[len(companyDataString)-2]
	company.Info.CompanyTeamName =companyUniqueID
	err = dB.Child("/Company/"+companyUniqueID).Update(&company)
	adminsCompany := AdminCompany{}
	adminsCompany.CompanyId = companyUniqueID
	adminsCompany.CompanyName = company.Info.CompanyName
	adminsCompany.CompanyStatus = helpers.StatusActive
	err = dB.Child("Admins/"+adminUniqueID+"/Company").Set(adminsCompany)
	if err != nil {
		log.Println(err)
		return false,company
	}
	return true,company
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

		return true
	}else{
		return false
	}
}
func (m *Admins)GetCompanyDetails(ctx context.Context, adminId string) (bool,Admins){
	companyAdmins := Admins{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("Admins/"+adminId).Value(&companyAdmins)
	if err != nil{
		log.Fatal(err)
		return false,companyAdmins
	}
     return true,companyAdmins
}
func(m *Admins) EditAdminDetails(ctx context.Context ,adminId string) (bool){
	log.Println("inside update")

	admin := Admins{}
	dB,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = dB.Child("Admins/"+adminId).Value(&admin)
	if err != nil{
		log.Fatal(err)
		return false
	}
	m.Settings.ThumbProfilePicture =admin.Settings.ThumbProfilePicture
	m.Settings.ProfilePicture =admin.Settings.ProfilePicture
	m.Settings.DateOfCreation = admin.Settings.DateOfCreation
	m.Settings.Status = admin.Settings.Status
	m.Info.LastName = admin.Info.LastName
	m.Info.Password = admin.Info.Password
	m.Company.CompanyId =admin.Company.CompanyId
	m.Company.CompanyName = admin.Company.CompanyName
	m.Company.CompanyStatus =admin.Company.CompanyStatus
	err = dB.Child("/Admins/"+adminId).Update(&m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}
func(m *Admins) EditAdminPassword(ctx context.Context ,adminId string,confirmPassword []byte) (bool){
	admin := Admins{}
	dB,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = dB.Child("Admins/"+adminId).Value(&admin)
	if err != nil{
		log.Fatal(err)
		return false
	}
	m.Settings.DateOfCreation = admin.Settings.DateOfCreation
	m.Settings.Status = admin.Settings.Status
	m.Settings.ProfilePicture =admin.Settings.ProfilePicture
	m.Settings.ThumbProfilePicture =admin.Settings.ThumbProfilePicture
	m.Info.LastName = admin.Info.LastName
	m.Info.Email =admin.Info.Email
	m.Info.FirstName = admin.Info.FirstName
	m.Info.PhoneNo = admin.Info.PhoneNo
	m.Company.CompanyId =admin.Company.CompanyId
	m.Company.CompanyName = admin.Company.CompanyName
	m.Company.CompanyStatus =admin.Company.CompanyStatus
	hashedPassword, err := bcrypt.GenerateFromPassword(confirmPassword, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false
	}
	m.Info.Password = hashedPassword
	err = dB.Child("/Admins/"+adminId).Update(&m)
	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}

func (m *Company)UpdateCompanyTeamName(ctx context.Context) (bool) {
	log.Println("in model")
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return  false
	}
	err = db.Child("Company").Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}


func IsEnteredAdminPasswordCorrect(ctx context.Context ,adminId string,enteredOldPassword []byte) (bool){
	admin :=Admins{}
	dB,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = dB.Child("Admins/"+adminId).Value(&admin)
	if err != nil{
		log.Fatal(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword(admin.Info.Password, enteredOldPassword)
	if err !=nil{
		log.Println(err)
		return false
	}
	return true

}
func (m *Admins)AdminDetails(ctx context.Context) (bool,map[string]Admins){
	companyAdmins := map[string]Admins{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("Admins/").Value(&companyAdmins)
	if err != nil{
		log.Fatal(err)
		return false,companyAdmins
	}
	return true,companyAdmins
}

/*-------------------------------Automatic Filling Of Country Related Details--------------------------*/

func GetAllCountryNameForFillDropDownList(ctx context.Context)(bool,map[string]Country){
	countryValues := map[string]Country{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("CountryDetails").Value(&countryValues)
	if err != nil{
		log.Fatal(err)
		return false,countryValues
	}
	return true,countryValues

}

func GetAllStatesByCountry(ctx context.Context,CountryName string,countryCode string)(bool,CountryData){

	countryValues := CountryData{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("CountryDetails/"+countryCode).Value(&countryValues)
	if err != nil{
		log.Fatal(err)
		return false,countryValues
	}
	return true,countryValues

}




