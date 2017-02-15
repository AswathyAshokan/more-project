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
		companyAdmins := models.CompanyAdmins{}
		info := models.Info{}
		settings := models.Settings{}
		info.FirstName = c.GetString("firstName")
		info.LastName = c.GetString("lastName")
		info.PhoneNo = c.GetString("phoneNo")
		info.Email = c.GetString("emailId")
		info.Password = []byte(c.GetString("password"))
		info.CompanyName = c.GetString("companyName")
		info.Address = c.GetString("address")
		info.State = c.GetString("state")
		info.ZipCode = c.GetString("zipCode")
		settings.DateCreated = time.Now().Unix()
		settings.Status = helpers.StatusActive
		companyAdmins.Info = info
		companyAdmins.Settings = settings
		log.Println("Registration Details:", companyAdmins)
		dbStatus := companyAdmins.AddUser(c.AppEngineCtx)
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

