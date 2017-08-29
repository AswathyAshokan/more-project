package models
import (
	"golang.org/x/net/context"
	"log"
	//"strings"
	//"reflect"
	//
	//"app/passporte/helpers"
	//"strconv"

)

func(m *Users) GetAllUsers(ctx context.Context) (bool,map[string]Users){
	userDetails := map[string]Users{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("Users").Value(&userDetails)
	//if err != nil {
	//	log.Println(err)
	//	return  false,userDetails
	//}
	return true,userDetails


}