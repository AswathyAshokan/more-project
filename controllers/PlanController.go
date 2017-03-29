package controllers

import (

	"app/passporte/viewmodels"

	"app/passporte/models"
	"encoding/json"
)

type PlanController struct {
	BaseController
}

//to Display Plan Details
func (c *PlanController) PlanDetails() {
	planViewModel := viewmodels.Plan{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	_, sessionStatus := SessionForPlan(w,r)
	planViewModel.SessionFlag = sessionStatus
	c.Data["vm"] = planViewModel
	c.TplName = "template/plan.html"
}

//For update company plan with newly selected company plan
func (c *PlanController) PlanUpdate() {
	w := c.Ctx.ResponseWriter
	r := c.Ctx.Request
	storedSession,_ := SessionForPlan(w,r)
	companyId := storedSession.CompanyId
	selectedCompanyPlan := c.GetString("companyPlan")
	company := models.Company{}
	company.Plan = selectedCompanyPlan
	dbStatus, _ := company.ChangeCompanyPlan(c.AppEngineCtx,companyId)
	switch dbStatus {
	case true :

		ClearSession(w)
		sessionValues := SessionValues{}
		sessionValues.AdminId = storedSession.AdminId
		sessionValues.AdminFirstName = storedSession.AdminFirstName
		sessionValues.AdminLastName = storedSession.AdminLastName
		sessionValues.AdminEmail = storedSession.AdminEmail
		sessionValues.CompanyId = storedSession.CompanyId
		sessionValues.CompanyName = storedSession.CompanyName
		sessionValues.CompanyTeamName = storedSession.CompanyTeamName
		sessionValues.CompanyPlan = selectedCompanyPlan

		SetSession(w, sessionValues)
		slices := []interface{}{"true", sessionValues.CompanyTeamName}
		sliceToClient, _ := json.Marshal(slices)
		w.Write(sliceToClient)
	case false:
		w.Write([]byte("false"))


	}



}


