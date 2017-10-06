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
	var keySlice []string
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	leave := models.LeaveRequests{}
	viewModel := viewmodels.LeaveViewModel{}
	dbStatus, leaveDetail := leave.GetAllLeaveRequest(c.AppEngineCtx)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(leaveDetail)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
	case false :
		log.Println(helpers.ServerConnectionError)
	}
	var keyForLeave []string
	for _, specifiedUserId := range keySlice {
		status, leaveDetailOfUser,_,_ := leave.GetAllLeaveRequestById(c.AppEngineCtx, specifiedUserId,companyId)
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
						companyUserLeaves := reflect.ValueOf(leaveDetailOfUser[key.String()].Company)
						for _, leaveKey := range companyUserLeaves.MapKeys() {
							if leaveKey.String() == companyId {
								tempValueSlice = append(tempValueSlice, leaveDetailOfUser[key.String()].Info.UserName+ "" + "(" + leaveDetailOfUser[key.String()].Company[leaveKey.String()].UserType + ")")
							}
						}
						startDate := strconv.FormatInt(int64(leaveDetailOfUser[key.String()].Info.StartDate), 10)
						tempValueSlice = append(tempValueSlice, startDate)
						endDate := strconv.FormatInt(int64(leaveDetailOfUser[key.String()].Info.EndDate), 10)
						tempValueSlice = append(tempValueSlice, endDate)
						numberOfDays := strconv.FormatInt(leaveDetailOfUser[key.String()].Info.NumberOfDays, 10)
						tempValueSlice = append(tempValueSlice, numberOfDays)
						tempValueSlice = append(tempValueSlice, leaveDetailOfUser[key.String()].Info.Reason)
								companyUserLeave := reflect.ValueOf(leaveDetailOfUser[key.String()].Company)
								for _, leaveKey := range companyUserLeave.MapKeys() {
									if leaveKey.String() == companyId {
										tempValueSlice = append(tempValueSlice, leaveDetailOfUser[key.String()].Company[leaveKey.String()].Status)
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
	log.Println("leave details",viewModel)
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

