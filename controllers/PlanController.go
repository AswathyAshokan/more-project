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
	planViewModel := viewmodels.Plan{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := SessionForPlan(w,r)
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = cookieToken.Decode("session", cookie.Value, &value); err == nil {
			log.Println("cookie", cookie)
			planViewModel.SessionFlag = true
			planViewModel.CompanyTeamName = storedSession.CompanyTeamName
		} else {
			log.Println("session cannot set")
		}
	} else {
		log.Println("log first")
		planViewModel.SessionFlag = false
	}
	log.Println()
	c.Data["vm"] = planViewModel
	c.TplName = "template/plan.html"
}

//For update company plan with newly selected company plan
func (c *PlanController) PlanUpdate() {
	w := c.Ctx.ResponseWriter
	r := c.Ctx.Request
	/*sessionValues := SessionValues{}*/
	storedSession := SessionForPlan(w,r)
	companyId := storedSession.CompanyId
	log.Println("compant id",companyId)
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


