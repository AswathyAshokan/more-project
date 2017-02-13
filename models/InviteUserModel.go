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

//Add new invite Users to database
func(m *InviteUser) AddInviteToDb(ctx context.Context)bool {
	//log.Println("values in model",this)
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("User").Push(m)
	if err != nil {
		log.Println(err)
		return  false
	}
	return true
}

//Fetch all the details of invite user from database
func(m *InviteUser) GetAllInviteUsersDetails(ctx context.Context) map[string]InviteUser {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]InviteUser{}
	err = db.Child("User").Value(&value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

//delete each invite user from database using invite UserId
func(m *InviteUser) DeleteInviteUserById(ctx context.Context, InviteUserId string) bool{
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

func(m *InviteUser) GetAllInviteUserForEdit(ctx context.Context, InviteUserId string) (InviteUser,bool){
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

func(m *InviteUser) UpdateInviteUserById(ctx context.Context,InviteUserKey string) (bool) {

	db,err :=GetFirebaseClient(ctx,"")
	log.Println("valueesss:", m)
	err = db.Child("/User/"+ InviteUserKey).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}

