/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
	"app/passporte/helpers"
	"reflect"
	"github.com/kjk/betterguid"
	"time"
	"math/rand"
)
type Group struct {
	Info 		GroupInfo
	Members	 	map[string]GroupMembers
	Settings 	GroupSettings
	Tasks		map[string] TasksGroup
}
type TasksGroup struct {
	TasksGroupStatus	string
}


type GroupMembers struct {
	MemberName	string
	Status 		string
}

type GroupInfo struct {
	GroupName       	string
	CompanyTeamName		string
}

type GroupSettings struct {
	DateOfCreation  	int64
	Status         	 	string
}

// Insert new groups to database
func(m *Group) AddGroupToDb(ctx context.Context) (bool){
	log.Println("add group")
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	UserOrGroup :=UserGroup{}
	UserStatus :=UserGroup{}
	UserOrGroupForUpdate :=UserGroup{}

	groupUniqueID := betterguid.New()
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 3)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	generatedString :=string(result)
	log.Println("genertedstring",generatedString)
	newGeneratedKey:=groupUniqueID[0:len(groupUniqueID)-1]+generatedString
	log.Println("newly gener",newGeneratedKey)
	groupUniqueID =newGeneratedKey




	err = db.Child("Group/"+groupUniqueID).Set(m)
	if err != nil {
		log.Println(err)
		return false
	}

	var UserGroupKey []string


	/*UserGroup.CompanyId = m.Info.CompanyTeamName
	UserGroup.DateOfCreation = m.Settings.DateOfCreation
	UserGroup.GroupName = m.Info.GroupName
	UserGroup.GroupStatus = m.Settings.Status*/
	UserOrGroup.GroupName = m.Info.GroupName
	UserOrGroup.CompanyId = m.Info.CompanyTeamName
	UserOrGroupForUpdate.GroupName=m.Info.GroupName
	UserOrGroupForUpdate.CompanyId = m.Info.CompanyTeamName
	UserOrGroupForUpdate.groupId=groupUniqueID

	dataValue := reflect.ValueOf(m.Members)
	UserOrGroup.groupId = groupUniqueID
	for _, key := range dataValue.MapKeys() {
		if UserOrGroupForUpdate.GroupName != "" &&UserOrGroupForUpdate.CompanyId !=""{
			err = db.Child("/Users/" + key.String() + "/Group/" + groupUniqueID).Set(UserOrGroup)
			UserGroupKey=append(UserGroupKey,"true")
			if err != nil {
				log.Println("w16")
				log.Println("Insertion error:", err)
				return false
			}
		}
	}

	if len(UserGroupKey) !=len(m.Members){
		log.Println("danger111111")
		dataValue := reflect.ValueOf(m.Members)
		for _, key := range dataValue.MapKeys() {
			err = db.Child("/Users/" + key.String() + "/Group/" + groupUniqueID).Value(&UserStatus)
			if len(UserStatus.GroupName) ==0{
				err = db.Child("/Users/" + key.String() + "/Group/" + groupUniqueID).Set(UserOrGroupForUpdate)
			}


		}

	}

	return  true
}

//Fetch all the details of group
func GetAllGroupDetails(ctx context.Context,companyTeamName string) (map[string]Group,bool){
	//user := User{}
	log.Println("w12")
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]Group{}
	err = db.Child("Group").OrderBy("Info/CompanyTeamName").EqualTo(companyTeamName).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value,false
	}
	return value,true


}

// Delete each group using group id
func (n *Group)DeleteGroup(ctx context.Context, groupKey string) bool{

	GroupDeletion := GroupSettings{}
	groupStatusUpdate := GroupSettings{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ groupKey +"/Settings/").Value(&groupStatusUpdate)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	GroupDeletion.Status = helpers.UserStatusDeleted
	GroupDeletion.DateOfCreation = groupStatusUpdate.DateOfCreation
	log.Println("delete",GroupDeletion)
	err = db.Child("/Group/"+ groupKey +"/Settings").Update(&GroupDeletion)
	if err != nil {
		log.Fatal(err)
		return  false
	}

	return  true

}

// To get all the keys of User
func GetUsersForDropdown(ctx context.Context,companyId string) (UsersCompany,bool) {
	db,err :=GetFirebaseClient(ctx,"")
	allUser := UsersCompany{}
	err = db.Child("Company/"+companyId+"/Users").Value(&allUser)
	if err != nil {
		log.Fatal(err)
		return allUser,false
	}
	err = db.Child("Users").Value(&allUser)
	log.Println("vvvv",allUser)
	if err != nil {
		log.Println(err)
		return allUser,false
	}
	return allUser,true
}

// for fill the dropdown list using name(users) in add group
func(m *Users) TakeGroupMemberName(ctx context.Context,companyTeamName string) ( map[string]CompanyUsers,bool) {
	log.Println("cp1")
	allUserDetails :=map[string]CompanyUsers{}

	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return allUserDetails,false
	}
	err = db.Child("Company/"+companyTeamName+"/Users").Value(&allUserDetails)
	if err != nil{
		log.Println("danger3")
		log.Fatal(err)
		return allUserDetails,false
	}

	return allUserDetails,true
}

//Collecting Group details using Id
func(m *Group) GetGroupDetailsById(ctx context.Context,groupKey string) (Group,bool){
	groupDetails :=  Group{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+groupKey).Value(&groupDetails)
	if err != nil {
		log.Fatal(err)
		return groupDetails, false
	}
	return groupDetails,true
}

//Update the group details after editing
func(m *Group) UpdateGroupDetails(ctx context.Context,groupKey string) (bool) {

	UserOrGroup :=UserGroup{}
	UserStatus :=UserGroup{}
	UserForUpdate :=UserGroup{}
	oldMembers := map[string]GroupMembers{}
	groupStatusDetails := GroupSettings{}
	db,err :=GetFirebaseClient(ctx,"")

	err = db.Child("/Group/"+ groupKey +"/Members/").Value(&oldMembers)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	var oldKeySlice []string
	OldUserKey := reflect.ValueOf(oldMembers)
	for _, oldKey := range OldUserKey.MapKeys(){
		oldKeySlice = append(oldKeySlice,oldKey.String())
	}


	err = db.Child("/Group/"+groupKey+"/Settings").Value(&groupStatusDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	m.Settings.DateOfCreation = groupStatusDetails.DateOfCreation
	m.Settings.Status = groupStatusDetails.Status
	err = db.Child("/Group/"+ groupKey).Set(&m)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	var newKeySlice []string
	memberData := reflect.ValueOf(m.Members)
	for _, key := range memberData.MapKeys(){
		newKeySlice = append(newKeySlice,key.String())

	}
	log.Println("newKeySlice",newKeySlice)
	log.Println("oldKeySlice",oldKeySlice)

	// enter group values into user in case of new users
	newUser := make([]string, 0)
	s_one := newKeySlice
	s_two := oldKeySlice
	for _, s := range s_one {
		if !inslice(s, s_two) {
			newUser = append(newUser, s)
		}
	}
	log.Println("result",newUser)
	UserOrGroup.CompanyId = m.Info.CompanyTeamName
	UserOrGroup.GroupName = m.Info.GroupName
	UserForUpdate.groupId=groupKey
	UserForUpdate.CompanyId=m.Info.CompanyTeamName
	UserForUpdate.GroupName =m.Info.GroupName

	var NewUserUpdateCount []string
	for i := 0;i<len(newUser);i++ {
		log.Println("new usersss",newUser[i])
		if UserOrGroup.CompanyId !="" &&UserForUpdate.GroupName !=""{
			err = db.Child("/Users/" +newUser[i] + "/Group/" + groupKey).Set(UserOrGroup)
			NewUserUpdateCount=append(NewUserUpdateCount,"true")

		}

		if err != nil {
			log.Println("w16")
			log.Println("Insertion error:", err)
			return false
		}
	}

	//delete GroupDetails from User when delete old user in updation

	oldUser := make([]string,0)
	for _,s := range oldKeySlice{
		if !inslice(s, newKeySlice) {
			oldUser = append(oldUser, s)
		}
	}
	var OldUserUpdateCount []string
	for i := 0;i<len(oldUser);i++ {
		log.Println("old usersssss",oldUser[i])
		err = db.Child("/Users/" +oldUser[i] + "/Group/" + groupKey).Remove()
		if err != nil {
			log.Println("w16")
			log.Println("Insertion error:", err)
			return false
		}
		err = db.Child("/Users/"+oldUser[i] +"/Settings/Notifications/GroupChat/"+groupKey).Remove()
		if err != nil {
			log.Println("w16")
			log.Println("Insertion error:", err)
			return false
		}
		OldUserUpdateCount=append(OldUserUpdateCount,"true")
	}

	/*for i := 0;i<len(newKeySlice);i++{
		for j := 0;j<len(oldKeySlice); j++ {
			if newKeySlice[i] != oldKeySlice[j]{
				log.Println("newKeySlice[i]",newKeySlice[i])
				//err = db.Child("Users/"+newKeySlice[i]+"/Group/"+groupKey).Remove()
			}
		}

	}*/
	//.....update in task


	//groupDetailForUpdation := map[string]Tasks{}
	//taskGroupForUpdate :=TaskGroup{}
	//taskGroupDetail :=TaskGroup{}
	//
	//err = db.Child("/Tasks/").Value(&groupDetailForUpdation)
	//dataValue := reflect.ValueOf(groupDetailForUpdation)
	//for _, key := range dataValue.MapKeys() {
	//	log.Println("hhhh")
	//	dataValueContact := reflect.ValueOf(groupDetailForUpdation[key.String()].UsersAndGroups.Group)
	//	for _, groupkey := range dataValueContact.MapKeys() {
	//		if  groupkey.String()== groupKey {
	//			log.Println("task id",key.String())
	//			err = db.Child("Tasks/" + key.String() + "/UsersAndGroups/Group/"+groupKey).Value(&taskGroupDetail)
	//			log.Println("contact inside task",taskGroupDetail)
	//			taskGroupForUpdate.Members = m.Members
	//			taskGroupForUpdate.GroupName = m.Info.GroupName
	//			taskGroupForUpdate.GroupStatus = taskGroupDetail.GroupStatus
	//			err = db.Child("Tasks/" + key.String() + "/UsersAndGroups/Group/"+groupKey).Update(&taskGroupForUpdate)
	//
	//		}
	//	}
	//}
	if err != nil {
		log.Fatal(err)
		return  false
	}



	//in case of missing users


	if len(NewUserUpdateCount)!=len(newUser){
		for i :=0;i<len(newUser);i++{

			err = db.Child("/Users/" +newUser[i] + "/Group/" + groupKey).Value(&UserStatus)
			if len(UserStatus.GroupName )==0&&len(UserStatus.CompanyId )==0{
				err = db.Child("/Users/" +newUser[i] + "/Group/" + groupKey).Set(UserForUpdate)
			}

		}
	}

	if len(OldUserUpdateCount)!=len(oldUser){
		for i:=0;i<len(oldUser);i++{
			err = db.Child("/Users/" +oldUser[i] + "/Group/" + groupKey).Remove()
			if err != nil {
				log.Println("w16")
				log.Println("Insertion error:", err)
				return false
			}
			err = db.Child("/Users/"+oldUser[i] +"/Settings/Notifications/GroupChat/"+groupKey).Remove()
			if err != nil {
				log.Println("w16")
				log.Println("Insertion error:", err)
				return false
			}

		}
	}
	return true

}

func inslice(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}

//check group name is already exist
func IsGroupNameUsed(ctx context.Context,groupName string)(bool) {
	groupDetails := map[string]Group{}

	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Child("Group").OrderBy("Info/GroupName").EqualTo(groupName).Value(&groupDetails)
	if err != nil {
		log.Fatal(err)
	}
	if len(groupDetails) == 0 {

		return true
	} else {
		dataValue := reflect.ValueOf(groupDetails)
		for _, key := range dataValue.MapKeys() {
			if groupDetails[key.String()].Settings.Status == helpers.StatusActive {
				return false
			}
		}
	}
	return true
}
func (m *Group) DeleteGroupFromDB(ctx context.Context, groupId string,TaskSlice []string)(bool)  {
	GroupMembers := map[string]GroupMembers{}
	groupDetailForUpdate :=TasksGroup{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	groupDetailForUpdate.TasksGroupStatus =helpers.StatusInActive
	for i:=0;i<len(TaskSlice);i++{
		log.Println(TaskSlice[i])
		err = dB.Child("/Group/"+ groupId+"/Tasks/"+TaskSlice[i]).Set(&groupDetailForUpdate)

	}
	groupSettingsUpdation := GroupSettings{}
	groupDeletion := GroupSettings{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ groupId+"/Settings").Value(&groupSettingsUpdation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	groupDeletion.Status = helpers.UserStatusDeleted
	groupDeletion.DateOfCreation = groupSettingsUpdation.DateOfCreation
	err = db.Child("Group/"+groupId+"/Settings").Update(&groupDeletion)
	if err != nil {
		log.Fatal(err)
		return  false
	}

	err = db.Child("/Group/"+ groupId +"/Members/").Value(&GroupMembers)
	if err != nil {
		log.Fatal(err)
		return  false
	}

	memberData := reflect.ValueOf(GroupMembers)
	for _, key := range memberData.MapKeys(){
		err = db.Child("Users/"+key.String()+"/Group/"+groupId).Remove()
		err = db.Child("/Users/"+key.String()+"/Settings/Notifications/GroupChat/"+groupId).Remove()
	}

	//taskGroupDetail :=TaskGroup{}
	//taskGroupForUpdate :=TaskGroup{}
	//for i:=0;i<len(TaskSlice);i++ {
	//	err = dB.Child("Tasks/" + TaskSlice[i]+"/UsersAndGroups/Group/"+groupId).Value(&taskGroupDetail)
	//	log.Println("details from task job",)
	//	taskGroupForUpdate.GroupStatus =helpers.StatusInActive
	//	taskGroupForUpdate.GroupName=taskGroupDetail.GroupName
	//	taskGroupForUpdate.Members =taskGroupDetail.Members
	//
	//	log.Println("fhsgjs",taskGroupForUpdate)
	//	err = dB.Child("Tasks/" + TaskSlice[i]+"/UsersAndGroups/Group/"+groupId).Update(&taskGroupForUpdate)
	//
	//}
	if err!=nil{
		log.Println("Deletion error:",err)
		return false
	}
	return true
}
func (m *TasksGroup) IsGroupUsedForTask( ctx context.Context, groupId string)(bool,map[string]TasksGroup)  {
	groupDetail := map[string]TasksGroup{}
	dB, err := GetFirebaseClient(ctx,"")
	if err!=nil{
		log.Println("Connection error:",err)
	}
	err = dB.Child("/Group/"+ groupId+"/Tasks/").Value(&groupDetail)
	if err!=nil{
		log.Println("Insertion error:",err)
		return false,groupDetail
	}
	log.Println("group detail",groupDetail)

	return true,groupDetail
}
func(m *Group) DeleteGroupFromDBForNonTask(ctx context.Context,groupId string) bool{
	groupSettingsUpdation := GroupSettings{}
	groupDeletion := GroupSettings{}
	GroupMembers := map[string]GroupMembers{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ groupId+"/Settings").Value(&groupSettingsUpdation)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	groupDeletion.Status = helpers.UserStatusDeleted
	groupDeletion.DateOfCreation = groupSettingsUpdation.DateOfCreation
	err = db.Child("Group/"+groupId+"/Settings").Update(&groupDeletion)
	if err != nil {
		log.Fatal(err)
		return  false
	}

	err = db.Child("/Group/"+ groupId +"/Members/").Value(&GroupMembers)
	if err != nil {
		log.Fatal(err)
		return  false
	}

	memberData := reflect.ValueOf(GroupMembers)
	log.Println("memberData",memberData)
	for _, key := range memberData.MapKeys(){
		err = db.Child("Users/"+key.String()+"/Group/"+groupId).Remove()
		err = db.Child("/Users/"+key.String()+"/Settings/Notifications/GroupChat/"+groupId).Remove()

	}
	return  true
}



