package controllers
import (

	"app/passporte/models"
	"app/passporte/viewmodels"

	"reflect"
	"app/passporte/helpers"
	"strconv"
	// "fmt"

	"log"
)
type TimeSheetController struct {
	BaseController
}

func (c *TimeSheetController)LoadTimeSheetDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	var keySlice []string
	var keySliceForActiveTask []string

	var keyForLog []string

	//var sliceForLeaveDetails []string
	viewModel := viewmodels.TimeSheetViewModel{}


	task := models.Tasks{}

		dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:

			dataValue := reflect.ValueOf(tasks)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			var logUserSlice [][]viewmodels.LogDetails

			var keySliceForActiveTaskCompletedUsers []string
			for _, taskKey := range keySlice {

				var tempSlice	[]string
				userValue := reflect.ValueOf(tasks[taskKey].UsersAndGroups.User)
				for _, key := range userValue.MapKeys() {



					if tasks[taskKey].UsersAndGroups.User[key.String()].UserTaskStatus == helpers.StatusCompleted&&tasks[taskKey].Settings.Status ==helpers.StatusActive {
						log.Println("task key",taskKey)
						keySliceForActiveTask = append(keySliceForActiveTask, taskKey)
						keySliceForActiveTaskCompletedUsers = append(keySliceForActiveTaskCompletedUsers, key.String())
						startDate := strconv.FormatInt(tasks[taskKey].Info.StartDate, 10)
						endDate := strconv.FormatInt(tasks[taskKey].Info.EndDate, 10)
						tempSlice = append(tempSlice,startDate)
						tempSlice = append(tempSlice,endDate)
						tempSlice = append(tempSlice,tasks[taskKey].Info.TaskName)
						tempSlice = append(tempSlice,taskKey)
						tempSlice = append(tempSlice,key.String())
						tempSlice =append(tempSlice,tasks[taskKey].UsersAndGroups.User[key.String()].FullName)
						log.Println("task deatils",tempSlice)
						viewModel.TaskDetails =append(viewModel.TaskDetails ,tempSlice)
					}

				}

				logDetails :=models.WorkLog{}
				//var duration []string
				dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)



				var userStructSlice []viewmodels.LogDetails

				switch dbStatus {

				case true:
					//var userName string
					logValue := reflect.ValueOf(logUserDetail)
					for _, key := range logValue.MapKeys() {
						keyForLog = append(keyForLog, key.String())
					}
					//var timeSliceForNew [][]string
					log.Println("log keyyy",keyForLog)
					log.Println("active user keyyyyyyyyy",keySliceForActiveTaskCompletedUsers)
					for _, logKey := range logValue.MapKeys() {

						//for j := 0; j < len(keySliceForActiveTaskCompletedUsers); j++ {
							if  logUserDetail[logKey.String()].LogDescription == "Work Started" || logUserDetail[logKey.String()].LogDescription == "End of work day"||logUserDetail[logKey.String()].LogDescription =="Completed" &&logUserDetail[logKey.String()].TaskID == taskKey{
								//for k:=0;k<len(keySliceForActiveTask);k++ {
								//	if keySliceForActiveTask[k]==logUserDetail[logKey.String()].TaskID{
										log.Println("task key ", logUserDetail[logKey.String()].TaskID)
										log.Println("task key 2", taskKey)
										//var timeSliceNew []string

										//if logUserDetail[keyForLog[i]].UserID == keySliceForActiveTaskCompletedUsers[j] {
										log.Println("loop executed")
										var userStruct viewmodels.LogDetails

										userStruct.LogTime = logUserDetail[logKey.String()].LogTime
										userStruct.TaskID = logUserDetail[logKey.String()].TaskID
										userStruct.Type = logUserDetail[logKey.String()].Type
										userStruct.UserID = logUserDetail[logKey.String()].UserID
										userStruct.UserName = logUserDetail[logKey.String()].UserName
										userStruct.LogDescription = logUserDetail[logKey.String()].LogDescription
										userStructSlice = append(userStructSlice, userStruct)
										//logTimeString := strconv.FormatInt(logUserDetail[logKey.String()].LogTime, 10)
										//timeSliceNew =append(timeSliceNew,logTimeString)
										//timeSliceNew =append(timeSliceNew,logUserDetail[logKey.String()].TaskID)
										//timeSliceNew =append(timeSliceNew,logUserDetail[logKey.String()].Type)
										//timeSliceNew =append(timeSliceNew,logUserDetail[logKey.String()].UserID)
										//timeSliceNew =append(timeSliceNew,logUserDetail[logKey.String()].UserName)
										//timeSliceNew =append(timeSliceNew,logUserDetail[logKey.String()].LogDescription)
										//viewModel.LogDetailsTask =append(viewModel.LogDetailsTask ,timeSliceNew)
										//timeSliceForNew=append(timeSliceForNew,timeSliceNew)


									//}






									//}


									//}
								//}


						}

					}
					logUserSlice = append(logUserSlice, userStructSlice)
					viewModel.LogArray =logUserSlice
					viewModel.Keys=keySliceForActiveTaskCompletedUsers
					log.Println("final",viewModel.LogArray)
					//viewModel.LogDetailsTask =timeSliceForNew








				}

				//leaveDetail


				//var keySliceForUser []string
				//var keySlice []string
				//var tempSliceForLeave []string
				//companyUsersForLeave := models.Company{}
				//leave := models.LeaveRequests{}
				//dbStatus, companyUserDetail := companyUsersForLeave.GetUsersForDropdownFromCompany(c.AppEngineCtx, companyTeamName)
				//
				//switch dbStatus {
				//case true:
				//	dataValue := reflect.ValueOf(companyUserDetail)
				//	for _, key := range dataValue.MapKeys() {
				//		dataValue := reflect.ValueOf(companyUserDetail[key.String()].Users)
				//		for _, userKey := range dataValue.MapKeys() {
				//			keySliceForUser = append(keySliceForUser, userKey.String())
				//
				//		}
				//	}
				//case false :
				//	log.Println(helpers.ServerConnectionError)
				//}
				//dbStatus, leaveDetail := leave.GetAllLeaveRequest(c.AppEngineCtx, keySliceForUser)
				//switch dbStatus {
				//case true:
				//	dataValue := reflect.ValueOf(leaveDetail)
				//	for _, key := range dataValue.MapKeys() {
				//		keySlice = append(keySlice, key.String())
				//		tempSliceForLeave = append(tempSliceForLeave,key.String())
				//
				//
				//
				//	}
				//	log.Println("leave key",tempSliceForLeave)
				//	for _, leaveKey := range tempSliceForLeave {
				//		for i:=0;i<len(keySliceForActiveTaskCompletedUsers);i++{
				//			if leaveKey == keySliceForActiveTaskCompletedUsers[i] {
				//				status, leaveDetailOfUser,_,_ := leave.GetAllLeaveRequestById(c.AppEngineCtx, leaveKey,companyTeamName)
				//				switch status {
				//				case true:
				//					dataValue := reflect.ValueOf(leaveDetailOfUser)
				//					for _, key := range dataValue.MapKeys() {
				//						if leaveDetailOfUser[key.String()].Settings.Status == "Accepted"{
				//							numberOfDays := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.NumberOfDays, 10)
				//							startDateOfLeave := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.StartDate, 10)
				//							endDateOfLeave := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.EndDate, 10)
				//							sliceForLeaveDetails=append(sliceForLeaveDetails,numberOfDays)
				//							sliceForLeaveDetails=append(sliceForLeaveDetails,startDateOfLeave)
				//							sliceForLeaveDetails=append(sliceForLeaveDetails,endDateOfLeave)
				//							sliceForLeaveDetails=append(sliceForLeaveDetails,leaveKey)
				//							sliceForLeaveDetails =append(sliceForLeaveDetails,key.String())
				//
				//						}
				//
				//					}
				//				case false:
				//					log.Println(helpers.ServerConnectionError)
				//				}
				//			}
				//		}
				//	}
				//case false :
				//	log.Println(helpers.ServerConnectionError)
				//}
			}
		}
	log.Println("hiiiiiiiiii")
	users :=models.Users{}
	dbStatus, companyUserDetail := users.GetAllUsers(c.AppEngineCtx)
	switch dbStatus{
	case true:
		dataValue := reflect.ValueOf(companyUserDetail)
		for _, key := range dataValue.MapKeys() {
			companyValue :=reflect.ValueOf(companyUserDetail[key.String()].WorkLocation)
			for _, companyKey := range companyValue.MapKeys() {
				if companyUserDetail[key.String()].WorkLocation[companyKey.String()].CompanyId ==companyTeamName{
					var tempSliceOfUser []string
					tempSliceOfUser=append(tempSliceOfUser,companyUserDetail[key.String()].WorkLocation[companyKey.String()].DailyStartDate)
					tempSliceOfUser=append(tempSliceOfUser,companyUserDetail[key.String()].WorkLocation[companyKey.String()].DailyEndDate)
					tempSliceOfUser=append(tempSliceOfUser,key.String())
					viewModel.UserStartTimeAndEndTime=append(viewModel.UserStartTimeAndEndTime,tempSliceOfUser)


				}

			}
		}
	case false:
		log.Println(helpers.ServerConnectionError)


	}




	log.Println("user dataaaaaaaaaa",viewModel.UserStartTimeAndEndTime)
	//log.Println("leave details of user",sliceForLeaveDetails)
	//viewModel.LeaveDetails=append(viewModel.LeaveDetails,sliceForLeaveDetails)
	viewModel.CompanyTeamName =companyTeamName
	viewModel.CompanyPlan = storedSession.CompanyPlan
	viewModel.AdminLastName =storedSession.AdminLastName
	viewModel.AdminFirstName =storedSession.AdminFirstName
	viewModel.ProfilePicture =storedSession.ProfilePicture
	c.Data["vm"] = viewModel
	c.TplName = "template/time-sheet.html"

	}

