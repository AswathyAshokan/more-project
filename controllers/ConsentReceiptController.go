package controllers
import (
	"time"
	"app/passporte/models"
	"app/passporte/helpers"
	"app/passporte/viewmodels"
	"reflect"
	"log"
	"strings"
)
type ConsentReceiptController struct {
	BaseController
}
func (c *ConsentReceiptController) AddConsentReceipt() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	consentView :=viewmodels.ConsentReceipt{}
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	consentData := models.ConsentReceipts{}
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		members := models.ConsentMembers{}
		consentData.Info.ReceiptName = c.GetString("recieptName")
		tempGroupId := c.GetStrings("selectedUserIds")
		tempGroupMembers := c.GetStrings("selectedUserNames")
		instructions := c.GetString("instructionsForUser")
		instructionSlice := strings.Split(instructions, ",")
		log.Println("firt tot work id ",instructions)
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		tempMembersMap := make(map[string]models.ConsentMembers)
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberName = tempGroupMembers[i]
			tempMembersMap[tempGroupId[i]] = members
		}
		consentData.Members = tempMembersMap
		dbStatus := consentData.AddConsentToDb(c.AppEngineCtx,instructionSlice)

		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		groupUser := models.Users{}
		var keySlice []string
		var allUserNames [] string
		allUserDetails, dbStatus := groupUser.TakeGroupMemberName(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(allUserDetails)

			for _, groupKey := range dataValue.MapKeys() {
				keySlice = append(keySlice, groupKey.String())
			}
			for _, k := range keySlice {
				if allUserDetails[k].Status != helpers.UserStatusDeleted {
					allUserNames = append(allUserNames, allUserDetails[k].FullName)
					consentView.GroupMembers = allUserNames
					consentView.GroupKey = keySlice
				}
			}
			consentView.CompanyTeamName = storedSession.CompanyTeamName
			consentView.CompanyPlan   =  storedSession.CompanyPlan
			consentView.AdminLastName =storedSession.AdminLastName
			consentView.AdminFirstName =storedSession.AdminFirstName
			consentView.ProfilePicture =storedSession.ProfilePicture
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		c.Data["vm"] = consentView
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-consentreceipt.html"
	}
}
func (c* ConsentReceiptController)LoadConsentReceipt(){

	//r := c.Ctx.Request
	//w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	//storedSession := ReadSession(w, r, companyTeamName)
	//customerViewModel := viewmodels.Customer{}
	dbStatus,allConsent:= models.GetAllConsentReceiptDetails(c.AppEngineCtx,companyTeamName)
	log.Println("hhhh",allConsent,dbStatus)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allConsent)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			if allConsent[k].Settings.Status != helpers.UserStatusDeleted {
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.ReceiptName)
				/*tempValueSlice = append(tempValueSlice, allConsent[k].Info)
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.)
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.ZipCode)
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.Email)
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.Phone)
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.ContactPerson)
				tempValueSlice = append(tempValueSlice,k)
				customerViewModel.Values=append(customerViewModel.Values,tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}

		}
		customerViewModel.Keys = keySlice
		customerViewModel.CompanyTeamName = storedSession.CompanyTeamName
		customerViewModel.CompanyPlan = storedSession.CompanyPlan
		customerViewModel.AdminFirstName =storedSession.AdminFirstName
		customerViewModel.AdminLastName =storedSession.AdminLastName
		customerViewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = customerViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/customer-details.html"
	*//*case false:
		log.Println(helpers.ServerConnectionError)
	}*/


			}
		}
	}

	c.Layout = "layout/layout.html"
	c.TplName = "template/consentreceipt-details.html"

}