/*Author: Sarath
Date:01/02/2017*/

package controllers

import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
	"app/passporte/helpers"
	"time"
	"reflect"
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
	log.Println("The userDetails stored in session:",storedSession)
	nfcDetails := models.NFC{}
	allNfcDetails := nfcDetails.GetAllNFCDetails(c.AppEngineCtx)
	log.Println("NFC details:", allNfcDetails)

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

	log.Println("KeySlice", keySlice)
	viewModel.Keys	= keySlice
	viewModel.CompanyTeamName = storedSession.CompanyTeamName


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
		log.Println("NFC Details:", nfc)
		dbStatus := nfc.AddNFC(c.AppEngineCtx)
		switch dbStatus{
		case false:
			w.Write([]byte("false"))
		case true:
			w.Write([]byte("true"))
		}
	}else{
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["vm"] = viewModel
		log.Println("VIEW_MODEL:",viewModel)
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-nfc.html"
	}
}

//Edit NFC Tag
func (c *NfcController)EditNFC(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	log.Println("here: ", companyTeamName)
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method =="POST"{
		nfcId := c.Ctx.Input.Param(":nfcId")
		nfc := models.NFC{}
		nfc.Info.CustomerName = c.GetString("customerName")
		nfc.Info.Site = c.GetString("site")
		nfc.Info.Location = c.GetString("location")
		nfc.Info.NFCNumber = c.GetString("nfcNumber")
		NfcUpdateStatus := nfc.UpdateNFCDetails(c.AppEngineCtx, nfcId)
		switch NfcUpdateStatus {
		case true:
			log.Println("CASR TRUE:")
			w.Write([]byte("true"))
		case false:
			log.Println("CASEFALSE")
			w.Write([]byte("false"))
		}

	}else {
		nfcId := c.Ctx.Input.Param(":nfcId")
		viewModel := viewmodels.EditNfcViewModel{}
		nfcDetails := models.NFC{}
		log.Println("NFC Id: ", nfcId)
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

			c.Data["array"] = viewModel
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
	log.Println("Controller:DeleteNFC()")
	key := c.GetString("Key")
	deleteStatus := models.DeleteNFC(c.AppEngineCtx, key)
	if deleteStatus == false {
		w.Write([]byte("false"))
	}else{
		w.Write([]byte("true"))
	}
}