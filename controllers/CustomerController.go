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
// add customer to database


func (c *CustomerController) AddCustomer() {
	customer := models.Customer{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter

	if r.Method == "POST" {

		customer.CustomerName = c.GetString("customername")
		context := appengine.NewContext(r)
		log.Infof(context, "Values of struct: %v", customer.CustomerName )
		customer.ContactPerson = c.GetString("contactperson")
		customer.Address = c.GetString("address")
		customer.Phone = c.GetString("phone")
		customer.Email = c.GetString("email")
		customer.State = c.GetString("state")
		customer.ZipCode = c.GetString("zipcode")
		log.Infof(context, "Values of struct: %v", customer)
		dbStatus := customer.AddToDb(c.AppEngineCtx)
		log.Infof(context, "fafh",dbStatus )

		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))



		}

		//http.Redirect(w, r, "/customer", 301)


	} else {
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-customer.html"

	}
}
// view details of customer from database


func (c *CustomerController) CustomerDetails() {
	r := c.Ctx.Request
	NewContext := appengine.NewContext(r)

	customer := models.Customer{}
	CustomerInfo := customer.DisplayCustomer(c.AppEngineCtx)
	CustomerDataValue := reflect.ValueOf(CustomerInfo)
	var CustomerValueSlice []models.Customer // to store data values from slice
	CustomerViewModel := viewmodels.Customer{}
	var CustomerKeySlice []string	// to store the key of a slice
	for _, CustomerKey := range CustomerDataValue.MapKeys() {
		CustomerKeySlice = append(CustomerKeySlice, CustomerKey.String())//to get keys
		CustomerValueSlice = append(CustomerValueSlice, CustomerInfo[CustomerKey.String()])//to get values
		CustomerViewModel.Customers = append(CustomerViewModel.Customers, CustomerInfo[CustomerKey.String()])

	}
	CustomerViewModel.CustomerKey = CustomerKeySlice
	log.Infof(NewContext, "key of", CustomerViewModel)
	c.Data["vm"] = CustomerViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/customer-details.html"
}



// delete each customer


func (c *CustomerController) DeleteCustomer() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	customerKey:=c.Ctx.Input.Param(":customerkey")
	NewContext := appengine.NewContext(r)
	customer := models.Customer{}
	DbStatus :=customer.DeleteCustomer(c.AppEngineCtx,customerKey)
	switch DbStatus {
	case true:
		http.Redirect(w, r, "/customer", 301)
	case false:
		log.Infof(NewContext,"failed")

	}



}


//edit profile of each users

func (c *CustomerController) EditCustomer() {
	customer := models.Customer{}
	r := c.Ctx.Request
	key:=c.Ctx.Input.Param(":customerkey")
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

func (c *CustomerController) ViewCustomer() {
	r := c.Ctx.Request
	//var Key int
	Groupkey := c.Ctx.Input.Param(":customerkey")
	exam := appengine.NewContext(r)
	log.Infof(exam, "iddddddddd: %v", Groupkey)
}


