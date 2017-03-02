/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
)
type Group struct {
	Info 		GroupInfo
	Members	 	map[string]GroupMembers
	Settings 	GroupSettings
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
	//log.Println("%s\n", v)
	//log.Println(reflect.TypeOf(v))
	return value,true


}

// Delete each group using group id
func(m *Group) DeleteGroup(ctx context.Context, GroupKey string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ GroupKey).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

// To get all the keys of User
func (m *Users)GetUsersForDropdown(ctx context.Context) (map[string]Users,bool) {
	db,err :=GetFirebaseClient(ctx,"")
	allUser := map[string]Users{}
	err = db.Child("Users").Value(&allUser)
	//err = db.Child("Users").Value(&allUser)
	if err != nil {
		log.Println(err)
		return allUser,false
	}
	return allUser,true


}

// for fill the dropdown list using name(users) in add group
func(m *Users) TakeGroupMemberName(ctx context.Context,groupKeySlice []string) ([]string, bool) {
	allUserDetails :=Users{}
	var allUserNames [] string
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return allUserNames, false
	}

	for i := 0; i <len(groupKeySlice); i++ {
		//err = db.Child("/Users/"+groupKeySlice[i]).Child("Info").Value(&v)
		err = db.Child("/Users/"+groupKeySlice[i]).Value(&allUserDetails)
		if err != nil{
			log.Fatal(err)
			return allUserNames, false
		}
		allUserNames = append(allUserNames, (allUserDetails.Info.FullName))

	}
	return allUserNames, true
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


	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ groupKey).Update(&m)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}

//check group name is already exist
func IsGroupNameUsed(ctx context.Context,groupName string)(bool) {
	log.Println("groupName",groupName)
	groupDetails := map[string]Group{}
	db, err := GetFirebaseClient(ctx, "")
	if err != nil {
		log.Println("No Db Connection!")
	}
	err = db.Child("Group").OrderBy("Info/GroupName").EqualTo(groupName).Value(&groupDetails)
	if err != nil {
		log.Fatal(err)
	}
	if len(groupDetails)==0{
		log.Println("map null:",groupDetails)
		return true
	}else{
		log.Println("map not null:",groupDetails)
		return false
	}

}


