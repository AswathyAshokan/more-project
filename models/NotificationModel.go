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
	Category	string
	WorkLocation	string
	WorkId 		string
}

type NotificationForLeave struct {
	EndDate        int64
	IsRead        bool
	LogTime        int64
	NumberOfDays    int64
	StartDate    int64
	UserName    string

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

func UpdateAllNotifications(ctx context.Context,companyTeamName string,UpdateIdArray []string,expiryId []string,userId []string)(bool) {

	upDateReadStatus := ExpiryNotification{}
	oldReadStatus :=map[string]ExpiryNotification{}
	expiryDetails := map[string]Expirations{}
	//CompanyUsers := map[string]CompanyUsers{}
	notificationValue := map[string]Notification{}
	notificationBeforeUpdate := map[string]Notification{}
	notificationUpdate :=Notification{}
	notificationUpdateSuccess :=Notification{}



	notificationValueForLeave := map[string]NotificationForLeave{}
	notificationBeforeUpdateForLeave := map[string]NotificationForLeave{}
	notificationUpdateForLeave :=NotificationForLeave{}
	notificationUpdateSuccessForLeave :=NotificationForLeave{}


	dB, err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Value(&notificationValue)
	notificationOfUser := reflect.ValueOf(notificationValue)
	for _, notificationUserKey := range notificationOfUser.MapKeys() {
		err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()).Value(&notificationBeforeUpdate)
		notifications := reflect.ValueOf(notificationBeforeUpdate)
		for _, notificationKey := range notifications.MapKeys() {
			err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Value(&notificationUpdate)
			if notificationUpdate.Category =="WorkLocation"{
				notificationUpdateSuccess.Date=notificationUpdate.Date
				notificationUpdateSuccess.TaskName=notificationUpdate.WorkLocation
				notificationUpdateSuccess.TaskLocation=notificationUpdate.WorkId
				notificationUpdateSuccess.IsRead=true
				notificationUpdateSuccess.IsSentMail=notificationUpdate.IsSentMail
				notificationUpdateSuccess.Message=notificationUpdate.Message
				notificationUpdateSuccess.UserName=notificationUpdate.UserName
				notificationUpdateSuccess.Category = notificationUpdate.Category
			} else {
				notificationUpdateSuccess.Date=notificationUpdate.Date
				notificationUpdateSuccess.TaskName=notificationUpdate.TaskName
				notificationUpdateSuccess.TaskLocation=notificationUpdate.TaskLocation
				notificationUpdateSuccess.IsRead=true
				notificationUpdateSuccess.IsSentMail=notificationUpdate.IsSentMail
				notificationUpdateSuccess.Message=notificationUpdate.Message
				notificationUpdateSuccess.UserName=notificationUpdate.UserName
				notificationUpdateSuccess.TaskId=notificationUpdate.TaskId
				notificationUpdateSuccess.Category = notificationUpdate.Category
			}
			err = dB.Child("Notifications/UserDelay/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Set(notificationUpdateSuccess)


		}
	}




	err = dB.Child("Notifications/LeaveRequests/"+ companyTeamName).Value(&notificationValueForLeave)
	notificationOfUserForLeave := reflect.ValueOf(notificationValueForLeave)
	for _, notificationUserKey := range notificationOfUserForLeave.MapKeys() {
		err = dB.Child("Notifications/LeaveRequests/"+ companyTeamName+"/"+notificationUserKey.String()).Value(&notificationBeforeUpdateForLeave)
		notifications := reflect.ValueOf(notificationBeforeUpdateForLeave)
		for _, notificationKey := range notifications.MapKeys() {
			err = dB.Child("Notifications/LeaveRequests/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Value(&notificationUpdateForLeave)
			notificationUpdateSuccessForLeave.NumberOfDays=notificationUpdateForLeave.NumberOfDays
			notificationUpdateSuccessForLeave.LogTime=notificationUpdateForLeave.LogTime
			notificationUpdateSuccessForLeave.EndDate=notificationUpdateForLeave.EndDate
			notificationUpdateSuccessForLeave.IsRead=true
			notificationUpdateSuccessForLeave.StartDate=notificationUpdateForLeave.StartDate
			notificationUpdateSuccessForLeave.UserName=notificationUpdateForLeave.UserName
			err = dB.Child("Notifications/LeaveRequests/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Set(notificationUpdateSuccessForLeave)


		}
	}


	for i :=0;i<len(userId);i++{
		err = dB.Child("/Expirations/" + userId[i]).Value(&expiryDetails)
		log.Println("expiryDetails",expiryDetails)
		for j:=0;j<len(expiryId);j++{
			err = dB.Child("/Expirations/"+userId[i]+"/"+expiryId[j]+"/Company/"+companyTeamName+"/NotificationShedules/").Value(&oldReadStatus)
			log.Println("oldReadStatus",oldReadStatus)
			log.Println("UpdateIdArray",UpdateIdArray)
			eachExpirationValues := reflect.ValueOf(oldReadStatus)
			for k:=0;k<len(UpdateIdArray);k++{
				for _, eachkey := range eachExpirationValues.MapKeys() {
					if UpdateIdArray[k] == eachkey.String(){
						upDateReadStatus.NotificationDate = oldReadStatus[UpdateIdArray[k]].NotificationDate
						upDateReadStatus.Category = oldReadStatus[UpdateIdArray[k]].Category
						upDateReadStatus.ExpiryId = oldReadStatus[UpdateIdArray[k]].ExpiryId
						upDateReadStatus.IsDeleted = oldReadStatus[UpdateIdArray[k]].IsDeleted
						upDateReadStatus.IsRead = true
						upDateReadStatus.LocalDate = oldReadStatus[UpdateIdArray[k]].LocalDate
						upDateReadStatus.IsViewed = oldReadStatus[UpdateIdArray[k]].IsViewed
						err = dB.Child("/Expirations/" + userId[i] + "/" + expiryId[j] + "/Company/" + companyTeamName + "/NotificationShedules/" +UpdateIdArray[k]).Set(upDateReadStatus)
						for i, v := range UpdateIdArray {
							if v == UpdateIdArray[k]  {
								UpdateIdArray =UpdateIdArray[:i+copy(UpdateIdArray[i:], UpdateIdArray[i+1:])]
								break
							}
						}


					}
				}

			}



		}

	}

	return true
}



func DeleteAllNotifications(ctx context.Context,companyTeamName string,UpdateIdArray []string,expiryId []string,userId []string)(bool) {

	oldReadStatus :=map[string]ExpiryNotification{}
	expiryDetails := map[string]Expirations{}
	dB, err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false
	}
	//err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Value(&notificationValue)
	err = dB.Child("Notifications/UserDelay/"+ companyTeamName).Remove()
	err = dB.Child("Notifications/LeaveRequests/"+ companyTeamName).Remove()
	log.Println("deleted sucessfully")
	if err != nil {
		log.Fatal(err)
		return false
	}

	for i :=0;i<len(userId);i++{
		err = dB.Child("/Expirations/" + userId[i]).Value(&expiryDetails)
		log.Println("expiryDetails",expiryDetails)
		for j:=0;j<len(expiryId);j++{
			err = dB.Child("/Expirations/"+userId[i]+"/"+expiryId[j]+"/Company/"+companyTeamName+"/NotificationShedules/").Value(&oldReadStatus)
			log.Println("oldReadStatus",oldReadStatus)
			log.Println("UpdateIdArray",UpdateIdArray)
			eachExpirationValues := reflect.ValueOf(oldReadStatus)
			for k:=0;k<len(UpdateIdArray);k++{
				for _, eachkey := range eachExpirationValues.MapKeys() {
					if UpdateIdArray[k] == eachkey.String(){
						err = dB.Child("/Expirations/" + userId[i] + "/" + expiryId[j] + "/Company/" + companyTeamName + "/NotificationShedules/" +UpdateIdArray[k]).Remove()
						for i, v := range UpdateIdArray {
							if v == UpdateIdArray[k]  {
								UpdateIdArray =UpdateIdArray[:i+copy(UpdateIdArray[i:], UpdateIdArray[i+1:])]
								break
							}
						}
					}
				}
			}
		}
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

			err = db.Child("/Expirations/"+keySlice[i]+"/"+k.String()+"/Company/"+companyTeamName+"/NotificationShedules/").Value(&expiryNotification)
			eachExpirationValues := reflect.ValueOf(expiryNotification)
			for _, eachkey := range eachExpirationValues.MapKeys() {

				err = db.Child("Users/"+keySlice[i]+"/Info/FullName").Value(&fullName)
				err = db.Child("/Expirations/"+keySlice[i]+"/"+k.String()+"/Info").Value(&typeOfExpiryData)


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
				expiryDate :=  strconv.FormatInt(int64(typeOfExpiryData.ExpirationDate), 10)
				expiryNotificationArray= append(expiryNotificationArray,expiryDate)
				expiryNotificationArray = append(expiryNotificationArray,eachkey.String())
				expiryNotificationArray = append(expiryNotificationArray,keySlice[i])
				//tempArray = append(tempArray,expiryNotificationArray)
				//expiryNotificationArray = expiryNotificationArray[:0]
				var condition=""
				if len(tempArray) == 0{

					tempArray = append(tempArray,expiryNotificationArray)
					condition ="true"
				}else {
					for i :=0;i<len(tempArray);i++{
						//log.Println("tempArray[i]",tempArray[i])
						for j:=0;j<len(tempArray[i]);j++{
							log.Println("date1",tempArray[i][4])
							log.Println("date2",unixDate)

							if tempArray[i][4] ==unixDate{
								condition ="true"
								//tempArray = append(tempArray,expiryNotificationArray)
								//break
								//log.Println("tempArray[i]",tempArray[i][4])
							}




						}
					}
				}
				log.Println("hhhhhhhhhhhhhh",condition)
				if condition ==""{
					tempArray = append(tempArray,expiryNotificationArray)
				}

				//expiryNotificationArray = expiryNotificationArray[:0]
			}

		}


		log.Println("nnnnnnn",tempArray)

	}
	return true,tempArray
}


//notification for leave

func GetAllNotificationsForLeave(ctx context.Context,companyTeamName string)(bool,map[string]NotificationForLeave) {
	notificationValue := map[string]NotificationForLeave{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Notifications/LeaveRequests/"+ companyTeamName).Value(&notificationValue)
	if err != nil {
		log.Fatal(err)
		return false, notificationValue
	}
	return true, notificationValue
}

func GetAllNotificationsOfUserForLeave(ctx context.Context,companyTeamName string,userKey string)(bool,map[string]NotificationForLeave) {
	notificationValue := map[string]NotificationForLeave{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Notifications/LeaveRequests/"+ companyTeamName+"/"+userKey).Value(&notificationValue)
	if err != nil {
		log.Fatal(err)
		return false, notificationValue
	}
	return true, notificationValue
}

