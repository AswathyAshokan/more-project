/*Created By Farsana*/
package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"

	"app/passporte/helpers"
	"log"
	"reflect"
)

type InviteUserController struct {
	BaseController
}

//Add new invite users to database
func (c *InviteUserController) AddInvitation() {
	user := models.InviteUser{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		user.FirstName = c.GetString("firstname")
		user.LastName = c.GetString("lastname")
		user.EmailId = c.GetString("emailid")
		user.UserType = c.GetString("usertype")
		user.DateOfCreation =(time.Now().UnixNano() / 1000000)
		user.Status = "inactive"
		dbStatus := user.AddInviteToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-invite-user.html"
	}
}

// fetch all the details of invite user from database
func (c *InviteUserController) InvitationDetails() {

	info,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(info)
		inviteUserViewModel := viewmodels.InviteUserViewModel{}
		var keySlice []string     //to store the keys of slice
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, info[k].FirstName)
			tempValueSlice = append(tempValueSlice, info[k].LastName)
			tempValueSlice = append(tempValueSlice, info[k].EmailId)
			tempValueSlice = append(tempValueSlice, info[k].UserType)
			tempValueSlice = append(tempValueSlice, info[k].Status)
			inviteUserViewModel.Values=append(inviteUserViewModel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		inviteUserViewModel.Keys = keySlice
		c.Data["vm"] = inviteUserViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/invite-user-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

//delete invite user details using invite user id
func (c *InviteUserController) DeleteInvitation() {
	w := c.Ctx.ResponseWriter
	InviteUserId :=c.Ctx.Input.Param(":inviteuserid")
	user := models.InviteUser{}
	result :=user.DeleteInviteUserById(c.AppEngineCtx, InviteUserId)
	switch result {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}

//edit profile of each invite user using invite user id
func (c *InviteUserController) EditInvitation() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	InviteUserId := c.Ctx.Input.Param(":inviteuserid")
	user := models.InviteUser{}
	if r.Method == "POST" {
		user.FirstName = c.GetString("firstname")
		user.LastName = c.GetString("lastname")
		user.EmailId = c.GetString("emailid")
		user.UserType = c.GetString("usertype")
		dbStatus :=user.UpdateInviteUserById(c.AppEngineCtx, InviteUserId)

		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		editResult, DbStatus := user.GetAllInviteUserForEdit(c.AppEngineCtx, InviteUserId)
		switch DbStatus {
		case true:
			invitationViewModel := viewmodels.InviteUserViewModel{}
			invitationViewModel.FirstName = editResult.FirstName
			invitationViewModel.LastName = editResult.LastName
			invitationViewModel.EmailId = editResult.EmailId
			invitationViewModel.UserType = editResult.UserType
			invitationViewModel.Status = editResult.Status
			invitationViewModel.PageType = helpers.SelectPageForEdit
			invitationViewModel.InviteId = InviteUserId
			c.Data["vm"] = invitationViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-invite-user.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}


