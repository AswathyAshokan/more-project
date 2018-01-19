/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"net/http"
	"encoding/json"
	"log"
	"reflect"
)

type LoginController struct {
	BaseController
}
func (c *LoginController) Root() {
	c.TplName = "template/root.html"

}

//function for checking the login details
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
			sessionValues.ProfilePicture =adminDetails.Settings.ThumbProfilePicture
			sessionValues.PaymentStatus = companyDetails.Settings.PaymentStatus
			SetSession(w, sessionValues)
			log.Println("hhhh",sessionValues.CompanyPlan)
			slices := []interface{}{"true", sessionValues.CompanyTeamName,sessionValues.AdminFirstName,sessionValues.PaymentStatus,sessionValues.CompanyPlan}
			sliceToClient, _ := json.Marshal(slices)
			log.Println("sliceToClient",sliceToClient)
			w.Write(sliceToClient)
		case false:
			dbStatus,superAdminDetails:= login.CheckSuperAdminLogin(c.AppEngineCtx)

			switch dbStatus {
			case true:

				dataValue := reflect.ValueOf(superAdminDetails)
				var keySlice []string
				for _, key := range dataValue.MapKeys() {
					keySlice = append(keySlice, key.String())
				}

				for _, k := range keySlice {
					sessionValueForSuperAdmin := SessionForAdminValues{}
					sessionValueForSuperAdmin.SuperAdminEmail = superAdminDetails[k].Info.Email
					sessionValueForSuperAdmin.SuperAdminId = k
					tempAdminFirstName := superAdminDetails[k].Info.FirstName
					tempAdminLastName := superAdminDetails[k].Info.LastName
					sessionValueForSuperAdmin.SuperAdminFullName = tempAdminFirstName+" "+tempAdminLastName
					SetSessionForSuperAdmin(w, sessionValueForSuperAdmin)
					slices := []interface{}{"SuperAdmin"}
					sliceToClient, _ := json.Marshal(slices)
					w.Write(sliceToClient)

				}
			case false:
				log.Println("false")
				w.Write([]byte("false"))
			}

		}

	} else {
		c.TplName = "template/login.html"
	}

}
//logout section
func (c *LoginController)Logout(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	ClearSession(w)
	http.Redirect(w, r, "/login", 302)
}

//logout section of super admin
func (c *LoginController)LogoutForSuperAdmin(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	ClearSessionForSuperAdmin(w)
	http.Redirect(w, r, "/login", 302)
}
