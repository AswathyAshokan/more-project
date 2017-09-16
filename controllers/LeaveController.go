package controllers
import (
	"app/passporte/models"
	"reflect"
	"strconv"
	"app/passporte/helpers"
	"app/passporte/viewmodels"
	"log"

)
type LeaveController struct {
	BaseController
}

func (c *LeaveController) LoadUserLeave() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	//var keySliceForUser []string
	var keySlice []string
	//var commonKey []string
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	//companyUsersForLeave := models.Company{}
	leave := models.LeaveRequests{}
	//dbStatus, companyUserDetail := companyUsersForLeave.GetUsersForDropdownFromCompany(c.AppEngineCtx, companyTeamName)
	viewModel := viewmodels.LeaveViewModel{}
	//switch dbStatus {
	//case true:
	//	dataValue := reflect.ValueOf(companyUserDetail)
	//	for _, key := range dataValue.MapKeys() {
	//		dataValue := reflect.ValueOf(companyUserDetail[key.String()].Users)
	//		for _, userKey := range dataValue.MapKeys() {
	//			keySliceForUser = append(keySliceForUser, userKey.String())
	//		}
	//	}
	//case false :
	//	log.Println(helpers.ServerConnectionError)
	//}
	dbStatus, leaveDetail := leave.GetAllLeaveRequest(c.AppEngineCtx)
	switch dbStatus {
	case true:
		log.Println("leve request",leaveDetail)
		dataValue := reflect.ValueOf(leaveDetail)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
	case false :
		log.Println(helpers.ServerConnectionError)
	}
	//compare two slice
	//for i := 0; i < 2; i++ {
	//	for _, sliceOfUser := range keySliceForUser {
	//		found := false
	//		for _, sliceOfLeaveRequest := range keySlice {
	//			if sliceOfUser == sliceOfLeaveRequest {
	//				found = true
	//				break
	//			}
	//		}
	//		// String not found. We add it to return slice
	//		if found {
	//			commonKey = append(commonKey, sliceOfUser)
	//		}
	//	}
	//	// Swap the slices, only if it was the first loop
	//	if i == 0 {
	//		keySliceForUser, keySlice = keySlice, keySliceForUser
	//	}
	//}
	////remove duplicate value of slice
	//duplicateKey := map[string]bool{}
	//
	//// Create a map of all unique elements.
	//for v := range commonKey {
	//	duplicateKey[commonKey[v]] = true
	//}

	// Place all keys from the map into a slice.
	//var userLeaveKey []string
	var keyForLeave []string
	//for key, _ := range duplicateKey {
	//	userLeaveKey = append(userLeaveKey, key)
	//}
	for _, specifiedUserId := range keySlice {
		status, leaveDetailOfUser,userDetail,userInvitation := leave.GetAllLeaveRequestById(c.AppEngineCtx, specifiedUserId,companyId)
		switch status {
		case true:
			log.Println("lvvvv",leaveDetailOfUser)
			dataValue := reflect.ValueOf(leaveDetailOfUser)
			for _, key := range dataValue.MapKeys() {
				companyLeaveData := reflect.ValueOf(leaveDetailOfUser[key.String()].Company)
				for _, CompanyKey := range companyLeaveData.MapKeys() {
					if CompanyKey.String() == companyId {
						keyForLeave = append(keyForLeave, key.String())
						var tempValueSlice []string
						inviteUser := reflect.ValueOf(userInvitation)
						for _, InviteKey := range inviteUser.MapKeys() {
							if userDetail.Email == userInvitation[InviteKey.String()].Email {
								tempValueSlice = append(tempValueSlice, userDetail.FullName + "" + "(" + userInvitation[InviteKey.String()].UserType + ")")
								break
							}
						}
						startDate := strconv.FormatInt(int64(leaveDetailOfUser[key.String()].Info.StartDate), 10)
						//startDate := time.Unix(leaveDetailOfUser[key.String()].Info.StartDate, 0)
						tempValueSlice = append(tempValueSlice, startDate)
						endDate := strconv.FormatInt(int64(leaveDetailOfUser[key.String()].Info.EndDate), 10)
						//endDate := time.Unix(leaveDetailOfUser[key.String()].Info.EndDate, 0)
						tempValueSlice = append(tempValueSlice, endDate)
						numberOfDays := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.NumberOfDays, 10)
						tempValueSlice = append(tempValueSlice, numberOfDays)
						tempValueSlice = append(tempValueSlice, leaveDetailOfUser[key.String()].Info.Reason)
						for _, InviteKeys := range inviteUser.MapKeys() {
							if userDetail.Email == userInvitation[InviteKeys.String()].Email {
								companyUserLeave := reflect.ValueOf(leaveDetailOfUser[key.String()].Company)
								for _, leaveKey := range companyUserLeave.MapKeys() {
									if leaveKey.String() == companyId {
										tempValueSlice = append(tempValueSlice, leaveDetailOfUser[key.String()].Company[leaveKey.String()].Status)
									}
								}
							}
						}
						tempValueSlice = append(tempValueSlice, key.String())
						tempValueSlice = append(tempValueSlice, specifiedUserId)
						viewModel.Values = append(viewModel.Values, tempValueSlice)
						tempValueSlice = tempValueSlice[:0]
					}
				}
			}
		case false :
			log.Println(helpers.ServerConnectionError)
		}
	}
	viewModel.AdminFirstName =storedSession.AdminFirstName
	viewModel.AdminLastName =storedSession.AdminLastName
	viewModel.CompanyPlan =storedSession.CompanyPlan
	viewModel.CompanyTeamName =storedSession.CompanyTeamName
	viewModel.ProfilePicture =storedSession.ProfilePicture
	viewModel.UserKeys =keySlice
	viewModel.Keys = keyForLeave
	c.Data["vm"] = viewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/leave-detail.html"
}

func (c *LeaveController) LoadAcceptUserLeave() {
	leaveKey := c.Ctx.Input.Param(":leaveKey")
	userKey := c.Ctx.Input.Param(":userKey")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r, companyTeamName)
	companyName :=storedSession.CompanyName
	leave := models.LeaveRequests{}

	status:= leave.AcceptLeaveRequestById(c.AppEngineCtx, leaveKey,userKey,companyTeamName,companyName)
	log.Println("sucess")
	switch status {
	case true:
		w.Write([]byte("true"))

	case false:
		w.Write([]byte("false"))

	}

}
func (c *LeaveController) LoadRejectUserLeave() {
	leaveKey := c.Ctx.Input.Param(":leaveKey")
	userKey := c.Ctx.Input.Param(":userKey")
	log.Println("leave key",leaveKey)
	log.Println("user key",userKey)
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r, companyTeamName)
	companyName :=storedSession.CompanyName
	leave := models.LeaveRequests{}
	status:= leave.RejectLeaveRequestById(c.AppEngineCtx, leaveKey,userKey,companyTeamName,companyName)
	log.Println("sucess")
	switch status {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))

	}

}

