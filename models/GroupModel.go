/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
)
type Group struct {

	GroupName       string
	DateOfCreation  int64
	Status          string
	Members	 	[]GroupMembers
}

type GroupMembers struct {
	MemberId	string
	MemberName	string
}

/*type UserInformation struct {
	FirstName string

}*/

// Insert new groups to database
func(m *Group) AddGroupToDb(ctx context.Context) (bool){
	//log.Println("values in model",m)
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
func GetAllGroupDetails(ctx context.Context) (map[string]Group,bool){
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]Group{}
	err = db.Child("Group").Value(&value)
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
func(m *Group) GetUsersForDropdown(ctx context.Context) (map[string]InviteUser,bool) {
	db,err :=GetFirebaseClient(ctx,"")
	allUser := map[string]InviteUser{}
	//err = db.Child("Users").Value(&v)
	err = db.Child("User").Value(&allUser)
	if err != nil {
		log.Fatal(err)
		return allUser,false
	}
	return allUser,true


}

// for fill the dropdown list using first name and second name(user) in add group
func(m *Group) TakeGroupMemberName(ctx context.Context,groupKeySlice []string) ([]string, bool) {
	allUserDetails := InviteUser{}
	var allUserNames [] string
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return allUserNames, false
	}

	for i := 0; i <len(groupKeySlice); i++ {
		//err = db.Child("/Users/"+groupKeySlice[i]).Child("Info").Value(&v)
		err = db.Child("/User/"+groupKeySlice[i]).Value(&allUserDetails)
		if err != nil{
			log.Fatal(err)
			return allUserNames, false
		}
		allUserNames = append(allUserNames, (allUserDetails.FirstName + " " + allUserDetails.LastName))

	}
	return allUserNames, true
}

//Collecting Group details using Id
func(m *Group) GetGroupDetailsById(ctx context.Context,groupKey string) (Group,bool){
	groupDetails := Group{}
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


