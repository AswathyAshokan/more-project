package models

import (
	"log"
	"golang.org/x/net/context"

)

type Plan struct {
	CompanyPlan	string
}


func(m Company) ChangeCompanyPlan(ctx context.Context,companyId string) (bool,Company){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("/Company/"+ companyId+"/Plan").Set(m.Plan)
	if err != nil {
		log.Println(err)
		return false,m
	}


	return  true,m
}
