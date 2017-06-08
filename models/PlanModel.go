package models

import (
	"log"
	"golang.org/x/net/context"
	"app/passporte/helpers"

)

type Plan struct {
	CompanyPlan	string
}


func(m Company) ChangeCompanyPlan(ctx context.Context,companyId string) (bool,Company){
	log.Println("gghhsjgshsjsj")
	companyDataFromCompany := Company{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/Company/"+companyId).Value(&companyDataFromCompany)
	if err != nil {
		log.Fatal(err)
		return  false,m
	}
	companyDataFromCompany.Settings.PaymentStatus = helpers.PaymentPaidStatus
	err = db.Child("/Company/"+ companyId).Update(&companyDataFromCompany)
	if err != nil {
		log.Fatal(err)
		return  false,m
	}

	err = db.Child("/Company/"+ companyId+"/Plan").Set(m.Plan)
	if err != nil {
		log.Println(err)
		return false,m
	}


	return  true,m
}
