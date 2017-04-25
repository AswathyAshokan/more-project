/*Created By Farsana*/
package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"log"
	"net/smtp"
	"html/template"
	"bytes"

	"reflect"
)

type InviteUserController struct {
	BaseController
}
type TemplateData struct{
	AdminName	string
	AdminEmail 	string
	InvitedUser	string
	CompanyName	string
}
//Add new invite users to database
func (c *InviteUserController) AddInvitation() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	inviteUser := models.EmailInvitation{}
	addViewModel := viewmodels.AddInviteUserViewModel{}
	if r.Method == "POST" {
		inviteUser.Info.CompanyAdmin = storedSession.AdminFirstName+" "+storedSession.AdminLastName
		inviteUser.Info.FirstName = c.GetString("firstname")
		inviteUser.Info.LastName = c.GetString("lastname")
		inviteUser.Info.Email = c.GetString("emailid")
		inviteUser.Info.UserType = c.GetString("usertype")
		inviteUser.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		inviteUser.Settings.Status = helpers.StatusInActive
		inviteUser.Settings.UserResponse = helpers.UserResponsePending
		inviteUser.Info.CompanyTeamName = storedSession.CompanyTeamName
		inviteUser.Info.CompanyId = storedSession.CompanyId
		inviteUser.Info.CompanyName = storedSession.CompanyName
		userFullName := storedSession.AdminFirstName+" "+storedSession.AdminLastName
		companyID := models.GetCompanyIdByCompanyTeamName(c.AppEngineCtx, companyTeamName)
		dbStatus := inviteUser.CheckEmailIdInDb(c.AppEngineCtx,companyID)
		log.Println("status",dbStatus)
		switch dbStatus {
		case true:
			dbStatus := inviteUser.AddInviteToDb(c.AppEngineCtx,companyID)
			switch dbStatus {
			case true:
				templateData := TemplateData{}
				templateData.AdminEmail = storedSession.AdminEmail
				templateData.AdminName = userFullName
				templateData.CompanyName = inviteUser.Info.CompanyName
				templateData.InvitedUser =  inviteUser.Info.FirstName
				t,err := template.ParseFiles("views/email/invite-email.html")
				if err != nil {
					log.Println(err)
				}
				buf := new(bytes.Buffer)
				if err = t.Execute(buf, templateData); err != nil {
					log.Println(err)
				}
				body := buf.String()
				from := "passportetest@gmail.com"
				to := inviteUser.Info.Email
				subject := "Subject: Passporte - Invitation\n"
				mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
				message := []byte(subject + mime + "\n" + body)
				if err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", "passportetest@gmail.com", "passporte123", "smtp.gmail.com"), from, []string{to}, []byte(message)); err != nil {
				log.Println(err)
				}
				w.Write([]byte("true"))
			case false:
				w.Write([]byte("false"))
			}
		case false:
			log.Println("condition failed and return false")
			w.Write([]byte("false"))
		}
	} else {
		companyPlan := storedSession.CompanyPlan
		if companyPlan == "family" {
			info, dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx, companyTeamName)
			switch dbStatus {
			case true:
				var count = 0
				var tempValueSlice []string
				var keySlice []string
				dataValue := reflect.ValueOf(info)
				var uniqueEmailSlice []string
				for _, key := range dataValue.MapKeys() {
					keySlice = append(keySlice, key.String())
				}
				for _, key := range keySlice{
					//check is email id is present in the slice
					if helpers.StringInSlice(info[key].Email, uniqueEmailSlice) == false {
						tempValueSlice = append(tempValueSlice, info[key].UserResponse)
						uniqueEmailSlice = append(uniqueEmailSlice, info[key].Email)//appent email id into slice
					}
					for i := 0; i < len(tempValueSlice); i++ {
						if tempValueSlice[i] == helpers.UserResponsePending || tempValueSlice[i] == helpers.UserResponseAccepted{
							count = count + 1
						}
					}
					for i := count; i < 4; i++ {
						addViewModel.AllowInvitations = true
					}
				}
			case false:
				log.Println("failed")
			}
		}else {
			addViewModel.AllowInvitations =true
		}
		addViewModel.CompanyTeamName = storedSession.CompanyTeamName
		addViewModel.CompanyPlan = storedSession.CompanyPlan
		addViewModel.AdminFirstName = storedSession.AdminFirstName
		addViewModel.AdminLastName = storedSession.AdminLastName
		addViewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = addViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-invite-user.html"
	}
}

// fetch all the details of invite user from database
func (c *InviteUserController) InvitationDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	inviteUserViewModel := viewmodels.InviteUserViewModel{}
	info,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx,companyId)
	var keySlice []string
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(info)
		//to store the keys of slice
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		for _, k := range keySlice {
			if info[k].UserResponse == helpers.UserResponseAccepted ||info[k].UserResponse == helpers.UserResponsePending{
				var tempValueSlice []string
				tempValueSlice = append(tempValueSlice, info[k].FirstName)
				tempValueSlice = append(tempValueSlice, info[k].LastName)
				tempValueSlice = append(tempValueSlice, info[k].Email)
				tempValueSlice = append(tempValueSlice, info[k].UserType)
				tempValueSlice = append(tempValueSlice,info[k].UserResponse)
				tempValueSlice = append(tempValueSlice, info[k].Status)
				inviteUserViewModel.Values = append(inviteUserViewModel.Values, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}
		}

	case false:
		log.Println(helpers.ServerConnectionError)
	}
	inviteUserViewModel.CompanyTeamName = storedSession.CompanyTeamName
	inviteUserViewModel.CompanyPlan = storedSession.CompanyPlan
	inviteUserViewModel.AdminFirstName = storedSession.AdminFirstName
	inviteUserViewModel.AdminLastName = storedSession.AdminLastName
	inviteUserViewModel.ProfilePicture =storedSession.ProfilePicture
	inviteUserViewModel.Keys = keySlice
	c.Data["vm"] = inviteUserViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/invite-user-details.html"
}

//delete invite user details using invite user id
func (c *InviteUserController) DeleteInvitation() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	InviteUserId :=c.Ctx.Input.Param(":inviteuserid")
	InviteUser := models.Invitation{}
	result := InviteUser.DeleteInviteUserById(c.AppEngineCtx, InviteUserId,companyTeamName)
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
	invitationViewModel := viewmodels.EditInviteUserViewModel{}
	InviteUserId := c.Ctx.Input.Param(":inviteuserid")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)

	if r.Method == "POST" {
		invitation := models.CompanyInvitations{}
		invitation.FirstName = c.GetString("firstname")
		invitation.LastName = c.GetString("lastname")
		invitation.Email = c.GetString("emailid")
		invitation.UserType = c.GetString("usertype")
		dbStatus := invitation.UpdateInviteUserById(c.AppEngineCtx, InviteUserId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		editResult, DbStatus := models.GetAllInviteUserForEdit(c.AppEngineCtx)
		log.Println("all", editResult)
		switch DbStatus {
		case true:

			var keySlice []string
			var keySliceById []string
			dataValue := reflect.ValueOf(editResult)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
				log.Println("all key", keySlice)
			}
			for _, keyIn := range keySlice {
				invitationData := models.GetInvitationById(c.AppEngineCtx, InviteUserId, keyIn)
				log.Println("log", invitationData)
				dataValueOfInvite := reflect.ValueOf(invitationData)
				for _, k := range dataValueOfInvite.MapKeys() {
					keySliceById = append(keySliceById, k.String())
				}
				for _, k := range keySliceById {
					if InviteUserId == k {
						invitationViewModel.FirstName = invitationData[k].Info.FirstName
						invitationViewModel.LastName = invitationData[k].Info.LastName
						invitationViewModel.UserType = invitationData[k].Info.UserType
						invitationViewModel.UserResponse = invitationData[k].Settings.UserResponse
						invitationViewModel.Status = invitationData[k].Settings.Status
						invitationViewModel.PageType = helpers.SelectPageForEdit
						invitationViewModel.CompanyTeamName = storedSession.CompanyTeamName
						invitationViewModel.CompanyPlan = storedSession.CompanyPlan
						invitationViewModel.AdminFirstName = storedSession.AdminFirstName
						invitationViewModel.AdminLastName = storedSession.AdminLastName
						invitationViewModel.EmailId = invitationData[k].Info.Email
						invitationViewModel.InviteId = InviteUserId
						log.Println("view", invitationViewModel)
						c.Data["vm"] = invitationViewModel
						c.Layout = "layout/layout.html"
						c.TplName = "template/add-invite-user.html"

					}
				}
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}

	}
}
//func (c *InviteUserController)CheckEmailInvitation(){
//	log.Println("inside")
//	w := c.Ctx.ResponseWriter
//	r := c.Ctx.Request
//	emailId :=  c.GetString("emailId")
//	log.Println("email in contro",emailId)
//	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
//	storedSession := ReadSession(w, r, companyTeamName)
//	companyId :=storedSession.CompanyId
//
//	isEmailUsed := models.CheckEmailIsUsedInvitation(c.AppEngineCtx, emailId,companyId)
//	if isEmailUsed == false {
//		w.Write([]byte("false"))
//	}else{
//		w.Write([]byte("true"))
//	}
//}


