
/* Author :Aswathy Ashok */

package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"
	"reflect"
	"app/passporte/helpers"
	"log"
	"bytes"
	"regexp"
	"strings"
	"strconv"
)

type TaskController struct {
	BaseController
}
/*Add task details to DB*/
func (c *TaskController)AddNewTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		task:=models.Tasks{}
		task.Info.TaskName= c.GetString("taskName")
		task.Job.JobName= c.GetString("jobName")
		task.Job.JobId = c.GetString("jobId")
		task.Customer.CustomerName = c.GetString("customerName")
		task.Customer.CustomerId =c.GetString("jobId")
		task.Info.TaskLocation =c.GetString("taskLocation")
		startDateString := c.GetString("startDateFomJs")
		endDateString :=c.GetString("endDateFromJs")
		layout := "01/02/2006 15:04"
		startDate, err := time.Parse(layout, startDateString)
		if err != nil {
			log.Println(err)
		}
		log.Println("Timeeeeeeeeeeeeeeeeeeeeeee: ", startDate.UTC().Unix())
		log.Println("Timeeeeeeeeeeeeeeeeeeeeeee: ", startDate.Unix())
		task.Info.StartDate = startDate.UTC().Unix()
		endDate, err := time.Parse(layout, endDateString)
		if err != nil {
			log.Println(err)
		}
		task.Info.EndDate = endDate.Unix()
		task.Info.TaskDescription = c.GetString("taskDescription")
		task.Info.UserNumber = c.GetString("minUsers")
		task.Info.Log = c.GetString("log")
		UserOrGroupIdArray := c.GetStrings("userOrGroup")
		UserOrGroupNameArray := c.GetStrings("userAndGroupName")
		tempContactName := c.GetStrings("contactName")
		tempContactId := c.GetStrings("contactId")
		task.Info.LoginType=c.GetString("loginType")
		task.Location.Latitude = c.GetString("latitude")
		task.Location.Longitude = c.GetString("longitude")
		FitToWork := c.GetString("addFitToWork")
		FitToWorkSlice := strings.Split(FitToWork, ",")
		task.Settings.DateOfCreation =time.Now().Unix()
		task.Settings.Status = helpers.StatusActive

		WorkBreak :=c.GetString("workExplosureText")
		WorkBreakSlice :=strings.Split(WorkBreak, ",")

		tempFitToWorkCheck :=c.GetString("fitToWorkCheck")
		if tempFitToWorkCheck =="on" {
			task.Settings.FitToWorkDisplayStatus ="EachTime"
		} else {
			task.Settings.FitToWorkDisplayStatus = "OnceADay"
		}
		task.Info.CompanyTeamName = storedSession.CompanyTeamName
		companyId :=storedSession.CompanyId
		userMap := make(map[string]models.TaskUser)
		groupMap := make(map[string]models.TaskGroup)
		groupNameAndDetails := models.TaskGroup{}
		userName :=models.TaskUser{}
		group := models.Group{}
		var keySliceForGroup [] string
		var MemberNameArray [] string
		groupMemberNameMap := make(map[string]models.GroupMemberName)
		//members := models.GroupMemberName{}

		for i := 0; i < len(UserOrGroupIdArray); i++ {
			tempName := UserOrGroupNameArray[i]
			tempId := UserOrGroupIdArray[i]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if((userOrGroupSelection[1]) == "User") {
				tempName = tempName[:len(tempName)-7]
				userName.FullName = tempName
				userName.Status =helpers.StatusActive
				userMap[tempId] = userName
			} else {
				tempName = tempName[:len(tempName)-8]
				groupNameAndDetails.GroupName = tempName

				//Getting member name from group


				groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, tempId)
				switch dbStatus {
				case true:

					memberData := reflect.ValueOf(groupDetails.Members)
					for _, key := range memberData.MapKeys() {
						keySliceForGroup = append(keySliceForGroup, key.String())
					}
					for i := 0; i < len(keySliceForGroup); i++ {
						MemberNameArray = append(MemberNameArray,groupDetails.Members[keySliceForGroup[i]].MemberName)
					}

				case false:
					log.Println(helpers.ServerConnectionError)
				}
				groupNameAndDetails.Members = groupMemberNameMap
				groupMap[tempId] = groupNameAndDetails
			}
		}

		task.UsersAndGroups.User = userMap
		task.UsersAndGroups.Group = groupMap
		contactMap := make(map[string]models.TaskContact)
		contact :=models.ContactUser{}
		taskContactDetail :=models.TaskContact{}
		for i := 0; i < len(tempContactId); i++ {

			 dbStatus,contactDetails := contact.RetrieveContactIdFromDB(c.AppEngineCtx, tempContactId[i])
			switch dbStatus {
			case true:
				taskContactDetail.ContactName = tempContactName[i]
				taskContactDetail.PhoneNumber = contactDetails.Info.PhoneNumber
				taskContactDetail.EmailId = contactDetails.Info.Email
				contactMap[tempContactId[i]] = taskContactDetail
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		}
		task.Contacts = contactMap


		//Add data to task DB
		dbStatus :=task.AddTaskToDB(c.AppEngineCtx,companyId,FitToWorkSlice,WorkBreakSlice)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}else {
		viewModel  := viewmodels.AddTaskViewModel{}
		companyUsers :=models.Company{}
		var keySlice []string
		var keySliceForGroupAndUser 	[]string
		var keySliceForContact		[]string
		//Getting Jobs
		dbStatus, allJobs := models.GetAllJobs(c.AppEngineCtx,companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(allJobs)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for _, k := range dataValue.MapKeys() {
				viewModel.JobNameArray   = append(viewModel.JobNameArray, allJobs[k.String()].Info.JobName)
				viewModel.JobCustomerNameArray = append(viewModel.JobCustomerNameArray, allJobs[k.String()].Customer.CustomerName)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}

		//Getting users and groups

		usersDetail :=models.Users{}
		dbStatus ,testUser:= companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx,companyTeamName)

		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(testUser)
			for _, key := range dataValue.MapKeys() {

				dataValue := reflect.ValueOf(testUser[key.String()].Users)
				for _, userKeys := range dataValue.MapKeys() {
					//viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKey.String()].FullName+" (User)")
					dbStatus := usersDetail.GetActiveUsersEmailForDropDown(c.AppEngineCtx, userKeys.String(),testUser[key.String()].Users[userKeys.String()].Email,companyTeamName)
					switch dbStatus {
					case true:
						viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKeys.String()].FullName+" (User)")
						keySliceForGroupAndUser = append(keySliceForGroupAndUser, userKeys.String())
					case false:
						log.Println(helpers.ServerConnectionError)


					}

				}
			}

			allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx,companyTeamName)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allGroups)
				for _, key := range dataValue.MapKeys() {
					var memberSlice []string
					keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
					viewModel.GroupNameArray = append(viewModel.GroupNameArray, allGroups[key.String()].Info.GroupName+" (Group)")

					// For selecting members while selecting a group in dropdown
					memberSlice = append(memberSlice, key.String())
					groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
					for _, memberKey := range groupDataValue.MapKeys()  {
						memberSlice = append(memberSlice, memberKey.String())
					}
					viewModel.GroupMembers = append(viewModel.GroupMembers, memberSlice)

				}
				viewModel.UserAndGroupKey=keySliceForGroupAndUser
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}

		//for getting all contact
		dbStatus, contacts := models.GetAllContact(c.AppEngineCtx,companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(contacts)
			for _, key := range dataValue.MapKeys() {
				keySliceForContact = append(keySliceForContact, key.String())
			}
			for _, k := range dataValue.MapKeys() {
				viewModel.ContactNameArray  = append(viewModel.ContactNameArray , contacts[k.String()].Info.Name)
			}
			viewModel.CompanyTeamName=storedSession.CompanyTeamName
			viewModel.CompanyPlan = storedSession.CompanyPlan
			viewModel.Key = keySlice
			viewModel.ContactKey=keySliceForContact
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		viewModel.AdminFirstName = storedSession.AdminFirstName
		viewModel.AdminLastName = storedSession.AdminLastName
		viewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-task.html"
	}

}

/* display all task details*/
func (c *TaskController)LoadTaskDetail() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	jobId := ""
	jobId = c.Ctx.Input.Param(":jobId")

	task := models.Tasks{}
	dbStatus, tasks := task.RetrieveTaskFromDB(c.AppEngineCtx,companyTeamName)
	viewModel := viewmodels.TaskDetailViewModel{}

	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(tasks)
		var keySlice []string
		var userKeySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			if tasks[k].Settings.Status == "Active" {
				var tempValueSlice []string

				tempJobAndCustomer := ""
				if tasks[k].Job.JobName != "" {
					var buffer bytes.Buffer
					buffer.WriteString(tasks[k].Job.JobName)
					buffer.WriteString(" (")
					buffer.WriteString(tasks[k].Customer.CustomerName)
					buffer.WriteString(")")
					tempJobAndCustomer = buffer.String()
					buffer.Reset()
				}
				tempValueSlice = append(tempValueSlice, "")
				tempValueSlice = append(tempValueSlice, tempJobAndCustomer)

				if !helpers.StringInSlice(tasks[k].Customer.CustomerName, viewModel.UniqueCustomerNames) && tasks[k].Customer.CustomerName != "" {
					viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames, tasks[k].Customer.CustomerName)
				}
				if jobId == tasks[k].Job.JobId {

					viewModel.SelectedJob = tasks[k].Job.JobName
					viewModel.SelectedCustomerForJob = tasks[k].Customer.CustomerName
				}
				if len(jobId) == 0 {
					viewModel.SelectedJob = ""
					viewModel.JobMatch = "true"

				}

				if !helpers.StringInSlice(tasks[k].Job.JobName, viewModel.UniqueJobNames) && tasks[k].Job.JobName != "" {
					viewModel.UniqueJobNames = append(viewModel.UniqueJobNames, tasks[k].Job.JobName)
				}
				//collecting fit to work from task
				fitToWorkDataValue := reflect.ValueOf(tasks[k].FitToWork)
				tempFitToWork := ""

				for _, fitToWorkKey := range fitToWorkDataValue.MapKeys() {

					var bufferFitToWork bytes.Buffer
					if len(tempFitToWork) == 0 {
						bufferFitToWork.WriteString(tasks[k].FitToWork[fitToWorkKey.String()].Description)
						tempFitToWork = bufferFitToWork.String()
						bufferFitToWork.Reset()
					} else {
						bufferFitToWork.WriteString(tempFitToWork)
						bufferFitToWork.WriteString(", ")
						bufferFitToWork.WriteString(tasks[k].FitToWork[fitToWorkKey.String()].Description)
						tempFitToWork = bufferFitToWork.String()
						bufferFitToWork.Reset()
					}
				}

				//displaying users
				usersDataValue := reflect.ValueOf(tasks[k].UsersAndGroups.User)

				//tempusersDataValue := ""
				var userCount =0
				for _, key := range usersDataValue.MapKeys() {
					userKeySlice = append(userKeySlice, key.String())
				}

				for _, userKey := range usersDataValue.MapKeys() {
					log.Println("user key",userKey)
					var userArray []string
					userArray =append(userArray,tasks[k].UsersAndGroups.User[userKey.String()].FullName)
					userArray =append(userArray,tasks[k].Settings.Status)
					//// using append
					// taskKeyCount :=len(k)
					//log.Println(keySlice)
					//log.Println("ggggggg",taskKeyCount)
					//	userKeyCount :=len(userKeySlice)
					//log.Println("hhhhh",userKeyCount)
					//	matrix := make([][]string, taskKeyCount,userKeyCount) // dim*dim matrix
					//	for i := 0; i < taskKeyCount; i++ {
					//		log.Println("cp1")
					//		matrix[i] = make([]string, 0, taskKeyCount)
					//		vector := make([]string, userKeyCount)
					//		for j := 0; j < userKeyCount; j++ {
					//			log.Println("cp2")
					//			vector[j] = tasks[k].UsersAndGroups.User[userKey.String()].FullName
					//			vector[j] =vector[j]+ tasks[k].Settings.Status
					//			matrix[i] = append(matrix[i], vector[j])
					//		}
					//
					//	}
					//	log.Println("hdhdshfhsfhshfgshgfhs",matrix)
					//

					//var bufferUser bytes.Buffer
					//if len(tempusersDataValue) == 0 {
					//	bufferUser.WriteString(tasks[k].UsersAndGroups.User[userKey.String()].FullName + " " + "(" + tasks[k].Settings.Status + ")")
					//	tempusersDataValue = bufferUser.String()
					//	bufferUser.Reset()
					//} else {
					//	bufferUser.WriteString(tempusersDataValue)
					//	bufferUser.WriteString(", ")
					//	bufferUser.WriteString(tasks[k].UsersAndGroups.User[userKey.String()].FullName + " " + "(" + tasks[k].Settings.Status + ")")
					//	tempusersDataValue = bufferUser.String()
					//	bufferUser.Reset()
					//}
					userCount =userCount+1
					viewModel.UserArray =append(viewModel.UserArray,userArray)
					userArray = userArray[:0]
				}

				tempuserCount := strconv.Itoa(userCount)
				//displaying contacts
				contactDataValue := reflect.ValueOf(tasks[k].Contacts)

				tempcontactDataValue := ""

				for _, contactKey := range contactDataValue.MapKeys() {

					var bufferContact bytes.Buffer
					if len(tempcontactDataValue) == 0 {
						bufferContact.WriteString(tasks[k].Contacts[contactKey.String()].ContactName)
						tempcontactDataValue = bufferContact.String()
						bufferContact.Reset()
					} else {
						bufferContact.WriteString(tempcontactDataValue)
						bufferContact.WriteString(", ")
						bufferContact.WriteString(tasks[k].Contacts[contactKey.String()].ContactName)
						tempcontactDataValue = bufferContact.String()
						bufferContact.Reset()
					}
				}

				tempValueSlice = append(tempValueSlice, tasks[k].Info.TaskName)
				tempValueSlice = append(tempValueSlice,tempuserCount)
				startTime := time.Unix(tasks[k].Info.StartDate, 0)
				startTimeOfTask := startTime.String()[11:16]
				startDate := time.Unix(tasks[k].Info.StartDate, 0).Format("2006/01/02")
				tempValueSlice = append(tempValueSlice, startDate+" "+"("+startTimeOfTask+")")
				endTime := time.Unix(tasks[k].Info.EndDate, 0)
				endTimeOfTask := endTime.String()[11:16]
				endDate := time.Unix(tasks[k].Info.EndDate, 0).Format("2006/01/02")
				tempValueSlice = append(tempValueSlice, endDate+" "+"("+endTimeOfTask+")")
				//tempValueSlice = append(tempValueSlice, tasks[k].Info.LoginType)
				//tempValueSlice = append(tempValueSlice, tasks[k].Info.UserNumber)
				//tempValueSlice = append(tempValueSlice, tasks[k].Settings.Status)
				//tempValueSlice = append(tempValueSlice, "")
				////tempValueSlice = append(tempValueSlice,  tempFitToWork)
				//tempValueSlice = append(tempValueSlice, tasks[k].Info.Log)
				//tempValueSlice = append(tempValueSlice, tempusersDataValue)
				//tempValueSlice = append(tempValueSlice, tempcontactDataValue)

				viewModel.Values = append(viewModel.Values, tempValueSlice)

				tempValueSlice = tempValueSlice[:0]
			}
		}
		//taskKeyCount =taskKeyCount+1
		viewModel.Keys = keySlice
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.CompanyPlan = storedSession.CompanyPlan
		viewModel.AdminFirstName = storedSession.AdminFirstName
		viewModel.AdminLastName = storedSession.AdminLastName
		viewModel.ProfilePicture =storedSession.ProfilePicture
		if  len(viewModel.SelectedJob) ==0 && len(jobId) !=0{

			viewModel.JobMatch ="false"
			viewModel.SelectedJob ="false"
		}
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/task-details.html"

		case false:
			log.Println(helpers.ServerConnectionError)

	}


}
/*delete task details from DB*/
func (c *TaskController)LoadDeleteTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	ReadSession(w, r, companyTeamName)
	companyId :=storedSession.CompanyId
	taskId :=c.Ctx.Input.Param(":taskId")
	task := models.Tasks{}
	dbStatus := task.DeleteTaskFromDB(c.AppEngineCtx, taskId,companyId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("true"))
	}
}

/* Edit task details*/
func (c *TaskController)LoadEditTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	taskId := c.Ctx.Input.Param(":taskId")

	if r.Method == "POST" {
		task := models.Tasks{}
		companyId :=storedSession.CompanyId
		task.Info.TaskName = c.GetString("taskName")
		task.Job.JobName = c.GetString("jobName")
		task.Job.JobId = c.GetString("jobId")
		task.Customer.CustomerName = c.GetString("customerName")
		task.Customer.CustomerId = c.GetString("jobId")
		task.Info.TaskLocation =c.GetString("taskLocation")
		startDateString := c.GetString("startDateFomJs")
		endDateString :=c.GetString("endDateFromJs")
		layout := "01/02/2006 15:04"
		startDate, err := time.Parse(layout, startDateString)
		if err != nil {
			log.Println(err)
		}
		task.Info.StartDate = startDate.Unix()
		endDate, err := time.Parse(layout, endDateString)
		if err != nil {
			log.Println(err)
		}
		task.Info.EndDate = endDate.Unix()
		task.Info.TaskDescription = c.GetString("taskDescription")
		task.Info.UserNumber = c.GetString("minUsers")
		task.Info.Log = c.GetString("log")
		UserOrGroupIdArray := c.GetStrings("userOrGroup")
		UserOrGroupNameArray := c.GetStrings("userAndGroupName")
		tempContactName := c.GetStrings("contactName")
		tempContactId := c.GetStrings("contactId")
		task.Info.LoginType = c.GetString("loginType")
		task.Location.Latitude = c.GetString("latitude")
		task.Location.Longitude = c.GetString("longitude")
		FitToWork := c.GetString("addFitToWork")
		FitToWorkSlice := strings.Split(FitToWork, ",")
		task.Settings.DateOfCreation = time.Now().Unix()
		tempFitToWorkCheck :=c.GetString("fitToWorkCheck")

		WorkBreak :=c.GetString("workExplosureText")
		WorkBreakSlice :=strings.Split(WorkBreak, ",")
		if tempFitToWorkCheck =="on" {
			task.Settings.FitToWorkDisplayStatus ="EachTime"
		} else {
			task.Settings.FitToWorkDisplayStatus = "OnceADay"
		}
		task.Settings.Status = helpers.StatusActive
		task.Info.CompanyTeamName = storedSession.CompanyTeamName
		userMap := make(map[string]models.TaskUser)
		groupMap := make(map[string]models.TaskGroup)
		groupNameAndDetails := models.TaskGroup{}
		userName := models.TaskUser{}
		group := models.Group{}
		var keySliceForGroup [] string
		var MemberNameArray [] string
		groupMemberNameMap := make(map[string]models.GroupMemberName)
		//members := models.GroupMemberName{}

		for i := 0; i < len(UserOrGroupIdArray); i++ {
			tempName := UserOrGroupNameArray[i]
			tempId := UserOrGroupIdArray[i]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if ((userOrGroupSelection[1]) == "User") {
				tempName = tempName[:len(tempName) - 7]
				userName.FullName = tempName
				userName.Status =helpers.StatusActive
				userMap[tempId] = userName
			} else {
				tempName = tempName[:len(tempName) - 8]
				groupNameAndDetails.GroupName = tempName

				//Getting member name from group


				groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, tempId)
				switch dbStatus {
				case true:

					memberData := reflect.ValueOf(groupDetails.Members)
					for _, key := range memberData.MapKeys() {
						keySliceForGroup = append(keySliceForGroup, key.String())
					}
					for i := 0; i < len(keySliceForGroup); i++ {
						MemberNameArray = append(MemberNameArray, groupDetails.Members[keySliceForGroup[i]].MemberName)
					}
				case false:
					log.Println(helpers.ServerConnectionError)
				}
				groupNameAndDetails.Members = groupMemberNameMap
				groupMap[tempId] = groupNameAndDetails
			}
		}

		task.UsersAndGroups.User = userMap
		task.UsersAndGroups.Group = groupMap
		contactMap := make(map[string]models.TaskContact)
		contact :=models.ContactUser{}
		taskContactDetail := models.TaskContact{}
		for i := 0; i < len(tempContactId); i++ {
			dbStatus,contactDetails := contact.RetrieveContactIdFromDB(c.AppEngineCtx, tempContactId[i])
			switch dbStatus {
			case true:
				taskContactDetail.ContactName = tempContactName[i]
				taskContactDetail.PhoneNumber = contactDetails.Info.PhoneNumber
				taskContactDetail.EmailId = contactDetails.Info.Email
				contactMap[tempContactId[i]] = taskContactDetail
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		}
		task.Contacts = contactMap

		//Add data to task DB
		dbStatus := task.UpdateTaskToDB(c.AppEngineCtx, taskId,companyId,FitToWorkSlice,WorkBreakSlice)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {

		viewModel := viewmodels.EditTaskViewModel{}
		task := models.Tasks{}
		companyUsers :=models.Company{}
		taskId := c.Ctx.Input.Param(":taskId")
		dbStatus, taskDetail := task.GetTaskDetailById(c.AppEngineCtx, taskId)
		switch dbStatus {
		case true:
			dbStatus, allJobs := models.GetAllJobs(c.AppEngineCtx,companyTeamName)
			var keySlice                        	[]string
			var keySliceForGroupAndUser        	[]string
			var keySliceForContact                	[]string
			var fitToWorkSlice			[]string
			var WorkBreak				[]string
			groupMember := models.Group{}

			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allJobs)
				for _, key := range dataValue.MapKeys() {
					keySlice = append(keySlice, key.String())
				}
				for _, k := range dataValue.MapKeys() {
					viewModel.JobNameArray = append(viewModel.JobNameArray, allJobs[k.String()].Info.JobName)
					viewModel.JobCustomerNameArray = append(viewModel.JobCustomerNameArray, allJobs[k.String()].Customer.CustomerName)
				}
			case false:
				log.Println(helpers.ServerConnectionError)
			}
			//getting all users for drop down
			usersDetail :=models.Users{}
			dbStatus ,testUser:= companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx,companyTeamName)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(testUser)
				for _, key := range dataValue.MapKeys() {

					dataValue := reflect.ValueOf(testUser[key.String()].Users)
					for _, userKeys := range dataValue.MapKeys() {
						//viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKey.String()].FullName+" (User)")
						dbStatus := usersDetail.GetActiveUsersEmailForDropDown(c.AppEngineCtx, userKeys.String(),testUser[key.String()].Users[userKeys.String()].Email,companyTeamName)
						switch dbStatus {
						case true:
							viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKeys.String()].FullName+" (User)")
							keySliceForGroupAndUser = append(keySliceForGroupAndUser, userKeys.String())
						case false:
							log.Println(helpers.ServerConnectionError)


						}

					}
				}

				allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx,companyTeamName)
				switch dbStatus {
				case true:
					dataValue := reflect.ValueOf(allGroups)
					for _, key := range dataValue.MapKeys() {
						var memberSlice []string
						keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
						viewModel.GroupNameArray = append(viewModel.GroupNameArray, allGroups[key.String()].Info.GroupName + "(Group)")

						// For selecting members while selecting a group in dropdown
						memberSlice = append(memberSlice, key.String())
						groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
						for _, memberKey := range groupDataValue.MapKeys()  {
							memberSlice = append(memberSlice, memberKey.String())
						}
						viewModel.GroupMembers = append(viewModel.GroupMembers, memberSlice)

					}
					viewModel.UserAndGroupKey = keySliceForGroupAndUser
				case false:
					log.Println(helpers.ServerConnectionError)
				}
			case false:
				log.Println(helpers.ServerConnectionError)
			}
			dbStatus, contacts := models.GetAllContact(c.AppEngineCtx,companyTeamName)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(contacts)
				for _, key := range dataValue.MapKeys() {
					keySliceForContact = append(keySliceForContact, key.String())
				}
				for _, k := range dataValue.MapKeys() {
					viewModel.ContactNameArray = append(viewModel.ContactNameArray, contacts[k.String()].Info.Name)
				}

				viewModel.ContactKey = keySliceForContact

				//contact name to edit
				 dbStatus,contactDetails := task.GetTaskDetailById(c.AppEngineCtx, taskId)
				switch dbStatus {
				case true:
					dataValue := reflect.ValueOf(contactDetails.Contacts)
					for _, key := range dataValue.MapKeys() {
						viewModel.ContactNameToEdit = append(viewModel.ContactNameToEdit, key.String())
					}

					//Selecting group name which is to be edited...

					dbStatus,groupDetails := task.GetTaskDetailById(c.AppEngineCtx, taskId)
					switch dbStatus {
					case true:
						dataValue := reflect.ValueOf(groupDetails.UsersAndGroups.Group)
						for _, key := range dataValue.MapKeys() {
							viewModel.GroupMembersAndUserToEdit = append(viewModel.GroupMembersAndUserToEdit,  key.String())
						}



						//selecting user name to edit
						dbStatus,groupDetails := task.GetTaskDetailById(c.AppEngineCtx, taskId)
						switch dbStatus {
						case true:
							dataValue := reflect.ValueOf(groupDetails.UsersAndGroups.User)
							for _, key := range dataValue.MapKeys() {
								viewModel.GroupMembersAndUserToEdit = append(viewModel.GroupMembersAndUserToEdit,  key.String())
								viewModel.UsersToEdit= append(viewModel.UsersToEdit,key.String())
							}
							groupValue :=reflect.ValueOf(groupDetails.UsersAndGroups.Group)
							for _, key := range groupValue.MapKeys() {
								viewModel.GroupMembersAndUserToEdit = append(viewModel.GroupMembersAndUserToEdit,  key.String())
								groupMemberDetail,_ := groupMember.GetGroupDetailsById(c.AppEngineCtx, key.String())
								groupMemberValue :=reflect.ValueOf(groupMemberDetail.Members)
								for _, key := range groupMemberValue.MapKeys() {
									viewModel.GroupsToEdit = append(viewModel.GroupsToEdit, key.String())
								}
							}
						case false:
							log.Println(helpers.ServerConnectionError)
						}
					case false:
						log.Println(helpers.ServerConnectionError)
					}
					//function to getting break details by task id
					taskWork :=models.TaskExposure{}
					dbStatus, taskWorkBreak := taskWork.GetTaskWorkBreakDetailById(c.AppEngineCtx, taskId)
					switch dbStatus {
					case true:
						workValue := reflect.ValueOf(taskWorkBreak)
						for _, key := range workValue.MapKeys() {
							WorkBreak = append(WorkBreak,taskWorkBreak[key.String()].BreakTime)
						}
					case false:
						log.Println(helpers.ServerConnectionError)
					}

						viewModel.Key = keySlice
						viewModel.PageType = helpers.SelectPageForEdit
						viewModel.JobName = taskDetail.Job.JobName
						viewModel.TaskName = taskDetail.Info.TaskName
						startDate := time.Unix(taskDetail.Info.StartDate, 0).Format("01/02/2006")
						viewModel.StartDate = startDate
						endDate := time.Unix(taskDetail.Info.EndDate, 0).Format("01/02/2006")
						viewModel.EndDate = endDate
						startTime := time.Unix(taskDetail.Info.StartDate, 0)
						startTimeOfTask := startTime.String()[11:16]
						viewModel.StartTime = startTimeOfTask
						endTime := time.Unix(taskDetail.Info.EndDate, 0)
						endTimeOfTask := endTime.String()[11:16]
						viewModel.EndTime = endTimeOfTask
						viewModel.TaskDescription = taskDetail.Info.TaskDescription
						viewModel.UserNumber = taskDetail.Info.UserNumber
						viewModel.Log = taskDetail.Info.Log
						dataValue = reflect.ValueOf(taskDetail.FitToWork)
						for _, key := range dataValue.MapKeys() {
							fitToWorkSlice = append(fitToWorkSlice,taskDetail.FitToWork[key.String()].Description)
						}

						viewModel.FitToWorkCheck=taskDetail.Settings.FitToWorkDisplayStatus
						viewModel.FitToWork = fitToWorkSlice
						viewModel.WorkBreak = WorkBreak
						viewModel.JobId = taskDetail.Job.JobId
						viewModel.TaskId = taskId
						viewModel.CompanyTeamName = storedSession.CompanyTeamName
						viewModel.CompanyPlan = storedSession.CompanyPlan
						viewModel.TaskLocation =taskDetail.Info.TaskLocation
						viewModel.LoginType =taskDetail.Info.LoginType
						viewModel.AdminFirstName = storedSession.AdminFirstName
						viewModel.AdminLastName = storedSession.AdminLastName
						viewModel.ProfilePicture =storedSession.ProfilePicture
						c.Data["vm"] = viewModel
						c.Layout = "layout/layout.html"
						c.TplName = "template/add-task.html"

					case false:
						log.Println(helpers.ServerConnectionError)
					}
				}

			case false:
				log.Println(helpers.ServerConnectionError)
			}

	}

}