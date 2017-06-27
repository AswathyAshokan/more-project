
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
	"fmt"

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
		task.Job.JobStatus =helpers.StatusActive
		jobIdForTask  :=task.Job.JobId
		task.Customer.CustomerName = c.GetString("customerName")
		task.Customer.CustomerId =c.GetString("jobId")
		task.Customer.CustomerStatus =helpers.StatusActive
		customerIdForTask :=task.Customer.CustomerId
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
		logInMinutes :=c.GetString("log")
		logInMinutesInString, err := strconv.ParseInt(logInMinutes, 10, 64)
		if err != nil {
			// handle error
			fmt.Println(err)

		}
		task.Info.LogTimeInMinutes = logInMinutesInString
		UserOrGroupIdArray := c.GetStrings("userOrGroup")
		UserOrGroupNameArray := c.GetStrings("userAndGroupName")
		tempContactName := c.GetStrings("contactName")
		tempContactId := c.GetStrings("contactId")
		task.Info.LoginType=c.GetString("loginType")
		task.Info.NFCTagID =c.GetString("nfcTagId")
		task.Location.Latitude = c.GetString("latitude")
		task.Location.Longitude = c.GetString("longitude")
		exposureTask := c.GetString("exposureBreakTime")
		exposureWorkTime := c.GetString("exposureWorkTime")
		log.Println("breakkkkk time",exposureTask)
		FitToWork := c.GetString("addFitToWork")
		FitToWorkSlice := strings.Split(FitToWork, ",")
		task.Settings.DateOfCreation =time.Now().Unix()
		task.Settings.Status = helpers.StatusActive
		task.Settings.TaskStatus =helpers.StatusPending

		//WorkBreak :=c.GetString("workExplosureText")
		TaskBreakTimeSlice :=strings.Split(exposureTask, ",")
		TaskWorkTimeSlice :=strings.Split(exposureWorkTime, ",")

		tempFitToWorkCheck :=c.GetString("fitToWorkCheck")
		if tempFitToWorkCheck =="on" {
			task.Settings.FitToWorkDisplayStatus ="EachTime"
		} else {
			task.Settings.FitToWorkDisplayStatus = "OnceADay"
		}
		task.Info.CompanyTeamName = storedSession.CompanyTeamName
		task.Info.CompanyName =storedSession.CompanyName
		companyId :=storedSession.CompanyId
		userMap := make(map[string]models.TaskUser)
		groupMap := make(map[string]models.TaskGroup)
		groupNameAndDetails := models.TaskGroup{}
		userName :=models.TaskUser{}
		group := models.Group{}
		var keySliceForGroup [] string
		var MemberNameArray [] string
		var groupKeySlice	[]string
		

		//groupMemberNameMap := make(map[string]models.GroupMemberName)
		//members := models.GroupMemberName{}
		groupMemberMap := make(map[string]models.GroupMemberName)
		groupMemberNameForTask :=models.GroupMemberName{}

		for i := 0; i < len(UserOrGroupIdArray); i++ {
			tempName := UserOrGroupNameArray[i]
			tempId := UserOrGroupIdArray[i]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if((userOrGroupSelection[1]) == "User") {
				tempName = tempName[:len(tempName)-7]
				userName.FullName = tempName
				userName.Status =helpers.StatusActive
				userName.UserTaskStatus =helpers.StatusPending
				userMap[tempId] = userName
			} else {
				tempName = tempName[:len(tempName)-8]
				groupNameAndDetails.GroupName = tempName
				groupNameAndDetails.GroupStatus =helpers.StatusActive
				//Getting member name from group


				groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, tempId)
				switch dbStatus {
				case true:

					memberData := reflect.ValueOf(groupDetails.Members)
					for _, key := range memberData.MapKeys() {
						keySliceForGroup =append(keySliceForGroup,key.String())
						MemberNameArray = append(MemberNameArray,groupDetails.Members[key.String()].MemberName)

					}

				case false:
					log.Println(helpers.ServerConnectionError)
				}
				for i := 0; i < len(keySliceForGroup); i++ {
					groupMemberNameForTask.MemberName =MemberNameArray[i]
					groupMemberMap[keySliceForGroup[i]] = groupMemberNameForTask

				}
				groupNameAndDetails.Members = groupMemberMap
				log.Println("hgjghrh",groupMemberMap)
				groupMap[tempId] = groupNameAndDetails
				groupKeySlice = append(groupKeySlice,tempId)

			}




		}


		log.Println("group detail",groupMap)
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
				taskContactDetail.ContactStatus =helpers.StatusActive
				contactMap[tempContactId[i]] = taskContactDetail
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		}
		task.Contacts = contactMap


		//Add data to task DB
		dbStatus :=task.AddTaskToDB(c.AppEngineCtx,companyId,FitToWorkSlice, TaskBreakTimeSlice,TaskWorkTimeSlice,tempContactId,groupKeySlice,jobIdForTask,customerIdForTask)
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
		var taskUserSlice [][]viewmodels.TaskUsers
		for _, k := range keySlice {
			if tasks[k].Settings.Status == "Active" {
				var tempValueSlice []string
				var minUserArray []string

				tempJobAndCustomer := ""
				if tasks[k].Job.JobName != "" {
					var buffer bytes.Buffer
					if tasks[k].Job.JobStatus =="Active"{
						buffer.WriteString(tasks[k].Job.JobName)
					} else{
						buffer.WriteString("")
					}

					buffer.WriteString(" (")
					if tasks[k].Customer.CustomerStatus =="Active"{
						buffer.WriteString(tasks[k].Customer.CustomerName)
					}else{
						buffer.WriteString("")
					}


					buffer.WriteString(")")
					tempJobAndCustomer = buffer.String()
					log.Println("ddfffdfdff",tempJobAndCustomer)
					if tempJobAndCustomer ==" ()"{
						log.Println("hhhhhh")
						tempJobAndCustomer=""
					}
					buffer.Reset()
				}
				tempValueSlice = append(tempValueSlice, "")
				tempValueSlice = append(tempValueSlice, tempJobAndCustomer)

				if !helpers.StringInSlice(tasks[k].Customer.CustomerName, viewModel.UniqueCustomerNames) && tasks[k].Customer.CustomerName != "" {
					viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames, tasks[k].Customer.CustomerName)
					log.Println("ggggg",viewModel.UniqueCustomerNames)
					log.Println("fffff",tasks[k].Customer.CustomerName)
				}else{
					viewModel.UniqueCustomerNames = append(viewModel.UniqueCustomerNames,"")
				}
				if jobId == tasks[k].Job.JobId {

					viewModel.SelectedJob = tasks[k].Job.JobName
					viewModel.SelectedCustomerForJob = tasks[k].Customer.CustomerName
				}
				if len(jobId) == 0 {
					viewModel.SelectedJob = ""
					viewModel.JobMatch = "true"

				}

				if !helpers.StringInSlice(tasks[k].Job.JobName, viewModel.UniqueJobNames) && tasks[k].Job.JobName != ""&& tasks[k].Job.JobStatus =="Active" {
					viewModel.UniqueJobNames = append(viewModel.UniqueJobNames, tasks[k].Job.JobName)
				}else{
					viewModel.UniqueJobNames = append(viewModel.UniqueJobNames,"")
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
				var userStructSlice []viewmodels.TaskUsers

				for _, userKey := range usersDataValue.MapKeys() {

					var userStruct viewmodels.TaskUsers
					userStruct.Name = tasks[k].UsersAndGroups.User[userKey.String()].FullName
					userStruct.Status = tasks[k].UsersAndGroups.User[userKey.String()].UserTaskStatus
					userStruct.TaskId =k
					userStructSlice = append(userStructSlice, userStruct)
					//userArray =append(userArray,tasks[k].UsersAndGroups.User[userKey.String()].FullName)
					//userArray =append(userArray,tasks[k].Settings.Status)
					//taskKeyCount :=len(keySlice)
					//userKeyCount :=len(userKeySlice)
					//var innerSlice []string

					//userArray =append(userArray,tasks[k].UsersAndGroups.User[userKey.String()].FullName)
					//userArray =append(userArray,tasks[k].UsersAndGroups.User[userKey.String()].Status)
					//outerSlice = append(outerSlice, userArray)
					userCount =userCount+1
				}
				taskUserSlice = append(taskUserSlice, userStructSlice)
				//viewModel.UserArray = outerSlice
				//log.Println("array",outerSlice)
				tempUserCount := strconv.Itoa(userCount)
				//displaying contacts
				contactDataValue := reflect.ValueOf(tasks[k].Contacts)

				tempContactDataValue := ""

				for _, contactKey := range contactDataValue.MapKeys() {

					var bufferContact bytes.Buffer
					if  tasks[k].Contacts[contactKey.String()].ContactStatus =="Active" {
						if len(tempContactDataValue) == 0 {
							bufferContact.WriteString(tasks[k].Contacts[contactKey.String()].ContactName)
							tempContactDataValue = bufferContact.String()
							bufferContact.Reset()
						} else {
							bufferContact.WriteString(tempContactDataValue)
							bufferContact.WriteString(", ")
							bufferContact.WriteString(tasks[k].Contacts[contactKey.String()].ContactName)
							tempContactDataValue = bufferContact.String()
							bufferContact.Reset()
						}
					}
				}

				tempValueSlice = append(tempValueSlice, tasks[k].Info.TaskName)
				tempValueSlice = append(tempValueSlice, tempUserCount)
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
				minUserArray = append(minUserArray,tasks[k].Info.UserNumber)
				minUserArray = append(minUserArray,tasks[k].Info.LoginType)
				minUserArray = append(minUserArray,k)
				viewModel.MinUserAndLoginTypeArray =append(viewModel.MinUserAndLoginTypeArray,minUserArray)
				minUserArray =minUserArray[:0]

			}
		}

		viewModel.UserArray = taskUserSlice


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
		logInMinutes :=c.GetString("log")
		logInMinutesInString, err := strconv.ParseInt(logInMinutes, 10, 64)
		if err != nil {
			// handle error
			fmt.Println(err)

		}
		task.Info.LogTimeInMinutes = logInMinutesInString
		UserOrGroupIdArray := c.GetStrings("userOrGroup")
		UserOrGroupNameArray := c.GetStrings("userAndGroupName")
		tempContactName := c.GetStrings("contactName")
		tempContactId := c.GetStrings("contactId")
		task.Info.LoginType = c.GetString("loginType")
		task.Info.NFCTagID =c.GetString("nfcTagId")
		task.Location.Latitude = c.GetString("latitude")
		task.Location.Longitude = c.GetString("longitude")
		FitToWork := c.GetString("addFitToWork")
		FitToWorkSlice := strings.Split(FitToWork, ",")
		task.Settings.DateOfCreation = time.Now().Unix()
		tempFitToWorkCheck :=c.GetString("fitToWorkCheck")

		//WorkBreak :=c.GetString("workExplosureText")
		//WorkBreakSlice :=strings.Split(WorkBreak, ",")
		exposureTask := c.GetString("exposureBreakTime")
		exposureWorkTime := c.GetString("exposureWorkTime")
		TaskBreakTimeSlice :=strings.Split(exposureTask, ",")
		TaskWorkTimeSlice :=strings.Split(exposureWorkTime, ",")
		if tempFitToWorkCheck =="on" {
			task.Settings.FitToWorkDisplayStatus ="EachTime"
		} else {
			task.Settings.FitToWorkDisplayStatus = "OnceADay"
		}
		task.Settings.Status = helpers.StatusActive
		task.Info.CompanyTeamName = storedSession.CompanyTeamName
		task.Info.CompanyName =storedSession.CompanyName
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
		dbStatus := task.UpdateTaskToDB(c.AppEngineCtx, taskId,companyId,FitToWorkSlice,TaskBreakTimeSlice,TaskWorkTimeSlice)
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
			var WorkTime                            []string
			var BreakTime				[]string
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
						if task.Contacts[key.String()].ContactStatus =="Active"{
							viewModel.ContactNameToEdit = append(viewModel.ContactNameToEdit, key.String())
						}
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
							breakHourInInt, err := strconv.Atoi(taskWorkBreak[key.String()].BreakDurationInMinutes)
							//workHourInInt, err := strconv.Atoi(taskWorkBreak[key.String()].WorkTime)
							if err != nil {
								// handle error
								log.Println(err)

							}
							var breakHours = breakHourInInt/60
							var breakMinutes =breakHourInInt %60
							var breakHourInString =string(breakHours)
							var breakMinutesInString =string(breakMinutes)
							var prependBreakHours =""
							var prependBreakMinutes =""
							if len(breakHourInString) ==1 {

								prependBreakHours = fmt.Sprintf("%02d", breakHours)
							} else {
								prependBreakHours = string(breakHours)
							}
							if len(breakMinutesInString) ==1 {

								prependBreakMinutes = fmt.Sprintf("%02d", breakMinutes)
							} else {
								prependBreakMinutes = string(breakMinutes)
							}
							breakTimeForTask :=prependBreakHours+":"+prependBreakMinutes
							log.Println("break time ",breakTimeForTask)
							BreakTime = append(BreakTime,breakTimeForTask)

							workHourInInt, err := strconv.Atoi(taskWorkBreak[key.String()].BreakStartTimeInMinutes)
							//workHourInInt, err := strconv.Atoi(taskWorkBreak[key.String()].WorkTime)
							if err != nil {
								// handle error
								log.Println(err)

							}
							var workHours = workHourInInt/60
							var workMinutes =workHourInInt %60
							var workHourInString =string(workHours)
							var workMinutesInString =string(workMinutes)
							var prependWorkHours =""
							var prependWorkMinutes =""
							if len(workHourInString) ==1 {

								prependWorkHours = fmt.Sprintf("%02d", workHours)
							} else {
								prependWorkHours = string(workHours)
							}
							if len(workMinutesInString) ==1 {

								prependWorkMinutes = fmt.Sprintf("%02d", workMinutes)
							} else {
								prependWorkMinutes = string(workMinutes)
							}
							workTimeForTask :=prependWorkHours+":"+prependWorkMinutes
							log.Println("break time ",workTimeForTask)

							WorkTime = append(WorkTime,workTimeForTask)

						}
					case false:
						log.Println(helpers.ServerConnectionError)
					}

						viewModel.Key = keySlice
						viewModel.PageType = helpers.SelectPageForEdit
						if taskDetail.Job.JobStatus == "Active"{
							viewModel.JobName = taskDetail.Job.JobName
						}

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
						viewModel.NFCTagId = taskDetail.Info.NFCTagID
						log.Println("logTime",taskDetail.Info.LogTimeInMinutes)
						logTimeOfUser := strconv.FormatInt(taskDetail.Info.LogTimeInMinutes,10)
						viewModel.Log = logTimeOfUser
						dataValue = reflect.ValueOf(taskDetail.FitToWork)
						for _, key := range dataValue.MapKeys() {
							fitToWorkSlice = append(fitToWorkSlice,taskDetail.FitToWork[key.String()].Description)
						}

						viewModel.FitToWorkCheck=taskDetail.Settings.FitToWorkDisplayStatus
						viewModel.FitToWork = fitToWorkSlice
						viewModel.WorkTime = WorkTime
						viewModel.BreakTime =BreakTime
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