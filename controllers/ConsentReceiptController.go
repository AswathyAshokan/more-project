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
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		tempMembersMap := make(map[string]models.ConsentMembers)
		for i := 0; i < len(tempGroupId); i++ {
			members.Status = helpers.StatusAccepted
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
	consentViewModel :=viewmodels.LoadConsent{}
	var innerTableData [][]viewmodels.ConsentStruct
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allConsent)
		var keySlice []string
		var instructionKeySlice []string
		var memberKeySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k :=range keySlice{
			var tempValueSlice []string
			if allConsent[k].Settings.Status != helpers.UserStatusDeleted {
				log.Println("in loop 1")
				tempValueSlice = append(tempValueSlice, "")
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.ReceiptName)
				consentViewModel.Values = append(consentViewModel.Values, tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			//}

		//}
		//for _, k := range keySlice {
			//if allConsent[k].Settings.Status != helpers.UserStatusDeleted {
				log.Println("in loop 2")
				var consentRecievedDetails []viewmodels.ConsentStruct
				var consentStruct viewmodels.ConsentStruct
				//consentStruct.Status = allConsent[k].Settings.Status
				memberDataValue :=reflect.ValueOf(allConsent[k].Members)
				// for get name of members from members map
				for _, membersKey := range memberDataValue.MapKeys() {
					memberKeySlice = append(memberKeySlice, membersKey.String())
				}
				for _, eachMemberKey := range memberKeySlice {
					if allConsent[k].Members[eachMemberKey].Status ==helpers.StatusAccepted{
						consentStruct.AcceptedUsers = append(consentStruct.AcceptedUsers,allConsent[k].Members[eachMemberKey].MemberName)
						consentStruct.AcceptedUsers = append(consentStruct.AcceptedUsers, k)
					} else {
						consentStruct.RejectedUsers = append(consentStruct.RejectedUsers,allConsent[k].Members[eachMemberKey].MemberName)
						consentStruct.AcceptedUsers = append(consentStruct.RejectedUsers, k)
					}
					//consentStruct.Status = allConsent[k].Members[eachMemberKey].Status
					//consentStruct.UserName = append(consentStruct.UserName,allConsent[k].Members[eachMemberKey].MemberName)
				}
				memberKeySlice = memberKeySlice[:0]

				consentStruct.UserKey = k
				instructionDataValue := reflect.ValueOf(allConsent[k].Instructions)
				// for get each instructions  from instruction map
				for _, consentKey := range instructionDataValue.MapKeys() {
					instructionKeySlice = append(instructionKeySlice, consentKey.String())
				}
				for _, eachKey := range instructionKeySlice {
					//consentStruct.Status = allConsent[k].Instructions[eachKey].Status
					consentStruct.InstructionArray = append(consentStruct.InstructionArray,allConsent[k].Instructions[eachKey].Description)
				}

				instructionKeySlice =instructionKeySlice[:0]
				consentRecievedDetails = append(consentRecievedDetails,consentStruct)
				innerTableData = append(innerTableData, consentRecievedDetails)
				consentViewModel.InnerContent = innerTableData
			}
		}
		consentViewModel.Keys = keySlice
		consentViewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["array"] = consentViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/consentreceipt-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

func (c *ConsentReceiptController) EditConsentReceipt() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	consentData := models.ConsentReceipts{}
	consentId := c.Ctx.Input.Param(":consentId")
	consentView :=viewmodels.EditConsentReceipt{}
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		members := models.ConsentMembers{}
		consentData.Info.ReceiptName = c.GetString("recieptName")
		consentData.Info.CompanyTeamName = companyTeamName
		tempGroupId := c.GetStrings("selectedUserIds")
		tempGroupMembers := c.GetStrings("selectedUserNames")
		instructions := c.GetString("instructionsForUser")
		instructionSlice := strings.Split(instructions, ",")
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		tempMembersMap := make(map[string]models.ConsentMembers)
		for i := 0; i < len(tempGroupId); i++ {
			members.MemberName = tempGroupMembers[i]
			tempMembersMap[tempGroupId[i]] = members
		}
		consentData.Members = tempMembersMap
		dbStatus := consentData.UpdateConsentDetails(c.AppEngineCtx,consentId,instructionSlice)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}else {
		groupUser := models.Users{}
		var keySlice []string
		var allUserNames [] string
		var instructionSlice []string
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
			consentDetails :=models.GetSelectedUsersName(c.AppEngineCtx,consentId)
			consentDataValues := reflect.ValueOf(consentDetails.Members)
			for _,UsersKey := range consentDataValues.MapKeys(){
				consentView.UserNameToEdit = append(consentView.UserNameToEdit,UsersKey.String())
			}
			instructionArrayDataValue := reflect.ValueOf(consentDetails.Instructions)

			for _,instructionKey := range instructionArrayDataValue.MapKeys(){
				instructionSlice = append(instructionSlice,instructionKey.String())
			}
			for _, eachInstructionKey := range instructionSlice {
				consentView.InstructionArrayToEdit = append(consentView.InstructionArrayToEdit,consentDetails.Instructions[eachInstructionKey].Description)
			}
			consentView.ReceiptName = consentDetails.Info.ReceiptName
			consentView.ConsentId  = consentId
			consentView.CompanyTeamName = storedSession.CompanyTeamName
			consentView.CompanyPlan   =  storedSession.CompanyPlan
			consentView.AdminLastName =storedSession.AdminLastName
			consentView.AdminFirstName =storedSession.AdminFirstName
			consentView.ProfilePicture =storedSession.ProfilePicture
			consentView.PageType=helpers.SelectPageForEdit
		case false:
			log.Println(helpers.ServerConnectionError)
		}

	}

	c.Data["vm"] = consentView
	c.Layout = "layout/layout.html"
	c.TplName = "template/add-consentreceipt.html"
}


func (c *ConsentReceiptController) DeleteConsentReceipt() {
	log.Println("haii iam there")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	consentId :=c.Ctx.Input.Param(":consentId")
	dbStatus :=models.DeleteConsentRecieptById(c.AppEngineCtx, consentId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}