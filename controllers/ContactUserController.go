
/* Author :Aswathy Ashok */

package controllers

import (

	//"github.com/astaxie/beegae"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"log"
	"time"
	"fmt"
	"reflect"

	//"app/go_appengine/goroot/src/go/doc/testdata"
	//"github.com/gorilla/mux"
	"app/passporte/helpers"
)

type ContactUserController struct {
	BaseController
}

func (c *ContactUserController)LoadContact() {
	r := c.Ctx.Request
	w :=c.Ctx.ResponseWriter
	if r.Method == "POST" {

		user:=models.ContactUser{}
		user.Name= c.GetString("name")
		user.State = c.GetString("state")
		user.Zipcode = c.GetString("zipcode")
		user.Email = c.GetString("emailAddress")
		user.PhoneNumber= c.GetString("phoneNumber")
		user.Address = c.GetString("address")
		user.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println(reflect.TypeOf(user.CurrentDate))
		user.Status = "Completed"
		dbStatus := user.AddContactToDB(c.AppEngineCtx)
		switch dbStatus {

		case true:
			w.Write([]byte("true"))

		case false:
			w.Write([]byte("false"))
		}
	}else {

		c.Layout = "layout/layout.html"
		c.TplName = "template/add-contacts.html"


	}


}
func (c *ContactUserController)LoadContactdetail() {
	user := models.ContactUser{}
	dbStatus, contact := user.RetrieveContactFromDB(c.AppEngineCtx)
	viewModel := viewmodels.ContactUserViewModel{}

	switch dbStatus {

	case true:
		//var valueSlice []models.User
		dataValue := reflect.ValueOf(contact)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())


		}
		// To perform the opertion you want
		for _, k := range keySlice {
			var tempValueSlice []string
			log.Println("hai")
			tempValueSlice = append(tempValueSlice, contact[k].Name)
			tempValueSlice = append(tempValueSlice, contact[k].Address)
			tempValueSlice = append(tempValueSlice, contact[k].State)
			tempValueSlice = append(tempValueSlice, contact[k].Zipcode)
			tempValueSlice = append(tempValueSlice, contact[k].Email)
			tempValueSlice = append(tempValueSlice, contact[k].PhoneNumber)
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		viewModel.Keys = keySlice
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/contacts-details.html"

	case false:

	}
}
func (c *ContactUserController)LoadDeleteContact() {


	contactId :=c.Ctx.Input.Param(":contactId")
	user := models.ContactUser{}
	dbStatus := user.DeleteContactFromDB(c.AppEngineCtx, contactId)
	w := c.Ctx.ResponseWriter
	switch dbStatus {

	case true:
		w.Write([]byte("true"))
	case false :
		w.Write([]byte("false"))
	}


}

func (c *ContactUserController)LoadEditContact() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		contactId := c.Ctx.Input.Param(":contactId")
		user:=models.ContactUser{}
		user.Name= c.GetString("name")
		user.State = c.GetString("state")
		user.Zipcode = c.GetString("zipcode")
		user.Email = c.GetString("emailAddress")
		user.PhoneNumber= c.GetString("phoneNumber")
		user.Address = c.GetString("address")
		user.CurrentDate =time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println(reflect.TypeOf(user.CurrentDate))
		user.Status = "Completed"
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
			viewModel.Name=contact.Name
			viewModel.Address =contact.Address
			viewModel.State =contact.State
			viewModel.ZipCode =contact.Zipcode
			viewModel.Email =contact.Email
			viewModel.PhoneNumber =contact.PhoneNumber
			viewModel.ContactId=contactId
			c.Data["array"] = viewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-contacts.html"
		case false:

		}

	}
}