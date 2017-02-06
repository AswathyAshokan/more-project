/*Created By Farsana*/

package models
import (
	"golang.org/x/net/context"
	"log"
	//"app/passporte/models"
)
type Group struct {

	GroupName string
	GroupMembers string
}
type UserInformation struct {
	Email string
	UserName string
}
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
	v := map[string]Group{}
	err = db.Child("Group").Value(&v)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("%s\n", v)
	//log.Println(reflect.TypeOf(v))
	return v


}
func(this *Group) DeleteGroup(ctx context.Context, groupKey string) bool{
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+ groupKey).Remove()
	if err != nil {
		log.Fatal(err)
		return  false
	}
	return  true

}

// for fill the dropdown list in add group
func(this *UserInformation) GetUsersForDropdown(ctx context.Context) map[string]UserInformation {
	db,err :=GetFirebaseClient(ctx,"")
	v := map[string]UserInformation{}
	err = db.Child("Users").Value(&v)
	if err != nil {
		log.Fatal(err)
	}
	return v


}


func(this *UserInformation) TakeGroupMemberName(ctx context.Context,groupKeySlice []string)UserInformation {
	//user := User{}
	db,err :=GetFirebaseClient(ctx,"")
	v :=UserInformation{}
	for i := 0; i <len(groupKeySlice) ; i++ {
		err = db.Child("/Users/"+groupKeySlice[i]).Child("Info").Value(&v)
	}

	if err != nil {
		log.Fatal(err)
	}

	return v
}



