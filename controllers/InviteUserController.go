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
		log.Println("cp1",storedSession.CompanyTeamName)
		inviteUser.Info.CompanyAdmin = storedSession.AdminFirstName+" "+storedSession.AdminLastName
		inviteUser.Info.FirstName = c.GetString("firstname")
		inviteUser.Info.LastName = c.GetString("lastname")
		inviteUser.Info.Email = c.GetString("emailid")
		inviteUser.Info.UserType = c.GetString("usertype")
		inviteUser.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		inviteUser.Settings.Status = helpers.StatusPending
		inviteUser.Info.CompanyTeamName = storedSession.CompanyTeamName
		inviteUser.Info.CompanyId = storedSession.CompanyId
		inviteUser.Info.CompanyName = storedSession.CompanyName
		userFullName := storedSession.AdminFirstName+" "+storedSession.AdminLastName
		companyID := models.GetCompanyIdByCompanyTeamName(c.AppEngineCtx, companyTeamName)
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
	} else {
		/*companyPlan := storedSession.CompanyPlan*/
		/*if companyPlan == "family" {
			_, dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx, companyTeamName)
			switch dbStatus {
			case true:
				var count = 0
				var tempValueSlice []string
				dataValue := reflect.ValueOf(info)
				var uniqueEmailSlice []string
				for _, key := range dataValue.MapKeys() {
					log.Println("key",info[key.String()])
					//check is email id is present in the slice
					if helpers.StringInSlice(info[key.String()].Info.Email, uniqueEmailSlice) == false {
						tempValueSlice = append(tempValueSlice, info[key.String()].Settings.Status)
						uniqueEmailSlice = append(uniqueEmailSlice, info[key.String()].Info.Email)//appent email id into slice
					}

								for i := 0; i < len(tempValueSlice); i++ {
					if tempValueSlice[i] == helpers.StatusPending || tempValueSlice[i] == helpers.StatusAccepted {
						count = count + 1
					}
				}
				for i := count; i < 4; i++ {
					addViewModel.AllowInvitations = true
				}

			case false:
				log.Println("failed")
			}
		}else {
			addViewModel.AllowInvitations =true
		}*/
		addViewModel.AllowInvitations =true
		addViewModel.CompanyTeamName = storedSession.CompanyTeamName
		addViewModel.CompanyPlan = storedSession.CompanyPlan
		addViewModel.AdminFirstName = storedSession.AdminFirstName
		addViewModel.AdminLastName = storedSession.AdminLastName
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
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, info[k].FirstName)
			tempValueSlice = append(tempValueSlice, info[k].LastName)
			tempValueSlice = append(tempValueSlice, info[k].Email)
			tempValueSlice = append(tempValueSlice, info[k].UserType)
			tempValueSlice = append(tempValueSlice, info[k].Status)
			inviteUserViewModel.Values=append(inviteUserViewModel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		//if invitationByEmail[key].Settings.Status == helpers.StatusActive ||invitationByEmail[key].Settings.Status == helpers.StatusAccepted||invitationByEmail[key].Settings.Status==helpers.StatusPending{

			}

	case false:
		log.Println(helpers.ServerConnectionError)
	}
	inviteUserViewModel.CompanyTeamName = storedSession.CompanyTeamName
	inviteUserViewModel.CompanyPlan = storedSession.CompanyPlan
	inviteUserViewModel.AdminFirstName = storedSession.AdminFirstName
	inviteUserViewModel.AdminLastName = storedSession.AdminLastName
	inviteUserViewModel.Keys = keySlice
	log.Println("view",inviteUserViewModel)
	c.Data["array"] = inviteUserViewModel
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
	/*r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	InviteUserId := c.Ctx.Input.Param(":inviteuserid")
	inviteUser := models.EmailInvitation{}
	if r.Method == "POST" {
		inviteUser.Info.CompanyName = storedSession.CompanyName
		inviteUser.Info.CompanyPlan = storedSession.CompanyPlan
		inviteUser.Info.CompanyTeamName= storedSession.CompanyTeamName
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
			invitationViewModel.AdminFirstName = storedSession.AdminFirstName
			invitationViewModel.AdminLastName = storedSession.AdminLastName
			c.Data["vm"] = invitationViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-invite-user.html"
		case false:*/
			log.Println(helpers.ServerConnectionError)
		/*}*/
	/*}*/
}


