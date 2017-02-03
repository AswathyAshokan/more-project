
/* Author :Aswathy Ashok */

package controllers

import (

	//"github.com/astaxie/beegae"
	"app/passporte/models"
	"app/passporte/viewmodels"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"time"
	"fmt"
	"reflect"

	//"app/go_appengine/goroot/src/go/doc/testdata"
	//"github.com/gorilla/mux"
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
		context := appengine.NewContext(r)
		log.Infof(context, "requested struct: %+v", user)
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
		r := c.Ctx.Request
		ce := appengine.NewContext(r)
		log.Infof(ce, "%s\n", contact)
		//var valueSlice []models.User
		dataValue := reflect.ValueOf(contact)
		var keySlice []string
		var valueSlice []models.ContactUser
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())


		}
		// To perform the opertion you want
		for _, k := range keySlice {
			valueSlice = append(valueSlice, contact[k])
			viewModel.User = append(viewModel.User, contact[k])
			viewModel.Key=keySlice



		}

		log.Infof(ce,"Key:", keySlice, "Value:", valueSlice)
		//log.Infof(ce,"Value: ", valueSlice)
		//log.Infof(ce,"Value: ", valueSlice)
		//mvVar := map["Name"].(string)
		//m := f.(map[string]interface{}
		//viewModel.Name = contact[result[i]].Name
		//viewModel.PhoneNumber = contact["PhoneNumber"]
		//viewModel.Email = contact["Email"]
		//viewModel.Address = contact["Address"]
		//viewModel.State = contact["State"]
		//viewModel.ZipCode = contact["ZipCode"]
		log.Infof(ce, "typeeee",(reflect.TypeOf(viewModel)))
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/contacts-details.html"

	case false:

	}
}
func (c *ContactUserController)LoadDeleteContact() {

	r := c.Ctx.Request
	context := appengine.NewContext(r)
	contactId :=c.Ctx.Input.Param(":contactId")
	log.Infof(context,"idddddddddd", contactId)
	user := models.ContactUser{}
	dbStatus := user.DeleteContactFromDB(c.AppEngineCtx, contactId)

	switch dbStatus {

	case true:
		c.Redirect("/contact", 302)
	case false :
	}


}
func (c *ContactUserController)LoadEditContact() {
	r := c.Ctx.Request
	context := appengine.NewContext(r)
	contactId :=c.Ctx.Input.Param(":contactId")
	log.Infof(context,"idddddddddd", contactId)
	viewModel := viewmodels.ContactUserViewModel{}
	user := models.ContactUser{}
	dbStatus,contact := user.RetrieveContactIdFromDB(c.AppEngineCtx, contactId)
	switch dbStatus {
	case true:
		viewModel.Name=contact.Name
		viewModel.Address =contact.Address
		viewModel.State =contact.State
		viewModel.ZipCode =contact.Zipcode
		viewModel.Email =contact.Email
		viewModel.PhoneNumber =contact.PhoneNumber
		c.Data["vm"] = viewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-contactsnew.html"
	case false:

	}

}
