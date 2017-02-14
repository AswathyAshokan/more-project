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
		loginStatus, adminDetails := login.CheckLogin(c.AppEngineCtx)
		switch loginStatus{
		case true:
			log.Println("Login Successful!")
			SetSession(w, adminDetails)
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
	log.Println("The username stored in session:",storedSession.Info.Email)
	log.Println("The lastName stored in session:",storedSession.Info.LastName)
}