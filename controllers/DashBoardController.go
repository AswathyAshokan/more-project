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
	"strconv"
	//"fmt"

	"app/passporte/helpers"

	"app/passporte/viewmodels"
	"encoding/json"
)

type DashBoardController struct {
	BaseController
}
func IsValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
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
	log.Println("sp1")
	if len(companyTaskDetails) !=0 {
		log.Println("sp2")
		switch dbStatus {

		case true:
			dataValue := reflect.ValueOf(companyTaskDetails)
			var keySlice []string
			var totalUserStatus string

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for _, k := range keySlice {
				log.Println("sp3")
				var tempSLiceForBarChart []string

				if companyTaskDetails[k].Status ==helpers.StatusActive{
					dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, k)
					switch dbStatus {
					case true:
						if taskDetail.Settings.Status ==helpers.StatusActive && taskDetail.Customer.CustomerStatus ==helpers.StatusActive &&taskDetail.Job.JobStatus ==helpers.StatusActive{
							var userStatus	[]string
							var userKeySlice []string
							pending :=0
							completed :=0
							log.Println("hhhhhh",taskDetail.Info.TaskName)
							dataValue := reflect.ValueOf(taskDetail.UsersAndGroups.User)
							for _, key := range dataValue.MapKeys() {
								userKeySlice = append(userKeySlice, key.String())
							}
							for _, userKeyForTask := range userKeySlice {
								userStatus = append(userStatus,taskDetail.UsersAndGroups.User[userKeyForTask].UserTaskStatus)
							}
							tempSLiceForBarChart=append(tempSLiceForBarChart,taskDetail.Info.TaskName)
							tempSLiceForBarChart=append(tempSLiceForBarChart,k)
							startDate := strconv.FormatInt(taskDetail.Info.StartDate, 10)
							tempSLiceForBarChart=append(tempSLiceForBarChart,startDate)
							endDate := strconv.FormatInt(taskDetail.Info.StartDate, 10)
							tempSLiceForBarChart=append(tempSLiceForBarChart,endDate)
							viewModel.TaskDetailForBarChart =append(viewModel.TaskDetailForBarChart,tempSLiceForBarChart)

							//dataValue := reflect.ValueOf(taskDetail.UsersAndGroups.User)
							for _, userKey := range userKeySlice {
								if taskDetail.UsersAndGroups.User[userKey].Status !=helpers.UserStatusDeleted{
									var userArrayForBarChart []string
									userArrayForBarChart =append(userArrayForBarChart,taskDetail.UsersAndGroups.User[userKey].FullName)
									userArrayForBarChart =append(userArrayForBarChart,taskDetail.UsersAndGroups.User[userKey].UserTaskStatus)
									userArrayForBarChart =append(userArrayForBarChart,userKey)
									userArrayForBarChart =append(userArrayForBarChart,k)
									viewModel.UserDetailForBarChart =append(viewModel.UserDetailForBarChart,userArrayForBarChart)
								}
							}
							//create array for bar chart display
							logDetails :=models.WorkLog{}
							//var logUserForBarChart []string
							var keyForLog []string
							dbStatus, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)
							switch dbStatus {
							case true:
								//var userName string
								logValue := reflect.ValueOf(logUserDetail)
								for _, key := range logValue.MapKeys() {
									keyForLog = append(keyForLog, key.String())
								}
								//var timeSliceForNew [][]string
								log.Println("log keyyy", keyForLog)

								for _, logKey := range logValue.MapKeys() {
									for i :=0;i<len(userKeySlice);i++{
										if logUserDetail[logKey.String()].UserID ==userKeySlice[i] && k==logUserDetail[logKey.String()].TaskID&&logUserDetail[logKey.String()].LogDescription=="Work Started"{

										}
									}
								}
							case false:
								log.Println(helpers.ServerConnectionError)
							}
							bool1 :=IsValueInList("Pending", userStatus)
							bool2 :=IsValueInList("Open", userStatus)
							bool3 :=IsValueInList("Accepted", userStatus)
							if (bool1==true ||bool2==true || bool3 ==true ){
								totalUserStatus = "Pending"
							}else{
								totalUserStatus ="Completed"
							}

							log.Println("total status",totalUserStatus)
							for i:=0;i<len(userStatus);i++{


								if userStatus[i] == "Pending" {
									pending++

								}else{
									completed++
								}
							}
							completedTaskPercentage := float32(completed)/float32(len(userStatus))*100
							pendingTaskPercentage  := float32(pending)/ float32(len(userStatus))*100
							taskSettings :=models.TaskSetting{}
							taskSettings.UpdateTaskStatus(c.AppEngineCtx, k,totalUserStatus,completedTaskPercentage,pendingTaskPercentage)
							log.Println(dbStatus)
						}

					case false:
						log.Println(helpers.ServerConnectionError)
					}
				}
			}
			totalCompletion :=0
			totalPending :=0
			var completedOrPendingKey []string

			for _, taskKey := range keySlice {
				log.Println("sp4")
				dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, taskKey)
				switch dbStatus {
				case true:
					if taskDetail.Settings.Status ==helpers.StatusActive && taskDetail.Customer.CustomerStatus ==helpers.StatusActive &&taskDetail.Job.JobStatus ==helpers.StatusActive {
						completedOrPendingKey =append(completedOrPendingKey,taskKey)

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
			if len(completedOrPendingKey) !=0 {
				completedTaskPercentageForViewModel := float32(totalCompletion)/float32(len(completedOrPendingKey))*100
				pendingTaskPercentageForViewModel  := float32(totalPending)/ float32(len(completedOrPendingKey))*100
				viewModel.CompletedTask =completedTaskPercentageForViewModel
				viewModel.PendingTask =pendingTaskPercentageForViewModel
			}
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
			if info[inviteUserKey].UserResponse == "Accepted"{
				acceptedUser++
			} else if info[inviteUserKey].UserResponse == "Pending"{
				pendingUser++
			}else {
				rejectedUser++
			}
		}
		acceptedUsersPercentageForViewModel := float32(acceptedUser)/float32(len(inviteKey))*100
		rejectedUsersPercentageForViewModel  := float32(pendingUser)/ float32(len(inviteKey))*100
		pendingUsersPercentageForViewModel  := float32(rejectedUser)/ float32(len(inviteKey))*100
		viewModel.AcceptedUsers =acceptedUsersPercentageForViewModel
		viewModel.RejectedUsers =rejectedUsersPercentageForViewModel
		viewModel.PendingUsers =pendingUsersPercentageForViewModel
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
			if allJobs[k.String()].Customer.CustomerStatus =="Active"&&allJobs[k.String()].Settings.Status == helpers.StatusActive{
				activeJobKey = append(activeJobKey, k.String())
				viewModel.JobNameArray   = append(viewModel.JobNameArray, allJobs[k.String()].Info.JobName)
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
			if tasks[taskKey.String()].Customer.CustomerStatus =="Active"&&tasks[taskKey.String()].Settings.Status==helpers.StatusActive && tasks[taskKey.String()].Job.JobStatus==helpers.StatusActive{
				var tempValueSlice []string
				tempValueSlice =append(tempValueSlice,tasks[taskKey.String()].Job.JobName)
				tempValueSlice =append(tempValueSlice,tasks[taskKey.String()].Info.TaskName)
				viewModel.TaskDetailArray =append(viewModel.TaskDetailArray,tempValueSlice)
			}



		}
	case false:
		log.Println(helpers.ServerConnectionError)
	}

	dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
	log.Println("hghghghghghghghfydtdftyftyfyt",notificationValue)
	var notificationCount=0
	switch dbStatus {
	case true:
		notificationOfUser := reflect.ValueOf(notificationValue)
		for _, notificationUserKey := range notificationOfUser.MapKeys() {
			dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
			log.Println("notificationUserValue",notificationUserValue)
			switch dbStatus {
			case true:
				notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
				for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
					var NotificationArray []string
					var workLocationDelayArray []string
					if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
						notificationCount=notificationCount+1;
					}

					if notificationUserValue[notificationUserKeyForSpecific.String()].Category =="WorkLocation"{
						log.Println("iam in ffffff condition")

						workLocationDelayArray =append(workLocationDelayArray,notificationUserKey.String())
						workLocationDelayArray = append(workLocationDelayArray,notificationUserKeyForSpecific.String())
						workLocationDelayArray = append(workLocationDelayArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
						workLocationDelayArray =append(workLocationDelayArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
						workLocationDelayArray =append(workLocationDelayArray,notificationUserValue[notificationUserKeyForSpecific.String()].WorkLocation)
						tempCategory := "WorkLocation"+"t!@#$%&*YTREFFDD"
						workLocationDelayArray =append(workLocationDelayArray,tempCategory)
						tempDate := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
						workLocationDelayArray =append(workLocationDelayArray,tempDate)
						log.Println("workLocationDelayArray",workLocationDelayArray)
						viewModel.NotificationArray=append(viewModel.NotificationArray,workLocationDelayArray)
					} else {
						NotificationArray =append(NotificationArray,notificationUserKey.String())
						NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
						date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
						NotificationArray =append(NotificationArray,date)
						log.Println("NotificationArray",NotificationArray)
						viewModel.NotificationArray=append(viewModel.NotificationArray,NotificationArray)
					}

				}
			case false:
			}
		}
	case false:
	}
	log.Println("notificationCount",notificationCount)
	viewModel.NotificationNumber =notificationCount


	//get notification for admin when  documents of users is expired
	dbStatus,expiryNotification := models.GetAllNotificationsOfExpiration(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		viewModel.DocumentExpiryNotification = expiryNotification
	case false:

	}
	log.Println("nottttttttttt",viewModel.NotificationArray)
	viewModel.NotificationNumber =notificationCount
	viewModel.Key = activeJobKey
	viewModel.JobArrayLength =len(activeJobKey)
	viewModel.CompanyTeamName =companyTeamName
	viewModel.CompanyPlan = storedSession.CompanyPlan
	viewModel.AdminLastName =storedSession.AdminLastName
	viewModel.AdminFirstName =storedSession.AdminFirstName
	viewModel.ProfilePicture =storedSession.ProfilePicture
	c.Data["vm"] = viewModel
	c.TplName = "template/dash-board.html"

}

func (c *DashBoardController)LoadBarChartForDashBord() {
	//r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	taskName := c.GetString("TaskName")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	log.Println("kkkkkk", taskName, companyTeamName)
	companyTask := models.TaskIdInfo{}
	task := models.Tasks{}
	//section for getting total task completion and pending details
	dbStatus, companyTaskDetails := companyTask.RetrieveTaskFromCompany(c.AppEngineCtx, companyTeamName)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(companyTaskDetails)
		var keySlice []string
		logDetails := models.WorkLog{}
		//var totalUserStatus string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		var barChart [][] string
		var allPendinUserArray  [][]string
		//var allRejectedUserArray [][]string
		//var allAcceptdUserArray	 [][]string
		var  allCompletedUserArray [][]string
		var keysValues [] string
		var tempArray []string
		var starEndDateArray [] string
		for _, k := range keySlice {
			log.Println("sp3")
			//var tempSLiceForBarChart []string
			if companyTaskDetails[k].Status == helpers.StatusActive {
				Status, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, k)
				switch Status {
				case true:
					dataValueOfUser := reflect.ValueOf(taskDetail.UsersAndGroups.User)
					if taskDetail.Settings.Status == helpers.StatusActive {
						if taskDetail.Info.TaskName == taskName {
							log.Println("iam in first if")

							starEndDateArray = append(starEndDateArray, strconv.FormatInt(int64(taskDetail.Info.StartDate), 10))
							starEndDateArray = append(starEndDateArray, strconv.FormatInt(int64(taskDetail.Info.EndDate), 10))
							/*barChart = append(barChart,starEndDateArray)
							starEndDateArray = starEndDateArray[:0]*/
							//log.Println("starEndDateArray",starEndDateArray)

							_, logUserDetail := logDetails.GetLogDetailOfUser(c.AppEngineCtx, companyTeamName)
							log.Println("logUserDetail", logUserDetail)

							logValue := reflect.ValueOf(logUserDetail)
							for _, logKey := range logValue.MapKeys() {
								log.Println("in first loop")

								for _, userkey := range dataValueOfUser.MapKeys() {
									keysValues = append(keysValues,userkey.String())
									//var acceptedUsers []string
									var UserDetailsForStartTask []string
									var UserDetailsForCompleted []string


									log.Println("in second loop")
									if taskDetail.UsersAndGroups.User[userkey.String()].Status != helpers.UserStatusDeleted {
										log.Println("iam in second if")

										if userkey.String() == logUserDetail[logKey.String()].UserID {
											log.Println("iam in third if")


											if logUserDetail[logKey.String()].LogDescription == "Task Started" {
												log.Println("iam in fourth if")
												UserDetailsForStartTask = append(UserDetailsForStartTask, logUserDetail[logKey.String()].LogDescription)
												UserDetailsForStartTask = append(UserDetailsForStartTask,strconv.FormatInt(int64(logUserDetail[logKey.String()].LogTime), 10) )
												UserDetailsForStartTask = append(UserDetailsForStartTask,logUserDetail[logKey.String()].UserID)
												barChart = append(barChart,UserDetailsForStartTask)
												UserDetailsForStartTask = UserDetailsForStartTask[:0]
											} else if logUserDetail[logKey.String()].LogDescription == helpers.StatusCompleted {

												//UserDetailsForCompleted = append(UserDetailsForCompleted, taskDetail.UsersAndGroups.User[userkey.String()].FullName)
												UserDetailsForCompleted = append(UserDetailsForCompleted, logUserDetail[logKey.String()].LogDescription)
												UserDetailsForCompleted = append(UserDetailsForCompleted,strconv.FormatInt(int64(logUserDetail[logKey.String()].LogTime), 10))
												UserDetailsForCompleted = append(UserDetailsForCompleted, userkey.String())
												allCompletedUserArray = append(allCompletedUserArray,UserDetailsForCompleted)
												UserDetailsForCompleted = UserDetailsForCompleted [:0]
											} /*else if(logUserDetail[logKey.String()].LogDescription == helpers.StatusAccepted){
												acceptedUsers = append(acceptedUsers,logUserDetail[logKey.String()].LogDescription)
												acceptedUsers = append(acceptedUsers,strconv.FormatInt(int64(logUserDetail[logKey.String()].LogTime), 10))
												acceptedUsers = append(acceptedUsers,userkey.String())
												allAcceptdUserArray = append(allAcceptdUserArray,acceptedUsers)
												//barChart = append(barChart,acceptedUsers)
												acceptedUsers = acceptedUsers[:0]
											}*/
										}


										//log.Println("UserDetailsForPendingTask",UserDetailsForPendingTask)

									}
								}
							}



						}
					}
					for _, userkey := range dataValueOfUser.MapKeys() {
						if taskDetail.UsersAndGroups.User[userkey.String()].Status != helpers.UserStatusDeleted {
							if taskDetail.Info.TaskName == taskName {
								//var rejectedUsers []string
								var UserDetailsForPendingTask []string
								if taskDetail.UsersAndGroups.User[userkey.String()].UserTaskStatus == helpers.StatusPending {
									UserDetailsForPendingTask = append(UserDetailsForPendingTask, taskDetail.UsersAndGroups.User[userkey.String()].UserTaskStatus)
									UserDetailsForPendingTask = append(UserDetailsForPendingTask, userkey.String())
									allPendinUserArray = append(allPendinUserArray, UserDetailsForPendingTask)
									//barChart = append(barChart,UserDetailsForPendingTask)
									UserDetailsForPendingTask = UserDetailsForPendingTask [:0]
								} /*else if taskDetail.UsersAndGroups.User[userkey.String()].UserTaskStatus == helpers.UserResponseRejected {
									rejectedUsers = append(rejectedUsers, taskDetail.UsersAndGroups.User[userkey.String()].UserTaskStatus)
									rejectedUsers = append(rejectedUsers, userkey.String())
									allRejectedUserArray = append(allRejectedUserArray, rejectedUsers)
									rejectedUsers = rejectedUsers[:0]
								}*/
							}
						}

					}

				/*slices := []interface{}{"true"}
				sliceToClient, _ := json.Marshal(slices)
				log.Println("sliceToClient",sliceToClient)*/

				case false:
					log.Println("iam in error cobdition of inner loop")
				}
			}
		}
		for i:=0;i<len(keysValues);i++{

			exists := false
			for v := 0; v < i; v++ {
				if keysValues[v] == keysValues[i] {
					exists = true
					break
				}
			}
			// If no previous element exists, append this one.
			if !exists {
				tempArray = append(tempArray, keysValues[i])
			}
		}
		totalUsers :=  len(tempArray)
		log.Println("pending users",allPendinUserArray)
		log.Println("barChart",barChart)
		log.Println("allCompletedUserArray",allCompletedUserArray)
		slices := []interface{}{"true",barChart,allCompletedUserArray,allPendinUserArray,starEndDateArray,totalUsers}
		sliceToClient, _ := json.Marshal(slices)
		log.Println("sliceToClient",sliceToClient)
		w.Write(sliceToClient)
	case false:
		log.Println("iam in error cobdition")
		w.Write([]byte("false"))
	}
}


