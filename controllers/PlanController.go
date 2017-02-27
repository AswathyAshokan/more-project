package controllers

import (

	"log"
	"app/passporte/viewmodels"

	"app/passporte/models"
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
	planViewModel := viewmodels.Plan{}
	planViewModel.CompanyTeamName = storedSession.CompanyTeamName
	if cookie, err := r.Cookie("session"); err == nil {
		log.Println("cookie",cookie)
		planViewModel.SessionFlag = true
	} else {
		planViewModel.SessionFlag = false
	}
	c.Data["vm"] = planViewModel
	c.TplName = "template/plan.html"
}


func (c *PlanController) PlanCheck() {
	w := c.Ctx.ResponseWriter
	r := c.Ctx.Request
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w,r,companyTeamName)
	companyId := storedSession.CompanyId

	selectedCompanyPlan := c.GetString("companyPlan")
	company := models.Company{}
	company.Plan = selectedCompanyPlan
	dbStatus := company.ChangeCompanyPlan(c.AppEngineCtx,companyId)
	switch dbStatus {
	case true :
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))


	}



}