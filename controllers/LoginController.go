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
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		log.Println("hai I am here")
		login := models.Login{}
		login.Email = c.GetString("email")
		login.Password = []byte(c.GetString("password"))
		log.Println(login)
		loginStatus := login.CheckLogin(c.AppEngineCtx)
		switch loginStatus{
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		c.TplName = "template/login.html"
	}

}