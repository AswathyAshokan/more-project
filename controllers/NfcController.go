/*Author: Sarath
Date:01/02/2017*/
package controllers

import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
	"app/passporte/helpers"
)

type NfcController struct {
	BaseController
}


//Display NFC Details
func (c *NfcController) NFCDetails(){
	nfcDetails := models.NFC{}
	data := nfcDetails.GetNFCDetails(c.AppEngineCtx)
	log.Println(data)
	var valueSlice []models.NFC
	var keySlice []string

	for key, value := range data {
		valueSlice = append(valueSlice, value)
		keySlice = append(keySlice, key)
	}
	log.Println("KeySlice", keySlice)
	log.Println("ValueSlice", valueSlice)

	viewModel := viewmodels.NfcViewModel{}
	viewModel.Values = valueSlice
	viewModel.Keys	= keySlice


	c.Data["vm"] = viewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/nfc-details.html"
}

//Add new NFC Tag
func (c *NfcController)AddNFC(){
	r := c.Ctx.Request
	if r.Method=="POST" {
		w := c.Ctx.ResponseWriter
		nfc := models.NFC{}
		nfc.CustomerName = c.GetString("customerName")
		nfc.Site = c.GetString("site")
		nfc.Location = c.GetString("location")
		nfc.NFCNumber = c.GetString("nfcNumber")
		log.Println("NFC Details:", nfc)
		dbStatus := nfc.AddNFC(c.AppEngineCtx)
		switch dbStatus{
		case false:
			w.Write([]byte("false"))
		case true:
			w.Write([]byte("true"))
		}
	}else{
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-nfc.html"
	}
}

//Edit NFC Tag
func (c *NfcController)EditNFC(){
	log.Println("EditNFC()")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method =="POST"{
		nfcId := c.Ctx.Input.Param(":nfcId")
		nfc := models.NFC{}
		nfc.CustomerName = c.GetString("customerName")
		nfc.Site = c.GetString("site")
		nfc.Location = c.GetString("location")
		nfc.NFCNumber = c.GetString("nfcNumber")
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
		log.Println("NFC Id: ", nfcId)
		editStatus, nfcDetails := nfcDetails.GetNFCDetailsById(c.AppEngineCtx, nfcId)
		switch editStatus{
		case true:
			viewModel.PageType = helpers.SelectPageForEdit
			viewModel.NfcId = nfcId
			viewModel.CustomerName = nfcDetails.CustomerName
			viewModel.Location = nfcDetails.Location
			viewModel.NFCNumber = nfcDetails.NFCNumber
			viewModel.Site = nfcDetails.Site

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
	w := c.Ctx.ResponseWriter
	log.Println("Controller:DeleteNFC()")
	key := c.GetString("Key")
	deleteStatus := models.DeleteNFC(c.AppEngineCtx, key)
	if deleteStatus == false {
		w.Write([]byte("false"))
	}else{
		w.Write([]byte("true"))
	}
}