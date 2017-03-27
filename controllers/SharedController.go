package controllers


/*import (

)*/

type SharedDocumentController struct {
	BaseController
}

func (c *SharedDocumentController) SharedDocuments() {
	/*r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("session :",storedSession)*/
	c.Layout = "layout/layout.html"
	c.TplName = "template/shareddocument.html"
}

