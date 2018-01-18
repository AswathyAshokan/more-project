/*Created By Aswathy*/

//created for displaying company user details
package controllers

import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"log"
)

type CompanyProfileController struct {
	BaseController
}


// fetch all the details of invite user from database
func (c *CompanyProfileController) CompanyProfileDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	userDetail :=models.Users{}
	storedSession := ReadSession(w, r, companyTeamName)
	companyId := storedSession.CompanyId
	log.Println("companyId",companyId)
	companyViewModel := viewmodels.CompanyProfile{}
	dbStatus,usersArray,KeyArray,userExpand,nextOfKin := userDetail.GetAllUsersDetails(c.AppEngineCtx,companyId)
	switch dbStatus {
	case true:
		log.Println("users details",usersArray)
		log.Println("key Array",KeyArray)
		companyViewModel.Values =usersArray
		companyViewModel.Keys =KeyArray
		companyViewModel.UserExpand =userExpand
		companyViewModel.NextOfKin =nextOfKin

	case false:
		log.Println(helpers.ServerConnectionError)
	}



	companyViewModel.CompanyTeamName = storedSession.CompanyTeamName
	companyViewModel.CompanyPlan = storedSession.CompanyPlan
	companyViewModel.AdminFirstName = storedSession.AdminFirstName
	companyViewModel.AdminLastName = storedSession.AdminLastName
	companyViewModel.ProfilePicture =storedSession.ProfilePicture

	c.Data["vm"] = companyViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/company-profile.html"
}






