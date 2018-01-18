
//created for pending works

package controllers

import (
	"app/passporte/viewmodels"
)

type ComingSoonController struct {
	BaseController
}

//to Display Plan Details
func (c *ComingSoonController) LoadComingSoonController() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)

	LoadPendingWork := viewmodels.LoadPendingWork{}
	LoadPendingWork.CompanyTeamName = storedSession.CompanyTeamName
	c.Data["vm"] = LoadPendingWork
	c.Layout = "layout/layout.html"
	c.TplName = "template/pendingWorks.html"
}
