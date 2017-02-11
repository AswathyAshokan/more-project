/*Created By Farsana*/
package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	//"google.golang.org/appengine/log"
	//"google.golang.org/appengine"
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
		log.Println(len(tempGroupMembers))
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberId = tempGroupId[i]
			members.MemberName = tempGroupMembers[i]
			group.Members = append(group.Members, members)
		}
		log.Println(tempGroupId, tempGroupMembers)
		log.Println(group)
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
	group := models.Group{}
	info := group.DisplayGroup(c.AppEngineCtx)
	dataValue := reflect.ValueOf(info)
	groupViewModel := viewmodels.GroupList{}
	var keySlice []string
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}
	for _, k := range keySlice {
		var tempValueSlice []string
		membersNumber := len(info[k].Members)
		tempValueSlice = append(tempValueSlice, info[k].GroupName)
		tempValueSlice = append(tempValueSlice, strconv.Itoa(membersNumber))
		tempUserNames := ""
		log.Println(len(tempUserNames))
		var buffer bytes.Buffer
		for i := 0; i < membersNumber; i++ {
			if len(tempUserNames) == 0{
				buffer.WriteString(info[k].Members[i].MemberName)
				tempUserNames = buffer.String()
				buffer.Reset()
			} else {
				buffer.WriteString(tempUserNames)
				buffer.WriteString(", ")
				buffer.WriteString(info[k].Members[i].MemberName)
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

		group.GroupName = c.GetString("groupName")
		//group.Members = c.GetStrings("selectedUserIds")
		log.Println("group data",group)
		dbStatus := group.UpdateGroupDetails(c.AppEngineCtx, groupId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))

		}


	} else {
		groupUser := models.Group{}
		groupViewModel := viewmodels.EditGroupViewModel{}
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
			groupViewModel.GroupMembers = GroupMemberName
			groupViewModel.GroupKey = groupKeySlice
			groupViewModel.PageType = helpers.SelectPageForEdit
		case false:
			log.Println(helpers.ServerConnectionError)
		}

		//Selecting Data which is to be edited...
		groupDetails, dbStatus := group.GetGroupDetailsById(c.AppEngineCtx, groupId)
		switch dbStatus {
		case true:
			log.Println(groupDetails)
			groupViewModel.GroupNameToEdit = groupDetails.GroupName
			//groupViewModel.GroupMembersToEdit = groupDetails.GroupMembers
			groupViewModel.PageType = helpers.SelectPageForEdit
			groupViewModel.GroupId = groupId
			c.Data["vm"] = groupViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-group.html"
		case false:
			log.Println(helpers.ServerConnectionError)

		}
	}




}


