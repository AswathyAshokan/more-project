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
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	companyTask :=models.TaskIdInfo{}
	task := models.Tasks{}
	//section for getting total task completion and pending details
	dbStatus, companyTaskDetails := companyTask.RetrieveTaskFromCompany(c.AppEngineCtx,companyTeamName)
	var jobKeySlice []string
	if len(companyTaskDetails) !=0 {
		switch dbStatus {

		case true:
			log.Println("lengthhhh",len(companyTaskDetails))

			dataValue := reflect.ValueOf(companyTaskDetails)
			var keySlice []string
			var taskKeySlice []string
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
			totalCompletion :=0
			totalPending :=0
			for _, taskKey := range keySlice {
				dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, taskKey)
				switch dbStatus {
				case true:
					if taskDetail.Settings.Status =="Active" && taskDetail.Customer.CustomerStatus =="Active" {
						taskKeySlice = append(taskKeySlice,taskKey)

						if taskDetail.Settings.TaskStatus == "Completed" {

							totalCompletion++
						} else {
							totalPending++
						}
					}
				case false:
					log.Println(helpers.ServerConnectionError)
				}

			}
			completedTaskPercentageForViewModel := float32(totalCompletion)/float32(len(taskKeySlice))*100
			pendingTaskPercentageForViewModel  := float32(totalPending)/ float32(len(taskKeySlice))*100
			viewModel.CompletedTask =completedTaskPercentageForViewModel
			viewModel.PendingTask =pendingTaskPercentageForViewModel
			log.Println("task completed",viewModel.CompletedTask )
			log.Println("task pending",viewModel.PendingTask)


		case false:
			log.Println(helpers.ServerConnectionError)
		}
	} else{
		viewModel.CompletedTask =0
		viewModel.PendingTask =0
	}


	companyInvitaion :=models.CompanyInvitations{}
	acceptedUser :=0
	rejectedUser :=0
	pendingUser :=0
	dbStatus,info := companyInvitaion.InviteUserFromCompany(c.AppEngineCtx,companyTeamName)
	var inviteKey []string
	var activeInviteUserKey []string
	if len(info) !=0{
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(info)
			//to store the keys of slice
			for _, key := range dataValue.MapKeys() {
				inviteKey = append(inviteKey, key.String())
			}
		case false :
			log.Println(helpers.ServerConnectionError)
		}
		for _, inviteUserKey := range inviteKey {
			if info[inviteUserKey].Status =="Active" {
				activeInviteUserKey = append(activeInviteUserKey,inviteUserKey)
				if info[inviteUserKey].UserResponse == "Accepted" {
					acceptedUser++
				} else if info[inviteUserKey].UserResponse == "Pending" {
					pendingUser++
				} else {
					rejectedUser++
				}
			}


		}


		acceptedUsersPercentageForViewModel := float32(acceptedUser)/float32(len(activeInviteUserKey))*100
		rejectedUsersPercentageForViewModel  := float32(pendingUser)/ float32(len(activeInviteUserKey))*100
		pendingUsersPercentageForViewModel  := float32(rejectedUser)/ float32(len(activeInviteUserKey))*100
		viewModel.AcceptedUsers =acceptedUsersPercentageForViewModel
		viewModel.RejectedUsers =rejectedUsersPercentageForViewModel
		viewModel.PendingUsers =pendingUsersPercentageForViewModel
		log.Println("accpeted",viewModel.AcceptedUsers)
		log.Println("rejected",viewModel.RejectedUsers)
		log.Println("pending",viewModel.PendingUsers)

	}else {
		viewModel.AcceptedUsers =0
		viewModel.RejectedUsers =0
		viewModel.PendingUsers =0
	}




	//get job details
	var activeJobKey []string

	dbStatus, allJobs := models.GetAllJobs(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allJobs)
		for _, key := range dataValue.MapKeys() {
			jobKeySlice = append(jobKeySlice, key.String())
		}
		for _, k := range dataValue.MapKeys() {
			if allJobs[k.String()].Customer.CustomerStatus =="Active" {
				activeJobKey = append(activeJobKey, k.String())
				viewModel.JobNameArray = append(viewModel.JobNameArray, allJobs[k.String()].Info.JobName)
				viewModel.JobCustomerNameArray = append(viewModel.JobCustomerNameArray, allJobs[k.String()].Customer.CustomerName)
			}
		}
	case false:
		log.Println(helpers.ServerConnectionError)
	}

	dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		taskValue := reflect.ValueOf(tasks)
		for _, taskKey := range taskValue.MapKeys() {
			var tempValueSlice []string
			log.Println("task key",taskKey.String())
			tempValueSlice =append(tempValueSlice,tasks[taskKey.String()].Job.JobName)
			tempValueSlice =append(tempValueSlice,tasks[taskKey.String()].Info.TaskName)
			viewModel.TaskDetailArray =append(viewModel.TaskDetailArray,tempValueSlice)


		}
	case false:
		log.Println(helpers.ServerConnectionError)
	}

	viewModel.Key = activeJobKey
	viewModel.JobArrayLength =len(activeJobKey)
	viewModel.CompanyTeamName =companyTeamName
	viewModel.CompanyPlan = storedSession.CompanyPlan
	viewModel.AdminLastName =storedSession.AdminLastName
	viewModel.AdminFirstName =storedSession.AdminFirstName
	viewModel.ProfilePicture =storedSession.ProfilePicture
	log.Println("company plan",viewModel.CompanyPlan)
	log.Println("admin last",viewModel.AdminLastName)
	log.Println("admin first",viewModel.AdminFirstName)
	log.Println("profile ",viewModel.ProfilePicture)
	c.Data["vm"] = viewModel
	c.TplName = "template/dash-board.html"

}
