/*Created By Farsana*/
package controllers

import (
	//"fmt"
	"app/passporte/models"
	//"app/passporte/viewmodels"
	"log"
	"reflect"
	"app/passporte/viewmodels"

	"app/passporte/helpers"
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

		customer.ContactPerson = c.GetString("contactperson")
		customer.Address = c.GetString("address")
		customer.Phone = c.GetString("phone")
		customer.Email = c.GetString("email")
		customer.State = c.GetString("state")
		customer.ZipCode = c.GetString("zipcode")

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
	customer := models.Customer{}
	CustomerViewModel := viewmodels.Customer{}
	info := customer.GetAllCustomerDetails(c.AppEngineCtx)
	dataValue := reflect.ValueOf(info)
	var keySlice []string
	for _, key := range dataValue.MapKeys() {
		keySlice = append(keySlice, key.String())
	}


	for _, k := range keySlice {
		var tempValueSlice []string
		tempValueSlice = append(tempValueSlice, info[k].CustomerName)
		tempValueSlice = append(tempValueSlice, info[k].Address)
		tempValueSlice = append(tempValueSlice, info[k].State)
		tempValueSlice = append(tempValueSlice, info[k].ZipCode)
		tempValueSlice = append(tempValueSlice, info[k].Email)
		tempValueSlice = append(tempValueSlice, info[k].Phone)
		tempValueSlice = append(tempValueSlice, info[k].ContactPerson)
		CustomerViewModel.Values=append(CustomerViewModel.Values,tempValueSlice)
		tempValueSlice = tempValueSlice[:0]
	}

	CustomerViewModel.Keys = keySlice
	c.Data["vm"] = CustomerViewModel
	c.Layout = "layout/layout.html"
	c.TplName = "template/customer-details.html"
}

// delete each customer

func (c *CustomerController) DeleteCustomer() {
	w := c.Ctx.ResponseWriter
	customerKey :=c.Ctx.Input.Param(":customerid")

	customer := models.Customer{}
	result :=customer.DeleteCustomerById(c.AppEngineCtx, customerKey)
	switch result {
	case true:
		w.Write([]byte("true"))
	case false:
		w.Write([]byte("false"))

	}

}

//edit profile of each users

func (c *CustomerController) EditCustomer() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	customer := models.Customer{}
	customerId := c.Ctx.Input.Param(":customerid")

	if r.Method == "POST" {

		customer.CustomerName = c.GetString("customername")
		customer.Address = c.GetString("address")
		customer.ContactPerson = c.GetString("contactperson")
		customer.Email= c.GetString("email")
		customer.Phone = c.GetString("phone")
		customer.ZipCode = c.GetString("zipcode")
		customer.State = c.GetString("state")
		log.Println("new name",customer.CustomerName)
		dbStatus :=customer.UpdateCustomerDetailsById(c.AppEngineCtx, customerId)

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
			customerViewModel := viewmodels.Customer{}
			customerViewModel.State= editResult.State
			customerViewModel.ZipCode = editResult.ZipCode
			customerViewModel.Email = editResult.Email
			customerViewModel.ContactPerson = editResult.ContactPerson
			customerViewModel.Address = editResult.Address
			customerViewModel.CustomerName= editResult.CustomerName
			customerViewModel.Phone= editResult.Phone
			customerViewModel.PageType = helpers.SelectPageForEdit
			customerViewModel.CustomerId = customerId
			c.Data["vm"] = customerViewModel
			c.Layout = "layout/layout.html"
			c.TplName = "template/add-customer.html"
		case false:


		}
	}

}









