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
	inviteUser := models.Invitation{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	addViewModel := viewmodels.AddInviteUserViewModel{}
	if r.Method == "POST" {
		inviteUser.Info.FirstName = c.GetString("firstname")
		inviteUser.Info.LastName = c.GetString("lastname")
		inviteUser.Info.EmailId = c.GetString("emailid")
		inviteUser.Info.UserType = c.GetString("usertype")
		inviteUser.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		inviteUser.Settings.Status = "inactive"
		inviteUser.Info.CompanyTeamName = storedSession.CompanyTeamName
		dbStatus := inviteUser.AddInviteToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		addViewModel.CompanyTeamName = inviteUser.Info.CompanyTeamName
		c.Data["vm"] = addViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-invite-user.html"
	}
}

// fetch all the details of invite user from database
func (c *InviteUserController) InvitationDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
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
			tempValueSlice = append(tempValueSlice, info[k].Info.FirstName)
			tempValueSlice = append(tempValueSlice, info[k].Info.LastName)
			tempValueSlice = append(tempValueSlice, info[k].Info.EmailId)
			tempValueSlice = append(tempValueSlice, info[k].Info.UserType)
			tempValueSlice = append(tempValueSlice, info[k].Settings.Status)
			inviteUserViewModel.Values=append(inviteUserViewModel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		inviteUserViewModel.Keys = keySlice
		inviteUserViewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["vm"] = inviteUserViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/invite-user-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

//delete invite user details using invite user id
func (c *InviteUserController) DeleteInvitation() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	_ = ReadSession(w, r)
	InviteUserId :=c.Ctx.Input.Param(":inviteuserid")
	InviteUser := models.Invitation{}
	result := InviteUser.DeleteInviteUserById(c.AppEngineCtx, InviteUserId)
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
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	InviteUserId := c.Ctx.Input.Param(":inviteuserid")
	inviteUser := models.Invitation{}
	if r.Method == "POST" {
		inviteUser.Info.FirstName = c.GetString("firstname")
		inviteUser.Info.LastName = c.GetString("lastname")
		inviteUser.Info.EmailId = c.GetString("emailid")
		inviteUser.Info.UserType = c.GetString("usertype")
		dbStatus := inviteUser.UpdateInviteUserById(c.AppEngineCtx, InviteUserId)

		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		editResult, DbStatus := inviteUser.GetAllInviteUserForEdit(c.AppEngineCtx, InviteUserId)
		switch DbStatus {
		case true:
			invitationViewModel := viewmodels.EditInviteUserViewModel{}
			invitationViewModel.FirstName = editResult.Info.FirstName
			invitationViewModel.LastName = editResult.Info.LastName
			invitationViewModel.EmailId = editResult.Info.EmailId
			invitationViewModel.UserType = editResult.Info.UserType
			invitationViewModel.Status = editResult.Settings.Status
			invitationViewModel.PageType = helpers.SelectPageForEdit
			invitationViewModel.InviteId = InviteUserId
			invitationViewModel.CompanyTeamName= storedSession.CompanyTeamName
			c.Data["vm"] = invitationViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-invite-user.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}


