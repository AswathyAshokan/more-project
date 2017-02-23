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
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	reflect.TypeOf(storedSession)
	viewModel := viewmodels.Plan{}
	viewModel.Email = storedSession.Info.Email
	viewModel.FirstName = storedSession.Info.FirstName
	viewModel.SecondName = storedSession.Info.LastName
	log.Println("ggg",viewModel)
	c.Data["vm"] = viewModel
	c.TplName = "template/plan.html"
}
