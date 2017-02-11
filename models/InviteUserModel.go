/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"

)
type InviteUser struct {

	FirstName string
	LastName string
	EmailId string
	UserType string
	Status string
	DateOfCreation int64
}
func(this *InviteUser) AddInviteToDb(ctx context.Context)bool {
	//log.Println("values in model",this)
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("User").Push(this)
	if err != nil {
		log.Println(err)
		return  false
	}
	return true
}

func(this *InviteUser) DisplayUser(ctx context.Context) map[string]InviteUser {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]InviteUser{}
	err = db.Child("User").Value(&value)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("%s\n", v)
	//log.Println(reflect.TypeOf(v))
	return value


}

//delete a field

func(this *InviteUser) DeleteInviteUser(ctx context.Context, InviteUserId string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/User/"+ InviteUserId).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

//edit a record

func(this *InviteUser) EditInviteUser(ctx context.Context, InviteUserId string) (InviteUser,bool){
	log.Println("invite user key:", InviteUserId)
	value := InviteUser{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/User/"+ InviteUserId).Value(&value)
	log.Println("values:",value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}

	return value,true

}

func(this *InviteUser) UpdateInviteUser(ctx context.Context,InviteUserKey string) (bool) {


	db,err :=GetFirebaseClient(ctx,"")
	log.Println("valueesss:",this)
	err = db.Child("/User/"+ InviteUserKey).Update(&this)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}

