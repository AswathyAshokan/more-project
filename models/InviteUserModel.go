/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
	"time"
	"app/passporte/helpers"
	"strings"
)
type Invitation struct {
	Info            inviteUser
	Settings        InviteSettings
}

type inviteUser struct {
	FirstName 		string
	LastName 		string
	Email	 		string
	UserType 		string
	CompanyTeamName		string
	CompanyName		string
}

type InviteSettings struct {
	Status 		string
	DateOfCreation  int64
}
type UserCompany struct{
	DateOfJoin	int64
	Status 		string
	CompanyTeamName	string
	CompanyName	string
}

//Add new invite Users to database
func(m *Invitation) AddInviteToDb(ctx context.Context, companyID string)bool {
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	invitedUserData, err := db.Child("Invitation").Push(m)
	if err != nil {
		log.Println(err)
		return  false
	}
	user := map[string]Users{}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(m.Info.Email).Value(&user)
	var userID string
	for key := range user{
		log.Println("KEY>>>",key,user[key])
		userID = key
	}
	companyUsers := CompanyUsers{}
	companyUsers.DateOfJoin = time.Now().Unix()
	companyUsers.Status = helpers.StatusActive

	err = db.Child("Company/"+companyID+"/Users/"+userID).Set(companyUsers)
	if err != nil {
		log.Println(err)
		return false
	}

	userCompany := UserCompany{}
	userCompany.Status = helpers.StatusActive
	userCompany.DateOfJoin = companyUsers.DateOfJoin
	userCompany.CompanyTeamName = m.Info.CompanyTeamName
	userCompany.CompanyName = m.Info.CompanyName

	err = db.Child("Users/"+userID+"/Company/"+companyID).Set(userCompany)
	if err != nil {
		log.Println(err)
		return false
	}

	invitedUserDataString := strings.Split(invitedUserData.String(),"/")
	invitedUserUniqueID := invitedUserDataString[len(invitedUserDataString)-2]
	m.Settings.Status = helpers.StatusAccepted

	err = db.Child("Invitation/"+invitedUserUniqueID+"/Settings/Status").Set(m.Settings.Status)
	return true
}

//Fetch all the details of invite user from database
func GetAllInviteUsersDetails(ctx context.Context,companyTeamName string) (map[string]Invitation,bool) {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]Invitation{}
	err = db.Child("Invitation").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&value)
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

func GetCompanyIdByCompanyTeamName(ctx context.Context, companyTeamName string)string{
	company := map[string]Company{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	if err :=  dB.Child("Company").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&company); err != nil {
		log.Println(err)
	}
	log.Println("HERE!!!!:",company)
	var companyID string
	for key := range company{
		log.Println("KEYSSSSSSSS",key,company[key])
		companyID = key
	}
	return companyID

}

