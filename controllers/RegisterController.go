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
	"net/smtp"
	"math/rand"
	"reflect"
	"encoding/json"

	"sort"
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
		company.Info.Country = c.GetString("country")
		company.Info.Number = c.GetString("number")
		company.Info.City = c.GetString("city")
		company.Info.Street = c.GetString("street")
		company.Info.State = c.GetString("state")
		company.Info.ZipCode = c.GetString("zipCode")
		company.Settings.Status = helpers.StatusActive
		company.Settings.DateOfCreation = currentTime
		company.Plan = c.GetString("planType")
		company.Settings.PaymentStatus = helpers.PaymentPendingStatus
		admin := models.Admins{}
		admin.Info.FirstName = c.GetString("firstName")
		admin.Info.LastName = c.GetString("lastName")
		admin.Info.Email = c.GetString("emailId")
		admin.Info.PhoneNo = c.GetString("phoneNo")
		admin.Info.Password = []byte(c.GetString("password"))
		admin.Settings.DateOfCreation = currentTime
		admin.Settings.Status = helpers.StatusActive
		log.Println("state",company.Info.State)
		if company.Plan == helpers.PlanBusiness{
			company.Settings.LimitedUsers ="4"
		}
		dbStatus,_:= admin.CreateAdminAndCompany(c.AppEngineCtx, company)
		log.Println("dbStatus",dbStatus)
		switch dbStatus{
		case false:
			w.Write([]byte("false"))
		case true:
			w.Write([]byte("true"))
		}
	} else {

		//var tempDialCode   []string
		countryViewModel := viewmodels.DisplayCountryDetails{}
		var keySlice []string
		dbStatus, allCountryValues := models.GetAllCountryNameForFillDropDownList(c.AppEngineCtx)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(allCountryValues)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for _, k := range keySlice{
				var tempAllCountry []string
				tempAllCountry = append(tempAllCountry, allCountryValues[k].CountryName)
				//countryViewModel.CountryName = tempAllCountry
				tempAllCountry= append(tempAllCountry,allCountryValues[k].DialCode)
				tempAllCountry =append(tempAllCountry, k)
				tempCountryNameArray :=  append(countryViewModel.CountryName,allCountryValues[k].CountryName)
				sort.Strings(tempCountryNameArray)
				countryViewModel.CountryName = tempCountryNameArray
				countryViewModel.CountryAllData = append(countryViewModel.CountryAllData,tempAllCountry)
			}
			countryViewModel.Key = keySlice

		case false:
			log.Println(helpers.ServerConnectionError)
		}
		c.Data["vm"] = countryViewModel
		c.TplName = "template/register.html"
	}
}

func (c *RegisterController)CheckEmail(){
	w := c.Ctx.ResponseWriter
	emailId := c.GetString("emailId")
	isEmailUsed := models.CheckEmailIsUsed(c.AppEngineCtx, emailId)
	if isEmailUsed == false {
		log.Println("truesss")
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
		log.Println("t1",admin.Info.FirstName )
		admin.Info.Email = c.GetString("emailId")
		log.Println("t2",admin.Info.Email)
		admin.Info.PhoneNo = c.GetString("phoneNumber")
		log.Println("t3",admin.Info.PhoneNo)
		var profilePicture string
		tempProfile :=c.GetString("profilePicture")
		if len(tempProfile) !=0 {
			parts := strings.Split(tempProfile, "/")
			parts = append(parts[:len(parts)-3], strings.Join(parts[len(parts)-3:], "%2F"))
			profilePicture = strings.Join(parts, "/")
			admin.Settings.ProfilePicture = profilePicture
		}


		tempThumbProfile :=c.GetString("thumbPicture")
		if len(tempThumbProfile) !=0{
			thumParts := strings.Split(tempThumbProfile, "/")
			thumParts = append(thumParts[:len(thumParts)-3], strings.Join(thumParts[len(thumParts)-3:], "%2F"))
			thumbPicture := strings.Join(thumParts, "/")
			admin.Settings.ThumbProfilePicture=thumbPicture
		}

		log.Println("profile picture",profilePicture)
		dbStatus,adminPicture := admin.EditAdminDetails(c.AppEngineCtx, adminId)
		sessionValues := SessionValues{}
		sessionValues.AdminId = adminId
		sessionValues.AdminFirstName = admin.Info.FirstName
		sessionValues.AdminLastName = admin.Info.LastName
		sessionValues.AdminEmail = admin.Info.Email
		sessionValues.CompanyId = admin.Company.CompanyId
		sessionValues.CompanyName = storedSession.CompanyName
		sessionValues.CompanyTeamName = storedSession.CompanyTeamName
		sessionValues.CompanyPlan = storedSession.CompanyPlan
		if len(profilePicture)!=0{
			sessionValues.ProfilePicture =profilePicture
			log.Println("valuesss",profilePicture)
		}else{
			sessionValues.ProfilePicture=adminPicture
			log.Println("admin",adminPicture)
		}

		sessionValues.PaymentStatus = storedSession.PaymentStatus
		SetSession(w, sessionValues)
		switch dbStatus {

		case true:
			log.Println("trueeee")
			w.Write([]byte("true"))
		case false:
			log.Println("false")
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
func (c *RegisterController) ForgotPassword(){
	//w := c.Ctx.ResponseWriter
	//r := c.Ctx.Request
	c.Layout = "layout/layout.html"
	c.TplName = "template/forgot-password.html"

}


func (c *RegisterController)CheckingEmailId(){
	w := c.Ctx.ResponseWriter

	//viewModel := viewmodels.ForgotPassword{}
	emailId := c.GetString("emailId")
	log.Println("used email",emailId)
	isEmailUsed := models.CheckEmailIsUsed(c.AppEngineCtx, emailId)
	log.Println("inside",isEmailUsed)
	if isEmailUsed == false {
		var r *rand.Rand
		r = rand.New(rand.NewSource(time.Now().UnixNano()))

		const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
		result := make([]byte, 8)
		for i := range result {
			result[i] = chars[r.Intn(len(chars))]
		}
		//const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		//b := make([]byte, 8)
		//for i := range b {
		//	b[i] = letterBytes[rand.Intn(len(letterBytes))]
		//}

		body :="Dear member, we received a request for password change .this is your automatic genereted key "+string(result)
		//+"Go to site to set your new password. The key will be active for 10 minutes"

			//"Regards,"+
			//"The Passporte team"
		from := "passportetest@gmail.com"
		to := emailId
		subject := "Subject: Passporte - Forgot Password\n"
		mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		message := []byte(subject + mime + "\n" + body)
		if err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", "passportetest@gmail.com", "passporte123", "smtp.gmail.com"), from, []string{to}, []byte(message)); err != nil {
			log.Println(err)
		}
		//w.Write([]byte("false,"))
		//w.Write([]byte(string(result)))
		slices := []interface{}{"false", string(result)}
		sliceToClient, _ := json.Marshal(slices)
		w.Write(sliceToClient)
	}else{

		w.Write([]byte("true"))
	}
}

func (c *RegisterController) ResetPassword() {
	//r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	admin := models.Admins{}
	confirmPassword := []byte(c.GetString("confirmpassword"))
	log.Println("confirmpassword",confirmPassword)
	emailId := (c.GetString("emailId"))
	log.Println("emailAddress",emailId)
	dbStatus,adminDetails := admin.AdminDetails(c.AppEngineCtx)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(adminDetails)
		for _, key := range dataValue.MapKeys() {
			if adminDetails[key.String()].Info.Email == emailId{
				dbStatus := admin.EditAdminPassword(c.AppEngineCtx, key.String(),[] byte(confirmPassword))
				switch dbStatus {
				case true:
					w.Write([]byte("true"))
				case false:
					w.Write([]byte("false"))
				}

			}

		}

		//w.Write([]byte("true"))
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

/*----------------------------GetStates-----------------------------------------*/
func (c *RegisterController) GetStates() {
	w := c.Ctx.ResponseWriter
	CountryName := c.GetString("countryName")
	countryCode := c.GetString("countryCode")
	log.Println("countryCode",countryCode)
	var keySlice []string
	var tempStateArray  []string
	dbStatus, allStatesOfSelectedCountry:= models.GetAllStatesByCountry(c.AppEngineCtx,CountryName,countryCode)
	switch dbStatus {
	case true:
		//tempStateArray = allStatesOfSelectedCountry.States
		countryValues := reflect.ValueOf(allStatesOfSelectedCountry.States)
		for _, key := range countryValues.MapKeys() {
			keySlice = append(keySlice,key.String())
		}
		for _, k := range keySlice {
			 tempStateArray = append(tempStateArray,allStatesOfSelectedCountry.States[k])
		}
		log.Println("tempStateArray",tempStateArray)
		slices := []interface{}{"true",tempStateArray}
		sliceToClient, _ := json.Marshal(slices)
		w.Write(sliceToClient)
	case false:
		w.Write([]byte("false"))

	}
}