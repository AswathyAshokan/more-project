/*Created By Farsana*/
package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"reflect"
	"log"
	"app/passporte/helpers"
	"time"
	"strconv"
	"bytes"
	"strings"

)

type GroupController struct {
	BaseController
}

// Add new groups to database
func (c *GroupController) AddGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	groupViewModel := viewmodels.AddGroupViewModel{}
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		log.Println("cp3")
		group := models.Group{}
		members := models.GroupMembers{}
		group.Info.GroupName = c.GetString("groupName")
		tempGroupId := c.GetStrings("selectedUserIds")
		group.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		group.Settings.Status = helpers.StatusActive
		tempGroupMembers := c.GetStrings("selectedUserNames")
		group.Info.CompanyTeamName = storedSession.CompanyTeamName
		tempMembersMap := make(map[string]models.GroupMembers)
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberName = tempGroupMembers[i]
			tempMembersMap[tempGroupId[i]] = members
		}
		group.Members = tempMembersMap
		dbStatus := group.AddGroupToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		groupUser := models.Users{}

		var keySlice []string
		var allUserNames [] string
		allUserDetails, dbStatus := groupUser.TakeGroupMemberName(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(allUserDetails)

			for _, groupKey := range dataValue.MapKeys() {
				keySlice = append(keySlice, groupKey.String())
			}
			for _, k := range keySlice {
				if allUserDetails[k].Status != helpers.UserStatusDeleted {
					allUserNames = append(allUserNames, allUserDetails[k].FullName)
					groupViewModel.GroupMembers = allUserNames
					groupViewModel.GroupKey = keySlice
				}
			}
			groupViewModel.CompanyTeamName = storedSession.CompanyTeamName
			groupViewModel.CompanyPlan = storedSession.CompanyPlan
			groupViewModel.GroupKey = keySlice
			groupViewModel.PageType = helpers.SelectPageForAdd
			groupViewModel.AdminFirstName = storedSession.AdminFirstName
			groupViewModel.AdminLastName = storedSession.AdminLastName
			groupViewModel.ProfilePicture = storedSession.ProfilePicture
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
	c.Data["vm"] = groupViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/add-group.html"
}

// show the details of whole group from database
func (c *GroupController) GroupDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allGroups)
		groupViewModel := viewmodels.GroupList{}

		//Collecting group keys
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		//collecting group details
		for _, groupKey := range keySlice {
			var tempValueSlice []string
			membersNumber := len(allGroups[groupKey].Members)

			groupUsers := allGroups[groupKey].Members
			userDataValue := reflect.ValueOf(groupUsers)

			// collecting group member keys
			var userKeySlice []string
			for _, userKey := range userDataValue.MapKeys() {
				userKeySlice = append(userKeySlice, userKey.String())
			}

			//collecting group member details
			tempUserNames := ""
			var buffer bytes.Buffer
			for _, userKey := range userKeySlice {
				if len(tempUserNames) == 0{
					buffer.WriteString(groupUsers[userKey].MemberName)
					tempUserNames = buffer.String()
					buffer.Reset()
				} else {
					buffer.WriteString(tempUserNames)
					buffer.WriteString(", ")
					buffer.WriteString(groupUsers[userKey].MemberName)
					tempUserNames = buffer.String()
					buffer.Reset()
				}
			}
			if allGroups[groupKey].Settings.Status != helpers.UserStatusDeleted{
				tempValueSlice = append(tempValueSlice, allGroups[groupKey].Info.GroupName)
				tempValueSlice = append(tempValueSlice, strconv.Itoa(membersNumber))
				tempValueSlice = append(tempValueSlice, tempUserNames)
				tempValueSlice = append(tempValueSlice,groupKey)
				groupViewModel.Values = append(groupViewModel.Values, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}

		}
		groupViewModel.Keys = keySlice
		groupViewModel.CompanyTeamName = storedSession.CompanyTeamName
		groupViewModel.CompanyPlan = storedSession.CompanyPlan
		groupViewModel.AdminFirstName = storedSession.AdminFirstName
		groupViewModel.AdminLastName = storedSession.AdminLastName
		groupViewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = groupViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/group-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

// To delete each group from database
func (c *GroupController) DeleteGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	groupId :=c.Ctx.Input.Param(":groupId")
	group := models.Group{}
	dbStatus :=group.DeleteGroup(c.AppEngineCtx, groupId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))

	}
}

//Edit the Group Details
func (c *GroupController) EditGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	groupId := c.Ctx.Input.Param(":groupId")
	group := models.Group{}

	if r.Method == "POST" {
		m := make(map[string]models.GroupMembers)
		members := models.GroupMembers{}
		group.Info.GroupName = c.GetString("groupName")
		group.Info.CompanyTeamName = companyTeamName
		tempGroupId := c.GetStrings("selectedUserIds")
		tempGroupMembers := c.GetStrings("selectedUserNames")
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberName = tempGroupMembers[i]
			m[tempGroupId[i]] = members
		}
		group.Members = m
		dbStatus := group.UpdateGroupDetails(c.AppEngineCtx, groupId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		groupUser := models.Users{}
		viewModel := viewmodels.EditGroupViewModel{}
		var allUserNames [] string
		var keySlice []string
		// Getting all Data for page load...
		allUserDetails, dbStatus := groupUser.TakeGroupMemberName(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(allUserDetails)

			for _, groupKey := range dataValue.MapKeys() {
				keySlice = append(keySlice, groupKey.String())
			}
			for _, k := range keySlice {
				if allUserDetails[k].Status != helpers.UserStatusDeleted {
					allUserNames = append(allUserNames, allUserDetails[k].FullName)
					viewModel.GroupMembers = allUserNames
					viewModel.GroupKey = keySlice
					viewModel.PageType = helpers.SelectPageForEdit
				}
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		//Selecting Data which is to be edited...
		groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupId)
		switch dbStatus {
		case true:
			viewModel.GroupNameToEdit = groupDetails.Info.GroupName
			memberData := reflect.ValueOf(groupDetails.Members)

			for _, selectedMemberKey := range memberData.MapKeys(){
				viewModel.GroupMembersToEdit = append(viewModel.GroupMembersToEdit, selectedMemberKey.String())
			}
			viewModel.PageType = helpers.SelectPageForEdit
			viewModel.CompanyTeamName = storedSession.CompanyTeamName
			viewModel.CompanyPlan = storedSession.CompanyPlan
			viewModel.GroupId = groupId
			viewModel.AdminFirstName = storedSession.AdminFirstName
			viewModel.AdminLastName = storedSession.AdminLastName
			viewModel.ProfilePicture = storedSession.ProfilePicture
			c.Data["vm"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-group.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}

func (c *GroupController)  GroupNameCheck(){
	w := c.Ctx.ResponseWriter
	groupName := c.GetString("groupName")
	pageType := c.Ctx.Input.Param(":type")
	oldName := c.Ctx.Input.Param(":oldName")
	if pageType == "edit" && strings.Compare(oldName, groupName) == 0 {
		w.Write([]byte("true"))
	} else {
		dbStatus := models.IsGroupNameUsed(c.AppEngineCtx, groupName)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}





}

//functions for dependency test

func (c *GroupController)LoadDeleteGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	log.Println("inside delete")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	groupId := c.Ctx.Input.Param(":groupId")
	user := models.TasksGroup{}
	dbStatus, groupDetail := user.IsGroupUsedForTask(c.AppEngineCtx, groupId)
	log.Println("status", dbStatus)
	log.Println(groupDetail)
	switch dbStatus {
	case true:
		log.Println("true")
		if len(groupDetail) != 0 {
			dataValue := reflect.ValueOf(groupDetail)
			for _, key := range dataValue.MapKeys() {
				if groupDetail[key.String()].TasksGroupStatus == helpers.StatusActive {
					log.Println("insideeee fgjgfjh")
					w.Write([]byte("true"))
					break
				} else {
					log.Println("false")
					w.Write([]byte("false"))
				}
			}
		} else {
			w.Write([]byte("false"))
		}
	case false :
		w.Write([]byte("false"))
	}
}
func (c *GroupController) DeleteGroupIfNotInTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	groupId := c.Ctx.Input.Param(":groupId")
	user :=models.Group{}
	log.Println("inside deletion of cotact")
	group :=models.TasksGroup{}
	var TaskSlice []string
	dbStatus,groupDetails := group.IsGroupUsedForTask(c.AppEngineCtx, groupId)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(groupDetails)
		for _, key := range dataValue.MapKeys() {
			TaskSlice = append(TaskSlice, key.String())
		}
		dbStatus := user.DeleteGroupFromDB(c.AppEngineCtx, groupId,TaskSlice)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false :
			w.Write([]byte("false"))
		}
	}
}



func (c *GroupController) RemoveGroupFromTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	groupId := c.Ctx.Input.Param(":groupId")
	log.Println("hiiii",groupId)
	//contact :=models.TasksContact{}
	//var TaskSlice []string
	//dbStatus,contactDetails := contact.IsContactUsedForTask(c.AppEngineCtx, contactId)
	//switch dbStatus {
	//case true:
	//	dataValue := reflect.ValueOf(contactDetails)
	//	for _, key := range dataValue.MapKeys() {
	//		TaskSlice=append(TaskSlice,key.String())
	//	}
	//
	//	dbStatus := contact.DeleteContactFromTask(c.AppEngineCtx, contactId, TaskSlice)
	//	switch dbStatus {
	//	case true:
	//		w.Write([]byte("true"))
	//	case false:
	//		w.Write([]byte("false"))
	//	}
	//case false:
	//	log.Println("false")
	user :=models.Group{}
	log.Println("group id",groupId)
	dbStatus := user.DeleteGroupFromDBForNonTask(c.AppEngineCtx, groupId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}
}



