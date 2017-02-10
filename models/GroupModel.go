/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
)
type Group struct {

	GroupName        	string
	GroupMembers	 	[]string
}
/*type UserInformation struct {
	FirstName string

}*/
func(this *Group) AddGroupToDb(ctx context.Context) (bool){
	//log.Println("values in model",this)
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("Group").Push(this)

	if err != nil {
		log.Println(err)
		return false
	}

	return  true
}

func(this *Group) DisplayGroup(ctx context.Context) map[string]Group{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	value := map[string]Group{}
	err = db.Child("Group").Value(&value)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("%s\n", v)
	//log.Println(reflect.TypeOf(v))
	return value


}
func(this *Group) DeleteGroup(ctx context.Context, GroupKey string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ GroupKey).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

// for fill the dropdown list in add group
func(this *Group) GetUsersForDropdown(ctx context.Context) map[string]InviteUser {
	db,err :=GetFirebaseClient(ctx,"")
	v := map[string]InviteUser{}
	//err = db.Child("Users").Value(&v)
	err = db.Child("User").Value(&v)
	if err != nil {
		log.Fatal(err)
	}
	return v


}


func(this *Group) TakeGroupMemberName(ctx context.Context,groupKeySlice []string) ([]string, bool) {
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

//editing purpose

func(this *Group) EditGroupDetails(ctx context.Context,groupKey string) (Group,bool){

	value := Group{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+groupKey).Value(&value)
	if err != nil {
		log.Fatal(err)
		return value , false
	}
	return value,true

}

func(this *Group) UpdateGroupDetails(ctx context.Context,groupKey string) (bool) {


	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ groupKey).Update(&this)

	if err != nil {
		log.Fatal(err)
		return  false
	}
	return true

}


