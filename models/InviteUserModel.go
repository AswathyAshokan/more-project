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
	_,err = db.Child("Invitation").Push(this)
	if err != nil {
		log.Println(err)
		return  false
	}
	return true
}

func(this *InviteUser) DisplayUser(ctx context.Context) map[string]InviteUser {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	v := map[string]InviteUser{}
	err = db.Child("Invitation").Value(&v)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("%s\n", v)
	//log.Println(reflect.TypeOf(v))
	return v


}

//delete a field

func(this *InviteUser) DeleteUser(ctx context.Context,InviteUserKey string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Invitation/"+InviteUserKey).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

//edit a record

func(this *InviteUser) EditUser(ctx context.Context,InviteUserKey string) (InviteUser,bool){
	value := InviteUser{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Invitation/"+InviteUserKey).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}
