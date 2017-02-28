package controllers

import (

	"log"
	"app/passporte/viewmodels"

	"app/passporte/models"
)

type PlanController struct {
	BaseController
}

//to Display Plan Details
func (c *PlanController) PlanDetails() {
	log.Println("cp1")
	planViewModel := viewmodels.Plan{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession, sessionStatus := SessionForPlan(w,r)
	switch sessionStatus {
	case true:
		planViewModel.CompanyTeamName = storedSession.CompanyTeamName
		planViewModel.SessionFlag = sessionStatus
	case false:
		planViewModel.SessionFlag = sessionStatus
	}
	c.Data["vm"] = planViewModel
	c.TplName = "template/plan.html"
}

//For update company plan with newly selected company plan
func (c *PlanController) PlanUpdate() {
	w := c.Ctx.ResponseWriter
	r := c.Ctx.Request
	/*sessionValues := SessionValues{}*/
	storedSession,_ := SessionForPlan(w,r)
	companyId := storedSession.CompanyId
	selectedCompanyPlan := c.GetString("companyPlan")
	company := models.Company{}
	company.Plan = selectedCompanyPlan
	dbStatus,companyPlanUpdate := company.ChangeCompanyPlan(c.AppEngineCtx,companyId)
	switch dbStatus {
	case true :
		log.Println(companyPlanUpdate)
		/*sessionValues.CompanyPlan = companyPlanUpdate.Plan
		SetSession(w, sessionValues)*/
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))


	}



}


