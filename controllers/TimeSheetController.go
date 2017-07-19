package controllers
import (

	"app/passporte/models"
	//"app/passporte/viewmodels"

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
	var keySliceForActiveTaskCompletedUsers []string
	var keyForLog []string
	var tempSlice	[]string
	//var tempValueSlice []string
	logDetails := models.WorkLog{}
	task := models.Tasks{}

		dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(tasks)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for _, k := range keySlice {
				keySliceForActiveTask = append(keySliceForActiveTask, k)
				userValue := reflect.ValueOf(tasks[k].UsersAndGroups.User)
				for _, key := range userValue.MapKeys() {
					if tasks[k].UsersAndGroups.User[key.String()].UserTaskStatus == helpers.StatusCompleted {
						log.Println("task key",k)
						keySliceForActiveTaskCompletedUsers = append(keySliceForActiveTaskCompletedUsers, key.String())
						startDate := strconv.FormatInt(tasks[k].Info.StartDate, 10)
						endDate := strconv.FormatInt(tasks[k].Info.EndDate, 10)
						tempSlice = append(tempSlice,startDate)
						tempSlice = append(tempSlice,endDate)
						tempSlice = append(tempSlice,tasks[k].Info.TaskName)
						log.Println("task deatils",tempSlice)
					}
				}

				dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)

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

								log.Println("log details",logUserDetail[k])
							}
						}
					}
				//tempValueSlice = append(tempValueSlice, logUserDetail[key.String()].UserName)
				}

				var keySliceForUser []string
				var keySlice []string
				var tempSliceForLeave []string
				//var commonKey []string
				//storedSession := ReadSession(w, r, companyTeamName)
				//companyId := storedSession.CompanyId
				companyUsersForLeave := models.Company{}
				//companyUserDetail := models.Company{}
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
						log.Println("leave",tempSliceForLeave)

					}
				case false :
					log.Println(helpers.ServerConnectionError)
				}
			}
		}
	c.Layout = "layout/layout.html"
	c.TplName = "template/time-sheet.html"

	}

