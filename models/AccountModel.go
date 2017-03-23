package models

import (
	"log"
	"golang.org/x/net/context"
)

func GetAllSuperAdminsDetails(ctx context.Context)(bool,map[string]SuperAdmins) {
	supreAdmin := map[string]SuperAdmins{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("SuperAdmins").Value(&supreAdmin)
	if err != nil{
		log.Fatal(err)
		return false,supreAdmin
	}
	return true,supreAdmin


}


func(m *SuperAdmins) EditSuperAdminDetails(ctx context.Context) (bool){
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("SuperAdmins").Push(m)

	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}