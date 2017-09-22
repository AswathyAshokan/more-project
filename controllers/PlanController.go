package controllers

import (

	"app/passporte/viewmodels"

	"app/passporte/models"
	"encoding/json"
	"log"


)

type PlanController struct {
	BaseController
}

//to Display Plan Details
func (c *PlanController) PlanDetails() {
	planViewModel := viewmodels.Plan{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	sessionValues, sessionStatus := SessionForPlan(w,r)
	planViewModel.SessionFlag = sessionStatus
	planViewModel.CompanyPlan = sessionValues.CompanyPlan
	planViewModel.CompanyTeamName =sessionValues.CompanyTeamName
	//dbStatus,notificationValue := models.GetAllNotifications(c.AppEngineCtx,companyTeamName)
	//var notificationCount=0
	//switch dbStatus {
	//case true:
	//
	//	notificationOfUser := reflect.ValueOf(notificationValue)
	//	for _, notificationUserKey := range notificationOfUser.MapKeys() {
	//		dbStatus,notificationUserValue := models.GetAllNotificationsOfUser(c.AppEngineCtx,companyTeamName,notificationUserKey.String())
	//		switch dbStatus {
	//		case true:
	//			notificationOfUserForSpecific := reflect.ValueOf(notificationUserValue)
	//			for _, notificationUserKeyForSpecific := range notificationOfUserForSpecific.MapKeys() {
	//				var NotificationArray []string
	//				if notificationUserValue[notificationUserKeyForSpecific.String()].IsRead ==false{
	//					notificationCount=notificationCount+1;
	//				}
	//				NotificationArray =append(NotificationArray,notificationUserKey.String())
	//				NotificationArray =append(NotificationArray,notificationUserKeyForSpecific.String())
	//				NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].UserName)
	//				NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].Message)
	//				NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskLocation)
	//				NotificationArray =append(NotificationArray,notificationUserValue[notificationUserKeyForSpecific.String()].TaskName)
	//				date := strconv.FormatInt(notificationUserValue[notificationUserKeyForSpecific.String()].Date, 10)
	//				NotificationArray =append(NotificationArray,date)
	//				planViewModel.NotificationArray=append(planViewModel.NotificationArray,NotificationArray)
	//
	//			}
	//		case false:
	//		}
	//	}
	//case false:
	//}
	//planViewModel.NotificationNumber=notificationCount
	c.Data["vm"] = planViewModel
	c.TplName = "template/plan.html"
}

//For update company plan with newly selected company plan
func (c *PlanController) PlanUpdate() {
	log.Println("haiiiii")
	w := c.Ctx.ResponseWriter
	r := c.Ctx.Request
	storedSession,_ := SessionForPlan(w,r)
	companyId := storedSession.CompanyId
	selectedCompanyPlan := c.GetString("companyPlan")
	company := models.Company{}
	company.Plan = selectedCompanyPlan
	dbStatus, _ := company.ChangeCompanyPlan(c.AppEngineCtx,companyId)
	switch dbStatus {
	case true :

		ClearSession(w)
		sessionValues := SessionValues{}
		sessionValues.AdminId = storedSession.AdminId
		sessionValues.AdminFirstName = storedSession.AdminFirstName
		sessionValues.AdminLastName = storedSession.AdminLastName
		sessionValues.AdminEmail = storedSession.AdminEmail
		sessionValues.CompanyId = storedSession.CompanyId
		sessionValues.CompanyName = storedSession.CompanyName
		sessionValues.CompanyTeamName = storedSession.CompanyTeamName
		sessionValues.CompanyPlan = selectedCompanyPlan

		SetSession(w, sessionValues)
		slices := []interface{}{"true", sessionValues.CompanyTeamName,sessionValues.CompanyPlan}
		sliceToClient, _ := json.Marshal(slices)
		w.Write(sliceToClient)

	case false:
		w.Write([]byte("false"))


	}



}


