/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"net/http"
	"encoding/json"
)

type LoginController struct {
	BaseController
}
func (c *LoginController) Root() {
	c.TplName = "template/root.html"

}
func (c *LoginController) Login() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {

		login := models.Login{}
		login.Email = c.GetString("email")
		login.Password = []byte(c.GetString("password"))
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
			slices := []interface{}{"true", sessionValues.CompanyTeamName}
			sliceToClient, _ := json.Marshal(slices)
			w.Write(sliceToClient)
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
	ClearSession(w)
	http.Redirect(w, r, "/login", 302)
}