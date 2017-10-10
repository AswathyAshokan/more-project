/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
	"strings"
	"reflect"

	"app/passporte/helpers"
	"strconv"

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
	if err != nil {
		log.Println("No Db Connection!")
	}
	dataValue := reflect.ValueOf(companyInvitation)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}

	for _, keyIn := range keySlice {
		err = dB.Child("Company/" + companyID + "/Invitation/" + keyIn).Value(&companyInvitaionCheck)
		if err != nil {
			log.Println("No Db Connection!")
		}
		if companyInvitaionCheck.Email == m.Info.Email &&( companyInvitaionCheck.UserResponse ==helpers.UserResponsePending ||companyInvitaionCheck.UserResponse == helpers.UserResponseAccepted) {
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





func(m *EmailInvitation) AddInviteToDb(ctx context.Context,companyID string,adminName string)bool {
	userDetails := map[string]Users{}
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

	//add invite detail to user for notification

	err = db.Child("Users").OrderBy("Info/Email").EqualTo(invitation.Email).Value(&userDetails)
	InviteNotification := reflect.ValueOf(userDetails)
	for _, userKey := range InviteNotification.MapKeys() {
		userInvitationDetail :=UserInvitations{}
		userInvitationDetail.Category ="InviteUser"
		userInvitationDetail.IsViewed =false
		userInvitationDetail.IsRead =false
		userInvitationDetail.CompanyAdmin =adminName
		userInvitationDetail.CompanyName =m.Info.CompanyName
		userInvitationDetail.Date =m.Settings.DateOfCreation
		userInvitationDetail.IsDeleted =false
		err = db.Child("/Users/"+userKey.String()+"/Settings/Notifications/Invitations/"+invitationUniqueID).Set(userInvitationDetail)
		if err!=nil{
			log.Println("Insertion error:",err)
			return false
		}
	}
 return true
}

//Fetch all the details of invite user from database
func GetAllInviteUsersDetails(ctx context.Context,companyId string) (map[string]CompanyInvitations, string,bool) {
	value :=map[string]CompanyInvitations{}
	limitValue := CompanySettings{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
		return  value,limitValue.LimitedUsers,false
	}
	err =db.Child("/Company/"+companyId+"/Settings").Value(&limitValue)
	if err != nil {
		log.Fatal(err)
		return value,limitValue.LimitedUsers,false
	}



	err = db.Child("/Company/"+companyId+"/Invitation").Value(&value)
	if err != nil {
		log.Fatal(err)
		return value,limitValue.LimitedUsers,false
	}
	//err =db.Child("/WorkLog/").Remove()
	//log.Println("work log deleted")

	return value,limitValue.LimitedUsers,true
}

func UpdateNoOfLimitedUser(ctx context.Context,companyId string,newLimitValues int)int{
	limitValue := CompanySettings{}
	updateLimitValue := CompanySettings{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err =db.Child("/Company/"+companyId+"/Settings/").Value(&limitValue)
	if err != nil {
		log.Fatal(err)
		//return limitValue.LimitedUsers
	}
	sumOfOldANdNewNoOfUsers,_:= strconv.Atoi(limitValue.LimitedUsers)
	newValues := sumOfOldANdNewNoOfUsers +newLimitValues
	newValuesOfString := strconv.Itoa(newValues)
	updateLimitValue.LimitedUsers =newValuesOfString

	updateLimitValue.DateOfCreation = limitValue.DateOfCreation
	updateLimitValue.PaymentStatus  = limitValue.PaymentStatus
	updateLimitValue.UserID = limitValue.UserID
	updateLimitValue.Status = limitValue.Status
	err =db.Child("/Company/"+companyId+"/Settings/").Update(&updateLimitValue)

	return newValues


}


//delete each invite user from database using invite UserId
func(m *Invitation) CheckJobIsAssigned(ctx context.Context, InviteUserId string,companyTeamName string) bool {
	companyData := map[string]Company{}
	TaskMap := map[string]UserTasks{}
	userDetails := map[string]Users{}
	invitationData := CompanyInvitations{}

	var keySlice []string
	var taskKeySlice []string
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Fatal(err)
		//return false
	}
	err = db.Child("Company").Value(&companyData)
	if err != nil {
		//return false
	}
	dataValue := reflect.ValueOf(companyData)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, key := range keySlice {
		err = db.Child("Company/" + key + "/Invitation/" + InviteUserId).Value(&invitationData)
		if err != nil {
			//return false
		}
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(invitationData.Email).Value(&userDetails)
	taskValues := reflect.ValueOf(userDetails)
	for _, taskKey := range taskValues.MapKeys() {
		taskKeySlice = append(taskKeySlice, taskKey.String())
	}
	for _, taskKey := range taskKeySlice {
		err = db.Child("Users/" + taskKey + "/Tasks").Value(&TaskMap)
	}
	if len(TaskMap) == 0 {
		//return true

	}
	var taskKeySlice11 []string
	companyIdData := reflect.ValueOf(TaskMap)
	for _,taskIdForGetCompanyId:=range companyIdData.MapKeys(){
		taskKeySlice11 = append(taskKeySlice11,taskIdForGetCompanyId.String())
	}
	for i:= 0;i<len(taskKeySlice11);i++{
		if TaskMap[taskKeySlice11[i]].CompanyId == companyTeamName{
			return false
		}
	}
	return true
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
	editInvitation :=EmailInvitation{}
	updateInvitation :=EmailInvitation{}
	value := CompanyInvitations{}
	CompanyDataForProfile :=CompanyDataForProfile{}
	userMap := map[string]Users{}
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

	formattedEmail := strings.Replace(value.Email, ".", "_", -1)
	err = db.Child("Invitation/"+formattedEmail+"/"+InviteUserId).Value(&editInvitation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
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
	err = db.Child("Invitation/"+formattedEmail+"/"+InviteUserId).Update(&updateInvitation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(value.Email).Value(&userMap)
	dataValue := reflect.ValueOf(userMap)
	for _, key := range dataValue.MapKeys() {
		UserCompany := UsersCompany{}
		UpdatedUserCompany := UsersCompany{}
		err = db.Child("Users/" + key.String() + "/Company/" + companyTeamName).Value(&UserCompany)
		UpdatedUserCompany.Status = UserCompany.Status
		UpdatedUserCompany.CompanyName = UserCompany.CompanyName
		UpdatedUserCompany.DateOfJoin = UserCompany.DateOfJoin
		UpdatedUserCompany.UserType = m.UserType
		CompanyDataForProfile.UserType = m.UserType
		CompanyDataForProfile.Status = UserCompany.Status
		err = db.Child("Users/" + key.String() + "/Company/" + companyTeamName).Set(&UpdatedUserCompany)
		err = db.Child("/PublicUserProfile/"+key.String()+"/Company/"+companyTeamName).Set(CompanyDataForProfile)
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
	editInvitation :=EmailInvitation{}
	updateInvitationFromInvitation :=EmailInvitation{}
	var groupKeySlice  []string
	var groupMembersDetails = GroupMembers{}
	fullGroup := map[string]Group{}
	groupDetails := map[string]GroupMembers{}
	updateMemberDetails := GroupMembers{}
	//userForGroupDeletion :=map[string]Users{}

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
	/*if err != nil {
		log.Fatal(err)
		return false
	}*/
	dataValue := reflect.ValueOf(userMap)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {

		workLocationValues := map[string]WorkLocation{}
		err = db.Child("WorkLocation").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&workLocationValues)
		if err != nil {
			log.Fatal(err)
			//return false
		}
		WorkLocationUseruu := WorkLocationUser{}
		updateUserStatus := WorkLocationUser{}
		CompanyDetailsOfExpiration := CompanyData{}
		updateCompantDeatilsFromExperation := CompanyData{}
		dataValue := reflect.ValueOf(workLocationValues)
		for _, key := range dataValue.MapKeys() {
			log.Println("key.String()",key.String())
			log.Println("k",k)
			err = db.Child("/WorkLocation/"+key.String()+"/Info/UsersAndGroupsInWorkLocation/User/"+k).Value(&WorkLocationUseruu)
			updateUserStatus.FullName = WorkLocationUseruu.FullName
			updateUserStatus.Status = helpers.UserStatusDeleted
			err = db.Child("/WorkLocation/"+key.String()+"/Info/UsersAndGroupsInWorkLocation/User/"+k).Update(&updateUserStatus)
		}
		eachExpiry :=map[string]Expirations{}
		err = db.Child("/Expirations/"+k).Value(&eachExpiry)
		log.Println("eachExpiry",eachExpiry)
		eachDataValues := reflect.ValueOf(eachExpiry)
		for _, expirationKey := range eachDataValues.MapKeys() {
			log.Println("expirationKey.String()",expirationKey.String())
			err = db.Child("/Expirations/"+k+"/"+expirationKey.String()+"/Company/"+companyTeamName).Value(&CompanyDetailsOfExpiration)
			log.Println("iam in experation model",CompanyDetailsOfExpiration.CompanyName,CompanyDetailsOfExpiration.CompanyStatus)
			updateCompantDeatilsFromExperation.CompanyName = CompanyDetailsOfExpiration.CompanyName
			updateCompantDeatilsFromExperation.CompanyStatus = helpers.UserStatusDeleted
			err = db.Child("/Expirations/"+k+"/"+expirationKey.String()+"/Company/"+companyTeamName).Update(&updateCompantDeatilsFromExperation)
		}
		var deletedKey string
		var tempGroupKeys []string
		err = db.Child("Group").Value(&fullGroup)
		if err != nil {
			log.Fatal(err)
			return  false
		}
		groupDataValue := reflect.ValueOf(fullGroup)
			for _, groupKey := range groupDataValue.MapKeys() {
				groupKeySlice = append(groupKeySlice, groupKey.String())
			}
			for _, eachGroupKey := range groupKeySlice {
				err = db.Child("/Group/"+ eachGroupKey+"/Members/").Value(&groupDetails)
				if err != nil {
					log.Fatal(err)
					return  false
				}
				groupMembersDataValue := reflect.ValueOf(groupDetails)
				for _, groupMembersKey := range groupMembersDataValue.MapKeys() {
					if k == groupMembersKey.String(){
						deletedKey = groupMembersKey.String()

					}
				}
				err = db.Child("/Group/"+ eachGroupKey+"/Members/"+deletedKey).Value(&groupMembersDetails)
				if err != nil {
					log.Fatal(err)
					return  false
				}
				if groupMembersDetails.Status !="" && groupMembersDetails.MemberName != ""{
					tempGroupKeys = append(tempGroupKeys,eachGroupKey)
				}
			}
		var groupKeySlice []string
		for i:=0;i<len(tempGroupKeys);i++{
			updateMemberDetails.Status = helpers.UserStatusDeleted
			updateMemberDetails.MemberName = groupMembersDetails.MemberName
			err = db.Child("/Group/"+ tempGroupKeys[i]+"/Members/"+deletedKey).Update(&updateMemberDetails)
			 if err != nil {
				 log.Fatal(err)
				 return  false
			 }
			 UserGroup :=UserGroup{}
			 err = db.Child("Users/"+deletedKey+"/Group/"+tempGroupKeys[i]).Value(&UserGroup)
			if UserGroup.CompanyId == companyTeamName{
				groupKeySlice = append(groupKeySlice,tempGroupKeys[i])
			}

		 }
		log.Println("groupKeySlice !!!!!!!!",groupKeySlice)
		for j:=0;j<len(groupKeySlice);j++{
			err = db.Child("Users/"+deletedKey+"/Group/"+groupKeySlice[j]).Remove()
		}
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
		updateInvitation.UserResponse = helpers.UserStatusDeleted
		updateInvitation.UserType = invitationData.UserType
		err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Update(&updateInvitation)
		if err != nil {
			log.Fatal(err)
			return false
		}
		formattedEmail := strings.Replace(invitationData.Email, ".", "_", -1)
		err = db.Child("Invitation/"+formattedEmail+"/"+InviteUserId).Value(&editInvitation)
		if err != nil {
			log.Fatal(err)
			return  false
		}
		updateInvitationFromInvitation.Info.Email = editInvitation.Info.Email
		updateInvitationFromInvitation.Info.CompanyAdmin = editInvitation.Info.CompanyAdmin
		updateInvitationFromInvitation.Info.CompanyId= editInvitation.Info.CompanyId
		updateInvitationFromInvitation.Info.CompanyName = editInvitation.Info.CompanyName
		updateInvitationFromInvitation.Info.CompanyTeamName = editInvitation.Info.CompanyTeamName
		updateInvitationFromInvitation.Info.FirstName = editInvitation.Info.FirstName
		updateInvitationFromInvitation.Info.LastName =  editInvitation.Info.LastName
		updateInvitationFromInvitation.Info.UserType = editInvitation.Info.UserType
		updateInvitationFromInvitation.Settings.DateOfCreation = editInvitation.Settings.DateOfCreation
		updateInvitationFromInvitation.Settings.UserResponse = helpers.UserStatusDeleted
		updateInvitationFromInvitation.Settings.Status = helpers.UserStatusDeleted
		err = db.Child("Invitation/"+formattedEmail+"/"+InviteUserId).Update(&updateInvitationFromInvitation)
		if err != nil {
			log.Fatal(err)
			return  false
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
		log.Println(err)
		return  false
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = db.Child("Users").OrderBy("Info/Email").EqualTo(value.Email).Value(&userMap)
	/*if err != nil {
		log.Fatal(err)
		return false
	}*/
	dataValue := reflect.ValueOf(userMap)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {
		err = db.Child("Users/" + k + "/Tasks").Value(&taskInUsersMap)
		if err != nil {
			log.Fatal(err)
			return false
		}
		taskDataValue := reflect.ValueOf(taskInUsersMap)
		for _,taskKey:= range taskDataValue.MapKeys(){
			taskKeySlice = append(taskKeySlice,taskKey.String())
		}
		for _, specificTaskKey := range taskKeySlice {
			if taskInUsersMap[specificTaskKey].CompanyId == companyTeamName {

				err = db.Child("Users/" + k + "/Tasks/" + specificTaskKey).Value(&eachTaskInUser)
				if err != nil {
					log.Fatal(err)
					return false
				}
				updateTask.CompanyId = eachTaskInUser.CompanyId
				updateTask.CustomerName = eachTaskInUser.CustomerName
				updateTask.DateOfCreation = eachTaskInUser.DateOfCreation
				updateTask.EndDate = eachTaskInUser.EndDate
				updateTask.JobName = eachTaskInUser.JobName
				updateTask.StartDate = eachTaskInUser.StartDate
				updateTask.Status = helpers.StatusInActive
				updateTask.TaskName = eachTaskInUser.TaskName
				err = db.Child("Users/" + k + "/Tasks/" + specificTaskKey).Update(&updateTask)
				if err != nil {
					log.Fatal(err)
					return false
				}
			}
			err = db.Child("Tasks/"+specificTaskKey+"/UsersAndGroups/User/"+k).Value(&usersInTask)
			if err != nil {
				log.Fatal(err)
				return false
			}
			updateUsersInTask.Status = helpers.UserStatusDeleted
			updateUsersInTask.FullName = usersInTask.FullName
			updateUsersInTask.UserTaskStatus = usersInTask.UserTaskStatus
			err = db.Child("Tasks/"+specificTaskKey+"/UsersAndGroups/User/"+k).Update(&updateUsersInTask)
			if err != nil {
				log.Fatal(err)
				return false
			}

		}


	}

	return true
}

func CheckStatusInInvitationOfCompany(ctx context.Context,InviteUserId string, companyTeamName string)(bool) {
	value :=CompanyInvitations{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
	}
	if value.UserResponse == helpers.UserResponsePending || value.UserResponse == helpers.UserResponseRejected{
		return false
	} else  {
		return true
	}
}

func DeleteInviteUserIfStatusIsPending(ctx context.Context,InviteUserId string,companyTeamName string)(bool) {
	value := CompanyInvitations{}
	updateStatus :=CompanyInvitations{}
	editInvitation :=EmailInvitation{}
	updateInvitationFromInvitation :=EmailInvitation{}
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
	updateStatus.Status = helpers.UserStatusDeleted
	updateStatus.UserType = value.UserType
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+ InviteUserId).Update(&updateStatus)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	formattedEmail := strings.Replace(value.Email, ".", "_", -1)
	err = db.Child("Invitation/"+formattedEmail+"/"+InviteUserId).Value(&editInvitation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	updateInvitationFromInvitation.Info.Email = editInvitation.Info.Email
	updateInvitationFromInvitation.Info.CompanyAdmin = editInvitation.Info.CompanyAdmin
	updateInvitationFromInvitation.Info.CompanyId= editInvitation.Info.CompanyId
	updateInvitationFromInvitation.Info.CompanyName = editInvitation.Info.CompanyName
	updateInvitationFromInvitation.Info.CompanyTeamName = editInvitation.Info.CompanyTeamName
	updateInvitationFromInvitation.Info.FirstName = editInvitation.Info.FirstName
	updateInvitationFromInvitation.Info.LastName =  editInvitation.Info.LastName
	updateInvitationFromInvitation.Info.UserType = editInvitation.Info.UserType
	updateInvitationFromInvitation.Settings.DateOfCreation = editInvitation.Settings.DateOfCreation
	updateInvitationFromInvitation.Settings.UserResponse = helpers.UserStatusDeleted
	updateInvitationFromInvitation.Settings.Status = helpers.UserStatusDeleted
	err = db.Child("Invitation/"+formattedEmail+"/"+InviteUserId).Update(&updateInvitationFromInvitation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}

/*
func CheckUserAssignedToGroup(ctx context.Context,InviteUserId string,companyTeamName string)bool  {
	log.Println("iiiiiiiiiii i  model")
	value :=CompanyUsers{}
	userMap := map[string]Users{}
	var groupKeySlice  []string
	var groupMembersDetails = GroupMembers{}
	fullGroup := map[string]Group{}
	groupDetails := map[string]GroupMembers{}
	//updateMemberDetails := GroupMembers{}
	var keySlice []string
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Child("Users").OrderBy("Info/Email").EqualTo(value.Email).Value(&userMap)
	*/
/*if err != nil {
		log.Fatal(err)
		return false
	}*//*

	dataValue := reflect.ValueOf(userMap)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {

		err = db.Child("Group").Value(&fullGroup)
		if err != nil {
			log.Fatal(err)
		}
		groupDataValue := reflect.ValueOf(fullGroup)
		for _, groupKey := range groupDataValue.MapKeys() {
			groupKeySlice = append(groupKeySlice, groupKey.String())
		}
		for _, eachGroupKey := range groupKeySlice {
			err = db.Child("/Group/" + eachGroupKey + "/Members/").Value(&groupDetails)
			if err != nil {
				log.Fatal(err)
			}

			groupMembersDataValue := reflect.ValueOf(groupDetails)
			for _, groupMembersKey := range groupMembersDataValue.MapKeys() {
				err = db.Child("/Group/" + eachGroupKey + "/Members/" + groupMembersKey.String()).Value(&groupMembersDetails)
				if err != nil {
					log.Fatal(err)
					return false
				}
				if k == groupMembersKey.String() {
					log.Println("jjjjjjjjjjjjoooooooooo")
					return true
				}

			}
		}
	}
	return false



}*/



/*
func RemoveUsersGroup(ctx context.Context,InviteUserId string,companyTeamName string)bool  {
	value :=CompanyUsers{}
	userMap := map[string]Users{}
	var groupKeySlice  []string
	var groupMembersDetails = GroupMembers{}
	fullGroup := map[string]Group{}
	groupDetails := map[string]GroupMembers{}
	updateMemberDetails := GroupMembers{}
	var keySlice []string
	//userCompany := map[string]UsersCompany{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("Company/"+companyTeamName+"/Invitation/"+InviteUserId).Value(&value)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Child("Users").OrderBy("Info/Email").EqualTo(value.Email).Value(&userMap)

*/
/*if err != nil {
		log.Fatal(err)
		return false
	}*//*


	dataValue := reflect.ValueOf(userMap)
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {
		err = db.Child("Group").Value(&fullGroup)
		if err != nil {
			log.Fatal(err)
		}
		groupDataValue := reflect.ValueOf(fullGroup)
		for _, groupKey := range groupDataValue.MapKeys() {
			groupKeySlice = append(groupKeySlice, groupKey.String())
		}
		log.Println("groupKeySlice",groupKeySlice)
		for _, eachGroupKey := range groupKeySlice {
			err = db.Child("/Group/" + eachGroupKey + "/Members/").Value(&groupDetails)
			if err != nil {
				log.Fatal(err)
			}
			groupMembersDataValue := reflect.ValueOf(groupDetails)
			for _, groupMembersKey := range groupMembersDataValue.MapKeys() {
				err = db.Child("/Group/" + eachGroupKey + "/Members/" + groupMembersKey.String()).Value(&groupMembersDetails)
				if err != nil {
					log.Fatal(err)
					//eturn false
				}
				if k == groupMembersKey.String() {
					updateMemberDetails.Status = helpers.UserStatusDeleted
					updateMemberDetails.MemberName = groupMembersDetails.MemberName
					err = db.Child("/Group/" + eachGroupKey + "/Members/" + groupMembersKey.String()).Update(&updateMemberDetails)
					if err != nil {
						log.Fatal(err)
						//return  false
					}
					return true

				}


			}
		}
	}
	return false



}
*/
