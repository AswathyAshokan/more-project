/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/helpers"
	"app/passporte/viewmodels"
	"log"
	"strings"

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
		dbStatus,_:= admin.CreateAdminAndCompany(c.AppEngineCtx, company)
		switch dbStatus{
		case false:

			w.Write([]byte("false"))
		case true:
			/*var keySlice string
			dataValue := reflect.ValueOf(companyDetails)
			for _, key := range dataValue.MapKeys() {
				keySlice = key.String()
			}
			company.Info.CompanyTeamName = keySlice
			log.Println("company",companyDetails)*/
			/*companyStatus :=companyDetails.UpdateCompanyTeamName(c.AppEngineCtx)
			switch companyStatus  {
			case true:

				log.Println("true")
			case false:
				log.Println("false")

			}*/
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
		tempProfile :=c.GetString("profilePicture")
		profilePicture :=strings.Replace(tempProfile, "/", "%2F", -2)
		//log.Println("dfhsfhshfsjfhhsfhhs",newone)
		admin.Settings.ProfilePicture = profilePicture
		//result := strings.Split(admin.Settings.ProfilePicture, "/")
		//
		//// Display all elements.
		//for i := range result {
		//	var urlOfPrifle []string
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"//")
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"/")
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"/")
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"/")
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"/")
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"/")
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"%2F")
		//	urlOfPrifle =append(urlOfPrifle,result[i]+"%2F")
		//	urlOfPrifle = urlOfPrifle[:0]
		//}
		//log.Println("profilecontoller",urlOfPrifle)
		tempThumbProfile :=c.GetString("thumbPicture")
		thumbPicture :=strings.Replace(tempThumbProfile, "/", "%2F", -2)
		admin.Settings.ThumbProfilePicture=thumbPicture
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
