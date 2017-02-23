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
		log.Println("hai I am here")
		login := models.Login{}
		login.Email = c.GetString("email")
		login.Password = []byte(c.GetString("password"))
		log.Println(login)
		loginStatus, adminDetails := login.CheckLogin(c.AppEngineCtx)
		switch loginStatus{
		case true:
			log.Println("Login Successful!")
			sessionValues := SessionValues{}
			sessionValues.AdminId = "Chumma oru Id"
			sessionValues.AdminFirstName = adminDetails.Info.FirstName
			sessionValues.AdminLastName = adminDetails.Info.LastName
			sessionValues.AdminEmail = adminDetails.Info.Email
			sessionValues.CompanyId = "xyz"
			sessionValues.CompanyName = "dddd"
			sessionValues.CompanyTeamName = "dsfdss"
			sessionValues.CompanyPlan = "Family"
			SetSession(w, sessionValues)
			w.Write([]byte("true"))
		case false:
			log.Println("Invalid Username or Password!")
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