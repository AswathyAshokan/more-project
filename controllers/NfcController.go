/*Author: Sarath
Date:01/02/2017*/

package controllers

import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"time"
	"reflect"
	"strconv"
)

type NfcController struct {
	BaseController
}


//Display NFC Details
func (c *NfcController) NFCDetails(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	viewModel := viewmodels.NfcViewModel{}
	nfcDetails := models.NFC{}
	allNfcDetails := nfcDetails.GetAllNFCDetails(c.AppEngineCtx, storedSession.CompanyTeamName)
	dataValue := reflect.ValueOf(allNfcDetails)

	var keySlice []string
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}

	for _, k := range keySlice {
		var tempValueSlice []string
		tempValueSlice = append(tempValueSlice, allNfcDetails[k].Info.CustomerName)
		tempValueSlice = append(tempValueSlice, allNfcDetails[k].Info.Site)
		tempValueSlice = append(tempValueSlice, allNfcDetails[k].Info.Location)
		tempValueSlice = append(tempValueSlice, allNfcDetails[k].Info.NFCNumber)
		viewModel.Values = append(viewModel.Values, tempValueSlice)
		tempValueSlice = tempValueSlice[:0]
	}

	viewModel.Keys	= keySlice
	viewModel.CompanyTeamName = storedSession.CompanyTeamName
	viewModel.CompanyPlan = storedSession.CompanyPlan
	viewModel.AdminFirstName = storedSession.AdminFirstName
	viewModel.AdminLastName = storedSession.AdminLastName
	viewModel.ProfilePicture =storedSession.ProfilePicture
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
					viewModel.NotificationArray=append(viewModel.NotificationArray,NotificationArray)

				}
			case false:
			}
		}
	case false:
	}
	viewModel.NotificationNumber=notificationCount
	c.Data["vm"] = viewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/nfc-details.html"
}

//Add new NFC Tag
func (c *NfcController)AddNFC(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	viewModel := viewmodels.NfcViewModel{}
	if r.Method=="POST" {
		w := c.Ctx.ResponseWriter
		nfc := models.NFC{}
		nfc.Info.CompanyTeamName = storedSession.CompanyTeamName
		nfc.Info.CustomerName = c.GetString("customerName")
		nfc.Info.Site = c.GetString("site")
		nfc.Info.Location = c.GetString("location")
		nfc.Info.NFCNumber = c.GetString("nfcNumber")
		nfc.Settings.Status  = helpers.StatusActive
		nfc.Settings.DateOfCreation = time.Now().Unix()
		dbStatus := nfc.AddNFC(c.AppEngineCtx)
		switch dbStatus{
		case false:
			w.Write([]byte("false"))
		case true:
			w.Write([]byte("true"))
		}
	}else{
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
						viewModel.NotificationArray=append(viewModel.NotificationArray,NotificationArray)

					}
				case false:
				}
			}
		case false:
		}
		viewModel.NotificationNumber=notificationCount
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.AdminFirstName = storedSession.AdminFirstName
		viewModel.AdminLastName = storedSession.AdminLastName
		viewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-nfc.html"
	}
}

//Edit NFC Tag
func (c *NfcController)EditNFC(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method =="POST"{
		nfcId := c.Ctx.Input.Param(":nfcId")
		nfc := models.NFC{}
		nfc.Info.CompanyTeamName = storedSession.CompanyTeamName
		nfc.Info.CustomerName = c.GetString("customerName")
		nfc.Info.Site = c.GetString("site")
		nfc.Info.Location = c.GetString("location")
		nfc.Info.NFCNumber = c.GetString("nfcNumber")
		NfcUpdateStatus := nfc.UpdateNFCDetails(c.AppEngineCtx, nfcId)
		switch NfcUpdateStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}else {
		nfcId := c.Ctx.Input.Param(":nfcId")
		viewModel := viewmodels.EditNfcViewModel{}
		nfcDetails := models.NFC{}
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
						viewModel.NotificationArray=append(viewModel.NotificationArray,NotificationArray)

					}
				case false:
				}
			}
		case false:
		}
		viewModel.NotificationNumber=notificationCount
		editStatus, nfcDetails := nfcDetails.GetNFCDetailsById(c.AppEngineCtx, nfcId)
		switch editStatus{
		case true:
			viewModel.PageType = helpers.SelectPageForEdit
			viewModel.NfcId = nfcId
			viewModel.CustomerName = nfcDetails.Info.CustomerName
			viewModel.Location = nfcDetails.Info.Location
			viewModel.NFCNumber = nfcDetails.Info.NFCNumber
			viewModel.Site = nfcDetails.Info.Site
			viewModel.CompanyTeamName = storedSession.CompanyTeamName
			viewModel.CompanyPlan = storedSession.CompanyPlan
			viewModel.AdminFirstName = storedSession.AdminFirstName
			viewModel.AdminLastName = storedSession.AdminLastName
			viewModel.ProfilePicture =storedSession.ProfilePicture
			c.Data["vm"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-nfc.html"
		case false:

		}
	}
}

/*func (c *NfcController)Datatable() {
	log.Println("hiiiii")
	w := c.Ctx.ResponseWriter
	nfcDetails := models.NFC{}
	data := nfcDetails.GetNFCDetails(c.AppEngineCtx)
	log.Println(data)
	var valueSlice []models.NFC
	var keySlice []string

	for key, value := range data {
		valueSlice = append(valueSlice, value)
		keySlice = append(keySlice, key)
	}
	log.Println("KeySlice:", keySlice)
	log.Println("ValueSlice:", valueSlice)
	jsonObject,_ := json.Marshal(valueSlice)
	//c.Ctx.ResponseWriter.Write(jsonObject)

	*//*viewModel := ViewModel{}
	viewModel.Values = valueSlice
	viewModel.Keys = keySlice
	log.Println("viewModel:",viewModel)
	return viewModel*//*
	//c.Data["vm"] = viewModel
	log.Println("json Object:",jsonObject)
	w.Write([]byte(jsonObject))

}*/

//Delete NFC Tag
func (c *NfcController)DeleteNFC(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	key := c.GetString("Key")
	deleteStatus := models.DeleteNFC(c.AppEngineCtx, key)
	if deleteStatus == false {
		w.Write([]byte("false"))
	}else{
		w.Write([]byte("true"))
	}
}