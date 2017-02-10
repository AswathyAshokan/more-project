/*Created By Farsana*/
package controllers
import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	//"google.golang.org/appengine/log"
	//"google.golang.org/appengine"
	"reflect"
	"net/http"
	"log"
	"app/passporte/helpers"
)

type GroupController struct {
	BaseController
}
// Add group to database


func (c *GroupController) AddGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {

		group := models.Group{}
		group.GroupName = c.GetString("groupName")
		group.GroupMembers = c.GetStrings("addUser")
		log.Println(group)
		dbStatus := group.AddGroupToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))


		}
	} else {
		/* for users cntain info
		groupUser := models.UserInformation{}
		GroupMembers :=groupUser.GetUsersForDropdown(c.AppEngineCtx)  // retrive all the keys of a users
		log.Print("ffffff", GroupMembers)
		groupDataValue := reflect.ValueOf(GroupMembers)	// To store data values of slice
		var groupKeySlice []string	// To store keys of the slice
		for _, groupKey := range groupDataValue.MapKeys() {
			groupKeySlice = append(groupKeySlice, groupKey.String())

		}
		log.Print("data",groupKeySlice)
		infoUser := models.UserInformation{}
		// for retrieve the names of the users
		GroupMemberName := infoUser.TakeGroupMemberName(c.AppEngineCtx,groupKeySlice)
		log.Print("cccccc", GroupMemberName)*/



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
			groupViewModel := viewmodels.Group{}
			groupViewModel.GroupMembers = GroupMemberName
			groupViewModel.GroupKey = groupKeySlice
			groupViewModel.PageType = helpers.SelectPageForAdd
			c.Data["GroupArray"] = groupViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-group.html"
		case false:
		}
	}

}
// show the details of whole group from database

func (c *GroupController) GroupDetails() {
	//r := c.Ctx.Request
	group := models.Group{}
	groupInfo := group.DisplayGroup(c.AppEngineCtx)
	GroupDataValue := reflect.ValueOf(groupInfo)
	var GroupValueSlice []models.Group
	GroupViewModel := viewmodels.Group{}
	var GroupKeySlice []string
	for _, GroupKey := range GroupDataValue.MapKeys() {
		GroupKeySlice = append(GroupKeySlice, GroupKey.String())//to get keys
		GroupValueSlice = append(GroupValueSlice, groupInfo[GroupKey.String()])//to get values
		GroupViewModel.Groups = append(GroupViewModel.Groups, groupInfo[GroupKey.String()])

	}
	GroupViewModel.GroupKey = GroupKeySlice
	c.Data["GroupArray"] = GroupViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/group-details.html"
}
// To delete each group from database

func (c *GroupController) DeleteGroup() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	GroupKey :=c.Ctx.Input.Param(":groupkey")
	group := models.Group{}
	dbStatus :=group.DeleteGroup(c.AppEngineCtx, GroupKey)
	switch dbStatus {
	case true:
		http.Redirect(w, r, "/group", 301)
	case false:
		log.Println("false")

	}

}

//Editing

func (c *GroupController) EditGroup() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	groupKey := c.Ctx.Input.Param(":groupkey")
	group := models.Group{}

	if r.Method == "POST" {

		group.GroupName = c.GetString("groupName")
		group.GroupMembers = c.GetStrings("addUser")
		dbStatus := group.UpdateGroupDetails(c.AppEngineCtx, groupKey)

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
		groupMemberName,dbStatus := group.TakeGroupMemberName(c.AppEngineCtx,groupKeySlice)
		switch dbStatus {
		case true:
			groupViewModel := viewmodels.EditGroupViewModel{}
			groupViewModel.GroupMembers = groupMemberName
			groupViewModel.GroupKey = groupKeySlice

		case false:

		}

		editResult, dbStatus := group.EditGroupDetails(c.AppEngineCtx, groupKey)
		switch dbStatus {
		case true:
			groupViewModel := viewmodels.EditGroupViewModel{}
			groupViewModel.GroupName = editResult.GroupName
			groupViewModel.GroupMembers = editResult.GroupMembers
			groupViewModel.PageType = helpers.SelectPageForEdit
			groupViewModel.GroupId = groupKey
			c.Data["vm"] = groupViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-group.html"
		case false:
			log.Println("failed")

		}

	}




}


