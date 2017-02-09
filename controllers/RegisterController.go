/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"log"
)

type RegisterController struct {
	BaseController
}

func (c *RegisterController) Register() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		company := models.Company{}
		info := models.Info{}
		info.FirstName = c.GetString("firstName")
		info.LastName = c.GetString("lastName")
		info.PhoneNo = c.GetString("phoneNo")
		info.Email = c.GetString("emailId")
		info.Password = []byte(c.GetString("password"))
		info.CompanyName = c.GetString("companyName")
		info.Address = c.GetString("address")
		info.State = c.GetString("state")
		info.ZipCode = c.GetString("zipCode")
		company.Info = info
		log.Println("Registration Details:", company)
		dbStatus := company.AddUser(c.AppEngineCtx)
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