package controllers

import (

	"log"
	"app/passporte/viewmodels"

	"reflect"
)

type PlanController struct {
	BaseController
}



func (c *PlanController) PlanDetails() {
	r := c.Ctx.Request
	//w := c.Ctx.ResponseWriter
	//storedSession := ReadSession(w, r)
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	reflect.TypeOf(storedSession)

	log.Println("session:", storedSession)
	viewModel := viewmodels.Plan{}
	//sessionValues := SessionValues{}

	if cookie, err := r.Cookie("session"); err == nil {
		log.Println("cookie",cookie)
		viewModel.SessionFlag = true
	} else {
		viewModel.SessionFlag = false
	}

	/*viewModel.Email = storedSession.AdminEmail
	viewModel.FirstName = storedSession.AdminFirstName
	viewModel.SecondName = storedSession.AdminLastName*/
	log.Println("ggg",viewModel)
	c.Data["vm"] = viewModel
	c.TplName = "template/plan.html"
}


func (c *PlanController) PlanCheck() {
	/*log.Println("fff")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)

	log.Println(storedSession.Info.CompanyName)*/
}