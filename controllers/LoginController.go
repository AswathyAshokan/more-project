/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
"app/passporte/models"
"log"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {
	r := c.Ctx.Request
	if r.Method == "POST" {
		log.Println("hai I am here")
		login := models.Login{}
		login.Email = c.GetString("email")
		login.Password = []byte(c.GetString("password"))
		log.Println(login)
		login.CheckLogin(c.AppEngineCtx)
	} else {
		c.TplName = "template/login.html"
	}

}