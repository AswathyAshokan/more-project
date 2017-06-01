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
		consentData.Info.CompanyName = storedSession.CompanyName
		tempUserId := c.GetStrings("selectedUserIds")
		tempMembers := c.GetStrings("selectedUserNames")
		instructions := c.GetString("instructionsForUser")
		instructionSlice := strings.Split(instructions, ",")
		consentData.Settings.DateOfCreation = (time.Now().UnixNano() / 1000000)
		consentData.Settings.Status = helpers.StatusActive
		tempMembersMap := make(map[string]models.ConsentMembers)
		for i := 0; i < len(tempUserId); i++ {
			members.UserResponse = helpers.UserResponsePending
			members.FullName = tempMembers[i]
			tempMembersMap[tempUserId[i]] = members
		}
		consentData.Instructions.Users = tempMembersMap
		dbStatus := consentData.AddConsentToDb(c.AppEngineCtx,instructionSlice,companyTeamName,tempUserId)
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
	dbStatus,allConsent:= models.GetAllConsentReceiptDetails(c.AppEngineCtx)
	consentViewModel :=viewmodels.LoadConsent{}
	var innerTableData [][]viewmodels.ConsentStruct
	switch dbStatus {
	case true:
		var keySlice []string
		var tempKeySlice []string
		var instructionKeySlice []string
		var userKeySlice []string
		//var instructionKeySlice []string
		//var memberKeySlice []string
		dataValue := reflect.ValueOf(allConsent)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k :=range keySlice{
			consentById :=models.GetSelectedUsersName(c.AppEngineCtx,k)
			consentDataValues :=reflect.ValueOf(consentById)
			for _, consentKey := range consentDataValues.MapKeys() {
				tempKeySlice = append(tempKeySlice, consentKey.String())
			}
			for _, eachKey :=range tempKeySlice{
				log.Println("key",eachKey)
				var tempValueSlice []string
				if consentById[eachKey].Settings.Status!= helpers.UserStatusDeleted {
					tempValueSlice = append(tempValueSlice, "")
					tempValueSlice = append(tempValueSlice, consentById[eachKey].Info.ReceiptName)
					tempValueSlice = append(tempValueSlice,eachKey)
					consentViewModel.Values = append(consentViewModel.Values, tempValueSlice)
					tempValueSlice = tempValueSlice[:0]
					var consentReceivedDetails []viewmodels.ConsentStruct
					var consentStruct viewmodels.ConsentStruct
					getInstructions := models.GetAllInstructionsById(c.AppEngineCtx,k,eachKey)
					instructionDataValues := reflect.ValueOf(getInstructions)
					for _, instructionKey := range instructionDataValues.MapKeys() {
						instructionKeySlice = append(instructionKeySlice, instructionKey.String())
					}
					for _, eachInstructionId :=range instructionKeySlice{
						log.Println("athyaavisham ulla id",eachInstructionId)
						consentStruct.InstructionArray = append(consentStruct.InstructionArray,getInstructions[eachInstructionId].Description)
						getUsers := models.GetAllUsersNameAndStatus(c.AppEngineCtx,k,eachKey,eachInstructionId)
						usersDataValues := reflect.ValueOf(getUsers)
						for _, usersKey := range usersDataValues.MapKeys() {
							userKeySlice = append(userKeySlice, usersKey.String())
						}
						for _, eachUsersId :=range userKeySlice{
							if getUsers[eachUsersId].UserResponse == helpers.UserResponseAccepted{
								consentStruct.AcceptedUsers =append(consentStruct.AcceptedUsers,getUsers[eachUsersId].FullName)
							} else if getUsers[eachUsersId].UserResponse == helpers.UserResponseRejected{
								consentStruct.RejectedUsers = append(consentStruct.RejectedUsers,getUsers[eachUsersId].FullName)
							} else {
								consentStruct.PendingUsers = append(consentStruct.PendingUsers,getUsers[eachUsersId].FullName)
							}
							eachUsersId = eachUsersId[:0]
						}

						userKeySlice = userKeySlice[:0]
					}
					instructionKeySlice =instructionKeySlice[:0]
					consentStruct.UserKey = eachKey
					consentReceivedDetails = append(consentReceivedDetails,consentStruct)
					innerTableData = append(innerTableData, consentReceivedDetails)
					consentViewModel.InnerContent = innerTableData
				}

			}

			/*var tempValueSlice []string
			if allConsent[k].Settings.Status != helpers.UserStatusDeleted {
				tempValueSlice = append(tempValueSlice, "")
				tempValueSlice = append(tempValueSlice, allConsent[k].Info.ReceiptName)
				tempValueSlice = append(tempValueSlice,k)
				consentViewModel.Values = append(consentViewModel.Values, tempValueSlice)
				log.Println("tempValeus",allConsent[k].Info.ReceiptName)
				tempValueSlice = tempValueSlice[:0]
				var consentReceivedDetails []viewmodels.ConsentStruct
				var consentStruct viewmodels.ConsentStruct
				memberDataValue :=reflect.ValueOf(allConsent[k].Instructions.Users)
				// for get name of members from members map
				for _, membersKey := range memberDataValue.MapKeys() {
					memberKeySlice = append(memberKeySlice, membersKey.String())
				}
				for _, eachMemberKey := range memberKeySlice {
					if allConsent[k].Instructions.Users[eachMemberKey].UserResponse ==helpers.StatusAccepted{
						consentStruct.AcceptedUsers = append(consentStruct.AcceptedUsers,allConsent[k].Instructions.Users[eachMemberKey].FullName)
					} else  if allConsent[k].Instructions.Users[eachMemberKey].UserResponse ==helpers.UserResponseRejected{
						consentStruct.RejectedUsers = append(consentStruct.RejectedUsers,allConsent[k].Instructions.Users[eachMemberKey].FullName)
					}
				}
				memberKeySlice = memberKeySlice[:0]
				consentStruct.UserKey = k
				instructionDataValue := reflect.ValueOf(allConsent[k].Instructions)
				//for get each instructions  from instruction map
				for _, consentKey := range instructionDataValue.MapKeys() {
					instructionKeySlice = append(instructionKeySlice, consentKey.String())
				}
				*//*for _, eachKey := range instructionKeySlice {
					consentStruct.InstructionArray = append(consentStruct.InstructionArray, allConsent[k].Instructions.Description)

				}*//*
				instructionKeySlice =instructionKeySlice[:0]
				consentReceivedDetails = append(consentReceivedDetails,consentStruct)
				innerTableData = append(innerTableData, consentReceivedDetails)
				consentViewModel.InnerContent = innerTableData
			}*/
		}
		consentViewModel.Keys = keySlice
		log.Println("bdbdbd",consentViewModel)
		consentViewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["array"] = consentViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/consentreceipt-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

/*func (c *ConsentReceiptController) EditConsentReceipt() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	consentData := models.ConsentReceipts{}
	consentId := c.Ctx.Input.Param(":consentId")
	consentView :=viewmodels.EditConsentReceipt{}
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		var keySlice []string
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
		instructionStatus := models.IsInstructionEdited(c.AppEngineCtx,instructionSlice,consentId)
		switch instructionStatus {
		case true:
			filterData := models.GetMemberStatus(c.AppEngineCtx,consentId)
			dataValue := reflect.ValueOf(members)

			for _, key := range dataValue.MapKeys() {
				keySlice = append(keySlice, key.String())
			}
			for _,k := range keySlice {
				for i := 0; i < len(tempGroupId); i++ {
					if tempGroupMembers[i] == filterData[k].MemberName {
						log.Println("cp1")
						members.MemberName = tempGroupMembers[i]
						log.Println("status :",filterData[k].Status)
						members.Status = filterData[k].Status
						tempMembersMap[tempGroupId[i]] = members
					} else {
						log.Println("cp2")
						members.MemberName = tempGroupMembers[i]
						members.Status = helpers.UserResponsePending
						tempMembersMap[tempGroupId[i]] = members
					}
					consentData.Members = tempMembersMap
				}

			}

		case false:
			for i := 0; i < len(tempGroupId); i++ {
				members.MemberName = tempGroupMembers[i]
				members.Status = helpers.UserResponsePending
				tempMembersMap[tempGroupId[i]] = members
			}
			consentData.Members = tempMembersMap
		}
		dbStatus := consentData.UpdateConsentDetails(c.AppEngineCtx,consentId,instructionSlice,tempGroupId,tempGroupMembers)
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
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	consentId :=c.Ctx.Input.Param(":consentId")
	dbStatus :=models.DeleteConsentReceiptById(c.AppEngineCtx, consentId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}*/
