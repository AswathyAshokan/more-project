package models

import (
	"log"
	"golang.org/x/net/context"
)


//Fetch all the details of invite user from database
func GetAllUsersDetail(ctx context.Context, userId string ) (map[string]Expirations,bool) {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]Expirations{}
	err = db.Child("Expirations").Value(&value)
	if err != nil {
		log.Fatal(err)
		return value,false
	}
	return value,true
}
