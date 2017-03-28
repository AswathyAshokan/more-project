package controllers


import (
	"log"
	"app/passporte/models"
)

type SharedDocumentController struct {
	BaseController
}

func (c *SharedDocumentController) SharedDocuments() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	userId  := c.Ctx.Input.Param("::inviteuserid")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("session :",storedSession)
	info,dbStatus := models.GetAllUsersDetail(c.AppEngineCtx,userId)
	switch dbStatus {
	case true:
		log.Println("true",info)
	case false:
		log.Println("false")
	}
	c.Layout = "layout/layout.html"
	c.TplName = "template/shareddocument.html"
}

