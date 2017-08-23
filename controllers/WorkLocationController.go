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
)

type WorkLocationcontroller struct {
	BaseController
}

func (c *WorkLocationcontroller) AddWorkLocation() {
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
	group := models.Group{}
	userName :=models.WorkLocationUser{}
	WorkLocation := models.WorkLocation{}
	if r.Method == "POST" {
		groupKeySliceForWorkLocationString := c.GetString("groupArrayElement")
		UserOrGroupNameArray :=c.GetStrings("userAndGroupName")
		taskLocation := c.GetString("taskLocation")
		groupKeySliceForWorkLocation :=strings.Split(groupKeySliceForWorkLocationString, ",")
		userIdArray := c.GetStrings("selectedUserNames")
		var groupKeySlice	[]string
		for j:=0;j<len(userIdArray);j++ {
			tempName := UserOrGroupNameArray[j]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if (userOrGroupSelection[1]) == "User" {
				log.Println("cp1")
				tempName = tempName[:len(tempName) - 7]
				userName.FullName = tempName
				userName.Status = helpers.StatusActive
				log.Println("tempId", userIdArray[j])
				userMap[userIdArray[j]] = userName
				log.Println("userMap", userMap)
				WorkLocation.Info.WorkLocation = taskLocation
				WorkLocation.Info.CompanyTeamName = companyTeamName
				WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
				WorkLocation.Settings.DateOfCreation = time.Now().Unix()
				WorkLocation.Settings.Status = helpers.StatusActive
				log.Println("userMap[tempId]", userMap)

			}
			if groupKeySliceForWorkLocation[0] != "" {
				for i := 0; i < len(groupKeySliceForWorkLocation); i++ {
					groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupKeySliceForWorkLocation[i])
					switch dbStatus {
					case true:
						groupNameAndDetails.GroupName = groupDetails.Info.GroupName
						memberData := reflect.ValueOf(groupDetails.Members)
						for _, key := range memberData.MapKeys() {
							keySliceForGroup = append(keySliceForGroup, key.String())
							log.Println("status",groupDetails.Members[key.String()].Status)
							if groupDetails.Members[key.String()].Status != helpers.UserStatusDeleted{
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
		dbStatus :=WorkLocation.AddWorkLocationToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}else {
		usersDetail :=models.Users{}
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
						workLocationViewmodel.GroupNameArray = append(workLocationViewmodel.GroupNameArray, testUser[key.String()].Users[userKeys.String()].FullName + " (User)")
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
						workLocationViewmodel.GroupNameArray = append(workLocationViewmodel.GroupNameArray, allGroups[key.String()].Info.GroupName+" (Group)")

						// For selecting members while selecting a group in dropdown
						memberSlice = append(memberSlice, key.String())
						groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
						for _, memberKey := range groupDataValue.MapKeys()  {
							memberSlice = append(memberSlice, memberKey.String())
						}
						workLocationViewmodel.GroupMembers = append(workLocationViewmodel.GroupMembers, memberSlice)
						log.Println("iam in trouble",workLocationViewmodel.GroupMembers)

					}


				}
				workLocationViewmodel.UserAndGroupKey=keySliceForGroupAndUser
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}

	}
	workLocationViewmodel.AdminFirstName = storedSession.AdminFirstName
	workLocationViewmodel.AdminLastName = storedSession.AdminLastName
	workLocationViewmodel.ProfilePicture =storedSession.ProfilePicture
	workLocationViewmodel.CompanyTeamName = companyTeamName
	/*c.Layout = "layout/layout.html"*/
	c.Data["vm"] = workLocationViewmodel
	c.TplName = "template/add-workLocation.html"
}

func (c *WorkLocationcontroller) LoadWorkLocation() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	workLocation,dbStatus:= models.GetAllWorkLocationDetails(c.AppEngineCtx,companyTeamName)
	log.Println("workLocation",workLocation)
	viewModel := viewmodels.LoadWorkLocationViewModel{}
	var workLocationUserSlice [][]viewmodels.WorkLocationUsers
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(workLocation)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			var userStructSlice []viewmodels.WorkLocationUsers

			if workLocation[k].Settings.Status != helpers.UserStatusDeleted{
				userDataValues :=  reflect.ValueOf(workLocation[k].Info.UsersAndGroupsInWorkLocation.User)

				for _,userKey :=range userDataValues.MapKeys(){
					var userStruct viewmodels.WorkLocationUsers
					userStruct.Name = workLocation[k].Info.UsersAndGroupsInWorkLocation.User[userKey.String()].FullName
					userStruct.UserKey = k
					userStructSlice = append(userStructSlice, userStruct)
				}
				workLocationUserSlice = append(workLocationUserSlice,userStructSlice)
				tempValueSlice =append(tempValueSlice,workLocation[k].Info.WorkLocation)
				tempValueSlice =append(tempValueSlice,k)
				viewModel.Values=append(viewModel.Values,tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}

		}
		viewModel.Users = workLocationUserSlice
		viewModel.Keys = keySlice
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
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	workLocationId := c.Ctx.Input.Param(":worklocationid")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("iiii",storedSession,workLocationId)
	groupNameAndDetails := models.WorkLocationGroup{}
	groupMemberNameForTask :=models.GroupMemberNameInWorkLocation{}
	groupMemberMap := make(map[string]models.GroupMemberNameInWorkLocation)
	userName :=models.WorkLocationUser{}
	WorkLocation := models.WorkLocation{}
	viewModelForEdit := viewmodels.EditWorkLocation{}
	companyUsers :=models.Company{}
	userMap := make(map[string]models.WorkLocationUser)
	var keySliceForGroup [] string
	var keySliceForGroupAndUser []string
	groupMap := make(map[string]models.WorkLocationGroup)
	var MemberNameArray [] string
	group := models.Group{}
	if r.Method == "POST" {
		groupKeySliceForWorkLocationString := c.GetString("groupArrayElement")
		UserOrGroupNameArray :=c.GetStrings("userAndGroupName")
		taskLocation := c.GetString("taskLocation")
		groupKeySliceForWorkLocation :=strings.Split(groupKeySliceForWorkLocationString, ",")
		userIdArray := c.GetStrings("selectedUserNames")
		var groupKeySlice	[]string
		for j:=0;j<len(userIdArray);j++ {
			tempName := UserOrGroupNameArray[j]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if (userOrGroupSelection[1]) == "User" {
				log.Println("cp1")
				tempName = tempName[:len(tempName) - 7]
				userName.FullName = tempName
				userName.Status = helpers.StatusActive
				log.Println("tempId", userIdArray[j])
				userMap[userIdArray[j]] = userName
				log.Println("userMap", userMap)
				WorkLocation.Info.WorkLocation = taskLocation
				WorkLocation.Info.CompanyTeamName = companyTeamName
				WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
				WorkLocation.Settings.DateOfCreation = time.Now().Unix()
				WorkLocation.Settings.Status = helpers.StatusActive
				log.Println("userMap[tempId]", userMap)

			}
			if groupKeySliceForWorkLocation[0] != "" {
				for i := 0; i < len(groupKeySliceForWorkLocation); i++ {
					groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupKeySliceForWorkLocation[i])
					switch dbStatus {
					case true:
						groupNameAndDetails.GroupName = groupDetails.Info.GroupName
						memberData := reflect.ValueOf(groupDetails.Members)
						for _, key := range memberData.MapKeys() {
							keySliceForGroup = append(keySliceForGroup, key.String())
							log.Println("status",groupDetails.Members[key.String()].Status)
							if groupDetails.Members[key.String()].Status != helpers.UserStatusDeleted{
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
		dbStatus :=WorkLocation.EditWorkLocationToDb(c.AppEngineCtx,workLocationId,companyTeamName)
		switch dbStatus {
		case true:
			log.Println("true")
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}






	} else{
		usersDetail :=models.Users{}
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
		log.Println("workLocation",workLocation,dbStatus)
		switch dbStatus {
		case true:
			viewModelForEdit.PageType = helpers.SelectPageForEdit
			userDataValues :=  reflect.ValueOf(workLocation.Info.UsersAndGroupsInWorkLocation.User)

			for _,userKey :=range userDataValues.MapKeys(){
				viewModelForEdit.UserNameToEdit = append(viewModelForEdit.UserNameToEdit,workLocation.Info.UsersAndGroupsInWorkLocation.User[userKey.String()].FullName)
				viewModelForEdit.UsersKey = append(viewModelForEdit.UsersKey,userKey.String())
			}
			viewModelForEdit.WorkLocation = workLocation.Info.WorkLocation

		case false:
			log.Println(helpers.ServerConnectionError)
		}


	}
	viewModelForEdit.CompanyTeamName = storedSession.CompanyTeamName
	viewModelForEdit.CompanyPlan = storedSession.CompanyPlan
	viewModelForEdit.AdminFirstName =storedSession.AdminFirstName
	viewModelForEdit.AdminLastName =storedSession.AdminLastName
	viewModelForEdit.ProfilePicture =storedSession.ProfilePicture
	c.Data["vm"] = viewModelForEdit
	c.TplName = "template/add-workLocation.html"

}
