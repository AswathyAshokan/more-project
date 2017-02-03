/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"log"
)

type RegisterController struct{
	BaseController
}


func (c *RegisterController) Register(){

	r := c.Ctx.Request
	if r.Method =="POST" {

		user := models.User{}
		user.FirstName = c.GetString("firstName")
		user.LastName = c.GetString("lastName")
		user.PhoneNo = c.GetString("phoneNo")
		user.Email = c.GetString("emailId")
		user.Password = []byte(c.GetString("password"))
		user.CompanyName = c.GetString("companyName")
		user.Address = c.GetString("address")
		user.State = c.GetString("state")
		user.ZipCode = c.GetString("zipCode")
		log.Println("Registration Details:", user)
		user.AddUser(c.AppEngineCtx)
	}else{
		//c.Layout = "layout/default-layout.html"
		c.TplName = "template/register.html"
	}



}