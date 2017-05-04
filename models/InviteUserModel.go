/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
	"strings"
	"reflect"

	"app/passporte/helpers"
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
func(m *EmailInvitation) CheckEmailIdInDb(ctx context.Context,companyID string)bool {
	companyInvitation := map[string]Company{}
	companyInvitaionCheck :=CompanyInvitations{}
	var keySlice []string
	var Condition =""


	dB, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err =  dB.Child("Company/"+companyID+"/Invitation").Value(&companyInvitation)
	dataValue := reflect.ValueOf(companyInvitation)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	log.Println("key slice",keySlice)
	for _, keyIn := range keySlice {
		log.Println("key", keyIn)
		err = dB.Child("Company/" + companyID + "/Invitation/" + keyIn).Value(&companyInvitaionCheck)
		log.Println("ggg",companyInvitaionCheck.Email,"from",m.Info.Email)
		if companyInvitaionCheck.Email == m.Info.Email &&( companyInvitaionCheck.UserResponse =="Pending" ||companyInvitaionCheck.UserResponse == "Accepted") {
			log.Println("bnvbnb")
			Condition = "true"
			break

		} else {
			Condition = "false"
		}
	}
	if Condition =="true"{

			return false
	} else{
			return true
		}

	return true
}


func(m *EmailInvitation) AddInviteToDb(ctx context.Context,companyID string)bool {
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	log.Println("condition false")
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
func(m *Invitation) CheckJobIsAssigned(ctx context.Context, InviteUserId string,companyTeamName string) bool {
	companyData := map[string]Company{}
	TaskMap := map[string]UserTasks{}
	value := map[string]Users{}
	invitationData := CompanyInvitations{}
	var keySlice []string
	var taskKeySlice []string
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("t2")
		log.Fatal(err)
		return false
	}
	err = db.Child("Company").Value(&companyData)
	if err != nil {
		log.Println("t3")
		return false
	}
	dataValue := reflect.ValueOf(companyData)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, key := range keySlice {
		err = db.Child("Company/" + key + "/Invitation/" + InviteUserId).Value(&invitationData)
		if err != nil {
			log.Println("t4")
			return false
		}
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(invitationData.Email).Value(&value)
	if err != nil {
		log.Println("t5")
		log.Fatal(err)
		return false
	}
	taskValues := reflect.ValueOf(value)
	for _, taskKey := range taskValues.MapKeys() {
		taskKeySlice = append(taskKeySlice, taskKey.String())
	}
	for _, taskKey := range taskKeySlice {
		err = db.Child("Users/" + taskKey + "/Tasks").Value(&TaskMap)
		if err != nil {
			log.Println("t6")
			log.Fatal(err)
		}

	}
	if len(TaskMap) == 0 {
		log.Println("t7")
		return true
	} else {
		log.Println("t8")
		return false
	}
}

//fetch all the details of users for editing purpose
func GetAllUserFormCompanyEdit(ctx context.Context,companyTeamName string,InviteUserId string) (CompanyInvitations,bool){
	value := CompanyInvitations{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}

// update the the profile of user by invite user id
func(m *CompanyInvitations) UpdateInviteUserById(ctx context.Context,InviteUserId string,companyTeamName string) (bool) {
	invitation :=map[string]Invitation{}
	editInvitation :=EmailInvitation{}
	updateInvitation :=EmailInvitation{}
	var keySlice []string

	value := CompanyInvitations{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return false
	}
	m.Status = value.Status
	m.UserResponse = value.UserResponse
	m.Email = value.Email
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+ InviteUserId).Update(&m)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	err = db.Child("Invitation").Value(&invitation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	dataValue := reflect.ValueOf(invitation)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {

		err = db.Child("Invitation/"+k+"/"+InviteUserId).Value(&editInvitation)
		if err != nil {
			log.Fatal(err)
			return  false
		}
		if m.Email == editInvitation.Info.Email {
			log.Println("k",k,"invitekey",InviteUserId)
			updateInvitation.Info.Email = editInvitation.Info.Email
			updateInvitation.Info.CompanyAdmin = editInvitation.Info.CompanyAdmin
			updateInvitation.Info.CompanyId= editInvitation.Info.CompanyId
			updateInvitation.Info.CompanyName = editInvitation.Info.CompanyName
			updateInvitation.Info.CompanyTeamName = editInvitation.Info.CompanyTeamName
			updateInvitation.Info.FirstName = m.FirstName
			updateInvitation.Info.LastName =  m.LastName
			updateInvitation.Info.UserType = m.UserType
			updateInvitation.Settings.DateOfCreation = editInvitation.Settings.DateOfCreation
			updateInvitation.Settings.UserResponse = editInvitation.Settings.UserResponse
			updateInvitation.Settings.Status = editInvitation.Settings.Status
			/*err = db.Child("Invitation/"+k+"/"+InviteUserId).Update(&updateInvitation)
			if err != nil {
				log.Fatal(err)
				return  false
			}*/

		}

	}
	return true

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
func DeleteInviteUserById(ctx context.Context,InviteUserId string,companyTeamName string)(bool) {
	invitationData :=CompanyInvitations{}
	updateInvitation :=CompanyInvitations{}
	usersInCompany :=CompanyUsers{}
	updateUsersInCompany := CompanyUsers{}
	value :=CompanyUsers{}
	userMap := map[string]Users{}
	updateCompanyStatus := UsersCompany{}
	companyInUsers :=UsersCompany{}
	var keySlice []string
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
		return  false
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(value.Email).Value(&userMap)
	if err != nil {
		log.Fatal(err)
		return false
	}
	dataValue := reflect.ValueOf(userMap)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {
		err = db.Child("Users/"+k+"/Company/"+companyTeamName).Value(&companyInUsers)
		if err != nil {
			log.Fatal(err)
			return false
		}
		updateCompanyStatus.CompanyName = companyInUsers.CompanyName
		updateCompanyStatus.DateOfJoin = companyInUsers.DateOfJoin
		updateCompanyStatus.Status = helpers.UserStatusDeleted
		err = db.Child("Users/"+k+"/Company/"+companyTeamName).Update(&updateCompanyStatus)
		if err != nil {
			log.Fatal(err)
			return false
		}
		err = db.Child("Company/"+companyTeamName+"/Users/"+k).Value(&usersInCompany)
		if err != nil {
			log.Fatal(err)
			return false
		}
		updateUsersInCompany.Status = helpers.UserStatusDeleted
		updateUsersInCompany.DateOfJoin = usersInCompany.DateOfJoin
		updateUsersInCompany.Email=usersInCompany.Email
		updateUsersInCompany.FullName = usersInCompany.FullName
		err = db.Child("Company/"+companyTeamName+"/Users/"+k).Update(&updateCompanyStatus)
		if err != nil {
			log.Fatal(err)
			return false
		}
		err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&invitationData)
		if err != nil {
			log.Fatal(err)
			return false
		}
		updateInvitation.Email = invitationData.Email
		updateInvitation.FirstName = invitationData.FirstName
		updateInvitation.LastName = invitationData.LastName
		updateInvitation.Status= helpers.UserStatusDeleted
		updateInvitation.UserResponse = invitationData.UserResponse
		updateInvitation.UserType = invitationData.UserType
		log.Println("delete",updateInvitation)
		err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Update(&updateInvitation)
		if err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}



//Remove  users from task for delete

func RemoveUsersFromTaskForDelete(ctx context.Context,companyTeamName  string,InviteUserId string)(bool) {
	value :=CompanyInvitations{}
	userMap := map[string]Users{}
	taskInUsersMap :=map[string]UserTasks{}
	eachTaskInUser :=UserTasks{}
	updateTask := UserTasks{}
	usersInTask :=TaskUser{}
	updateUsersInTask := TaskUser{}
	var keySlice []string
	var taskKeySlice []string
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println("e2")
		log.Println(err)
		return  false
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Println("e3")
		log.Fatal(err)
		return false
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(value.Email).Value(&userMap)
	if err != nil {
		log.Println("e4")
		log.Fatal(err)
		return false
	}
	dataValue := reflect.ValueOf(userMap)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {
		err = db.Child("Users/" + k + "/Tasks").Value(&taskInUsersMap)
		if err != nil {
			log.Println("e5")
			log.Fatal(err)
			return false
		}
		taskDataValue := reflect.ValueOf(taskInUsersMap)
		for _,taskKey:= range taskDataValue.MapKeys(){
			taskKeySlice = append(taskKeySlice,taskKey.String())
		}
		for _, specificTaskKey := range taskKeySlice {
			err = db.Child("Users/" + k + "/Tasks/"+specificTaskKey).Value(&eachTaskInUser)
			log.Println("tasks ",eachTaskInUser)
			if err != nil {
				log.Println("e6")
				log.Fatal(err)
				return false
			}
			updateTask.CompanyId = eachTaskInUser.CompanyId
			updateTask.CustomerName = eachTaskInUser.CustomerName
			updateTask.DateOfCreation = eachTaskInUser.DateOfCreation
			updateTask.EndDate = eachTaskInUser.EndDate
			updateTask.JobName = eachTaskInUser.JobName
			updateTask.StartDate = eachTaskInUser.StartDate
			updateTask.Status =helpers.UserStatusDeleted
			updateTask.TaskName = eachTaskInUser.TaskName
			err = db.Child("Users/" + k + "/Tasks/"+specificTaskKey).Update(&updateTask)
			if err != nil {
				log.Println("e8")
				log.Fatal(err)
				return false
			}
			err = db.Child("Tasks/"+specificTaskKey+"/UsersAndGroups/User/"+k).Value(&usersInTask)
			if err != nil {
				log.Println("e9")
				log.Fatal(err)
				return false
			}
			updateUsersInTask.Status = helpers.UserStatusDeleted
			updateUsersInTask.FullName = usersInTask.FullName
			err = db.Child("Tasks/"+specificTaskKey+"/UsersAndGroups/User/"+k).Update(&updateUsersInTask)
			if err != nil {
				log.Println("e10")
				log.Fatal(err)
				return false
			}
		}

	}

	return true
}

func CheckStatusInInvitationOfCompany(ctx context.Context,InviteUserId string, companyTeamName string)(bool) {
	log.Println("id",InviteUserId)
	value :=CompanyInvitations{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("value",value.UserResponse)
	if value.UserResponse == helpers.UserResponsePending || value.UserResponse == helpers.UserResponseRejected{
		return false
	} else  {
		return true
	}
}

func DeleteInviteUserIfStatusIsPending(ctx context.Context,InviteUserId string,companyTeamName string)(bool) {
	value := CompanyInvitations{}
	updateStatus :=CompanyInvitations{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return false
	}
	updateStatus.UserResponse = helpers.UserStatusDeleted
	updateStatus.Email = value.Email
	updateStatus.FirstName = value.FirstName
	updateStatus.LastName = value.LastName
	updateStatus.Status = value.Status
	updateStatus.UserType = value.UserType
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+ InviteUserId).Update(&updateStatus)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}

