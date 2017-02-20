/*Created By Farsana*/
package controllers

import (
	"app/passporte/models"
	"log"
	"reflect"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"time"
	"strings"
)

type CustomerController struct {
	BaseController
}

// add new customer to database
func (c *CustomerController) AddCustomer() {
	customer := models.Customers{}
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		customer.Info.CustomerName = c.GetString("customername")
		customer.Info.ContactPerson = c.GetString("contactperson")
		customer.Info.Address = c.GetString("address")
		customer.Info.Phone = c.GetString("phone")
		customer.Info.Email = c.GetString("email")
		customer.Info.State = c.GetString("state")
		customer.Info.ZipCode = c.GetString("zipcode")
		customer.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		customer.Settings.Status = "inactive"
		dbStatus := customer.AddCustomersToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-customer.html"
	}
}

//Display all the details of customer
func (c *CustomerController) CustomerDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("The userDetails stored in session:",storedSession)
	CustomerViewModel := viewmodels.Customer{}
	allCustomer,dbStatus:= models.GetAllCustomerDetails(c.AppEngineCtx)
	log.Println("view",allCustomer)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allCustomer)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.CustomerName)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Address)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.State)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.ZipCode)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Email)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Phone)
			tempValueSlice = append(tempValueSlice, allCustomer[k].Info.ContactPerson)
			CustomerViewModel.Values=append(CustomerViewModel.Values,tempValueSlice)
			tempValueSlice = tempValueSlice[:0]
		}
		CustomerViewModel.Keys = keySlice
		c.Data["vm"] = CustomerViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/customer-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

// delete each customer using customer id
func (c *CustomerController) DeleteCustomer() {
	w := c.Ctx.ResponseWriter
	customerKey :=c.Ctx.Input.Param(":customerid")
	customer := models.Customers{}
	dbStatus :=customer.DeleteCustomerById(c.AppEngineCtx, customerKey)
	switch dbStatus {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))
	}
}

//edit profile of each users using customer id
func (c *CustomerController) EditCustomer() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	customerDetails := models.CustomerData{}
	customer := models.Customers{}
	customerId := c.Ctx.Input.Param(":customerid")
	if r.Method == "POST" {
		customerDetails.CustomerName = c.GetString("customername")
		customerDetails.Address = c.GetString("address")
		customerDetails.ContactPerson = c.GetString("contactperson")
		customerDetails.Email= c.GetString("email")
		customerDetails.Phone = c.GetString("phone")
		customerDetails.ZipCode = c.GetString("zipcode")
		customerDetails.State = c.GetString("state")

		log.Println("new name", customerDetails.CustomerName)
		dbStatus := customer.UpdateCustomerDetailsById(c.AppEngineCtx, customerId)

		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		editResult, DbStatus := customer.EditCustomer(c.AppEngineCtx, customerId)
		switch DbStatus {
		case true:
			customerViewModel := viewmodels.EditCustomerViewModel{}
			customerViewModel.State= editResult.Info.State
			customerViewModel.ZipCode = editResult.Info.ZipCode
			customerViewModel.Email = editResult.Info.Email
			customerViewModel.ContactPerson = editResult.Info.ContactPerson
			customerViewModel.Address = editResult.Info.Address
			customerViewModel.CustomerName= editResult.Info.CustomerName
			customerViewModel.Phone= editResult.Info.Phone
			customerViewModel.PageType = helpers.SelectPageForEdit
			customerViewModel.CustomerId = customerId
			c.Data["vm"] = customerViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-customer.html"
		case false:
			log.Println(helpers.ServerConnectionError)
		}
	}
}
func (c *CustomerController)  CustomerNameCheck(){
	w := c.Ctx.ResponseWriter
	customerName := c.GetString("customername")
	pageType := c.Ctx.Input.Param(":type")
	oldName := c.Ctx.Input.Param(":oldName")
	if pageType == "edit" && strings.Compare(oldName, customerName) == 0 {
		w.Write([]byte("true"))
	} else {
		dbStatus := models.IsCustomerNameUsed(c.AppEngineCtx,customerName)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}



}

