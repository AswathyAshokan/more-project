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
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	log.Println("session:", storedSession)
	planViewMidel := viewmodels.Plan{}
	planViewMidel.CompanyTeamName = storedSession.CompanyTeamName
	if cookie, err := r.Cookie("session"); err == nil {
		log.Println("cookie",cookie)
		planViewMidel.SessionFlag = true
	} else {
		planViewMidel.SessionFlag = false
	}
	c.Data["vm"] = planViewMidel
	c.TplName = "template/plan.html"
}


func (c *PlanController) PlanCheck() {
	/*r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	planViewMidel := viewmodels.Plan{}
	*/
	log.Println("ghdgfhdgfdhg")
}