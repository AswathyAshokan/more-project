package controllers

import (
	"app/passporte/models"
	"log"
	"net/http"
)

type ByPassController struct {
	BaseController
}

/*Func for session bypass*/
func (c *ByPassController)ByPass() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	login := models.Login{}
	login.Email = "rahul.r@cynere.com"
	login.Password = []byte("Password321")
	log.Println(login)
	loginStatus, adminDetails, companyDetails, adminId := login.CheckLogin(c.AppEngineCtx)
	switch loginStatus{
	case true:
		sessionValues := SessionValues{}
		sessionValues.AdminId = adminId
		sessionValues.AdminFirstName = adminDetails.Info.FirstName
		sessionValues.AdminLastName = adminDetails.Info.LastName
		sessionValues.AdminEmail = adminDetails.Info.Email
		sessionValues.CompanyId = adminDetails.Company.CompanyId
		sessionValues.CompanyName = companyDetails.Info.CompanyName
		sessionValues.CompanyTeamName = companyDetails.Info.CompanyTeamName
		sessionValues.CompanyPlan = companyDetails.Plan
		SetSession(w, sessionValues)

		initialLink := "/" + sessionValues.CompanyTeamName + "/task/add"

		http.Redirect(w, r, initialLink, 302)

	case false:
		log.Println("Bypass Failed")


	}
}