/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
	"strings"
	"reflect"
)
type Invitation struct {
 	Email map[string]EmailInvitation
}
type EmailInvitation struct {
	Info            inviteUser
	Settings        InviteSettings
}

type inviteUser struct {
	FirstName 		string
	LastName 		string
	UserType 		string
	CompanyTeamName		string
	Email 			string
	CompanyName		string
	/*CompanyPlan		string*/
	CompanyAdmin            string
	CompanyId   		string
}

type InviteSettings struct {
	Status 		string
	UserResponse    string
	DateOfCreation  int64
}
type UserCompany struct{
	DateOfJoin	int64
	Status 		string
	CompanyTeamName	string
	CompanyName	string
}

//Add new invite Users to database
func(m *EmailInvitation) AddInviteToDb(ctx context.Context,companyID string)bool {
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	//Dots containing in email id replaced into underscore because firebase does not allow email id as a child in which containing dot
	formattedEmail := strings.Replace(m.Info.Email, ".", "_", -1)
	invitationData,err := db.Child("Invitation").Child(formattedEmail).Push(m)
	if err != nil {
		log.Println(err)
		return  false
	}
	invitationDataString := strings.Split(invitationData.String(),"/")
	invitationUniqueID := invitationDataString[len(invitationDataString)-2]
	invitation := CompanyInvitations{}
	invitation.FirstName = m.Info.FirstName
	invitation.LastName = m.Info.LastName
	invitation.UserResponse = m.Settings.UserResponse
	invitation.Status = m.Settings.Status
	invitation.UserType = m.Info.UserType
	invitation.Email = m.Info.Email
	err = db.Child("/Company/"+companyID+"/Invitation/"+invitationUniqueID).Set(invitation)
	if err != nil {
		log.Println(err)
		return  false
	}
	return true
}

func GetInvitationByEmailId(ctx context.Context,email string,companyTeamName string)(map[string]EmailInvitation,bool) {
	value := map[string]EmailInvitation{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
		return  value,false
	}

	//dB.Child("Admins").OrderBy("Info/Email").EqualTo(m.Email).Value(&admins)

	err = db.Child("Company").Child("Invitation").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value,false
	}
	return value,true

}




//Fetch all the details of invite user from database
func GetAllInviteUsersDetails(ctx context.Context,companyId string) (map[string]CompanyInvitations,bool) {
	value :=map[string]CompanyInvitations{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
		return  value,false
	}
	err = db.Child("/Company/"+companyId+"/Invitation").Value(&value)
	if err != nil {
		log.Fatal(err)
		return value,false
	}
	return value,true
}

//delete each invite user from database using invite UserId
func(m *Invitation) DeleteInviteUserById(ctx context.Context, InviteUserId string,companyTeamName string) bool{
	companyData := map[string]Company{}
	invitationData := map[string]CompanyInvitations{}
	var keySlice []string
	
	db,err := GetFirebaseClient(ctx,"")
	if err !=nil{
		log.Fatal(err)
		return false
	}
	err = db.Child("Company").Value(&companyData)
	if err !=nil{
		return false
	}


	dataValue := reflect.ValueOf(companyData)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, key := range keySlice{
		err = db.Child("Company/"+key+"/Invitation").Value(&invitationData)
		/*inviteValues :=reflect.ValueOf(invitationData)
		for _, k := range dataValue.MapKeys() {
			inviteKeySlice = append(inviteKeySlice, k.String())
		}
		for _, k := range inviteKeySlice{

		}*/

	}
	/*InvitationSettingsUpdate := InviteSettings{}
	InvitationDeletion := InviteSettings{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return  false
	}
	err = db.Child("/Invitation/"+ InviteUserId+"/Settings").Value(&InvitationSettingsUpdate)
	if err != nil {
		log.Fatal(err)
		return false
	}
	InvitationDeletion.DateOfCreation = InvitationSettingsUpdate.DateOfCreation
	InvitationDeletion.Status = helpers.StatusInActive

	err = db.Child("/Invitation/"+ InviteUserId+"/Settings").Update(&InvitationDeletion)
	if err != nil {
		log.Fatal(err)
		return  false
	}*/
	return  true


}

//fetch all the details of users for editing purpose
func GetAllInviteUserForEdit(ctx context.Context) (map[string]Invitation,bool){
	value := map[string]Invitation{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("Invitation").Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}

// update the the profile of user by invite user id
func(m *CompanyInvitations) UpdateInviteUserById(ctx context.Context,InviteUserId string) (bool) {

	/*InvitationFromComapy := InviteSettings{}*/
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("company").Child("/Invitation/"+ InviteUserId).Update(&m)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	/*err = db.Child("/Company/Invitation/"+InviteUserId).Value(&InvitationFromComapy)*/
	/*err = db.Child("/Invitation/"+ InviteUserId+"/Settings").Value(&InvitationSettingsDetails)
	if err != nil {
		log.Fatal(err)
		return false
	}
	*//*m.Settings.Status = InvitationSettingsDetails.Status
	m.Settings.DateOfCreation = InvitationSettingsDetails.DateOfCreation*//*
	err = db.Child("/Invitation/"+ InviteUserId).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true*/
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
	var companyID string
	for key := range company{
		companyID = key
	}
	return companyID

}

func(m *Invitation) GetUsersStatus(ctx context.Context, companyTeamName string)(map[string]Invitation,bool) {

	value := map[string]Invitation{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("Invitation").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return  value,true
}
func (m *Invitation)IsEmailIdUnique(ctx context.Context,emailIdCheck string)(bool) {
	invitationDetails := map[string]Invitation{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	if err :=  dB.Child("Invitation").OrderBy("Info/Email").EqualTo(emailIdCheck).Value(&invitationDetails); err != nil {
		log.Fatal(err)
	}
	if len(invitationDetails)==0{
		return true
	}else{
		return false
	}

}

func GetInvitationById( ctx context.Context,InviteUserId string,key string)(map[string]EmailInvitation)  {
	value := map[string]EmailInvitation{}
	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = dB.Child("/Invitation/"+key).Value(&value)
	if err != nil {
		return value
	}
	return value
}



func CheckEmailIsUsedInvitation(ctx context.Context, emailId string) bool{
	/*companyInvitation := map[string]Company{}

	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	if err =  dB.Child("Company").OrderBy("Info/Email").EqualTo(emailId).Value(&companyAdmins); err != nil {
		log.Fatal(err)
	}
	if len(companyAdmins)==0{

		return true
	}else{

		return false
	}*/
	return false
}

