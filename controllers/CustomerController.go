/*Created By Farsana*/
package controllers

import (
	//"fmt"
	"app/passporte/models"
	//"app/passporte/viewmodels"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"reflect"
	"app/passporte/viewmodels"
	"net/http"
)


type CustomerController struct {
	BaseController
}

func (c *CustomerController) AddCustomer() {
	customer := models.Customer{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {

		customer.CustomerName = c.GetString("customername")
		exam := appengine.NewContext(r)
		log.Infof(exam, "Values of struct: %v", customer.CustomerName )
		customer.ContactPerson = c.GetString("contactperson")
		customer.Address = c.GetString("address")
		customer.Phone = c.GetString("phone")
		customer.Email = c.GetString("email")
		customer.State = c.GetString("state")
		customer.ZipCode = c.GetString("zipcode")
		exam = appengine.NewContext(r)
		log.Infof(exam, "Values of struct: %v", customer)
		customer.AddToDb(c.AppEngineCtx)
		http.Redirect(w, r, "/customer-details.html", 302)

	} else {
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-customer.html"

	}
}

func (c *CustomerController) CustomerDetails() {
	r := c.Ctx.Request
	exam := appengine.NewContext(r)

	customer := models.Customer{}
	result := customer.DisplayCustomer(c.AppEngineCtx)
	dataValue := reflect.ValueOf(result)
	var valueSlice []models.Customer
	viewmodel := viewmodels.Customer{}
	var keySlice []string
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())//to get keys
		valueSlice = append(valueSlice, result[key.String()])//to get values
		viewmodel.Customers = append(viewmodel.Customers, result[key.String()])

	}
	viewmodel.Key=keySlice
	log.Infof(exam, "key of",viewmodel)
	c.Data["vm"] = viewmodel
	c.Layout = "layout/layout.html"
	c.TplName = "template/customer-details.html"
}



// delete each customer


func (c *CustomerController) CustomerDelete() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	key:=c.Ctx.Input.Param(":Key")
	exam := appengine.NewContext(r)
	customer := models.Customer{}
	result :=customer.DeleteCustomer(c.AppEngineCtx,key)
	switch result {
	case true:
		http.Redirect(w, r, "/customer-details", 302)
	case false:
		log.Infof(exam,"failed")

	}



}


//edit profile of each users

func (c *CustomerController) CustomerEdit() {
	customer := models.Customer{}
	r := c.Ctx.Request
	key:=c.Ctx.Input.Param(":Key")
	exam := appengine.NewContext(r)
	result,DbStatus :=customer.EditCustomer(c.AppEngineCtx,key)
	switch DbStatus {
	case true:
		viewmodel := viewmodels.Customer{}
		viewmodel.CustomerName = result.CustomerName
		viewmodel.ContactPerson = result.ContactPerson
		//viewmodel.EmailId = result.EmailId
		//viewmodel.UserType = result.UserType
		c.Data["vm"] = viewmodel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-user.html"
	case false:
		log.Infof(exam,"failed")

	}
	log.Infof(exam,"jhfjsgjgj: %+v",result)
}

//view the user

func (c *CustomerController) CustomerView() {
	r := c.Ctx.Request
	//var Key int
	key := c.Ctx.Input.Param(":Key")
	exam := appengine.NewContext(r)
	log.Infof(exam, "iddddddddd: %v", key)
}


