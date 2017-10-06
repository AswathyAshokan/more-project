package controllers


import (
	"log"
	"app/passporte/models"
	"reflect"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"strconv"
)
type SharedDocumentController struct {
	BaseController
}

func (c *SharedDocumentController) LoadSharedDocuments() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	var expiryKeySlice []string
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	userId := c.Ctx.Input.Param(":inviteuserid")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("session :", storedSession)
	documentsViewModels := viewmodels.SharedDocument{}
	info, dbStatus := models.GetAllInvitationDetail(c.AppEngineCtx, userId)

	if info.UserResponse == helpers.StatusAccepted {
		switch dbStatus {
		case true:
			tempEmailId := info.Email
			UserDetails:= models.GetAllUserDetail(c.AppEngineCtx, tempEmailId)
			log.Println("UserDetails",UserDetails)
			/*switch expiryStatus {
			case true:*/
				var keySlice []string
				dataValue := reflect.ValueOf(UserDetails)
				for _, key := range dataValue.MapKeys() {
					keySlice = append(keySlice, key.String())
				}
				for _, specifiedUserId := range keySlice {
					log.Println("specifiedUserId",specifiedUserId)
					expiry, status,Name := models.GetExpireDetailsOfUser(c.AppEngineCtx, specifiedUserId)
					log.Println("expiry", expiry)
					switch status {
					case true:
						dataValue := reflect.ValueOf(expiry)

						for _, key := range dataValue.MapKeys() {
							expiryKeySlice = append(expiryKeySlice, key.String())
						}
						for _, k := range expiryKeySlice {
							var tempValueSlice []string
							if expiry[k].Info.Mode == "Public" {
								tempValueSlice = append(tempValueSlice, expiry[k].Info.Description)
								expirationDate := strconv.FormatInt(int64(expiry[k].Info.ExpirationDate), 10)
								tempValueSlice = append(tempValueSlice, expirationDate)
								tempValueSlice = append(tempValueSlice,Name)
								tempValueSlice = append(tempValueSlice, expiry[k].Info.DocumentId)

								documentsViewModels.Values = append(documentsViewModels.Values, tempValueSlice)
								tempValueSlice = tempValueSlice[:0]
							}

						}
					case false :
						log.Println(helpers.ServerConnectionError)
					}

				}
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		}


		documentsViewModels.Keys = expiryKeySlice
		documentsViewModels.CompanyTeamName = storedSession.CompanyTeamName
		documentsViewModels.CompanyPlan = storedSession.CompanyPlan
		c.Data["vm"] = documentsViewModels
		c.TplName = "template/shareddocument.html"


}
func (c *SharedDocumentController) LoadSharedDocumentsAllSharedDocuments() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	var expiryKeySlice []string
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	log.Println("session :", storedSession)
	documentsViewModels := viewmodels.SharedDocument{}
	_, _ ,_,expiryKeySlice,AllSharedDocument:= models.GetAllSharedDocumentsByCompany(c.AppEngineCtx, companyTeamName)
	log.Println("AllSharedDocument",AllSharedDocument)
	/*switch dbStatus {
	case true:

		log.Println("expiry",expiry,Name)
		//if expiry.Info.Mode =="Public"{
		//	var tempValueSlice []string
		//	tempValueSlice = append(tempValueSlice, expiry.Info.Description)
		//	tempValueSlice = append(tempValueSlice, time.Unix(expiry.Info.ExpirationDate, 0).Format("01/02/2006"))
		//	tempValueSlice = append(tempValueSlice, Name)
		//	tempValueSlice = append(tempValueSlice, expiry.Info.DocumentId)
			documentsViewModels.Values = AllSharedDocument
		//
		//}
		log.Println("expiryKeySlice",expiryKeySlice)
		log.Println("alllll",documentsViewModels.Values)
	case false :
		log.Println(helpers.ServerConnectionError)
	}*/



	documentsViewModels.Values = AllSharedDocument
	documentsViewModels.Keys = expiryKeySlice
	documentsViewModels.CompanyTeamName = storedSession.CompanyTeamName
	documentsViewModels.CompanyPlan = storedSession.CompanyPlan
	c.Data["vm"] = documentsViewModels
	c.TplName = "template/shareddocument.html"
}






