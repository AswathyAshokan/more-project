package models

import (
	"log"
	"golang.org/x/net/context"
	"reflect"

	"strconv"
)
type NotificationForTask struct {
	Date		int64
	IsRead		bool
	IsSentMail	bool
	Message		string
	TaskId		string
	TaskLocation	string
	TaskName	string
	UserName	string
	Category	string
	Mode		string

}
type NotificationForWorkLocation struct {
	Category 	string
	Date		int64
	IsSentMail	bool
	Message		string
	WorkLocation	string
	WorkId 		string
	UserName	string
	IsRead		bool
	Mode  		string

}

type NotificationForLeave struct {
	EndDate        int64
	IsRead        bool
	LogTime        int64
	NumberOfDays    int64
	StartDate    int64
	UserName    string

}


func GetAllNotifications(ctx context.Context,companyTeamName string)(bool,map[string]NotificationForTask) {
	notificationValue := map[string]NotificationForTask{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Notifications/UserDelay/TaskLocation/"+companyTeamName).Value(&notificationValue)
	if err != nil {
		log.Fatal(err)
		//return false, notificationValue
	}
	return true, notificationValue
}

func GetAllWorkLocation(ctx context.Context,companyTeamName string)(bool,map[string]NotificationForWorkLocation){

	notification := map[string]NotificationForWorkLocation{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("/Notifications/UserDelay/WorkLocation/"+companyTeamName).Value(&notification)
	if err != nil {
		log.Fatal(err)
		//return false, notification
	}
	return true, notification
}

func GetAllNotificationsOfWorkLOcation(ctx context.Context,companyTeamName string,userKey string) (bool,map[string]NotificationForWorkLocation) {

	workLocationNotificationValues := map[string]NotificationForWorkLocation{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Notifications/UserDelay/WorkLocation/"+ companyTeamName+"/"+userKey).Value(&workLocationNotificationValues)
	if err !=nil{
		return false,workLocationNotificationValues
	}
	return true,workLocationNotificationValues

}

func GetAllNotificationsOfUser(ctx context.Context,companyTeamName string,userKey string)(bool,map[string]NotificationForTask) {
	notificationValue := map[string]NotificationForTask{}
	dB, err := GetFirebaseClient(ctx,"")
	err = dB.Child("Notifications/UserDelay/TaskLocation/"+ companyTeamName+"/"+userKey).Value(&notificationValue)
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

	oldWorkLocationNotificationValues := map[string]NotificationForWorkLocation{}


	updateTaskNotification :=NotificationForTask{}

	taskNoificationValues := map[string]NotificationForTask {}

	//updatedNotification := NotificationForTask {}





	notificationValueForLeave := map[string]NotificationForLeave{}
	notificationBeforeUpdateForLeave := map[string]NotificationForLeave{}
	notificationUpdateForLeave :=NotificationForLeave{}
	notificationUpdateSuccessForLeave :=NotificationForLeave{}

	dB, err := GetFirebaseClient(ctx,"")
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = dB.Child("/Notifications/UserDelay/TaskLocation/"+ companyTeamName).Value(&taskNoificationValues)
	notificationOfUserTask := reflect.ValueOf(taskNoificationValues)
	for _, taskUserKey := range notificationOfUserTask.MapKeys() {
		log.Println("ll1",taskUserKey.String())
		oldTaskNotification :=map[string]NotificationForTask {}
		err = dB.Child("/Notifications/UserDelay/TaskLocation/"+ companyTeamName+"/"+taskUserKey.String()).Value(&oldTaskNotification)
		log.Println("lklklkll4",oldTaskNotification)
		notifications := reflect.ValueOf(oldTaskNotification)
		for _, notificationKey := range notifications.MapKeys() {
			oldTaskNotificationStruct := NotificationForTask {}
			err = dB.Child("/Notifications/UserDelay/TaskLocation/"+ companyTeamName+"/"+taskUserKey.String()+"/"+notificationKey.String()).Value(&oldTaskNotificationStruct)
			if !(oldTaskNotificationStruct.IsRead){
				updateTaskNotification.Date=oldTaskNotificationStruct.Date
				updateTaskNotification.TaskName=oldTaskNotificationStruct.TaskName
				updateTaskNotification.TaskLocation=oldTaskNotificationStruct.TaskLocation
				updateTaskNotification.TaskId = oldTaskNotificationStruct.TaskId
				updateTaskNotification.IsRead=true
				updateTaskNotification.IsSentMail=oldTaskNotificationStruct.IsSentMail
				updateTaskNotification.Mode = oldTaskNotificationStruct.Mode
				updateTaskNotification.Message=oldTaskNotificationStruct.Message
				updateTaskNotification.UserName=oldTaskNotificationStruct.UserName
				updateTaskNotification.Category = oldTaskNotificationStruct.Category
				err = dB.Child("Notifications/UserDelay/TaskLocation/"+ companyTeamName+"/"+taskUserKey.String()+"/"+notificationKey.String()).Set(updateTaskNotification)

			}
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


	err = dB.Child("/Notifications/UserDelay/WorkLocation/"+ companyTeamName).Value(&oldWorkLocationNotificationValues)
	notificationOfUser := reflect.ValueOf(oldWorkLocationNotificationValues)
	for _, notificationUserKey := range notificationOfUser.MapKeys() {
		log.Println("keysss ",notificationUserKey)
		eachOldWorkLOcation := map[string]NotificationForWorkLocation{}
		err = dB.Child("/Notifications/UserDelay/WorkLocation/"+ companyTeamName+"/"+notificationUserKey.String()).Value(&eachOldWorkLOcation)
		notifications := reflect.ValueOf(eachOldWorkLOcation)
		for _, notificationKey := range notifications.MapKeys() {
			log.Println("notification key",notificationKey)
			workNotificationUpdateSuccess :=NotificationForWorkLocation{}
			err = dB.Child("/Notifications/UserDelay/WorkLocation/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Value(&workNotificationUpdateSuccess)
			log.Println("***********************************",workNotificationUpdateSuccess.Mode)
			updatedWorkNotification := NotificationForWorkLocation{}
			if !(workNotificationUpdateSuccess.IsRead){

				updatedWorkNotification.Date=workNotificationUpdateSuccess.Date
				updatedWorkNotification.WorkId=workNotificationUpdateSuccess.WorkId
				updatedWorkNotification.WorkLocation=workNotificationUpdateSuccess.WorkLocation
				updatedWorkNotification.IsRead=true
				updatedWorkNotification.Mode = workNotificationUpdateSuccess.Mode
				updatedWorkNotification.IsSentMail=workNotificationUpdateSuccess.IsSentMail
				updatedWorkNotification.Message=workNotificationUpdateSuccess.Message
				updatedWorkNotification.UserName=workNotificationUpdateSuccess.UserName
				updatedWorkNotification.Category = workNotificationUpdateSuccess.Category
				log.Println("iam here >>>>>>>>>>>>>>")
			}
			err = dB.Child("Notifications/UserDelay/WorkLocation/"+ companyTeamName+"/"+notificationUserKey.String()+"/"+notificationKey.String()).Set(updatedWorkNotification)
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
	err = dB.Child("Notifications/UserDelay/TaskLocation/"+ companyTeamName).Remove()
	err = dB.Child("Notifications/UserDelay/WorkLocation/"+ companyTeamName).Remove()
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

