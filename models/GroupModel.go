/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
	"app/passporte/helpers"
	"reflect"
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
	_,err = db.Child("Group").Push(m)

	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}

//Fetch all the details of group
func GetAllGroupDetails(ctx context.Context,companyTeamName string) (map[string]Group,bool){
	//user := User{}
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
	allUserDetails :=map[string]CompanyUsers{}

	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return allUserDetails,false
	}
	err = db.Child("Company/"+companyTeamName+"/Users").Value(&allUserDetails)
	if err != nil{
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

	groupStatusDetails := GroupSettings{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+groupKey+"/Settings").Value(&groupStatusDetails)
	if err != nil {
		log.Fatal(err)
		return  false
	}
	m.Settings.DateOfCreation = groupStatusDetails.DateOfCreation
	m.Settings.Status = groupStatusDetails.Status
	err = db.Child("/Group/"+ groupKey).Update(&m)


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
	return true

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

	groupDetailForUpdate :=TasksGroup{}
	dB, err := GetFirebaseClient(ctx,"")

	if err!=nil{
		log.Println("Connection error:",err)
	}
	groupDetailForUpdate.TasksGroupStatus =helpers.StatusInActive
	for i:=0;i<len(TaskSlice);i++{
		log.Println(TaskSlice[i])
		err = dB.Child("/Group/"+ groupId+"/Tasks/"+TaskSlice[i]).Update(&groupDetailForUpdate)

	}
	taskGroupDetail :=TaskGroup{}
	taskGroupForUpdate :=TaskGroup{}
	for i:=0;i<len(TaskSlice);i++ {
		err = dB.Child("Tasks/" + TaskSlice[i]+"/UsersAndGroups/Group/"+groupId).Value(&taskGroupDetail)
		log.Println("details from task job",)
		taskGroupForUpdate.GroupStatus =helpers.StatusInActive
		taskGroupForUpdate.GroupName=taskGroupDetail.GroupName
		taskGroupForUpdate.Members =taskGroupDetail.Members

		log.Println("fhsgjs",taskGroupForUpdate)
		err = dB.Child("Tasks/" + TaskSlice[i]+"/UsersAndGroups/Group/"+groupId).Update(&taskGroupForUpdate)

	}
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
	log.Println("id",groupId)
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
	return  true
}



