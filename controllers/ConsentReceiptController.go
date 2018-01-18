package controllers
import (
	"time"
	"app/passporte/models"
	"app/passporte/helpers"
	"app/passporte/viewmodels"
	"reflect"
	"log"
	"strings"

	"regexp"
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
	userName :=models.ConsentUser{}
	/*groupNameAndDetails := models.ConsentGroup{}
	groupMemberNameForConsent :=models.GroupMemberNameInConsent{}
	groupMemberMap := make(map[string]models.GroupMemberNameInConsent)
	groupMap := make(map[string]models.ConsentGroup)*/
	userMap := make(map[string]models.ConsentUser)
	/*var keySliceForGroup [] string
	var MemberNameArray [] string*/
	//var keySliceForGroup [] string
	//var MemberNameArray [] string
	if r.Method == "POST" {
		//members := models.ConsentMembers{}
		consentData.Info.ReceiptName = c.GetString("recieptName")
		consentData.Info.CompanyName = storedSession.CompanyName
		groupKeySliceForConsentString := c.GetString("groupArrayElement")
		log.Println("groupKeySliceForConsentString",groupKeySliceForConsentString)
		groupKeySliceForConsent :=strings.Split(groupKeySliceForConsentString, ",")
		log.Println("groupKeySliceForConsent",groupKeySliceForConsent)
		UserOrGroupNameArray :=c.GetStrings("userAndGroupName")
		log.Println("UserOrGroupNameArray",UserOrGroupNameArray)
		userIdArray := c.GetStrings("selectedUserNames")
		log.Println("userIdArray",userIdArray)
		instructions := c.GetString("instructionsForUser")
		log.Println("insrucionssssssss",instructions)
		instructionsFromUser := strings.Split(instructions, "/@@,")
		sliceLastValue := instructionsFromUser[len(instructionsFromUser)-1]
		SliceLastValuesWithOutAnySymbol := strings.Split(sliceLastValue, "/@@")
		instructionsFromUser = instructionsFromUser[:len(instructionsFromUser)-1]
		instructionsFromUser = append(instructionsFromUser, SliceLastValuesWithOutAnySymbol[0])
		instructionSlice := instructionsFromUser
		log.Println("instructionSlice",instructionSlice)
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		/*group := models.Group{}
		var groupKeySlice	[]string*/
		for j:=0;j<len(userIdArray);j++ {
			tempName := UserOrGroupNameArray[j]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if (userOrGroupSelection[1]) == "User" {
				tempName = tempName[:len(tempName) - 7]
				userName.FullName = tempName
				userName.UserResponse = helpers.StatusActive
				log.Println("tempId", userIdArray[j])
				userMap[userIdArray[j]] = userName
			}


		}
		//log.Println("group map",groupMap)
		consentData.Instructions.User= userMap
		dbStatus := consentData.AddConsentToDb(c.AppEngineCtx, instructionSlice, companyTeamName)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {
		var keySliceForGroupAndUser []string
		companyUsers :=models.Company{}
		usersDetail := models.Users{}
		dbStatus, testUser := companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx, companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(testUser)
			for _, key := range dataValue.MapKeys() {

				dataValue := reflect.ValueOf(testUser[key.String()].Users)
				for _, userKeys := range dataValue.MapKeys() {
					//viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKey.String()].FullName+" (User)")
					dbStatus := usersDetail.GetActiveUsersEmailForDropDown(c.AppEngineCtx, userKeys.String(), testUser[key.String()].Users[userKeys.String()].Email, companyTeamName)
					switch dbStatus {
					case true:
						consentView.GroupNameArray = append(consentView.GroupNameArray, testUser[key.String()].Users[userKeys.String()].FullName + " (User)")
						keySliceForGroupAndUser = append(keySliceForGroupAndUser, userKeys.String())
					case false:
						log.Println(helpers.ServerConnectionError)
					}
				}
			}

		case false:
			log.Println(helpers.ServerConnectionError)
		}
		allGroups, dbStatusOfGroup := models.GetAllGroupDetails(c.AppEngineCtx, companyTeamName)
		switch dbStatusOfGroup {
		case true:
			dataValueOfGroup := reflect.ValueOf(allGroups)
			for _, key := range dataValueOfGroup.MapKeys() {
				if allGroups[key.String()].Settings.Status == "Active" {
					var memberSlice []string

					keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
					consentView.GroupNameArray = append(consentView.GroupNameArray, allGroups[key.String()].Info.GroupName + " (Group)")

					// For selecting members while selecting a group in dropdown
					memberSlice = append(memberSlice, key.String())
					groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
					for _, memberKey := range groupDataValue.MapKeys()  {
						if allGroups[key.String()].Members[memberKey.String()].Status != helpers.UserStatusDeleted{
							memberSlice = append(memberSlice, memberKey.String())
						}
					}
					consentView.GroupMembers = append(consentView.GroupMembers, memberSlice)
					log.Println("iam in trouble", consentView.GroupMembers)

				}
				log.Println("o111")

			}
			log.Println("o12")

		case false:
			log.Println(helpers.ServerConnectionError)
		}

		consentView.UserAndGroupKeyForConsent = keySliceForGroupAndUser
		log.Println("all keys of user and group",keySliceForGroupAndUser)
		log.Println("all group details",consentView.GroupMembers)
		consentView.CompanyTeamName = storedSession.CompanyTeamName
		consentView.CompanyPlan = storedSession.CompanyPlan
		consentView.AdminLastName = storedSession.AdminLastName
		consentView.AdminFirstName = storedSession.AdminFirstName
		consentView.ProfilePicture = storedSession.ProfilePicture
		log.Println("endddddddd")
		c.Data["vm"] = consentView
		log.Println("all view",consentView)
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
							users := getInstructions[instructionKeyString].User
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
	userName :=models.ConsentUser{}
	storedSession := ReadSession(w, r, companyTeamName)
	userMap := make(map[string]models.ConsentUser)
	consentView :=viewmodels.EditConsentReceipt{}
	consentData := models.ConsentReceipts{}
	if r.Method == "POST" {
		consentData.Info.ReceiptName = c.GetString("recieptName")
		consentData.Info.CompanyName = storedSession.CompanyName
		UserOrGroupNameArray :=c.GetStrings("userAndGroupName")
		userIdArray := c.GetStrings("selectedUserNames")
		log.Println("userIdArray",userIdArray)
		instructions := c.GetString("instructionsForUser")
		instructionsFromUser := strings.Split(instructions, "/@@,")
		sliceLastValue := instructionsFromUser[len(instructionsFromUser)-1]
		SliceLastValuesWithOutAnySymbol := strings.Split(sliceLastValue, "/@@")
		instructionsFromUser = instructionsFromUser[:len(instructionsFromUser)-1]
		instructionsFromUser = append(instructionsFromUser, SliceLastValuesWithOutAnySymbol[0])
		instructionSlice := instructionsFromUser
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		for j:=0;j<len(userIdArray);j++ {
			tempName := UserOrGroupNameArray[j]
			//userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			//userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			//if (userOrGroupSelection[1]) == "User" {
			tempName = tempName[:len(tempName) - 7]
			userName.FullName = tempName
			userName.UserResponse = helpers.StatusActive
			userMap[userIdArray[j]] = userName
			//}

		}

		consentData.Instructions.User= userMap
		log.Println("user map from controller to models",consentData.Instructions.User)
		/*instructionStatus := models.IsInstructionEdited(c.AppEngineCtx,instructionSlice,consentId,companyTeamName)
		switch instructionStatus {
		case true:*/
		log.Println("iam in true condition")
		dbStatus := consentData.UpdateConsentDataIfInstructionNotChanged(c.AppEngineCtx,consentId,instructionSlice,companyTeamName)
		switch dbStatus {

		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

		log.Println("true nnn")
		/*case false:
			log.Println("iam in false condition")
			dbStatus := consentData.UpdateConsentDetailsIfInstructionChanged(c.AppEngineCtx,consentId,instructionSlice,companyTeamName)
			switch dbStatus {
			case true:
				w.Write([]byte("true"))
			case false:
				w.Write([]byte("false"))
			}
		}*/
	}else {
		var Instructions []string
		companyUsers :=models.Company{}
		var keySliceForGroupAndUser []string
		usersDetail :=models.Users{}
		dbStatus ,testUser:= companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx,companyTeamName)
		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(testUser)
			for _, key := range dataValue.MapKeys() {

				dataValue := reflect.ValueOf(testUser[key.String()].Users)
				for _, userKeys := range dataValue.MapKeys() {
					//viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKey.String()].FullName+" (User)")
					dbStatus := usersDetail.GetActiveUsersEmailForDropDown(c.AppEngineCtx, userKeys.String(), testUser[key.String()].Users[userKeys.String()].Email, companyTeamName)
					switch dbStatus {
					case true:
						consentView.GroupNameArray = append(consentView.GroupNameArray, testUser[key.String()].Users[userKeys.String()].FullName + " (User)")
						keySliceForGroupAndUser = append(keySliceForGroupAndUser, userKeys.String())
					case false:
						log.Println(helpers.ServerConnectionError)
					}
				}
			}

		case false:
			log.Println(helpers.ServerConnectionError)
		}
		allGroups, dbStatusOfGroup := models.GetAllGroupDetails(c.AppEngineCtx,companyTeamName)
		switch dbStatusOfGroup {
		case true:
			dataValueOfGroup := reflect.ValueOf(allGroups)
			for _, key := range dataValueOfGroup.MapKeys() {
				if allGroups[key.String()].Settings.Status =="Active"{
					var memberSlice []string

					keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
					consentView.GroupNameArray = append(consentView.GroupNameArray, allGroups[key.String()].Info.GroupName+" (Group)")

					// For selecting members while selecting a group in dropdown
					memberSlice = append(memberSlice, key.String())
					groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
					for _, memberKey := range groupDataValue.MapKeys()  {
						if allGroups[key.String()].Members[memberKey.String()].Status != helpers.UserStatusDeleted{
							memberSlice = append(memberSlice, memberKey.String())
						}
					}
					consentView.GroupMembers = append(consentView.GroupMembers, memberSlice)
				}
			}
			consentView.UserAndGroupKey=keySliceForGroupAndUser
		case false:
			log.Println(helpers.ServerConnectionError)
		}
		consentDetails := models.GetEachConsentByCompanyId(c.AppEngineCtx, consentId, companyTeamName)
		allInstructions := models.GetAllInstructionsFromConsent(c.AppEngineCtx, consentId, companyTeamName)
		dataValueOfInstruction := reflect.ValueOf(allInstructions)
		for _, instructionKey := range dataValueOfInstruction.MapKeys() {
			var UserName []string
			var selectedUserKey []string
			Instructions = append(Instructions, allInstructions[instructionKey.String()].Description)
			userDataValue := reflect.ValueOf(allInstructions[instructionKey.String()].User)
			for _, userKey := range userDataValue.MapKeys() {
				UserName = append(UserName, allInstructions[instructionKey.String()].User[userKey.String()].FullName)
				selectedUserKey = append(selectedUserKey, userKey.String())
			}
			consentView.SelectedUsersKey = selectedUserKey
			consentView.UsersKey = selectedUserKey
			log.Println("seleceted keyss",consentView.SelectedUsersKey)
			consentView.UserNameToEdit = UserName
		}
		consentView.UserAndGroupKeyForConsent = keySliceForGroupAndUser
		consentView.InstructionArrayToEdit = Instructions
		consentView.ReceiptName = consentDetails.Info.ReceiptName
		consentView.ConsentId = consentId
		consentView.CompanyTeamName = storedSession.CompanyTeamName
		consentView.CompanyPlan = storedSession.CompanyPlan
		consentView.AdminLastName = storedSession.AdminLastName
		consentView.AdminFirstName = storedSession.AdminFirstName
		consentView.ProfilePicture = storedSession.ProfilePicture
		consentView.PageType = helpers.SelectPageForEdit

	}
	c.Data["vm"] = consentView
	c.Layout = "layout/layout.html"
	c.TplName = "template/add-consentreceipt.html"

}




















