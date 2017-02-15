/*Created By Farsana*/
package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"reflect"
	"log"
	"app/passporte/helpers"
	"strconv"
	"bytes"
)

type GroupController struct {
	BaseController
}

// Add new groups to database
func (c *GroupController) AddGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		group := models.Group{}
		members := models.GroupMembers{}
		group.GroupName = c.GetString("groupName")
		tempGroupId := c.GetStrings("selectedUserIds")
		tempGroupMembers := c.GetStrings("selectedUserNames")
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberId = tempGroupId[i]
			members.MemberName = tempGroupMembers[i]
			group.Members = append(group.Members, members)
		}
		dbStatus := group.AddGroupToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		groupUser := models.Group{}
		GroupMembers,dbStatus :=groupUser.GetUsersForDropdown(c.AppEngineCtx)  // retrive all the keys of a users
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(GroupMembers)	// To store data values of slice
			var keySlice []string	// To store keys of the slice
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			group := models.Group{}
			GroupMemberName,dbStatus:= group.TakeGroupMemberName(c.AppEngineCtx, keySlice)
			switch dbStatus {
			case true:
				groupViewModel := viewmodels.AddGroupViewModel{}
				groupViewModel.GroupMembers = GroupMemberName
				groupViewModel.GroupKey = keySlice
				groupViewModel.PageType = helpers.SelectPageForAdd
				c.Data["vm"] = groupViewModel
				c.Layout = "layout/layout.html"
				c.TplName = "template/add-group.html"
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}

// show the details of whole group from database
func (c *GroupController) GroupDetails() {
	//r := c.Ctx.Request
	allGroups,dbStatus := models.GetAllGroupDetails(c.AppEngineCtx)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allGroups)
		groupViewModel := viewmodels.GroupList{}
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			membersNumber := len(allGroups[k].Members)
			tempValueSlice = append(tempValueSlice, allGroups[k].GroupName)
			tempValueSlice = append(tempValueSlice, strconv.Itoa(membersNumber))
			tempUserNames := ""
			var buffer bytes.Buffer
			for i := 0; i < membersNumber; i++ {
				if len(tempUserNames) == 0{
					buffer.WriteString(allGroups[k].Members[i].MemberName)
					tempUserNames = buffer.String()
					buffer.Reset()
				} else {
					buffer.WriteString(tempUserNames)
					buffer.WriteString(", ")
					buffer.WriteString(allGroups[k].Members[i].MemberName)
					tempUserNames = buffer.String()
					buffer.Reset()
				}
			}
			tempValueSlice = append(tempValueSlice, tempUserNames)
			groupViewModel.Values = append(groupViewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		groupViewModel.Keys = keySlice
		c.Data["vm"] = groupViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/group-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

// To delete each group from database
func (c *GroupController) DeleteGroup() {
	w := c.Ctx.ResponseWriter
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
	groupId := c.Ctx.Input.Param(":groupId")
	group := models.Group{}

	if r.Method == "POST" {
		members := models.GroupMembers{}
		group.GroupName = c.GetString("groupName")
		tempGroupId := c.GetStrings("selectedUserIds")
		tempGroupMembers := c.GetStrings("selectedUserNames")
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberId = tempGroupId[i]
			members.MemberName = tempGroupMembers[i]
			group.Members = append(group.Members, members)
		}
		log.Println("group data",group)
		dbStatus := group.UpdateGroupDetails(c.AppEngineCtx, groupId)
		switch dbStatus {
		case true:
			//http.Redirect(w,r,"/group",301)
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))

		}
	} else {
		groupUser := models.Group{}
		viewModel := viewmodels.EditGroupViewModel{}
		GroupMembers,dbStatus :=groupUser.GetUsersForDropdown(c.AppEngineCtx)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(GroupMembers)	// To store data values of slice
			var keySlice []string	// To store keys of the slice
			for _, groupKey := range dataValue.MapKeys() {
				keySlice = append(keySlice, groupKey.String())

			}
			group := models.Group{}

			// Getting all Data for page load...
			GroupMemberName,dbStatus:= group.TakeGroupMemberName(c.AppEngineCtx, keySlice)
			switch dbStatus {
			case true:
				viewModel.GroupMembers = GroupMemberName
				viewModel.GroupKey = keySlice
				viewModel.PageType = helpers.SelectPageForEdit
			case false:
				log.Println(helpers.ServerConnectionError)
			}

			//Selecting Data which is to be edited...
			groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupId)
			switch dbStatus {
			case true:
				log.Println(groupDetails)
				viewModel.GroupNameToEdit = groupDetails.GroupName
				for i :=0; i<len(groupDetails.Members); i++{
					viewModel.GroupMembersToEdit = append(viewModel.GroupMembersToEdit, groupDetails.Members[i].MemberId)
				}
				//groupViewModel.GroupMembersToEdit = groupDetails.GroupMembers[]
				viewModel.PageType = helpers.SelectPageForEdit
				viewModel.GroupId = groupId
				c.Data["vm"] = viewModel
				c.Layout = "layout/layout.html"
				c.TplName = "template/add-group.html"
			case false:
				log.Println(helpers.ServerConnectionError)

			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}


