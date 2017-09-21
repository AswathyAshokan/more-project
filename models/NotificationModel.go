package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"

)
type Notification struct {
	Date		int64
	IsRead		bool
	IsSentMail	bool
	Message		string
	TaskId		string
	TaskLocation	string
	TaskName	string
	UserName	string
}

func GetAllNotifications(ctx context.Context,companyTeamName string)(bool,map[string]Notification) {
	notificationValue := map[string]Notification{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Value(&notificationValue)
	if err != nil {
		log.Fatal(err)
		return false, notificationValue
	}
	return true, notificationValue
}
func GetAllNotificationsOfUser(ctx context.Context,companyTeamName string,userKey string)(bool,map[string]Notification) {
	notificationValue := map[string]Notification{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+userKey).Value(&notificationValue)
	if err != nil {
		log.Fatal(err)
		return false, notificationValue
	}
	return true, notificationValue
}

func UpdateAllNotifications(ctx context.Context,companyTeamName string)(bool) {
	notificationValue := map[string]Notification{}
	notificationBeforeUpdate := map[string]Notification{}
	notificationUpdate :=Notification{}
	notificationUpdateSuccess :=Notification{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Value(&notificationValue)
	notificationOfUser := reflect.ValueOf(notificationValue)
	for _, notificationUserKey := range notificationOfUser.MapKeys() {
		err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()).Value(&notificationBeforeUpdate)
		notifications := reflect.ValueOf(notificationBeforeUpdate)
		for _, notificationKey := range notifications.MapKeys() {
			err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Value(&notificationUpdate)
			notificationUpdateSuccess.Date=notificationUpdate.Date
			notificationUpdateSuccess.TaskName=notificationUpdate.TaskName
			notificationUpdateSuccess.TaskLocation=notificationUpdate.TaskLocation
			notificationUpdateSuccess.IsRead=true
			notificationUpdateSuccess.IsSentMail=notificationUpdate.IsSentMail
			notificationUpdateSuccess.Message=notificationUpdate.Message
			notificationUpdateSuccess.UserName=notificationUpdate.UserName
			notificationUpdateSuccess.TaskId=notificationUpdate.TaskId
			log.Println("loollllll",notificationUpdateSuccess);
			err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Set(notificationUpdateSuccess)


		}
	}
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}



func DeleteAllNotifications(ctx context.Context,companyTeamName string)(bool) {
	//notificationValue := map[string]Notification{}
	//notificationBeforeUpdate := map[string]Notification{}

	dB, err := GetFirebaseClient(ctx,"")
	//err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Value(&notificationValue)
	err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Remove()

	//notificationOfUser := reflect.ValueOf(notificationValue)
	//for _, notificationUserKey := range notificationOfUser.MapKeys() {
	//	err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()).Value(&notificationBeforeUpdate)
	//	notifications := reflect.ValueOf(notificationBeforeUpdate)
	//	for _, notificationKey := range notifications.MapKeys() {
	//
	//		err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Remove()
	//
	//
	//	}
	//}
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}