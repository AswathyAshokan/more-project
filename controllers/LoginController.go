/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"log"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		login := models.Login{}
		login.Email = c.GetString("email")
		login.Password = []byte(c.GetString("password"))
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
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))

		}

	} else {
		c.TplName = "template/login.html"
	}

}



func (c *LoginController)Logout(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	ClearSession(w,r)
	storedSession := ReadSession(w, r)
	log.Println("session:", storedSession)
	http.Redirect(w, r, "/", 302)
}