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
		consentData.Info.CompanyTeamName = companyTeamName
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

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	dbStatus,allConsent:= models.GetAllConsentReceiptDetails(c.AppEngineCtx,companyTeamName)
	log.Println("hhhh",allConsent,dbStatus)
	consentViewModel :=viewmodels.LoadConsent{}
	var innerTableData [][]viewmodels.ConsentStruct
	var consentRecievedDatails []viewmodels.ConsentStruct
	var consentStruct viewmodels.ConsentStruct
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allConsent)
		var keySlice []string
		var instructionKeySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			if allConsent[k].Settings.Status != helpers.UserStatusDeleted {
				tempValueSlice = append(tempValueSlice,"")
				tempValueSlice = append(tempValueSlice,allConsent[k].Info.ReceiptName)
				//tempValueSlice = append(tempValueSlice,k)
				consentViewModel.Values = append(consentViewModel.Values,tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
				log.Println("name",allConsent[k].Info.ReceiptName)
				log.Println("yyy",allConsent[k].Instructions)
				consentStruct.UserName =allConsent[k].Info.ReceiptName
				consentStruct.UserKey = k
				instructionDataValue := reflect.ValueOf(allConsent[k].Instructions)
				for _, consentKey := range instructionDataValue.MapKeys() {
					instructionKeySlice = append(instructionKeySlice, consentKey.String())
				}
				for _, eachKey := range instructionKeySlice {
					//var instructionArray []string
					consentStruct.InstructionArray = append(consentStruct.InstructionArray,allConsent[k].Instructions[eachKey].Description)
				}
				consentRecievedDatails= append(consentRecievedDatails,consentStruct)
				innerTableData = append(innerTableData,consentRecievedDatails)
				consentViewModel.InnerContent = innerTableData
				innerTableData = innerTableData[:0]

			}


		}
		log.Println("view model",consentRecievedDatails)
		consentViewModel.Keys = keySlice
		consentViewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["array"] = consentViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/consentreceipt-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
	/*
	c.Layout = "layout/layout.html"
	c.TplName = "template/consentreceipt-details.html"
	*/

}