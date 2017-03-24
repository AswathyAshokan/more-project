
/* Author :Aswathy Ashok */

package controllers

import (
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
	"time"
	"fmt"
	"reflect"
	"app/passporte/helpers"
)

type ContactUserController struct {
	BaseController
}

/* Add contact detail to DB*/
func (c *ContactUserController)AddNewContact() {
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		user:=models.ContactUser{}
		user.Info.Name= c.GetString("name")
		user.Info.State = c.GetString("state")
		user.Info.ZipCode = c.GetString("zipcode")
		user.Info.Email = c.GetString("emailAddress")
		user.Info.PhoneNumber= c.GetString("phoneNumber")
		user.Info.Address = c.GetString("address")
		user.Settings.DateOfCreation =time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println(reflect.TypeOf(user.Settings.DateOfCreation))
		user.Settings.Status = helpers.StatusActive
		user.Info.CompanyTeamName = storedSession.CompanyTeamName
		dbStatus := user.AddContactToDB(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}else {
		viewModel := viewmodels.ContactUserViewModel{}
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.CompanyPlan = storedSession.CompanyPlan
		viewModel.PageType = helpers.SelectPageForAdd
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-contacts.html"
	}


}

/*Display all contact detail*/
func (c *ContactUserController)DisplayContactDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	dbStatus, contact := models.GetAllContact(c.AppEngineCtx,companyTeamName)
	viewModel := viewmodels.ContactUserViewModel{}

	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(contact)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())


		}
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, contact[k].Info.Name)
			tempValueSlice = append(tempValueSlice, contact[k].Info.Address)
			tempValueSlice = append(tempValueSlice, contact[k].Info.State)
			tempValueSlice = append(tempValueSlice, contact[k].Info.ZipCode)
			tempValueSlice = append(tempValueSlice, contact[k].Info.Email)
			tempValueSlice = append(tempValueSlice, contact[k].Info.PhoneNumber)
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		viewModel.CompanyTeamName = storedSession.CompanyTeamName
		viewModel.CompanyPlan = storedSession.CompanyPlan
		viewModel.Keys = keySlice
		viewModel.PageType=helpers.SelectPageForAdd
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/contacts-details.html"

	case false:

		log.Println(helpers.ServerConnectionError)
	}
}

/*Function for delete contact from DB*/
func (c *ContactUserController)LoadDeleteContact() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)


	contactId :=c.Ctx.Input.Param(":contactId")
	user := models.ContactUser{}
	dbStatus := user.DeleteContactFromDB(c.AppEngineCtx, contactId)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}


}
/*To perform edit operation*/
func (c *ContactUserController)LoadEditContact() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	if r.Method == "POST" {
		contactId := c.Ctx.Input.Param(":contactId")
		user:=models.ContactUser{}
		user.Info.Name= c.GetString("name")
		user.Info.State = c.GetString("state")
		user.Info.ZipCode = c.GetString("zipcode")
		user.Info.Email = c.GetString("emailAddress")
		user.Info.PhoneNumber= c.GetString("phoneNumber")
		user.Info.Address = c.GetString("address")
		user.Settings.DateOfCreation =time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println(reflect.TypeOf(user.Settings.DateOfCreation))
		user.Settings.Status = helpers.StatusActive
		user.Info.CompanyTeamName = storedSession.CompanyTeamName
		dbStatus := user.UpdateContactToDB(c.AppEngineCtx,contactId)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		contactId := c.Ctx.Input.Param(":contactId")
		viewModel := viewmodels.ContactUserViewModel{}
		contact :=models.ContactUser{}
		dbStatus,contact := contact.RetrieveContactIdFromDB(c.AppEngineCtx, contactId)
		switch dbStatus {
		case true:
			viewModel.PageType = helpers.SelectPageForEdit
			log.Println("page type",viewModel.PageType);
			viewModel.Name=contact.Info.Name
			viewModel.Address =contact.Info.Address
			viewModel.State =contact.Info.State
			viewModel.ZipCode =contact.Info.ZipCode
			viewModel.Email =contact.Info.Email
			viewModel.PhoneNumber =contact.Info.PhoneNumber
			viewModel.ContactId=contactId
			viewModel.CompanyTeamName = storedSession.CompanyTeamName
			viewModel.CompanyPlan = storedSession.CompanyPlan
			c.Data["vm"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-contacts.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}