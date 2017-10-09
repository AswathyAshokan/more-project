/*Created By Farsana*/
package controllers

import (
	"app/passporte/models"
	"time"
	"app/passporte/viewmodels"
	"log"
	"reflect"
	"app/passporte/helpers"
	"strings"
)

type CompanyFileUploadController struct {
	BaseController
}


// fetch all the details of invite user from database
func (c *CompanyFileUploadController) CompanyFileUpload() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	companyFile :=models.CompanyFileUpload{}
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	log.Println("companyId",companyId)
	companyViewModel := viewmodels.CompanyFileUpload{}
	if r.Method == "POST" {
		log.Println("insideeeeeeeeeee");
		companyFile.Info.FolderName = c.GetString("folderName")
		log.Println("folder name",companyFile.Info.FolderName)
		companyFile.Info.FileName = c.GetString("fileName")
		uploadedFileUrl :=c.GetString("downloadUrl")
		if len(uploadedFileUrl) !=0{
			thumParts := strings.Split(uploadedFileUrl, "/")
			thumParts = append(thumParts[:len(thumParts)-4], strings.Join(thumParts[len(thumParts)-4:], "%2F"))
			thumbPicture := strings.Join(thumParts, "/")
			companyFile.Info.DocumentUrl=thumbPicture
		}
		companyFile.Settings.DateOfCreation =time.Now().Unix()
		companyFile.Settings.Status = helpers.StatusActive

		dbStatus := companyFile.AddCompanyDocument(c.AppEngineCtx,companyId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}

	companyViewModel.CompanyTeamName = storedSession.CompanyTeamName
	companyViewModel.CompanyPlan = storedSession.CompanyPlan
	companyViewModel.AdminFirstName = storedSession.AdminFirstName
	companyViewModel.AdminLastName = storedSession.AdminLastName
	companyViewModel.ProfilePicture =storedSession.ProfilePicture
	companyViewModel.PageType=helpers.SelectPageForAdd

	c.Data["vm"] = companyViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/company-fileUpload.html"
}

func (c *CompanyFileUploadController) CompanyFileUploadEdit() {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	documentId := c.Ctx.Input.Param(":documentId")
	companyFile :=models.CompanyFileUpload{}
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	log.Println("companyId",companyId)
	companyViewModel := viewmodels.CompanyFileUpload{}
	if r.Method == "POST" {
		log.Println("insideeeeeeeeeee");
		companyFile.Info.FolderName = c.GetString("folderName")
		log.Println("folder name",companyFile.Info.FolderName)
		companyFile.Info.FileName = c.GetString("fileName")
		uploadedFileUrl :=c.GetString("downloadUrl")
		if len(uploadedFileUrl) !=0{
			thumParts := strings.Split(uploadedFileUrl, "/")
			thumParts = append(thumParts[:len(thumParts)-4], strings.Join(thumParts[len(thumParts)-4:], "%2F"))
			thumbPicture := strings.Join(thumParts, "/")
			companyFile.Info.DocumentUrl=thumbPicture
		}
		companyFile.Settings.DateOfCreation =time.Now().Unix()
		companyFile.Settings.Status = helpers.StatusActive

		dbStatus := companyFile.AddCompanyDocument(c.AppEngineCtx,companyId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}


	//get document details for edit
	dbStatus,documentDetails := models.GetCompanyDocumentById(c.AppEngineCtx,companyId,documentId)
	switch dbStatus {
	case true:
		companyViewModel.FolderName=documentDetails.Info.FolderName
		companyViewModel.FileName =documentDetails.Info.FileName
		companyViewModel.DownloadUrl =documentDetails.Info.DocumentUrl
		companyViewModel.PageType =helpers.SelectPageForEdit


	case false:
		log.Println("error")
	}

	//get folderName
	var folderName []string
	dbStatus, companyFileDetails := models.GetAllCompanyDocument(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:

		companyValue := reflect.ValueOf(companyFileDetails)
		for _, companyKey := range companyValue.MapKeys() {
			folderName=append(folderName,companyFileDetails[companyKey.String()].Info.FolderName)

		}
	case false:
		log.Println("error")
	}
	result := removeDuplicates(folderName)
	log.Println("result",result)
	companyViewModel.FolderNameArray=result
	companyViewModel.DocumentId=documentId


	companyViewModel.CompanyTeamName = storedSession.CompanyTeamName
	companyViewModel.CompanyPlan = storedSession.CompanyPlan
	companyViewModel.AdminFirstName = storedSession.AdminFirstName
	companyViewModel.AdminLastName = storedSession.AdminLastName
	companyViewModel.ProfilePicture =storedSession.ProfilePicture

	c.Data["vm"] = companyViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/company-fileUpload.html"
}
func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func (c *CompanyFileUploadController) CompanyFileUploadEditWithOutChange() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	documentId := c.Ctx.Input.Param(":documentId")
	companyFile :=models.CompanyFileUpload{}
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	log.Println("companyId",companyId)
	if r.Method == "POST" {
		log.Println("insideeeeeeeeeee");
		companyFile.Info.FolderName = c.GetString("folderName")
		log.Println("folder name",companyFile.Info.FolderName)
		companyFile.Info.FileName = c.GetString("fileName")
		companyFile.Settings.Status = helpers.StatusActive

		uploadedFileUrl :=c.GetString("downloadUrl")
		if len(uploadedFileUrl) !=0{
			thumParts := strings.Split(uploadedFileUrl, "/")
			thumParts = append(thumParts[:len(thumParts)-4], strings.Join(thumParts[len(thumParts)-4:], "%2F"))
			thumbPicture := strings.Join(thumParts, "/")
			companyFile.Info.DocumentUrl=thumbPicture
		}

		dbStatus := companyFile.EditCompanyIdWithoutChange(c.AppEngineCtx,companyId,documentId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}

}







func (c *CompanyFileUploadController) CompanyFileUploadDetail() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	//companyFile :=models.CompanyFileUpload{}
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	log.Println("companyId",companyId)
	companyViewModel := viewmodels.CompanyFileUpload{}
	dbStatus, companyFileDetails := models.GetAllCompanyDocument(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(companyFileDetails)
		for _, companyDocId := range dataValue.MapKeys() {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, companyFileDetails[companyDocId.String()].Info.FileName) // ("+companyFileDetails[companyDocId.String()].Info.FolderName+")")
			tempValueSlice =append(tempValueSlice,companyFileDetails[companyDocId.String()].Info.DocumentUrl)
			//tempValueSlice =append(tempValueSlice,companyDocId.String())
			companyViewModel.Values = append(companyViewModel.Values, tempValueSlice)
			companyViewModel.Keys =append(companyViewModel.Keys,companyDocId.String())
		}
	}
	companyViewModel.CompanyTeamName = storedSession.CompanyTeamName
	companyViewModel.CompanyPlan = storedSession.CompanyPlan
	companyViewModel.AdminFirstName = storedSession.AdminFirstName
	companyViewModel.AdminLastName = storedSession.AdminLastName
	companyViewModel.ProfilePicture =storedSession.ProfilePicture

	c.Data["vm"] = companyViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/company-fileUpload-details.html"
}

func (c *CompanyFileUploadController)CompanyFileUploadDelete() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	log.Println("inside delete")
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	documentId := c.Ctx.Input.Param(":documentId")
	dbStatus := models.DeleteCompanyDocument(c.AppEngineCtx,companyTeamName, documentId)
	log.Println("statusssssss", dbStatus)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}
}









