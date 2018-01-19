/*Created By Farsana*/
package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"log"
	"html/template"
	"bytes"

	"reflect"
	"strconv"
	"google.golang.org/appengine/urlfetch"
	"gopkg.in/sendgrid/sendgrid-go.v2"

	"google.golang.org/appengine"


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
		inviteUser.Settings.DateOfCreation = time.Now().Unix()
		log.Println("inviteUser.Settings.DateOfCreation",inviteUser.Settings.DateOfCreation)
		inviteUser.Settings.Status = helpers.StatusInActive
		inviteUser.Settings.UserResponse = helpers.UserResponsePending
		inviteUser.Info.CompanyTeamName = storedSession.CompanyTeamName
		inviteUser.Info.CompanyId = storedSession.CompanyId
		inviteUser.Info.CompanyName = storedSession.CompanyName
		userFullName := storedSession.AdminFirstName+" "+storedSession.AdminLastName
		companyID := storedSession.CompanyTeamName

		dbStatus := inviteUser.CheckEmailIdInDb(c.AppEngineCtx,companyID)
		switch dbStatus {
		case true:
			Status := inviteUser.AddInviteToDb(c.AppEngineCtx,companyID,userFullName)
			switch Status {
			case true:
				log.Println("true add")
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
				//body := buf.String()
				//from := "aswathy.a@cynere.com"
				//to := inviteUser.Info.Email
				//subject := "Subject: Passporte - Invitation\n"
				//mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
				//message := []byte(subject + mime + "\n" + body)
				//if err := smtp.SendMail("smtp.gmail.com:465", smtp.PlainAuth("", "aswathy.a@cynere.com", "aswathyashok", "smtp.gmail.com"), from, []string{to}, []byte(message)); err != nil {
				//	log.Println(err)
				//}



				key := "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"
				sg := sendgrid.NewSendGridClientWithApiKey(key)

				// must change the net/http client to not use default transport
				ctx := appengine.NewContext(r)
				client := urlfetch.Client(ctx)
				sg.Client = client // <-- now using urlfetch, "overriding" default

				message := sendgrid.NewMail()
				message.AddTo(inviteUser.Info.Email)
				message.SetFrom("passportetest@gmail.com")
				message.SetSubject("Passporte - Invitation")
				message.SetHTML( buf.String())

				if e := sg.Send(message); e == nil {
					log.Println("lllllll")
				} else {

					log.Println("error",e)
				}

				w.Write([]byte("true"))
			case false:
				w.Write([]byte("false in Add"))
			}
		case false:
			log.Println("condition failed and return false")
			w.Write([]byte("false"))
		}
	} else {
		//limitValues := 4



		companyPlan := storedSession.CompanyPlan
		log.Println("plan",companyPlan)
		if companyPlan == helpers.PlanBusiness {
			info, limitedUsers ,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx, companyTeamName)
			switch dbStatus {
			case true:
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

				}
				var count = 0
				for i := 0; i < len(tempValueSlice); i++ {
					if tempValueSlice[i] == helpers.UserResponsePending || tempValueSlice[i] == helpers.UserResponseAccepted{
						count = count + 1
						log.Println("inside loop",count)
					}
				}
				newNumberOfUsers,_ := strconv.Atoi(limitedUsers)
				if count <newNumberOfUsers{
					addViewModel.AllowInvitations = true
				}else{
					addViewModel.AllowInvitations = false
				}
			case false:
				log.Println("failed")
			}
		}else {

			addViewModel.AllowInvitations =true
		}
		log.Println("hai iam there",addViewModel.AllowInvitations)
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
	log.Println("time now",time.Now())
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	inviteUserViewModel := viewmodels.InviteUserViewModel{}
	info,limitedUser,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx,companyId)
	var keySlice []string
	switch dbStatus {
	case true:
		log.Println("limitedUser",limitedUser)
		dataValue := reflect.ValueOf(info)
		//to store the keys of slice
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}

		for _, k := range keySlice {
			if info[k].Status != helpers.UserStatusDeleted &&info[k].Status !=helpers.UserStatusDeleted{
				var tempValueSlice []string
				tempValueSlice = append(tempValueSlice, info[k].FirstName)
				tempValueSlice = append(tempValueSlice, info[k].LastName)
				tempValueSlice = append(tempValueSlice, info[k].Email)
				tempValueSlice = append(tempValueSlice, info[k].UserType)
				tempValueSlice = append(tempValueSlice,info[k].UserResponse)
				tempValueSlice = append(tempValueSlice,k)
				inviteUserViewModel.Values = append(inviteUserViewModel.Values, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}

		}

	case false:
		log.Println(helpers.ServerConnectionError)
	}

	dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
	var notificationCount=0
	switch dbStatus {
	case true:

		notificationOfUser := reflect.ValueOf(notificationValue)
		for _, notificationUserKey := range notificationOfUser.MapKeys() {
			dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
			switch dbStatus {
			case true:
				notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
				for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
					var NotificationArray []string
					if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
						notificationCount=notificationCount+1;
					}
					NotificationArray =append(NotificationArray,notificationUserKey.String())
					NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
					date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
					NotificationArray =append(NotificationArray,date)
					inviteUserViewModel.NotificationArray=append(inviteUserViewModel.NotificationArray,NotificationArray)

				}
			case false:
			}
		}
	case false:
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
	log.Println("cp1")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	InviteUserId :=c.Ctx.Input.Param(":inviteuserid")
	InviteUser := models.Invitation{}
	result := InviteUser.CheckJobIsAssigned(c.AppEngineCtx, InviteUserId,companyTeamName)
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

		dbStatus := invitation.UpdateInviteUserById(c.AppEngineCtx, InviteUserId,companyTeamName)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		editViewResult, DbStatus := models.GetAllUserFormCompanyEdit(c.AppEngineCtx,companyTeamName,InviteUserId)
		log.Println("all", editViewResult)
		switch DbStatus {
		case true:


			if editViewResult.UserResponse !=helpers.UserStatusDeleted{
				invitationViewModel.FirstName = editViewResult.FirstName
				invitationViewModel.LastName = editViewResult.LastName
				invitationViewModel.UserType = editViewResult.UserType
				invitationViewModel.UserResponse = editViewResult.UserResponse
				invitationViewModel.PageType = helpers.SelectPageForEdit
				invitationViewModel.CompanyTeamName = storedSession.CompanyTeamName
				invitationViewModel.CompanyPlan = storedSession.CompanyPlan
				invitationViewModel.AdminFirstName = storedSession.AdminFirstName
				invitationViewModel.AdminLastName = storedSession.AdminLastName
				invitationViewModel.EmailId = editViewResult.Email
				invitationViewModel.InviteId = InviteUserId
				c.Data["vm"] = invitationViewModel
				c.Layout = "layout/layout.html"
				c.TplName = "template/add-invite-user.html"
			}

		case false:
			log.Println(helpers.ServerConnectionError)
		}

	}
}

//delete user from task
func (c *InviteUserController) RemoveUserFromTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	InviteUserId := c.Ctx.Input.Param(":inviteuserid")
	dbStatus := models.RemoveUsersFromTaskForDelete(c.AppEngineCtx,companyTeamName, InviteUserId)
	log.Println("get escaped  1",dbStatus)
	switch dbStatus {
	case true:
		result := models.DeleteInviteUserById(c.AppEngineCtx, InviteUserId, companyTeamName)
		switch result {
		case true:
			w.Write([]byte("true"))

		case false:
			log.Println("true for my false life")
		}
	case false:
		log.Println("false")
	}
}

//delete user if not used in task
func (c *InviteUserController) DeleteUserIfNotInTask() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	InviteUserId := c.Ctx.Input.Param(":inviteuserid")
	companyInvitationStatus := models.CheckStatusInInvitationOfCompany(c.AppEngineCtx, InviteUserId, companyTeamName)

	switch companyInvitationStatus {
	case true:
		result := models.DeleteInviteUserById(c.AppEngineCtx, InviteUserId, companyTeamName)
		log.Println("getttt2",result)
		switch result {
		case true:
			w.Write([]byte("true"))
		case false:
			log.Println("true for my false life")
			w.Write([]byte("false"))
		}
	case false:
		status :=models.DeleteInviteUserIfStatusIsPending(c.AppEngineCtx, InviteUserId, companyTeamName)
		switch status {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}


}

//add invite user details after upgradation of plan
func (c *InviteUserController) AddInvitationByUpgradationOfPlan() {

	log.Println("haiiiosnjskhsdhf")
	numberOfUsers := c.Ctx.Input.Param(":numberOfUsers")
	log.Println("numberOfUsers",numberOfUsers)
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
		inviteUser.Settings.DateOfCreation = time.Now().Unix()
		log.Println("inviteUser.Settings.DateOfCreation",inviteUser.Settings.DateOfCreation)
		inviteUser.Settings.Status = helpers.StatusInActive
		inviteUser.Settings.UserResponse = helpers.UserResponsePending
		inviteUser.Info.CompanyTeamName = storedSession.CompanyTeamName
		inviteUser.Info.CompanyId = storedSession.CompanyId
		inviteUser.Info.CompanyName = storedSession.CompanyName
		userFullName := storedSession.AdminFirstName+" "+storedSession.AdminLastName
		companyID := storedSession.CompanyTeamName

		dbStatus := inviteUser.CheckEmailIdInDb(c.AppEngineCtx,companyID)
		switch dbStatus {
		case true:
			Status := inviteUser.AddInviteToDb(c.AppEngineCtx,companyID,userFullName)
			switch Status {
			case true:
				log.Println("true add")
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
				//body := buf.String()
				//from := "aswathy.a@cynere.com"
				//to := inviteUser.Info.Email
				//subject := "Subject: Passporte - Invitation\n"
				//mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
				//message := []byte(subject + mime + "\n" + body)
				//if err := smtp.SendMail("smtp.gmail.com:465", smtp.PlainAuth("", "aswathy.a@cynere.com", "aswathyashok", "smtp.gmail.com"), from, []string{to}, []byte(message)); err != nil {
				//	log.Println(err)
				//}
				key := "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"
				sg := sendgrid.NewSendGridClientWithApiKey(key)

				// must change the net/http client to not use default transport
				ctx := appengine.NewContext(r)
				client := urlfetch.Client(ctx)
				sg.Client = client // <-- now using urlfetch, "overriding" default

				message := sendgrid.NewMail()
				message.AddTo(inviteUser.Info.Email)
				message.SetFrom("passportetest@gmail.com")
				message.SetSubject("Passporte - Invitation")
				message.SetHTML( buf.String())

				if e := sg.Send(message); e == nil {
					log.Println("lllllll")
				} else {

					log.Println("error",e)
				}


				w.Write([]byte("true"))
			case false:
				w.Write([]byte("false in Add"))
			}
		case false:
			log.Println("condition failed and return false")
			w.Write([]byte("false"))
		}
	} else {
		newLimitValues,_:= strconv.Atoi(numberOfUsers)
		companyPlan := storedSession.CompanyPlan
		limitedValueOfUsers :=  strconv.Itoa(newLimitValues)
		storedSessionForPayment := ReadSessionForPayment(w, r)
		log.Println("number of userssssssssssssss",storedSessionForPayment.NumberOfUsers)
		if companyPlan == helpers.PlanBusiness && limitedValueOfUsers==storedSessionForPayment.NumberOfUsers {
			updatedNoOfUsers := models.UpdateNoOfLimitedUser(c.AppEngineCtx, companyTeamName,newLimitValues)
			ClearSessionForPayment(w)
			info, _ ,dbStatus := models.GetAllInviteUsersDetails(c.AppEngineCtx, companyTeamName)
			switch dbStatus {
			case true:
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

				}
				var count = 0
				for i := 0; i < len(tempValueSlice); i++ {
					if tempValueSlice[i] == helpers.UserResponsePending || tempValueSlice[i] == helpers.UserResponseAccepted{
						count = count + 1
						log.Println("inside loop",count)
					}
				}
				if count <updatedNoOfUsers{
					log.Println("true")
					addViewModel.AllowInvitations = true
				}else{
					log.Println("fa;se")
					addViewModel.AllowInvitations = false
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








