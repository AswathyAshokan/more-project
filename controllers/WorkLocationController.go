package controllers
import (
	"app/passporte/models"
	"log"
	"reflect"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"regexp"
	"time"
	"strings"
	"strconv"
	"fmt"
)

type WorkLocationcontroller struct {
	BaseController
}


func (c *WorkLocationcontroller) AddWorkLocation() {
	log.Println("w1")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	companyUsers :=models.Company{}
	var keySliceForGroupAndUser []string
	workLocationViewmodel := viewmodels.AddLocationViewModel{}
	userMap := make(map[string]models.WorkLocationUser)
	groupNameAndDetails := models.WorkLocationGroup{}
	groupMemberNameForTask :=models.GroupMemberNameInWorkLocation{}
	groupMemberMap := make(map[string]models.GroupMemberNameInWorkLocation)
	groupMap := make(map[string]models.WorkLocationGroup)
	var keySliceForGroup [] string
	var MemberNameArray [] string
	var notificationCount = 0
	group := models.Group{}
	userName :=models.WorkLocationUser{}
	WorkLocation := models.WorkLocation{}
	if r.Method == "POST" {
		exposureTask := c.GetString("exposureBreakTime")
		exposureWorkTime := c.GetString("exposureWorkTime")
		log.Println("exposureTask",exposureTask)
		log.Println("exposureWorkTime",exposureWorkTime)
		fitToWorkName :=c.GetString("fitToWorkName")
		groupKeySliceForWorkLocationString := c.GetString("groupArrayElement")
		UserOrGroupNameArray :=c.GetStrings("userAndGroupName")
		taskLocation := c.GetString("taskLocation")
		dailyEndTime := c.GetString("dailyEndTimeString")
		dailyStartTime := c.GetString("dailyStartTimeString")
		groupKeySliceForWorkLocation :=strings.Split(groupKeySliceForWorkLocationString, ",")
		userIdArray := c.GetStrings("selectedUserNames")
		latitude := c.GetString("latitudeId")
		longitude := c.GetString("longitudeId")
		startDate := c.GetString("startDateTimeStamp")
		endDate := c.GetString("endDateTimeStamp")
		startDateInt , err := strconv.ParseInt(startDate, 10, 64)
		endDateInt, err := strconv.ParseInt(endDate, 10, 64)
		loginType := c.GetString("loginType")
		log.Println("loginType",loginType)
		logInMinutes :=c.GetString("logInMinutes")
		logInMinutesInString, err := strconv.ParseInt(logInMinutes, 10, 64)
		WorkLocationBreakTimeSlice :=strings.Split(exposureTask, ",")
		WorkLocationTimeSlice :=strings.Split(exposureWorkTime, ",")
		WorkLocation.Info.NFCTagID =c.GetString("nfcTagForTask")


		layout := "01/02/2006 15:04"
		startDateInUnix, err := time.Parse(layout, dailyStartTime)
		if err != nil {
			log.Println(err)
		}
		//task.Info.StartDate = startDate.UTC().Unix()
		endDateInUnix, err := time.Parse(layout, dailyEndTime)
		if err != nil {
			log.Println(err)
		}
		//task.Info.EndDate = endDate.Unix()

		tempFitToWorkCheck :=c.GetString("fitToWorkCheck")
		log.Println("tempFitToWorkCheck",tempFitToWorkCheck)
		if tempFitToWorkCheck =="on" {
			WorkLocation.Settings.FitToWorkDisplayStatus ="EachTime"
		} else {
			WorkLocation.Settings.FitToWorkDisplayStatus = "OnceADay"
		}

		var groupKeySlice	[]string
		for j:=0;j<len(userIdArray);j++ {
			tempName := UserOrGroupNameArray[j]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if (userOrGroupSelection[1]) == "User" {
				tempName = tempName[:len(tempName) - 7]
				userName.FullName = tempName
				userName.Status = helpers.StatusActive
				log.Println("tempId", userIdArray[j])
				userMap[userIdArray[j]] = userName
				log.Println("userMap", userMap)
				WorkLocation.Info.WorkLocation = taskLocation
				WorkLocation.Info.CompanyTeamName = companyTeamName
				WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
				WorkLocation.Info.Latitude =latitude
				WorkLocation.Info.Longitude =longitude
				WorkLocation.Info.StartDate =startDateInt
				WorkLocation.Info.EndDate =endDateInt
				WorkLocation.Info.DailyStartDate = startDateInUnix.Unix()
				WorkLocation.Info.DailyEndDate =endDateInUnix.Unix()
				WorkLocation.Settings.DateOfCreation = time.Now().Unix()
				WorkLocation.Settings.Status = helpers.StatusActive
				WorkLocation.Info.LoginType = loginType
				WorkLocation.Info.LogTimeInMinutes = logInMinutesInString
				log.Println("loginType",loginType)
				log.Println("userMap[tempId]", userMap)

			}
			if groupKeySliceForWorkLocation[0] != "" {
				log.Println("w5")
				for i := 0; i < len(groupKeySliceForWorkLocation); i++ {
					groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupKeySliceForWorkLocation[i])
					switch dbStatus {
					case true:
						groupNameAndDetails.GroupName = groupDetails.Info.GroupName
						memberData := reflect.ValueOf(groupDetails.Members)
						for _, key := range memberData.MapKeys() {

							log.Println("status",groupDetails.Members[key.String()].Status)
							if groupDetails.Members[key.String()].Status != helpers.UserStatusDeleted{
								keySliceForGroup = append(keySliceForGroup, key.String())
								MemberNameArray = append(MemberNameArray, groupDetails.Members[key.String()].MemberName)
								break
							}
						}
					case false:
						log.Println(helpers.ServerConnectionError)
					}
					for i := 0; i < len(keySliceForGroup); i++ {
						groupMemberNameForTask.MemberName = MemberNameArray[i]
						groupMemberMap[keySliceForGroup[i]] = groupMemberNameForTask
					}
					groupNameAndDetails.Members = groupMemberMap
					groupMap[string(groupKeySliceForWorkLocation[i])] = groupNameAndDetails
					groupKeySlice = append(groupKeySlice, string(groupKeySliceForWorkLocation[i]))
				}
			}
		}
		WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
		if groupKeySliceForWorkLocation[0] !="" {
			WorkLocation.Info.UsersAndGroupsInWorkLocation.Group = groupMap
		}

		companyName := storedSession.CompanyName
		dbStatus :=WorkLocation.AddWorkLocationToDb(c.AppEngineCtx,companyTeamName,fitToWorkName,WorkLocationBreakTimeSlice,WorkLocationTimeSlice,companyName)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}else {
		usersDetail := models.Users{}
		dbStatus, testUser := companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(testUser)
			for _, key := range dataValue.MapKeys() {

				dataValue := reflect.ValueOf(testUser[key.String()].Users)
				for _, userKeys := range dataValue.MapKeys() {
					//viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKey.String()].FullName+" (User)")
					dbStatus := usersDetail.GetActiveUsersEmailForDropDown(c.AppEngineCtx, userKeys.String(), testUser[key.String()].Users[userKeys.String()].Email, companyTeamName)
					switch dbStatus {
					case true:
						workLocationViewmodel.GroupNameArray = append(workLocationViewmodel.GroupNameArray, testUser[key.String()].Users[userKeys.String()].FullName + " (User)")
						keySliceForGroupAndUser = append(keySliceForGroupAndUser, userKeys.String())
					case false:
						log.Println(helpers.ServerConnectionError)
					}
				}
			}
			allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx, companyTeamName)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allGroups)
				for _, key := range dataValue.MapKeys() {
					if allGroups[key.String()].Settings.Status == "Active" {
						var memberSlice []string

						keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
						workLocationViewmodel.GroupNameArray = append(workLocationViewmodel.GroupNameArray, allGroups[key.String()].Info.GroupName + " (Group)")

						// For selecting members while selecting a group in dropdown
						memberSlice = append(memberSlice, key.String())
						groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
						for _, memberKey := range groupDataValue.MapKeys() {
							memberSlice = append(memberSlice, memberKey.String())
						}
						workLocationViewmodel.GroupMembers = append(workLocationViewmodel.GroupMembers, memberSlice)
						log.Println("iam in trouble", workLocationViewmodel.GroupMembers)

					}

				}
				workLocationViewmodel.UserAndGroupKey = keySliceForGroupAndUser
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}

		workLocationValues := models.IsWorkAssignedToUser(c.AppEngineCtx, companyTeamName)
		log.Println("allWorkLocationData", workLocationValues)
		dataValue := reflect.ValueOf(workLocationValues)

		for _, key := range dataValue.MapKeys() {
			log.Println("alredy key", key.String())
			if workLocationValues[key.String()].Settings.Status != helpers.StatusInActive {
				userDataValues := reflect.ValueOf(workLocationValues[key.String()].Info.UsersAndGroupsInWorkLocation.User)
				for _, userKey := range userDataValues.MapKeys() {
					var tempUserArray []string
					log.Println("alredy strat", workLocationValues[key.String()].Info.StartDate)
					log.Println("alredy end", workLocationValues[key.String()].Info.EndDate)
					startDateFromDbInInt := strconv.FormatInt(int64(workLocationValues[key.String()].Info.StartDate), 10)
					endDateFromDbInInt := strconv.FormatInt(workLocationValues[key.String()].Info.EndDate, 10)
					tempUserArray = append(tempUserArray, userKey.String())
					tempUserArray = append(tempUserArray, startDateFromDbInInt)
					tempUserArray = append(tempUserArray, endDateFromDbInInt)
					workLocationViewmodel.DateValues = append(workLocationViewmodel.DateValues, tempUserArray)
				}
			}

		}
		dbStatus, notificationValue := models.GetAllNotifications(c.AppEngineCtx, companyTeamName)

		switch dbStatus {
		case true:

			notificationOfUser := reflect.ValueOf(notificationValue)
			for _, notificationUserKey := range notificationOfUser.MapKeys() {
				dbStatus, notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx, companyTeamName, notificationUserKey.String())
				switch dbStatus {
				case true:
					notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
					for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
						var NotificationArray []string
						if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead == false {
							notificationCount = notificationCount + 1;
						}
						NotificationArray = append(NotificationArray, notificationUserKey.String())
						NotificationArray = append(NotificationArray, notificationUserKeyForSpecific.String())
						NotificationArray = append(NotificationArray, notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
						NotificationArray = append(NotificationArray, notificationUserValue[notificationUserKeyForSpecific.String()].Message)
						NotificationArray = append(NotificationArray, notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
						NotificationArray = append(NotificationArray, notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
						date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
						NotificationArray = append(NotificationArray, date)
						workLocationViewmodel.NotificationArray = append(workLocationViewmodel.NotificationArray, NotificationArray)

					}
				case false:
				}
			}
		case false:
		}


		//getting fit to work
		var keySliceForFitToWork        []string
		var tempKeySliceFitToWork        []string
		var tempInstructionKeySlice        []string
		var instructionDescription        []string
		var fitToWorkStructSlice        []viewmodels.TaskFitToWork
		var tempfitToWorkStructSlice        [][]viewmodels.TaskFitToWork
		var fitToWorkStruct viewmodels.TaskFitToWork
		fitToWorkById := models.GetSelectedCompanyName(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			fitToWorkDataValues := reflect.ValueOf(fitToWorkById)
			for _, fitToWorkKey := range fitToWorkDataValues.MapKeys() {
				tempKeySliceFitToWork = append(tempKeySliceFitToWork, fitToWorkKey.String())
			}
			for _, eachKey := range tempKeySliceFitToWork {
				if fitToWorkById[eachKey].Settings.Status == helpers.StatusActive {
					keySliceForFitToWork = append(keySliceForFitToWork, eachKey)
					workLocationViewmodel.FitToWorkArray = append(workLocationViewmodel.FitToWorkArray, fitToWorkById[eachKey].FitToWorkName)
					getInstructions := models.GetAllInstructionsOfFitToWorkById(c.AppEngineCtx, companyTeamName, eachKey)
					fitToWorkInstructionValues := reflect.ValueOf(getInstructions)
					for _, fitToWorkInstructionKey := range fitToWorkInstructionValues.MapKeys() {
						tempInstructionKeySlice = append(tempInstructionKeySlice, fitToWorkInstructionKey.String())
					}

					for _, eachInstructionKey := range tempInstructionKeySlice {
						instructionDescription = append(instructionDescription, getInstructions[eachInstructionKey].Description)

					}
					fitToWorkStruct.FitToWorkName = fitToWorkById[eachKey].FitToWorkName
					fitToWorkStruct.Instruction = instructionDescription
					fitToWorkStruct.InstructionKey = tempInstructionKeySlice
					fitToWorkStructSlice = append(fitToWorkStructSlice, fitToWorkStruct)
				}
			}
			workLocationViewmodel.FitToWorkKey = keySliceForFitToWork
			tempfitToWorkStructSlice = append(tempfitToWorkStructSlice, fitToWorkStructSlice)
			workLocationViewmodel.FitToWorkForTask = tempfitToWorkStructSlice
		case false:
			log.Println(helpers.ServerConnectionError)
		}

		workLocationViewmodel.NotificationNumber = notificationCount
		log.Println("tempUserArray", workLocationViewmodel.DateValues)
		workLocationViewmodel.AdminFirstName = storedSession.AdminFirstName
		workLocationViewmodel.AdminLastName = storedSession.AdminLastName
		workLocationViewmodel.ProfilePicture = storedSession.ProfilePicture
		workLocationViewmodel.CompanyTeamName = storedSession.CompanyTeamName
		/*c.Layout = "layout/layout.html"*/
		c.Data["vm"] = workLocationViewmodel
		c.TplName = "template/add-workLocation.html"
	}
}

func (c *WorkLocationcontroller) LoadWorkLocation() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	workLocation,dbStatus:= models.GetAllWorkLocationDetails(c.AppEngineCtx,companyTeamName)
	viewModel := viewmodels.LoadWorkLocationViewModel{}
	var workLocationUserSlice [][]viewmodels.WorkLocationUsers
	var workLocationExposureSlice [][]viewmodels.WorkExposure
	var KeyValues []string
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(workLocation)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			var minUserArray  []string
			if workLocation[k].Settings.Status != helpers.StatusInActive{
				log.Println("workLocation[k].Info",workLocation[k].Info)
				userDataValues :=  reflect.ValueOf(workLocation[k].Info.UsersAndGroupsInWorkLocation.User)
				var userStructSlice []viewmodels.WorkLocationUsers
				for _,userKey :=range userDataValues.MapKeys(){

					var userStruct viewmodels.WorkLocationUsers
					if workLocation[k].Info.UsersAndGroupsInWorkLocation.User[userKey.String()].Status!=helpers.UserStatusDeleted{
						userStruct.Name = workLocation[k].Info.UsersAndGroupsInWorkLocation.User[userKey.String()].FullName
						userStruct.UserKey = k
					}
					userStructSlice = append(userStructSlice, userStruct)
				}
				workLocationUserSlice = append(workLocationUserSlice,userStructSlice)
				log.Println("work location user slice",userStructSlice)
				tempValueSlice = append(tempValueSlice,"")
				tempValueSlice =append(tempValueSlice,workLocation[k].Info.WorkLocation)
				tempValueSlice = append(tempValueSlice,strconv.FormatInt(int64(workLocation[k].Info.StartDate),10))
				tempValueSlice = append(tempValueSlice,strconv.FormatInt(int64(workLocation[k].Info.EndDate),10))
				tempValueSlice =append(tempValueSlice,k)
				KeyValues = append(KeyValues,k)
				viewModel.Values=append(viewModel.Values,tempValueSlice)
				tempValueSlice = tempValueSlice[:0]

				minUserArray = append(minUserArray,workLocation[k].Info.LoginType)
				LogTimeInMinutes := strconv.FormatInt(workLocation[k].Info.LogTimeInMinutes, 10)
				minUserArray = append(minUserArray,LogTimeInMinutes)
				minUserArray =append(minUserArray,workLocation[k].FitToWork.Info.FitToWorkName)
				minUserArray = append(minUserArray,k)
				viewModel.MinUserAndLoginTypeArray =append(viewModel.MinUserAndLoginTypeArray,minUserArray)
				minUserArray =minUserArray[:0]
				dbStatus, taskExposureDetails := models.GetWorkLocationBreakDetailById(c.AppEngineCtx,k)
				switch dbStatus {
				case true:
					exposureValue := reflect.ValueOf(taskExposureDetails)
					var exposureStructSlice []viewmodels.WorkExposure
					for _, key := range exposureValue.MapKeys() {

						var exposureStruct viewmodels.WorkExposure
						exposureStruct.BreakMinute =taskExposureDetails[key.String()].BreakDurationInMinutes
						exposureStruct.WorkingHour =taskExposureDetails[key.String()].BreakStartTimeInMinutes
						exposureStruct.TaskId =k
						exposureStructSlice =append(exposureStructSlice,exposureStruct)

					}
					workLocationExposureSlice = append(workLocationExposureSlice, exposureStructSlice)
				case false:
				}
				viewModel.ExposureArray =workLocationExposureSlice

			}
		}

		log.Println("viewModel.ExposureArray",viewModel.MinUserAndLoginTypeArray)
		dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
		var notificationCount=0
		switch dbStatus {
		case true:

			notificationOfUser := reflect.ValueOf(notificationValue)
			for _, notificationUserKey := range notificationOfUser.MapKeys() {
				dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
				switch dbStatus {
				case true:
					notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
					for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
						var NotificationArray []string
						if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
							notificationCount=notificationCount+1;
						}
						NotificationArray =append(NotificationArray,notificationUserKey.String())
						NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
						NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
						date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
						NotificationArray =append(NotificationArray,date)
						viewModel.NotificationArray=append(viewModel.NotificationArray,NotificationArray)

					}
				case false:
				}
			}
		case false:
		}
		viewModel.NotificationNumber=notificationCount
		viewModel.Users = workLocationUserSlice
		viewModel.Keys = KeyValues
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.CompanyPlan = storedSession.CompanyPlan
		viewModel.AdminFirstName =storedSession.AdminFirstName
		viewModel.AdminLastName =storedSession.AdminLastName
		viewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/worklocation-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}



func (c *WorkLocationcontroller) EditWorkLocation() {

	log.Println("inside  edittttttt");
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	workLocationId := c.Ctx.Input.Param(":worklocationid")
	storedSession := ReadSession(w, r, companyTeamName)
	companyUsers :=models.Company{}
	var keySliceForGroupAndUser []string
	//workLocationViewmodel := viewmodels.AddLocationViewModel{}
	userMap := make(map[string]models.WorkLocationUser)
	groupNameAndDetails := models.WorkLocationGroup{}
	groupMemberNameForTask :=models.GroupMemberNameInWorkLocation{}
	groupMemberMap := make(map[string]models.GroupMemberNameInWorkLocation)
	groupMap := make(map[string]models.WorkLocationGroup)
	var keySliceForGroup [] string
	var MemberNameArray [] string
	var notificationCount = 0
	group := models.Group{}
	viewModelForEdit :=viewmodels.EditWorkLocation{}
	usersDetail :=models.Users{}
	userName :=models.WorkLocationUser{}
	WorkLocation := models.WorkLocation{}
	if r.Method == "POST" {
		exposureTask := c.GetString("exposureBreakTime")
		exposureWorkTime := c.GetString("exposureWorkTime")
		log.Println("exposureTask",exposureTask)
		log.Println("exposureWorkTime",exposureWorkTime)
		fitToWorkName :=c.GetString("fitToWorkName")
		groupKeySliceForWorkLocationString := c.GetString("groupArrayElement")
		UserOrGroupNameArray :=c.GetStrings("userAndGroupName")
		taskLocation := c.GetString("taskLocation")
		dailyEndTime := c.GetString("dailyEndTimeString")
		dailyStartTime := c.GetString("dailyStartTimeString")
		groupKeySliceForWorkLocation :=strings.Split(groupKeySliceForWorkLocationString, ",")
		userIdArray := c.GetStrings("selectedUserNames")
		latitude := c.GetString("latitudeId")
		longitude := c.GetString("longitudeId")
		startDate := c.GetString("startDateTimeStamp")
		endDate := c.GetString("endDateTimeStamp")
		startDateInt , err := strconv.ParseInt(startDate, 10, 64)
		endDateInt, err := strconv.ParseInt(endDate, 10, 64)
		loginType := c.GetString("loginType")
		log.Println("loginType",loginType)
		logInMinutes :=c.GetString("logInMinutes")
		logInMinutesInString, err := strconv.ParseInt(logInMinutes, 10, 64)
		WorkLocationBreakTimeSlice :=strings.Split(exposureTask, ",")
		WorkLocationTimeSlice :=strings.Split(exposureWorkTime, ",")
		WorkLocation.Info.NFCTagID =c.GetString("nfcTagForTask")
		log.Println("WorkLocation.Info.NFCTagID",WorkLocation.Info.NFCTagID)


		layout := "01/02/2006 15:04"
		startDateInUnix, err := time.Parse(layout, dailyStartTime)
		if err != nil {
			log.Println(err)
		}
		//task.Info.StartDate = startDate.UTC().Unix()
		endDateInUnix, err := time.Parse(layout, dailyEndTime)
		if err != nil {
			log.Println(err)
		}
		//task.Info.EndDate = endDate.Unix()

		tempFitToWorkCheck :=c.GetString("fitToWorkCheck")
		log.Println("tempFitToWorkCheck",tempFitToWorkCheck)
		if tempFitToWorkCheck =="on" {
			WorkLocation.Settings.FitToWorkDisplayStatus ="EachTime"
		} else {
			WorkLocation.Settings.FitToWorkDisplayStatus = "OnceADay"
		}
		log.Println("userIdArray",userIdArray)

		var tempArray []string
		for i:=0;i<len(userIdArray);i++{

			exists := false
			for v := 0; v < i; v++ {
				if userIdArray[v] == userIdArray[i] {
					exists = true
					break
				}
			}
			// If no previous element exists, append this one.
			if !exists {
				tempArray = append(tempArray, userIdArray[i])
			}
		}

		var groupKeySlice	[]string
		for j:=0;j<len(tempArray);j++ {
			log.Println("iam in inner loop")
			tempName := UserOrGroupNameArray[j]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if (userOrGroupSelection[1]) == "User" {
				tempName = tempName[:len(tempName) - 7]
				userName.FullName = tempName
				userName.Status = helpers.StatusActive
				log.Println("tempId", userIdArray[j])
				userMap[userIdArray[j]] = userName
				log.Println("userMap", userMap)
				WorkLocation.Info.WorkLocation = taskLocation
				WorkLocation.Info.CompanyTeamName = companyTeamName
				WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
				WorkLocation.Info.Latitude =latitude
				WorkLocation.Info.Longitude =longitude
				WorkLocation.Info.StartDate =startDateInt
				WorkLocation.Info.EndDate =endDateInt
				WorkLocation.Info.DailyStartDate = startDateInUnix.Unix()
				WorkLocation.Info.DailyEndDate =endDateInUnix.Unix()
				WorkLocation.Settings.DateOfCreation = time.Now().Unix()
				WorkLocation.Settings.Status = helpers.StatusActive
				WorkLocation.Info.LoginType = loginType
				WorkLocation.Info.LogTimeInMinutes = logInMinutesInString
				log.Println("loginType",loginType)
				log.Println("userMap[tempId]", userMap)

			}
			if groupKeySliceForWorkLocation[0] != "" {
				log.Println("w5")
				for i := 0; i < len(groupKeySliceForWorkLocation); i++ {
					groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupKeySliceForWorkLocation[i])
					switch dbStatus {
					case true:
						groupNameAndDetails.GroupName = groupDetails.Info.GroupName
						memberData := reflect.ValueOf(groupDetails.Members)
						for _, key := range memberData.MapKeys() {

							log.Println("status",groupDetails.Members[key.String()].Status)
							if groupDetails.Members[key.String()].Status != helpers.UserStatusDeleted{
								keySliceForGroup = append(keySliceForGroup, key.String())
								MemberNameArray = append(MemberNameArray, groupDetails.Members[key.String()].MemberName)
								break
							}
						}
					case false:
						log.Println(helpers.ServerConnectionError)
					}
					for i := 0; i < len(keySliceForGroup); i++ {
						groupMemberNameForTask.MemberName = MemberNameArray[i]
						groupMemberMap[keySliceForGroup[i]] = groupMemberNameForTask
					}
					groupNameAndDetails.Members = groupMemberMap
					groupMap[string(groupKeySliceForWorkLocation[i])] = groupNameAndDetails
					groupKeySlice = append(groupKeySlice, string(groupKeySliceForWorkLocation[i]))
				}
			}
		}
		log.Println("uu1")
		WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
		log.Println("uu2",WorkLocation.Info.UsersAndGroupsInWorkLocation.User)
		if groupKeySliceForWorkLocation[0] !="" {
			log.Println("uu2")
			WorkLocation.Info.UsersAndGroupsInWorkLocation.Group = groupMap
			log.Println("uu3",WorkLocation.Info.UsersAndGroupsInWorkLocation.Group )
		}
		log.Println("op1")
		companyName :=storedSession.CompanyName
		dbStatus :=WorkLocation.EditWorkLocationToDb(c.AppEngineCtx,workLocationId,companyTeamName,fitToWorkName,WorkLocationBreakTimeSlice,WorkLocationTimeSlice,companyName)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else{

		dbStatus ,testUser:= companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx,companyTeamName)

		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(testUser)
			for _, key := range dataValue.MapKeys() {

				dataValue := reflect.ValueOf(testUser[key.String()].Users)
				for _, userKeys := range dataValue.MapKeys() {
					//viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKey.String()].FullName+" (User)")
					dbStatus := usersDetail.GetActiveUsersEmailForDropDown(c.AppEngineCtx, userKeys.String(), testUser[key.String()].Users[userKeys.String()].Email, companyTeamName)
					switch dbStatus {
					case true:
						viewModelForEdit.GroupNameArray = append(viewModelForEdit.GroupNameArray, testUser[key.String()].Users[userKeys.String()].FullName + " (User)")
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
					if allGroups[key.String()].Settings.Status =="Active"{
						var memberSlice []string

						keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
						viewModelForEdit.GroupNameArray = append(viewModelForEdit.GroupNameArray, allGroups[key.String()].Info.GroupName+" (Group)")

						// For selecting members while selecting a group in dropdown
						memberSlice = append(memberSlice, key.String())
						groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
						for _, memberKey := range groupDataValue.MapKeys()  {
							memberSlice = append(memberSlice, memberKey.String())
						}
						viewModelForEdit.GroupMembers = append(viewModelForEdit.GroupMembers, memberSlice)
						log.Println("iam in trouble",viewModelForEdit.GroupMembers)

					}

					//viewModelForEdit.WorkLocation = allGroups[key.String()]
				}

				viewModelForEdit.WorkLogId = workLocationId
				viewModelForEdit.UserAndGroupKey=keySliceForGroupAndUser
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		workLocation,dbStatus:= models.GetAllWorkLocationDetailsByWorkId(c.AppEngineCtx,workLocationId)
		switch dbStatus {
		case true:
			viewModelForEdit.PageType = helpers.SelectPageForEdit
			userDataValues :=  reflect.ValueOf(workLocation.Info.UsersAndGroupsInWorkLocation.User)

			for _,userKey :=range userDataValues.MapKeys(){
				viewModelForEdit.UserNameToEdit = append(viewModelForEdit.UserNameToEdit,workLocation.Info.UsersAndGroupsInWorkLocation.User[userKey.String()].FullName)
				viewModelForEdit.UsersKey = append(viewModelForEdit.UsersKey,userKey.String())
			}
			startTime := time.Unix(workLocation.Info.DailyStartDate, 0)
			startTimeOfWorkLocation := startTime.String()[11:16]
			endTime := time.Unix(workLocation.Info.DailyEndDate,0)
			endTimeOfWorkLocation := endTime.String()[11:16]
			log.Println("endTimeOfWorkLocation",endTimeOfWorkLocation)
			log.Println("startTimeOfWorkLocation",startTimeOfWorkLocation)
			viewModelForEdit.FitToWorkCheck = workLocation.Settings.FitToWorkDisplayStatus
			viewModelForEdit.LatitudeForEditing = workLocation.Info.Latitude
			viewModelForEdit.LongitudeForEditing = workLocation.Info.Longitude
			viewModelForEdit.WorkLocation = workLocation.Info.WorkLocation
			viewModelForEdit.DailyStartTime = startTimeOfWorkLocation
			viewModelForEdit.DailyEndTime = endTimeOfWorkLocation
			viewModelForEdit.StartDate = workLocation.Info.StartDate
			viewModelForEdit.EndDate = workLocation.Info.EndDate
			viewModelForEdit.LoginType = workLocation.Info.LoginType
			viewModelForEdit.LogInMin = workLocation.Info.LogTimeInMinutes
			viewModelForEdit.FitToWorkName = workLocation.FitToWork.Info.FitToWorkName
			viewModelForEdit.NFCTagId = workLocation.Info.NFCTagID

		case false:
			log.Println(helpers.ServerConnectionError)
		}

	}


	workLocationValues := models.IsWorkAssignedToUser(c.AppEngineCtx,companyTeamName)
	log.Println("allWorkLocationData",workLocationValues)
	dataValue := reflect.ValueOf(workLocationValues)

	for _, key := range dataValue.MapKeys() {
		log.Println("alredy key",key.String())
		if workLocationValues[key.String()].Settings.Status != helpers.StatusInActive{
			userDataValues :=  reflect.ValueOf(workLocationValues[key.String()].Info.UsersAndGroupsInWorkLocation.User)
			for _,userKey :=range userDataValues.MapKeys(){
				var tempUserArray []string
				log.Println("alredy strat",workLocationValues[key.String()].Info.StartDate)
				log.Println("alredy end",workLocationValues[key.String()].Info.EndDate )
				startDateFromDbInInt:= strconv.FormatInt(int64(workLocationValues[key.String()].Info.StartDate), 10)
				endDateFromDbInInt:=strconv.FormatInt(workLocationValues[key.String()].Info.EndDate, 10)
				tempUserArray = append(tempUserArray,userKey.String())
				tempUserArray = append(tempUserArray, startDateFromDbInInt)
				tempUserArray = append(tempUserArray, endDateFromDbInInt)
				viewModelForEdit.DateValues = append(viewModelForEdit.DateValues,tempUserArray)
			}
		}

	}
	dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
	//var notificationCount =0
	switch dbStatus {
	case true:

		notificationOfUser := reflect.ValueOf(notificationValue)
		for _, notificationUserKey := range notificationOfUser.MapKeys() {
			dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
			switch dbStatus {
			case true:
				notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
				for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
					var NotificationArray []string
					if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
						notificationCount=notificationCount+1;
					}
					NotificationArray =append(NotificationArray,notificationUserKey.String())
					NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
					date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
					NotificationArray =append(NotificationArray,date)
					viewModelForEdit.NotificationArray=append(viewModelForEdit.NotificationArray,NotificationArray)

				}
			case false:
			}
		}
	case false:
	}



	var keySliceForFitToWork        []string
	var tempKeySliceFitToWork        []string
	var tempInstructionKeySlice        []string
	var instructionDescription        []string
	var fitToWorkStructSlice        []viewmodels.TaskFitToWork
	var tempfitToWorkStructSlice        [][]viewmodels.TaskFitToWork
	var fitToWorkStruct viewmodels.TaskFitToWork
	fitToWorkById := models.GetSelectedCompanyName(c.AppEngineCtx, companyTeamName)
	switch dbStatus {
	case true:
		fitToWorkDataValues := reflect.ValueOf(fitToWorkById)
		for _, fitToWorkKey := range fitToWorkDataValues.MapKeys() {
			tempKeySliceFitToWork = append(tempKeySliceFitToWork, fitToWorkKey.String())
		}
		for _, eachKey := range tempKeySliceFitToWork {
			if fitToWorkById[eachKey].Settings.Status == helpers.StatusActive {
				keySliceForFitToWork = append(keySliceForFitToWork, eachKey)
				viewModelForEdit.FitToWorkArray = append(viewModelForEdit.FitToWorkArray, fitToWorkById[eachKey].FitToWorkName)
				getInstructions := models.GetAllInstructionsOfFitToWorkById(c.AppEngineCtx, companyTeamName, eachKey)
				fitToWorkInstructionValues := reflect.ValueOf(getInstructions)
				for _, fitToWorkInstructionKey := range fitToWorkInstructionValues.MapKeys() {
					tempInstructionKeySlice = append(tempInstructionKeySlice, fitToWorkInstructionKey.String())
				}

				for _, eachInstructionKey := range tempInstructionKeySlice {
					instructionDescription = append(instructionDescription, getInstructions[eachInstructionKey].Description)

				}
				fitToWorkStruct.FitToWorkName = fitToWorkById[eachKey].FitToWorkName
				fitToWorkStruct.Instruction = instructionDescription
				fitToWorkStruct.InstructionKey = tempInstructionKeySlice
				fitToWorkStructSlice = append(fitToWorkStructSlice, fitToWorkStruct)
			}
		}
		viewModelForEdit.FitToWorkKey = keySliceForFitToWork
		tempfitToWorkStructSlice = append(tempfitToWorkStructSlice, fitToWorkStructSlice)
		viewModelForEdit.FitToWorkForTask = tempfitToWorkStructSlice
	case false:
		log.Println(helpers.ServerConnectionError)
	}

	//var fitToWorkSlice			[]string
	var WorkTime                            []string
	var BreakTime				[]string
	dbStatus, WorkBreak := models.GetWorkLocationBreakDetailById(c.AppEngineCtx, workLocationId)
	switch dbStatus {
	case true:
		workValue := reflect.ValueOf(WorkBreak)
		for _, key := range workValue.MapKeys() {
			breakHourInInt, err := strconv.Atoi(WorkBreak[key.String()].BreakDurationInMinutes)
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
			BreakTime = append(BreakTime,breakTimeForTask)

			workHourInInt, err := strconv.Atoi(WorkBreak[key.String()].BreakStartTimeInMinutes)
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

			WorkTime = append(WorkTime,workTimeForTask)

		}
	case false:
		log.Println(helpers.ServerConnectionError)
	}

	viewModelForEdit.BreakTime =BreakTime
	viewModelForEdit.WorkTime = WorkTime
	viewModelForEdit.NotificationNumber=notificationCount
	viewModelForEdit.CompanyTeamName = storedSession.CompanyTeamName
	viewModelForEdit.CompanyPlan = storedSession.CompanyPlan
	viewModelForEdit.AdminFirstName =storedSession.AdminFirstName
	viewModelForEdit.AdminLastName =storedSession.AdminLastName
	viewModelForEdit.ProfilePicture =storedSession.ProfilePicture
	log.Println("viewModelForEdit break time",viewModelForEdit.BreakTime)
	log.Println("WorkTime",WorkTime)

	c.Data["vm"] = viewModelForEdit
	c.TplName = "template/add-workLocation.html"

}


func (c *WorkLocationcontroller) DeleteWorkLocation() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	workLocationId := c.Ctx.Input.Param(":workLocationId")
	log.Println("workLocationId",workLocationId)
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println(storedSession)
	dbStatus := models.DeleteWorkLog(c.AppEngineCtx,workLocationId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}


