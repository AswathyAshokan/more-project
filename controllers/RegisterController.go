/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"log"
	"time"
	"app/passporte/helpers"
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
		company.Info.TeamName = c.GetString("teamName")
		company.Info.Address = c.GetString("address")
		company.Info.State = c.GetString("state")
		company.Info.ZipCode = c.GetString("zipCode")
		company.Settings.Status = helpers.StatusActive
		company.Settings.DateOfCreation = currentTime
		company.Plan = helpers.PlanFamily
		admin := models.Admins{}
		admin.Info.CompanyName = company.Info.CompanyName
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
	log.Println("Checking isEmailUsed:",emailId)
	isEmailUsed := models.CheckEmailIsUsed(c.AppEngineCtx, emailId)
	if isEmailUsed == false {
		w.Write([]byte("false"))
	}else{
		w.Write([]byte("true"))
	}
}

