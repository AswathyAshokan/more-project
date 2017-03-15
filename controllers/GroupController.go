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
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		log.Println("cp3")
		group := models.Group{}
		members := models.GroupMembers{}
		group.Info.GroupName = c.GetString("groupName")
		tempGroupId := c.GetStrings("selectedUserIds")
		group.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		group.Settings.Status = "inactive"
		tempGroupMembers := c.GetStrings("selectedUserNames")
		log.Println("group details :",group)
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
		companyUsers :=models.Company{}
		groupViewModel := viewmodels.AddGroupViewModel{}
		var keySlice []string
		dbStatus,GroupMembers :=companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx,companyTeamName)  // retrive all the keys of a users
		switch dbStatus {

		case true:
			dataValue := reflect.ValueOf(GroupMembers)
			for _, key := range dataValue.MapKeys() {
				dataValue := reflect.ValueOf(GroupMembers[key.String()].Users)
				for _, userKey := range dataValue.MapKeys() {
					groupViewModel.GroupMembers   = append(groupViewModel.GroupMembers ,GroupMembers[key.String()].Users[userKey.String()].FullName)
					keySlice = append(keySlice, userKey.String())
				}

			}
				groupViewModel.CompanyTeamName = storedSession.CompanyTeamName
				groupViewModel.CompanyPlan = storedSession.CompanyPlan
				groupViewModel.GroupKey = keySlice
				groupViewModel.PageType = helpers.SelectPageForAdd
				c.Data["vm"] = groupViewModel
				c.Layout = "layout/layout.html"
				c.TplName = "template/add-group.html"
			case false:
				log.Println(helpers.ServerConnectionError)
			}

	}
}

// show the details of whole group from database
func (c *GroupController) GroupDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("The userDetails stored in session:",storedSession)
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
			tempValueSlice = append(tempValueSlice, allGroups[groupKey].Info.GroupName)
			tempValueSlice = append(tempValueSlice, strconv.Itoa(membersNumber))
			tempValueSlice = append(tempValueSlice, tempUserNames)
			groupViewModel.Values = append(groupViewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		groupViewModel.Keys = keySlice
		groupViewModel.CompanyTeamName = storedSession.CompanyTeamName
		groupViewModel.CompanyPlan = storedSession.CompanyPlan
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
	log.Println("cp1")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	groupId := c.Ctx.Input.Param(":groupId")
	log.Println("team name",companyTeamName )
	group := models.Group{}

	if r.Method == "POST" {
		log.Println("cp3")
		m := make(map[string]models.GroupMembers)
		members := models.GroupMembers{}
		group.Info.GroupName = c.GetString("groupName")
		group.Info.CompanyTeamName=companyTeamName
		tempGroupId := c.GetStrings("selectedUserIds")
		tempGroupMembers := c.GetStrings("selectedUserNames")
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberName = tempGroupMembers[i]
			m[tempGroupId[i]] = members
		}
		group.Members = m
		dbStatus := group.UpdateGroupDetails(c.AppEngineCtx, groupId)
		log.Println("cp4",dbStatus)

		switch dbStatus {
		case true:
			//http.Redirect(w,r,"/group",301)
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		log.Println("cp2")
		groupUser := models.Users{}
		viewModel := viewmodels.EditGroupViewModel{}
		GroupMembers, dbStatus := groupUser.GetUsersForDropdown(c.AppEngineCtx)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(GroupMembers)        // To store data values of slice
			var keySlice []string        // To store keys of the slice
			for _, groupKey := range dataValue.MapKeys() {
				keySlice = append(keySlice, groupKey.String())
			}

			// Getting all Data for page load...
			GroupMemberName, dbStatus := groupUser.TakeGroupMemberName(c.AppEngineCtx, keySlice)
			switch dbStatus {
			case true:
				viewModel.GroupMembers = GroupMemberName
				viewModel.GroupKey = keySlice
				viewModel.PageType = helpers.SelectPageForEdit
			case false:
				log.Println(helpers.ServerConnectionError)
			}

			//Selecting Data which is to be edited...
			log.Println("group key11 :",groupId)
			groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupId)
			switch dbStatus {
			case true:
				viewModel.GroupNameToEdit = groupDetails.Info.GroupName
				memberData := reflect.ValueOf(groupDetails.Members)

				for _, selectedMemberKey := range memberData.MapKeys() {
					viewModel.GroupMembersToEdit = append(viewModel.GroupMembersToEdit, selectedMemberKey.String())
				}
				viewModel.PageType = helpers.SelectPageForEdit
				viewModel.CompanyTeamName = storedSession.CompanyTeamName
				viewModel.CompanyPlan = storedSession.CompanyPlan
				viewModel.GroupId = groupId
				c.Data["vm"] = viewModel
				c.Layout = "layout/layout.html"
				c.TplName = "template/add-group.html"
			case false:
				log.Println(helpers.ServerConnectionError)
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		}
	}
}

func (c *GroupController)  GroupNameCheck(){
	w := c.Ctx.ResponseWriter
	//groupName :=c.Ctx.Input.Param("groupName")
	groupName := c.GetString("groupName")
	pageType := c.Ctx.Input.Param(":type")
	oldName := c.Ctx.Input.Param(":oldName")
	log.Println("groupName:",groupName)
	log.Println("pageType:",pageType)
	log.Println("oldName:",oldName)

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




