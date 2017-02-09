/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
//"app/passporte/models"
//"log"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {
	r := c.Ctx.Request
	if r.Method == "POST" {
		//user := models.User{}
		//user.Email = c.GetString("email")
		//user.Password = []byte(c.GetString("password"))
		//log.Println(user)
	} else {
		c.TplName = "template/login.html"
	}

}