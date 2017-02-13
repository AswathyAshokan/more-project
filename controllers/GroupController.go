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

// Add group to database
func (c *GroupController) AddGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		log.Println("I am Here")
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
		GroupMembers :=groupUser.GetUsersForDropdown(c.AppEngineCtx)  // retrive all the keys of a users
		groupDataValue := reflect.ValueOf(GroupMembers)	// To store data values of slice
		var groupKeySlice []string	// To store keys of the slice
		for _, groupKey := range groupDataValue.MapKeys() {
			groupKeySlice = append(groupKeySlice, groupKey.String())

		}
		group := models.Group{}
		GroupMemberName,dbStatus:= group.TakeGroupMemberName(c.AppEngineCtx,groupKeySlice)
		switch dbStatus {
		case true:
			groupViewModel := viewmodels.AddGroupViewModel{}
			groupViewModel.GroupMembers = GroupMemberName
			groupViewModel.GroupKey = groupKeySlice
			groupViewModel.PageType = helpers.SelectPageForAdd
			c.Data["vm"] = groupViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-group.html"
		case false:
			log.Println("Server connection error")
		}
	}

}

// show the details of whole group from database
func (c *GroupController) GroupDetails() {
	//r := c.Ctx.Request
	allGroups := models.GetAllGroupDetails(c.AppEngineCtx)
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
		log.Println(len(tempUserNames))
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
}
// To delete each group from database

func (c *GroupController) DeleteGroup() {

	log.Println("deleteion")
	w := c.Ctx.ResponseWriter
	groupId :=c.Ctx.Input.Param(":groupId")
	group := models.Group{}
	dbStatus :=group.DeleteGroup(c.AppEngineCtx, groupId)
	log.Println("deletion completed")
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))

	}

}

//Editing

func (c *GroupController) EditGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	groupId := c.Ctx.Input.Param(":groupId")
	group := models.Group{}

	if r.Method == "POST" {
		log.Println("1")
		members := models.GroupMembers{}
		group.GroupName = c.GetString("groupName")
		log.Println("2", group.GroupName)
		tempGroupId := c.GetStrings("selectedUserIds")
		log.Println("3", tempGroupId)
		tempGroupMembers := c.GetStrings("selectedUserNames")
		log.Println("4", tempGroupMembers)
		for i := 0; i < len(tempGroupId); i++ {
			log.Println("5")
			members.MemberId = tempGroupId[i]
			log.Println("6")
			members.MemberName = tempGroupMembers[i]
			log.Println("7")
			group.Members = append(group.Members, members)
			log.Println("8")
		}
		log.Println("9")
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
		GroupMembers :=groupUser.GetUsersForDropdown(c.AppEngineCtx)  // retrive all the keys of a users
		groupDataValue := reflect.ValueOf(GroupMembers)	// To store data values of slice
		var groupKeySlice []string	// To store keys of the slice
		for _, groupKey := range groupDataValue.MapKeys() {
			groupKeySlice = append(groupKeySlice, groupKey.String())

		}
		group := models.Group{}

		// Getting all Data for page load...
		GroupMemberName,dbStatus:= group.TakeGroupMemberName(c.AppEngineCtx,groupKeySlice)
		switch dbStatus {
		case true:
			viewModel.GroupMembers = GroupMemberName
			viewModel.GroupKey = groupKeySlice
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
	}

}


