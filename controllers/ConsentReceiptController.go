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
	consentView := viewmodels.ConsentReceipt{}
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	consentData := models.ConsentReceipts{}
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		members := models.ConsentMembers{}
		consentData.Info.ReceiptName = c.GetString("recieptName")
		consentData.Info.CompanyName = storedSession.CompanyName
		tempUserId := c.GetStrings("selectedUserIds")
		log.Println("selected users key",tempUserId)
		tempMembers := c.GetStrings("selectedUserNames")
		log.Println("selecetd members name",tempMembers)
		instructions := c.GetString("instructionsForUser")
		instructionsFromUser := strings.Split(instructions, "/@@,")
		sliceLastValue := instructionsFromUser[len(instructionsFromUser)-1]
		SliceLastValuesWithOutAnySymbol := strings.Split(sliceLastValue, "/@@")
		instructionsFromUser = instructionsFromUser[:len(instructionsFromUser)-1]
		instructionsFromUser = append(instructionsFromUser, SliceLastValuesWithOutAnySymbol[0])
		instructionSlice := instructionsFromUser
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		tempMembersMap := make(map[string]models.ConsentMembers)
		for i := 0; i < len(tempUserId); i++ {
			members.UserResponse = helpers.UserResponsePending
			members.FullName = tempMembers[i]
			tempMembersMap[tempUserId[i]] = members
		}
		consentData.Instructions.Users = tempMembersMap
		dbStatus := consentData.AddConsentToDb(c.AppEngineCtx, instructionSlice, companyTeamName, tempUserId)
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

			//for _, groupKey := range dataValue.MapKeys() {
			//	keySlice = append(keySlice, groupKey.String())
			//}
			for _, groupKey := range dataValue.MapKeys() {
				if allUserDetails[groupKey.String()].Status != helpers.UserStatusDeleted {
					keySlice = append(keySlice, groupKey.String())
					allUserNames = append(allUserNames, allUserDetails[groupKey.String()].FullName)
					consentView.GroupMembers = allUserNames
					consentView.GroupKey = keySlice
				}
			}
			consentView.CompanyTeamName = storedSession.CompanyTeamName
			consentView.CompanyPlan = storedSession.CompanyPlan
			consentView.AdminLastName = storedSession.AdminLastName
			consentView.AdminFirstName = storedSession.AdminFirstName
			consentView.ProfilePicture = storedSession.ProfilePicture
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
	dbStatus,allConsent:= models.GetAllConsentReceiptDetails(c.AppEngineCtx)
	consentViewModel :=viewmodels.LoadConsent{}
	switch dbStatus {
	case true:
		var keySlice []string
		var tempKeySlice []string
		dataValue := reflect.ValueOf(allConsent)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k :=range keySlice {
			if k == companyTeamName {
				consentById := models.GetSelectedUsersName(c.AppEngineCtx, k)
				consentDataValues := reflect.ValueOf(consentById)
				for _, consentKey := range consentDataValues.MapKeys() {
					tempKeySlice = append(tempKeySlice, consentKey.String())
				}
				for _, eachKey := range tempKeySlice {
					log.Println("key", eachKey)
					var tempValueSlice []string

					if consentById[eachKey].Settings.Status != helpers.UserStatusDeleted {

						tempValueSlice = append(tempValueSlice, "")
						tempValueSlice = append(tempValueSlice, consentById[eachKey].Info.ReceiptName)
						tempValueSlice = append(tempValueSlice, eachKey)
						consentViewModel.Values = append(consentViewModel.Values, tempValueSlice)
						tempValueSlice = tempValueSlice[:0]

						getInstructions := models.GetAllInstructionsById(c.AppEngineCtx, k, eachKey)
						log.Println("getInstructions", getInstructions)
						for _, instructionKey := range reflect.ValueOf(getInstructions).MapKeys() {
							var consentStructVM viewmodels.ConsentStruct
							var instructionKeySlice []string
							instructionKeyString := instructionKey.String()
							consentStructVM.InstructionKey = eachKey
							instructionKeySlice = append(instructionKeySlice, instructionKeyString)
							consentStructVM.Description = getInstructions[instructionKeyString].Description
							users := getInstructions[instructionKeyString].Users
							for _, userKey := range reflect.ValueOf(users).MapKeys() {
								userKeyString := userKey.String()
								if users[userKeyString].UserResponse == helpers.UserResponseAccepted {
									consentStructVM.AcceptedUsers = append(consentStructVM.AcceptedUsers, users[userKeyString].FullName)
								} else if users[userKeyString].UserResponse == helpers.UserResponseRejected {
									consentStructVM.RejectedUsers = append(consentStructVM.RejectedUsers, users[userKeyString].FullName)
								} else {
									// Pending
									consentStructVM.PendingUsers = append(consentStructVM.PendingUsers, users[userKeyString].FullName)
								}
							}
							consentViewModel.InnerContent = append(consentViewModel.InnerContent, consentStructVM)
						}
					}

				}
			}
		}

		consentViewModel.Keys = keySlice
		consentViewModel.CompanyTeamName = storedSession.CompanyTeamName
		consentViewModel.CompanyPlan = storedSession.CompanyPlan
		consentViewModel.AdminFirstName = storedSession.AdminFirstName
		consentViewModel.AdminLastName = storedSession.AdminLastName
		consentViewModel.ProfilePicture =storedSession.ProfilePicture
		consentViewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["vm"] = consentViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/consentreceipt-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}


func (c *ConsentReceiptController) DeleteConsentReceipt() {
	log.Println("hhhooooooo")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	consentId :=c.Ctx.Input.Param(":consentId")
	dbStatus :=models.DeleteConsentReceiptById(c.AppEngineCtx, consentId,companyTeamName)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}

func (c *ConsentReceiptController) EditConsentReceipt() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	consentId := c.Ctx.Input.Param(":consentId")
	storedSession := ReadSession(w, r, companyTeamName)
	consentData := models.ConsentReceipts{}
	consentView :=viewmodels.EditConsentReceipt{}
	if r.Method == "POST" {
		members := models.ConsentMembers{}
		consentData.Info.ReceiptName = c.GetString("recieptName")
		consentData.Info.CompanyName = storedSession.CompanyName
		tempGroupId := c.GetStrings("selectedUserIds")
		tempGroupMembers := c.GetStrings("selectedUserNames")
		instructions := c.GetString("instructionsForUser")
		instructionsFromUser := strings.Split(instructions, "/@@,")
		sliceLastValue := instructionsFromUser[len(instructionsFromUser)-1]
		SliceLastValuesWithOutAnySymbol := strings.Split(sliceLastValue, "/@@")
		instructionsFromUser = instructionsFromUser[:len(instructionsFromUser)-1]
		instructionsFromUser = append(instructionsFromUser, SliceLastValuesWithOutAnySymbol[0])
		instructionSlice := instructionsFromUser
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		tempMembersMap := make(map[string]models.ConsentMembers)
		for i := 0; i < len(tempGroupId); i++ {
			members.FullName = tempGroupMembers[i]
			members.UserResponse = helpers.UserResponsePending
			tempMembersMap[tempGroupId[i]] = members
		}
		consentData.Instructions.Users = tempMembersMap
		instructionStatus := models.IsInstructionEdited(c.AppEngineCtx,instructionSlice,consentId,companyTeamName)
		switch instructionStatus {
		case true:
			dbStatus := consentData.UpdateConsentDataIfInstructionNotChanged(c.AppEngineCtx,consentId,instructionSlice,tempGroupId,tempGroupMembers,companyTeamName)
			switch dbStatus {
			case true:
				w.Write([]byte("true"))
			case false:
				w.Write([]byte("false"))
			}

			log.Println("true nnn")
		case false:
			dbStatus := consentData.UpdateConsentDetailsIfInstructionChanged(c.AppEngineCtx,consentId,instructionSlice,tempGroupId,tempGroupMembers,companyTeamName)
			log.Println("dbStatus",dbStatus)
			switch dbStatus {
			case true:
				w.Write([]byte("true"))
			case false:
				w.Write([]byte("false"))
			}
		}

	}else {

		var allUserNames [] string
		allUsers := models.Users{}
		allUserDetails, dbStatus := allUsers.TakeGroupMemberName(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			var Instructions []string

			var keySlice []string
			dataValue := reflect.ValueOf(allUserDetails)

			//for _, groupKey := range dataValue.MapKeys() {
			//	keySlice = append(keySlice, groupKey.String())
			//}
			for _, groupKey := range dataValue.MapKeys() {
				if allUserDetails[groupKey.String()].Status != helpers.UserStatusDeleted {
					keySlice = append(keySlice, groupKey.String())
					allUserNames = append(allUserNames, allUserDetails[groupKey.String()].FullName)
					consentView.GroupMembers = allUserNames
					consentView.GroupKey = keySlice
				}
			}
			consentDetails :=models.GetEachConsentByCompanyId(c.AppEngineCtx,consentId,companyTeamName)
			allInstructions := models.GetAllInstructionsFromConsent(c.AppEngineCtx,consentId,companyTeamName)
			dataValueOfInstruction := reflect.ValueOf(allInstructions)
			for _, instructionKey:=range dataValueOfInstruction.MapKeys(){
				var UserName []string
				var selectedUserKey []string
				Instructions = append(Instructions,allInstructions[instructionKey.String()].Description)
				userDataValue := reflect.ValueOf(allInstructions[instructionKey.String()].Users)
				for _, userKey := range userDataValue.MapKeys() {
					UserName = append(UserName, allInstructions[instructionKey.String()].Users[userKey.String()].FullName)
					selectedUserKey = append(selectedUserKey,userKey.String())
				}
				consentView.SelectedUsersKey = selectedUserKey
				consentView.UserNameToEdit = UserName
			}
			consentView.InstructionArrayToEdit = Instructions
			consentView.ReceiptName = consentDetails.Info.ReceiptName
			consentView.ConsentId  = consentId
			consentView.CompanyTeamName = storedSession.CompanyTeamName
			consentView.CompanyPlan   =  storedSession.CompanyPlan
			consentView.AdminLastName =storedSession.AdminLastName
			consentView.AdminFirstName =storedSession.AdminFirstName
			consentView.ProfilePicture =storedSession.ProfilePicture
			consentView.PageType=helpers.SelectPageForEdit
		case false:

		}
	}
	c.Data["vm"] = consentView
	c.Layout = "layout/layout.html"
	c.TplName = "template/add-consentreceipt.html"

}








