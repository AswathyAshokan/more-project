package models

import (
	"log"
	"golang.org/x/net/context"
)

// get all registered company in passporte for super admin
func (m *Company)GetAllRegisteredCompanyDetails(ctx context.Context)(bool,map[string]Company)  {
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

	log.Println("same data:",adminDetailsById)
	return true,adminDetailsById
}
