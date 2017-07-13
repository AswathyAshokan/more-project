package controllers

import (

	"app/passporte/models"
	"app/passporte/viewmodels"
	"strings"
	//"app/go_appengine/goroot/src/log"
	"app/passporte/helpers"
	"reflect"
	"log"
)

type FitToWorkController struct {
	BaseController
}
func (c *FitToWorkController)AddNewFitToWork() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	fitToWorkView := viewmodels.FitToWork{}
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	fitToWorkView.CompanyTeamName = storedSession.CompanyTeamName
	fitToWorkView.CompanyPlan = storedSession.CompanyPlan
	fitToWorkView.AdminLastName = storedSession.AdminLastName
	fitToWorkView.AdminFirstName = storedSession.AdminFirstName
	fitToWorkView.ProfilePicture = storedSession.ProfilePicture
	fitToWorkData := models.FitToWork{}
	if r.Method == "POST" {
		fitToWorkData.FitToWorkName = c.GetString("fitWorkName")
		fitToWorkData.Settings.Status =helpers.StatusActive
		instructions := c.GetString("instructionsForUser")
		instructionsFromUser := strings.Split(instructions, "/@@,")
		sliceLastValue := instructionsFromUser[len(instructionsFromUser)-1]
		SliceLastValuesWithOutAnySymbol := strings.Split(sliceLastValue, "/@@")
		instructionsFromUser = instructionsFromUser[:len(instructionsFromUser)-1]
		instructionsFromUser = append(instructionsFromUser, SliceLastValuesWithOutAnySymbol[0])
		instructionSlice := instructionsFromUser
		log.Println("instruction",instructionSlice)
		dbStatus := fitToWorkData.AddFitToWorkToDb(c.AppEngineCtx, instructionSlice, companyTeamName)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	} else {

		c.Data["vm"] = fitToWorkView
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-fit-work.html"
	}


}
func (c* FitToWorkController)LoadFitToWork(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	dbStatus, allFitToWork := models.GetAllFitToWorkDetails(c.AppEngineCtx)
	fitToWorkViewModel :=viewmodels.FitToWork{}
	switch dbStatus {
	case true:
		var keySlice []string
		var tempKeySlice []string
		dataValue := reflect.ValueOf(allFitToWork)
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k :=range keySlice {
			if k == companyTeamName {

				fitToWorkById := models.GetSelectedCompanyName(c.AppEngineCtx, k)
				fitToWorkDataValues := reflect.ValueOf(fitToWorkById)
				for _, fitToWorkKey := range fitToWorkDataValues.MapKeys() {
					tempKeySlice = append(tempKeySlice, fitToWorkKey.String())
				}
				for _, eachKey := range tempKeySlice {
					log.Println("key", eachKey)
					var tempValueSlice []string

					if fitToWorkById[eachKey].Settings.Status == helpers.StatusActive {
						tempValueSlice = append(tempValueSlice, "")
						tempValueSlice = append(tempValueSlice, fitToWorkById[eachKey].FitToWorkName)
						tempValueSlice = append(tempValueSlice, eachKey)
						fitToWorkViewModel.Values = append(fitToWorkViewModel.Values, tempValueSlice)
						tempValueSlice = tempValueSlice[:0]

						getInstructions := models.GetAllInstructionsOfFitToWorkById(c.AppEngineCtx, k, eachKey)
						log.Println("getInstructions", getInstructions)
						for _, instructionKey := range reflect.ValueOf(getInstructions).MapKeys() {
							var fitToWorkVM viewmodels.FitToWorkStruct
							var instructionKeySlice []string
							instructionKeyString := instructionKey.String()
							fitToWorkVM.InstructionKey = eachKey
							instructionKeySlice = append(instructionKeySlice, instructionKeyString)
							fitToWorkVM.Description = getInstructions[instructionKeyString].Description
							fitToWorkViewModel.InnerContent = append(fitToWorkViewModel.InnerContent, fitToWorkVM)
						}
					}

				}
			}
		}
		fitToWorkViewModel.Keys = keySlice
		fitToWorkViewModel.CompanyTeamName = storedSession.CompanyTeamName
		fitToWorkViewModel.CompanyPlan = storedSession.CompanyPlan
		fitToWorkViewModel.AdminFirstName = storedSession.AdminFirstName
		fitToWorkViewModel.AdminLastName = storedSession.AdminLastName
		fitToWorkViewModel.ProfilePicture =storedSession.ProfilePicture
		fitToWorkViewModel.CompanyTeamName = storedSession.CompanyTeamName
		c.Data["vm"] = fitToWorkViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/fit-to-work-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

func (c *FitToWorkController) EditFitToWork() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	fitToWorkId := c.Ctx.Input.Param(":fitToWorkId")
	storedSession := ReadSession(w, r, companyTeamName)
	fitToWorkData := models.FitToWork{}
	fitToWorkView :=viewmodels.EditFitToWork{}
	if r.Method == "POST" {

		fitToWorkData.FitToWorkName = c.GetString("fitWorkName")
		fitToWorkData.Settings.Status =helpers.StatusActive
		instructions := c.GetString("instructionsForUser")
		instructionsFromUser := strings.Split(instructions, "/@@,")
		sliceLastValue := instructionsFromUser[len(instructionsFromUser)-1]
		SliceLastValuesWithOutAnySymbol := strings.Split(sliceLastValue, "/@@")
		instructionsFromUser = instructionsFromUser[:len(instructionsFromUser)-1]
		instructionsFromUser = append(instructionsFromUser, SliceLastValuesWithOutAnySymbol[0])
		instructionSlice := instructionsFromUser
		log.Println("instruction",instructionSlice)
		dbStatus := fitToWorkData.UpdateFitToWorkToDb(c.AppEngineCtx, instructionSlice, companyTeamName,fitToWorkId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	}else {
		var Instructions []string
		fitToWorkDetails :=models.GetEachFitToWorkByCompanyId(c.AppEngineCtx, fitToWorkId,companyTeamName)
		allInstructions := models.GetAllInstructionsFromFitToWork(c.AppEngineCtx,fitToWorkId,companyTeamName)
		dataValueOfInstruction := reflect.ValueOf(allInstructions)
		for _, instructionKey:=range dataValueOfInstruction.MapKeys(){
			Instructions = append(Instructions,allInstructions[instructionKey.String()].Description)
		}
		fitToWorkView.InstructionArrayToEdit = Instructions
		fitToWorkView.FitToWorkName = fitToWorkDetails.FitToWorkName
		fitToWorkView.FitToWorkId  = fitToWorkId
		fitToWorkView.CompanyTeamName = storedSession.CompanyTeamName
		fitToWorkView.CompanyPlan   =  storedSession.CompanyPlan
		fitToWorkView.AdminLastName =storedSession.AdminLastName
		fitToWorkView.AdminFirstName =storedSession.AdminFirstName
		fitToWorkView.ProfilePicture =storedSession.ProfilePicture
		fitToWorkView.PageType=helpers.SelectPageForEdit


	}
	c.Data["vm"] = fitToWorkView
	c.Layout = "layout/layout.html"
	c.TplName = "template/add-fit-work.html"

}

func (c *FitToWorkController) DeleteFitToWork() {
	log.Println("hhhooooooo")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
	fitToWorkId :=c.Ctx.Input.Param(":fitToWorkId")
	dbStatus :=models.DeleteFitToWorkById(c.AppEngineCtx, fitToWorkId,companyTeamName)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}