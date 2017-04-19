/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/helpers"
	"app/passporte/viewmodels"
	"log"

)

type RegisterController struct {
	BaseController
}

//Register new Company Admin
func (c *RegisterController) Register() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		currentTime := time.Now().Unix()

		company := models.Company{}
		company.Info.CompanyName = c.GetString("companyName")
		company.Info.CompanyTeamName = c.GetString("teamName")
		company.Info.Address = c.GetString("address")
		company.Info.State = c.GetString("state")
		company.Info.ZipCode = c.GetString("zipCode")
		company.Settings.Status = helpers.StatusActive
		company.Settings.DateOfCreation = currentTime
		company.Plan = helpers.PlanBusiness
		admin := models.Admins{}
		admin.Info.FirstName = c.GetString("firstName")
		admin.Info.LastName = c.GetString("lastName")
		admin.Info.Email = c.GetString("emailId")
		admin.Info.PhoneNo = c.GetString("phoneNo")
		admin.Info.Password = []byte(c.GetString("password"))
		admin.Settings.DateOfCreation = currentTime
		admin.Settings.Status = helpers.StatusActive
		dbStatus := admin.CreateAdminAndCompany(c.AppEngineCtx, company)
		switch dbStatus{
		case false:
			w.Write([]byte("false"))
		case true:
			w.Write([]byte("true"))
		}
	} else {
		c.TplName = "template/register.html"
	}
}

func (c *RegisterController)CheckEmail(){
	w := c.Ctx.ResponseWriter
	emailId := c.GetString("emailId")
	isEmailUsed := models.CheckEmailIsUsed(c.AppEngineCtx, emailId)
	if isEmailUsed == false {
		w.Write([]byte("false"))
	}else{
		w.Write([]byte("true"))
	}
}

type Storage struct {
	Token	string
	RefreshToken string
	Bucket string
	APIKey string
}

func (c *RegisterController) EditProfile() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	ReadSession(w, r, companyTeamName)
	adminId :=storedSession.AdminId
	plan :=storedSession.CompanyPlan
	admin := models.Admins{}
	if r.Method == "POST" {
		admin.Info.FirstName = c.GetString("name")
		admin.Info.Email = c.GetString("emailId")
		admin.Info.PhoneNo = c.GetString("phoneNumber")
		admin.Settings.ProfilePicture = c.GetString("profilePicture")
		admin.Settings.ThumbProfilePicture=c.GetString("thumbPicture")
		dbStatus := admin.EditAdminDetails(c.AppEngineCtx, adminId)

		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}else {

		viewModel := viewmodels.EditProfileViewModel{}

		dbStatus,adminDetail:= admin.GetCompanyDetails(c.AppEngineCtx, adminId)
		switch dbStatus {
		case true:
			viewModel.Email = adminDetail.Info.Email
			viewModel.FirstName =adminDetail.Info.FirstName
			viewModel.LastName = adminDetail.Info.LastName
			viewModel.PhoneNo = adminDetail.Info.PhoneNo
			viewModel.CompanyTeamName =companyTeamName
			viewModel.CompanyPlan =plan
			viewModel.AdminFirstName = storedSession.AdminFirstName
			viewModel.AdminLastName = storedSession.AdminLastName
			viewModel.ProfilePicture =storedSession.ProfilePicture
			c.Data["vm"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/edit-profile.html"

		case false:
			log.Println(helpers.ServerConnectionError)
		}

	}
}
func (c *RegisterController) ChangeAdminsPassword() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	ReadSession(w, r, companyTeamName)
	adminId :=storedSession.AdminId
	admin := models.Admins{}
	if r.Method == "POST" {
		confirmPassword := (c.GetString("confirmpassword"))
		log.Println(confirmPassword)
		dbStatus := admin.EditAdminPassword(c.AppEngineCtx, adminId,[] byte(confirmPassword))
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}
}
func (c *RegisterController) OldAdminPasswordCheck(){w := c.Ctx.ResponseWriter
	r := c.Ctx.Request
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	ReadSession(w, r, companyTeamName)
	adminId :=storedSession.AdminId
	enteredOldPassword := (c.GetString("oldPassword"))
	dbStatus := models.IsEnteredAdminPasswordCorrect(c.AppEngineCtx,adminId,[] byte(enteredOldPassword) )
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}
