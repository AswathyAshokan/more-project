package controllers


import (
	"log"
	"app/passporte/models"
	"reflect"
	"app/passporte/viewmodels"
	"strconv"
	"app/passporte/helpers"
)

type SharedDocumentController struct {
	BaseController
}

func (c *SharedDocumentController) LoadSharedDocuments() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	userId :=c.Ctx.Input.Param(":inviteuserid")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("session :",storedSession)
	documentsViewModels := viewmodels.SharedDocument{}
	info,dbStatus := models.GetAllInvitationDetail(c.AppEngineCtx,userId)
	var expiryKeySlice []string
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
			for _, specifiedUserId := range keySlice {
				expiry,status := models.GetExpireDetailsOfUser(c.AppEngineCtx,specifiedUserId)
				switch status {
				case true:
					dataValue := reflect.ValueOf(expiry)

					for _, key := range dataValue.MapKeys() {
						expiryKeySlice = append(expiryKeySlice, key.String())
					}
					for _, k := range expiryKeySlice {
						var tempValueSlice []string
						tempValueSlice = append(tempValueSlice, expiry[k].Info.Description)
						tempValueSlice = append(tempValueSlice, strconv.FormatInt(expiry[k].Info.ExpirationDate, 10))
						documentsViewModels.Values=append(documentsViewModels.Values,tempValueSlice)
						log.Println("viewmodel :",documentsViewModels)
						tempValueSlice = tempValueSlice[:0]
					}
				case false :
					log.Println(helpers.ServerConnectionError)
				}
				documentsViewModels.Keys= expiryKeySlice
				c.Data["array"] = documentsViewModels
				c.Layout = "layout/layout.html"
				c.TplName = "template/shareddocument.html"
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	case false:
		log.Println(helpers.ServerConnectionError)
	}

}

