package models
import (
	"golang.org/x/net/context"
	"log"
	//"strings"
	//"reflect"
	//
	//"app/passporte/helpers"
	//"strconv"

)

type TimeSheet struct{
	DailyEndTime  	int64
	DailyStartTime	int64
	Date		string
	LogTime		int64
	TaskEndTime     map[string]TaskEndTimeForTimeSheet
	TaskId		string
	TaskName	string
	TaskStartTime	map[string]TaskStartTimeForTimeSheet
	UserName	string
	WorkId		string
	WorkStartTime   map[string]WorkStartTimeForTimeSheet
	WorkEndTime     map[string]WorkEndTimeForTimeSheet
	TaskDateFrom	int64
	TaskDateTo 	int64
	WorkLocation    string



}
type TaskEndTimeForTimeSheet struct{
	Time		int64
}
type TaskStartTimeForTimeSheet struct{
	Time		int64
}
type WorkStartTimeForTimeSheet struct{
	Time		int64
}
type WorkEndTimeForTimeSheet struct{
	Time		int64
}


func(m *Users) GetAllUsers(ctx context.Context) (bool,map[string]Users){
	userDetails := map[string]Users{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("Users").Value(&userDetails)
	//if err != nil {
	//	log.Println(err)
	//	return  false,userDetails
	//}
	return true,userDetails


}
func(m *TimeSheet) RetrieveTimeSheetFromDB(ctx context.Context,companyTeamName string) (bool,map[string]TimeSheet){
	timeSheetDetails := map[string]TimeSheet{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	err = db.Child("TimeSheet/"+companyTeamName).Value(&timeSheetDetails)
	//if err != nil {
	//	log.Println(err)
	//	return  false,userDetails
	//}
	return true,timeSheetDetails

}
func(m *TimeSheet) RetrieveTimeSheetUserFromDB(ctx context.Context,companyTeamName string,userId string) (bool,map[string]TimeSheet){
	timeSheetDetails := map[string]TimeSheet{}
	db,err :=GetFirebaseClient(ctx,"")
	if err != nil {
		log.Println(err)
	}
	log.Println("user idddddd",userId)
	err = db.Child("TimeSheet/"+companyTeamName+"/"+userId).Value(&timeSheetDetails)
	//if err != nil {
	//	log.Println(err)
	//	return  false,userDetails
	//}
	return true,timeSheetDetails

}
