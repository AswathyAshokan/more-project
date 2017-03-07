/*Created By Farsana*/
package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"log"
	"reflect"
	"net/smtp"
	"html/template"
	"bytes"
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
	log.Println("session VAlues :",storedSession)
	inviteUser := models.Invitation{}
	addViewModel := viewmodels.AddInviteUserViewModel{}
	if r.Method == "POST" {
		inviteUser.Info.FirstName = c.GetString("firstname")
		inviteUser.Info.LastName = c.GetString("lastname")
		inviteUser.Info.Email = c.GetString("emailid")
		inviteUser.Info.UserType = c.GetString("usertype")
		inviteUser.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		inviteUser.Settings.Status = helpers.StatusPending
		inviteUser.Info.CompanyTeamName = storedSession.CompanyTeamName
		inviteUser.Info.CompanyName = storedSession.CompanyName
		userFullname := storedSession.AdminFirstName+" "+storedSession.AdminLastName
		companyID := models.GetCompanyIdByCompanyTeamName(c.AppEngineCtx, companyTeamName)
		dbStatus := inviteUser.AddInviteToDb(c.AppEngineCtx,companyID)
		switch dbStatus {
		case true:

			templateData := TemplateData{}
			templateData.AdminEmail = storedSession.AdminEmail
			templateData.AdminName = userFullname
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
			from := "sarathcynere@gmail.com"
			to := inviteUser.Info.Email
			subject := "Subject: Passporte - Invitation\n"
			mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
			message := []byte(subject + mime + "\n" + body)
			if err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", "sarathcynere@gmail.com", "cynereinfotech", "smtp.gmail.com"), from, []string{to}, []byte(message)); err != nil {
				log.Println(err)
			}
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		info,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx,companyTeamName)
		switch dbStatus {
		case true:
			var count = 0
			companyPlan := storedSession.CompanyPlan
			var tempValueSlice []string
			if companyPlan == "family" {

				dataValue := reflect.ValueOf(info)

				var uniqueEmailSlice []string
				for _, key := range dataValue.MapKeys() {
					if helpers.StringInSlice(info[key.String()].Info.Email, uniqueEmailSlice) == false {
						tempValueSlice = append(tempValueSlice, info[key.String()].Settings.Status)
						uniqueEmailSlice = append(uniqueEmailSlice, info[key.String()].Info.Email)
					}

				}
				for i := 0; i < len(tempValueSlice); i++ {
					if tempValueSlice[i] == helpers.StatusPending || tempValueSlice[i] == helpers.StatusAccepted {
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

		addViewModel.CompanyTeamName = storedSession.CompanyTeamName
		addViewModel.CompanyPlan = storedSession.CompanyPlan
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
	info,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx,companyTeamName)
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
			tempValueSlice = append(tempValueSlice, info[k].Info.Email)
			tempValueSlice = append(tempValueSlice, info[k].Info.UserType)
			tempValueSlice = append(tempValueSlice, info[k].Settings.Status)
			inviteUserViewModel.Values=append(inviteUserViewModel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		inviteUserViewModel.Keys = keySlice
		inviteUserViewModel.CompanyTeamName = storedSession.CompanyTeamName
		inviteUserViewModel.CompanyPlan = storedSession.CompanyPlan
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
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
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
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	InviteUserId := c.Ctx.Input.Param(":inviteuserid")
	inviteUser := models.Invitation{}
	if r.Method == "POST" {
		inviteUser.Info.FirstName = c.GetString("firstname")
		inviteUser.Info.LastName = c.GetString("lastname")
		inviteUser.Info.Email = c.GetString("emailid")
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
			invitationViewModel.EmailId = editResult.Info.Email
			invitationViewModel.UserType = editResult.Info.UserType
			invitationViewModel.Status = editResult.Settings.Status
			invitationViewModel.PageType = helpers.SelectPageForEdit
			invitationViewModel.InviteId = InviteUserId
			invitationViewModel.CompanyTeamName= storedSession.CompanyTeamName
			invitationViewModel.CompanyPlan = storedSession.CompanyPlan
			c.Data["vm"] = invitationViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-invite-user.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}


