/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"

)
type Invitation struct {
	Info            inviteUser
	Settings        InviteSettings
}

type inviteUser struct {
	FirstName 		string
	LastName 		string
	EmailId 		string
	UserType 		string
	CompanyTeamName		string
}

type InviteSettings struct {
	Status 		string
	DateOfCreation  int64
}


//Add new invite Users to database
func(m *Invitation) AddInviteToDb(ctx context.Context)bool {
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("Invitation").Push(m)
	if err != nil {
		log.Println(err)
		return  false
	}
	return true
}

//Fetch all the details of invite user from database
func GetAllInviteUsersDetails(ctx context.Context) (map[string]Invitation,bool) {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]Invitation{}
	err = db.Child("Invitation").Value(&value)
	if err != nil {
		log.Fatal(err)
		return value,false
	}
	return value,true
}

//delete each invite user from database using invite UserId
func(m *Invitation) DeleteInviteUserById(ctx context.Context, InviteUserId string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Invitation/"+ InviteUserId).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

//fetch all the details of users for editing purpose
func(m *Invitation) GetAllInviteUserForEdit(ctx context.Context, InviteUserId string) (Invitation,bool){
	value := Invitation{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Invitation/"+ InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}

// update the the profile of user by invite user id
func(m *Invitation) UpdateInviteUserById(ctx context.Context,InviteUserKey string) (bool) {

	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Invitation/"+ InviteUserKey).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}

