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




