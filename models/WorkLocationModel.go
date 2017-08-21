package models
import (
	"log"
	"golang.org/x/net/context"

)
type WorkLocation struct {
	Info 		WorkLocationInfo
	Settings 	WorkLocationSettings
}
type WorkLocationInfo struct {
	WorkLocation       	string
	UsersAndGroupsInWorkLocation		UsersAndGroupsInWork
}

type WorkLocationSettings struct {
	DateOfCreation  	int64
	Status         	 	string
}
type UsersAndGroupsInWork struct {
	User 		map[string]WorkLocationUser
	Group 		map[string]WorkLocationGroup

}
type WorkLocationUser struct {
	FullName	string
	Status		string
}
type WorkLocationGroup struct{
	GroupName	string
	Members	 	map[string]GroupMemberNameInWorkLocation
}
type  GroupMemberNameInWorkLocation struct {
	MemberName	string

}


func(m *WorkLocation) AddWorkLocationToDb(ctx context.Context) (bool){
	log.Println("add group")
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Child("WorkLocation").Push(m)

	if err != nil {
		log.Println(err)
		return false
	}
	return  true
}


func(m *Group) GetGroupDetailsForWorkLocation(ctx context.Context,groupKey string) (Group,bool){
	log.Println("cp4")
	groupDetails :=  Group{}
	db,err :=GetFirebaseClient(ctx,"")
	err = db.Child("/Group/"+groupKey).Value(&groupDetails)
	if err != nil {
		log.Println("cp5")
		log.Fatal(err)
		return groupDetails, false
	}
	log.Println("cp6",groupDetails)
	return groupDetails,true
}

