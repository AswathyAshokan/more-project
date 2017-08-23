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
	//r := c.Ctx.Request
	//w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	var keySlice []string
	var keySliceForActiveTask []string

	var keyForLog []string

	var sliceForLeaveDetails []string
	viewModel := viewmodels.TimeSheetViewModel{}
	var userStructSlice []viewmodels.LogDetails
	var logUserSlice [][]viewmodels.LogDetails

	task := models.Tasks{}

		dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			log.Println("inside time sheet")
			dataValue := reflect.ValueOf(tasks)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}

			log.Println("tasks",keySlice)
			for _, taskKey := range keySlice {
				var keySliceForActiveTaskCompletedUsers []string
				var tempSlice	[]string
				userValue := reflect.ValueOf(tasks[taskKey].UsersAndGroups.User)
				for _, key := range userValue.MapKeys() {


					keySliceForActiveTask = append(keySliceForActiveTask, taskKey)
					if tasks[taskKey].UsersAndGroups.User[key.String()].UserTaskStatus == helpers.StatusCompleted&&tasks[taskKey].Settings.Status ==helpers.StatusActive {
						log.Println("task key",taskKey)
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

				log.Println("key slice for active task completed user",keySliceForActiveTaskCompletedUsers)
				logDetails :=models.WorkLog{}
				//var duration []string
				dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)
				log.Println("log deatils",logUserDetail)

				switch dbStatus {
				case true:
					//var userName string
					logValue := reflect.ValueOf(logUserDetail)
					for _, key := range logValue.MapKeys() {
						keyForLog = append(keyForLog, key.String())
					}
					for i := 0; i < len(keySliceForActiveTaskCompletedUsers); i++ {
						for _, k := range keyForLog {
							if logUserDetail[k].UserID == keySliceForActiveTaskCompletedUsers[i] {
								if  logUserDetail[k].LogDescription == "Work Started" || logUserDetail[k].LogDescription == "End of work day"||logUserDetail[k].LogDescription =="Completed" &&logUserDetail[k].TaskID == taskKey{
									var userStruct viewmodels.LogDetails
									userStruct.LogTime=logUserDetail[k].LogTime
									userStruct.TaskID = logUserDetail[k].TaskID
									userStruct.Type = logUserDetail[k].Type
									userStruct.UserID = logUserDetail[k].UserID
									userStruct.UserName = logUserDetail[k].UserName
									userStruct.LogDescription = logUserDetail[k].LogDescription
									userStructSlice = append(userStructSlice, userStruct)
								}

							}
						}
					}
					logUserSlice = append(logUserSlice, userStructSlice)
					log.Println("log details fromhdjjhjhsdjjh",logUserSlice)
					viewModel.LogArray =logUserSlice
				}
				//leaveDetail


				var keySliceForUser []string
				var keySlice []string
				var tempSliceForLeave []string
				companyUsersForLeave := models.Company{}
				leave := models.LeaveRequests{}
				dbStatus, companyUserDetail := companyUsersForLeave.GetUsersForDropdownFromCompany(c.AppEngineCtx, companyTeamName)

				switch dbStatus {
				case true:
					dataValue := reflect.ValueOf(companyUserDetail)
					for _, key := range dataValue.MapKeys() {
						dataValue := reflect.ValueOf(companyUserDetail[key.String()].Users)
						for _, userKey := range dataValue.MapKeys() {
							keySliceForUser = append(keySliceForUser, userKey.String())

						}
					}
				case false :
					log.Println(helpers.ServerConnectionError)
				}
				dbStatus, leaveDetail := leave.GetAllLeaveRequest(c.AppEngineCtx, keySliceForUser)
				switch dbStatus {
				case true:
					dataValue := reflect.ValueOf(leaveDetail)
					for _, key := range dataValue.MapKeys() {
						keySlice = append(keySlice, key.String())
						tempSliceForLeave = append(tempSliceForLeave,key.String())



					}
					log.Println("leave key",tempSliceForLeave)
					for _, leaveKey := range tempSliceForLeave {
						for i:=0;i<len(keySliceForActiveTaskCompletedUsers);i++{
							if leaveKey == keySliceForActiveTaskCompletedUsers[i] {
								status, leaveDetailOfUser,_,_ := leave.GetAllLeaveRequestById(c.AppEngineCtx, leaveKey,companyTeamName)
								switch status {
								case true:
									dataValue := reflect.ValueOf(leaveDetailOfUser)
									for _, key := range dataValue.MapKeys() {
										if leaveDetailOfUser[key.String()].Settings.Status == "Accepted"{
											numberOfDays := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.NumberOfDays, 10)
											startDateOfLeave := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.StartDate, 10)
											endDateOfLeave := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.EndDate, 10)
											sliceForLeaveDetails=append(sliceForLeaveDetails,numberOfDays)
											sliceForLeaveDetails=append(sliceForLeaveDetails,startDateOfLeave)
											sliceForLeaveDetails=append(sliceForLeaveDetails,endDateOfLeave)
											sliceForLeaveDetails=append(sliceForLeaveDetails,leaveKey)
											sliceForLeaveDetails =append(sliceForLeaveDetails,key.String())

										}

									}
								case false:
									log.Println(helpers.ServerConnectionError)
								}
							}
						}
					}
				case false :
					log.Println(helpers.ServerConnectionError)
				}
			}
		}
	log.Println("leave details of user",sliceForLeaveDetails)
	viewModel.LeaveDetails=append(viewModel.LeaveDetails,sliceForLeaveDetails)
	viewModel.CompanyTeamName =companyTeamName
	c.Data["vm"] = viewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/time-sheet.html"

	}

