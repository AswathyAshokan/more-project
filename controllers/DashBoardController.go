package controllers
import (
	"app/passporte/models"
	//"time"
	//"app/passporte/viewmodels"
	"reflect"
	//"app/passporte/helpers"
	"log"
	//"bytes"
	//"regexp"
	//"strconv"
	//"fmt"

	"app/passporte/helpers"

	"app/passporte/viewmodels"
)

type DashBoardController struct {
	BaseController
}
func (c *DashBoardController)LoadDashBoard() {
	viewModel  := viewmodels.DashBoardViewModel{}
	//r := c.Ctx.Request
	//w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	companyTask :=models.TaskIdInfo{}
	task := models.Tasks{}
	//section for getting total task completion and pending details
	dbStatus, companyTaskDetails := companyTask.RetrieveTaskFromCompany(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:

		dataValue := reflect.ValueOf(companyTaskDetails)
		var keySlice []string
		var totalUserStatus string

		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {


			dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, k)
			switch dbStatus {
			case true:
				var userStatus	[]string
				var userKeySlice []string
				pending :=0
				completed :=0
				dataValue := reflect.ValueOf(taskDetail.UsersAndGroups.User)
				for _, key := range dataValue.MapKeys() {
					userKeySlice = append(userKeySlice, key.String())
				}
				for _, k := range userKeySlice {

					userStatus = append(userStatus,taskDetail.UsersAndGroups.User[k].UserTaskStatus)
				}
				log.Println(userStatus)
				for i:=0;i<len(userStatus);i++{


					if userStatus[i] == "Pending" {

						totalUserStatus ="Pending"
						break
					}else{
						totalUserStatus ="Completed"
					}

				}
				for i:=0;i<len(userStatus);i++{


					if userStatus[i] == "Pending" {
						pending++

					}else{
						completed++
					}
				}
				log.Println("length",len(userStatus))
				completedTaskPercentage := float32(completed)/float32(len(userStatus))*100
				pendingTaskPercentage  := float32(pending)/ float32(len(userStatus))*100
				taskSettings :=models.TaskSetting{}
				taskSettings.UpdateTaskStatus(c.AppEngineCtx, k,totalUserStatus,completedTaskPercentage,pendingTaskPercentage)
				log.Println(dbStatus)

			case false:
				log.Println(helpers.ServerConnectionError)
			}


			}
		for _, taskKey := range keySlice {
			dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, taskKey)
			switch dbStatus {
			case true:
				totalCompletion :=0
				totalPending :=0
				if taskDetail.Settings.TaskStatus == "Completed"{
					totalCompletion++
				}else{
					totalPending++
				}
				completedTaskPercentageForViewModel := float32(totalCompletion)/float32(len(keySlice))*100
				pendingTaskPercentageForViewModel  := float32(totalPending)/ float32(len(keySlice))*100
				viewModel.CompletedTask =completedTaskPercentageForViewModel
				viewModel.PendingTask =pendingTaskPercentageForViewModel
			case false:
				log.Println(helpers.ServerConnectionError)
			}

		}


	case false:
		log.Println(helpers.ServerConnectionError)
	}
 //get total invited users
	//info,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx,companyTeamName)
	//var inviteKey []string
	//switch dbStatus {
	//case true:
	//	dataValue := reflect.ValueOf(info)
	//	//to store the keys of slice
	//	for _, key := range dataValue.MapKeys() {
	//		inviteKey = append(inviteKey, key.String())
	//	}


	c.Data["vm"] = viewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/dash-board.html"

}
