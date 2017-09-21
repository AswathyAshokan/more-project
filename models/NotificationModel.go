package models

import (
	"log"
	"golang.org/x/net/context"

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