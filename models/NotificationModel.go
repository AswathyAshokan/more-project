package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"

	"strconv"
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
	dB, err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false
	}
	//err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Value(&notificationValue)
	err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Remove()
	log.Println("deleted sucessfully")
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetAllNotificationsOfExpiration(ctx context.Context,companyTeamName string) (bool,[][]string) {

	//companyDataOfExpiry := map[string]CompanyData{}
	expiryNotification := map[string]ExpiryNotification{}
	expiryDetails := map[string]Expirations{}
	CompanyUsers := map[string]CompanyUsers{}
	typeOfExpiryData :=  ExpirationInfo{}
	var tempArray [][]string
	db, err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false,tempArray
	}
	err = db.Child("Company/"+companyTeamName+"/Users/").Value(&CompanyUsers)
	if err != nil {
		log.Fatal(err)
		return false,tempArray
	}
	dataValue := reflect.ValueOf(CompanyUsers)
	var keySlice []string
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice,key.String())
	}
	var fullName string
	for i:= 0;i<len(keySlice);i++{

		err = db.Child("/Expirations/"+keySlice[i]).Value(&expiryDetails)
		expiryDataValues := reflect.ValueOf(expiryDetails)

		for _, k := range expiryDataValues.MapKeys() {
			err = db.Child("Users/"+keySlice[i]+"/Info/FullName").Value(&fullName)
			err = db.Child("/Expirations/"+keySlice[i]+"/"+k.String()+"/Info").Value(&typeOfExpiryData)
			err = db.Child("/Expirations/"+keySlice[i]+"/"+k.String()+"/Company/"+companyTeamName+"/NotificationShedules/").Value(&expiryNotification)
			eachExpirationValues := reflect.ValueOf(expiryNotification)
			for _, eachkey := range eachExpirationValues.MapKeys() {
				var expiryNotificationArray []string
				expiryNotificationArray = append(expiryNotificationArray, expiryNotification[eachkey.String()].Category)
				expiryNotificationArray = append(expiryNotificationArray, expiryNotification[eachkey.String()].ExpiryId)
				str := strconv.FormatBool(expiryNotification[eachkey.String()].IsRead)
				expiryNotificationArray = append(expiryNotificationArray,str )
				expiryNotificationArray = append(expiryNotificationArray, expiryNotification[eachkey.String()].LocalDate)
				unixDate := strconv.FormatInt(int64(expiryNotification[eachkey.String()].NotificationDate), 10)
				expiryNotificationArray = append(expiryNotificationArray, unixDate)
				expiryNotificationArray = append(expiryNotificationArray,fullName)
				expiryNotificationArray = append(expiryNotificationArray,typeOfExpiryData.Type)

				tempArray = append(tempArray,expiryNotificationArray)
				expiryNotificationArray = expiryNotificationArray[:0]
			}

		}


		log.Println("nnnnnnn",expiryNotification)

	}
	return true,tempArray
}