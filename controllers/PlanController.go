package controllers

import (

	"log"
	"app/passporte/viewmodels"

)

type PlanController struct {
	BaseController
}



func (c *PlanController) PlanDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("The userDetails stored in session:",storedSession)
	log.Println("session:", storedSession)
	viewModel := viewmodels.Plan{}
	if cookie, err := r.Cookie("session"); err == nil {
		log.Println("cookie",cookie)
		viewModel.SessionFlag = true
	} else {
		viewModel.SessionFlag = false
	}
	c.Data["vm"] = viewModel
	c.TplName = "template/plan.html"
}


func (c *PlanController) PlanCheck() {
	log.Println("ghdgfhdgfdhg")
}