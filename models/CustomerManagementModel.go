package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"
	"app/passporte/helpers"
)

// get all registered company in passporte for super admin
func GetAllRegisteredCompanyDetails(ctx context.Context)(bool,map[string]Company)  {
	companyDetails := map[string]Company{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {

		log.Println("No Db Connection!")
	}
	err = dB.Child("Company").Value(&companyDetails)
	if err != nil{
		log.Println("hhh")
		log.Fatal(err)
		return false,companyDetails
	}
	return true,companyDetails

}

// Retrieve data of admin from database

func GetAdminDetailsById(ctx context.Context,adminKeyFromCompany []string) (bool, Admins) {
	adminDetailsById := Admins{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}

	for i := 0; i<len(adminKeyFromCompany); i++{
		err = dB.Child("/Admins/"+adminKeyFromCompany[i]).Value(&adminDetailsById)
		if err != nil{
			log.Fatal(err)
		}
	}
	return true,adminDetailsById
}

// Delete selected record from database

func DeleteCustomerManagementData(ctx context.Context,customerManagementId string)(bool){
	companyDataFromCompany := Company{}
	companyDataFromAdmin := map[string]Admins{}
	tempCompanyData := Admins{}
	dB,err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println("No Db COnnection!")
	}

	err = dB.Child("/Company/"+customerManagementId).Value(&companyDataFromCompany)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	companyDataFromCompany.Settings.Status = helpers.StatusInActive
	err = dB.Child("/Company/"+ customerManagementId).Update(&companyDataFromCompany)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	err = dB.Child("Admins").OrderBy("Company/CompanyId").EqualTo(customerManagementId).Value(&companyDataFromAdmin)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	var keySlice []string
	dataValue := reflect.ValueOf(companyDataFromAdmin)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	//update company status from admin model
	for _, k := range keySlice {
		tempCompanyData.Company.CompanyStatus = helpers.StatusInActive
		tempCompanyData.Company.CompanyId = companyDataFromAdmin[k].Company.CompanyId
		tempCompanyData.Company.CompanyName = companyDataFromAdmin[k].Company.CompanyName
		tempCompanyData.Info.Email = companyDataFromAdmin[k].Info.Email
		tempCompanyData.Info.FirstName = companyDataFromAdmin[k].Info.FirstName
		tempCompanyData.Info.LastName = companyDataFromAdmin[k].Info.LastName
		tempCompanyData.Info.Password = companyDataFromAdmin[k].Info.Password
		tempCompanyData.Info.PhoneNo = companyDataFromAdmin[k].Info.PhoneNo
		tempCompanyData.Settings.DateOfCreation = companyDataFromAdmin[k].Settings.DateOfCreation
		tempCompanyData.Settings.Status = companyDataFromAdmin[k].Settings.Status
		adminKey :=k
		err = dB.Child("/Admins/"+ adminKey).Update(&tempCompanyData)
		if err != nil {
			log.Fatal(err)
			return  false
		}
	}
	return true


}

