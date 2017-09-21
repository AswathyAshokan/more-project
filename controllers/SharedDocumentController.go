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
	dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
	var notificationCount=0
	switch dbStatus {
	case true:

		notificationOfUser := reflect.ValueOf(notificationValue)
		for _, notificationUserKey := range notificationOfUser.MapKeys() {
			dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
			switch dbStatus {
			case true:
				notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
				for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
					var NotificationArray []string
					if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
						notificationCount=notificationCount+1;
					}
					NotificationArray =append(NotificationArray,notificationUserKey.String())
					NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
					date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
					NotificationArray =append(NotificationArray,date)
					documentsViewModels.NotificationArray=append(documentsViewModels.NotificationArray,NotificationArray)

				}
			case false:
			}
		}
	case false:
	}
	documentsViewModels.NotificationNumber=notificationCount
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

	dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
	var notificationCount=0
	switch dbStatus {
	case true:

		notificationOfUser := reflect.ValueOf(notificationValue)
		for _, notificationUserKey := range notificationOfUser.MapKeys() {
			dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
			switch dbStatus {
			case true:
				notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
				for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
					var NotificationArray []string
					if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
						notificationCount=notificationCount+1;
					}
					NotificationArray =append(NotificationArray,notificationUserKey.String())
					NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
					NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
					date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
					NotificationArray =append(NotificationArray,date)
					documentsViewModels.NotificationArray=append(documentsViewModels.NotificationArray,NotificationArray)

				}
			case false:
			}
		}
	case false:
	}
	documentsViewModels.NotificationNumber=notificationCount
	documentsViewModels.Values = AllSharedDocument
	documentsViewModels.Keys = expiryKeySlice
	documentsViewModels.CompanyTeamName = storedSession.CompanyTeamName
	documentsViewModels.CompanyPlan = storedSession.CompanyPlan
	c.Data["vm"] = documentsViewModels
	c.TplName = "template/shareddocument.html"
}






