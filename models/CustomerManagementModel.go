package models

import (
	"log"
	"golang.org/x/net/context"
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

func DeleteCustomerManagementData(ctx context.Context,customerManagementId string)(bool,Company,map[string]Admins){
	companyData := Company{}
	adminData := map[string]Admins{}
	dB,err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println("No Db COnnection!")
	}

	err = dB.Child("/Company/"+customerManagementId).Value(&companyData)
	log.Println("cp1")
	if err != nil {
		log.Println("cp2")
		log.Fatal(err)
		return  false,companyData,adminData
	} else {
		log.Println("cp3")
		err = dB.Child("/Admins/"+customerManagementId).Value(&adminData)
		if err != nil {
			log.Fatal(err)
			return  false,companyData,adminData
		}
		log.Println("cp4")
	}
	log.Println("cp5")
	return true,companyData,adminData


}

//function for update company status From Company
func (m *Company)UpdateCompanyStatusToInactive(ctx context.Context,customerManagementId string)(bool)  {
	dB,err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println("No Db COnnection!")
	}
	err = dB.Child("/Company/"+ customerManagementId).Update(&m)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true
}

//function for update company status from Admin
func (m *Admins)UpdateAdminStatusToInactive(ctx context.Context,customerManagementId string)(bool)  {
	dB,err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println("No Db COnnection!")
	}
	err = dB.Child("/Admins/"+ customerManagementId).Update(&m)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true
}

