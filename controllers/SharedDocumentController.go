package controllers


import (
	"log"
	"app/passporte/models"
	"reflect"

)

type SharedDocumentController struct {
	BaseController
}

func (c *SharedDocumentController) SharedDocuments() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	userId :=c.Ctx.Input.Param(":inviteuserid")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("session :",storedSession)
	/*documentsViewModels := viewmodels.SharedDocument{}*/
	info,dbStatus := models.GetAllInvitationDetail(c.AppEngineCtx,userId)
	switch dbStatus {
	case true:
		tempEmailId := info.Email
		UserDetails,expiryStatus  := models.GetAllUserDetail(c.AppEngineCtx,tempEmailId)
		switch expiryStatus {
		case true:
			var keySlice []string
			dataValue := reflect.ValueOf(UserDetails)
			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			/*for _, specifiedUserId := range keySlice {
				expiry,status := models.GetExpireDetailsOfUser(c.AppEngineCtx,specifiedUserId)
				switch status {
				case true:
					log.Println("dfdfdsafs:",expiry)
					dataValue := reflect.ValueOf(expiry)
					for _, key := range dataValue.MapKeys() {
						keySlice = append(keySlice, key.String())
					}
					for _, k := range keySlice {
						documentsViewModels.DateOfExpiry = expiry[k].Info.ExpirationDate
						documentsViewModels.Description = expiry[k].Info.Description
						*//*documentsViewModels.DocumentLocation = expiry[k].Info.*//*
						log.Println("valuesss:",documentsViewModels)


					}
				case false :
					log.Println("false")


				}
			}*/
		case false:
			log.Println("false")
		}
	case false:
		log.Println("false")
	}
	/*c.Data["vm"] = documentsViewModels*/
	c.Layout = "layout/layout.html"
	c.TplName = "template/shareddocument.html"
}

