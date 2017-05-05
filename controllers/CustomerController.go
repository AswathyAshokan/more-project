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
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	customer := models.Customers{}
	addViewModel := viewmodels.AddCustomerViewModel{}
	if r.Method == "POST" {
		customer.Info.CustomerName = c.GetString("customername")
		customer.Info.ContactPerson = c.GetString("contactperson")
		customer.Info.Address = c.GetString("address")
		customer.Info.Phone = c.GetString("phone")
		customer.Info.Email = c.GetString("email")
		customer.Info.State = c.GetString("state")
		customer.Info.ZipCode = c.GetString("zipcode")
		customer.Info.CompanyTeamName = storedSession.CompanyTeamName
		customer.Settings.DateOfCreation =(time.Now().UnixNano() / 1000000)
		customer.Settings.Status = helpers.StatusActive
		dbStatus := customer.AddCustomersToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}

	} else {
		addViewModel.CompanyTeamName = storedSession.CompanyTeamName
		addViewModel.CompanyPlan   =  storedSession.CompanyPlan
		addViewModel.AdminLastName =storedSession.AdminLastName
		addViewModel.AdminFirstName =storedSession.AdminFirstName
		addViewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = addViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/add-customer.html"
	}
}

//Display all the details of customer
func (c *CustomerController) CustomerDetails() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	customerViewModel := viewmodels.Customer{}
	allCustomer,dbStatus:= models.GetAllCustomerDetails(c.AppEngineCtx,companyTeamName)
	switch dbStatus {
	case true:
		dataValue := reflect.ValueOf(allCustomer)
		var keySlice []string
		for _, key := range dataValue.MapKeys() {
			keySlice = append(keySlice, key.String())
		}
		for _, k := range keySlice {
			var tempValueSlice []string
			if allCustomer[k].Settings.Status != helpers.UserStatusDeleted{
				tempValueSlice = append(tempValueSlice, allCustomer[k].Info.CustomerName)
				tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Address)
				tempValueSlice = append(tempValueSlice, allCustomer[k].Info.State)
				tempValueSlice = append(tempValueSlice, allCustomer[k].Info.ZipCode)
				tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Email)
				tempValueSlice = append(tempValueSlice, allCustomer[k].Info.Phone)
				tempValueSlice = append(tempValueSlice, allCustomer[k].Info.ContactPerson)
				tempValueSlice = append(tempValueSlice,k)
				customerViewModel.Values=append(customerViewModel.Values,tempValueSlice)
				tempValueSlice = tempValueSlice[:0]
			}

		}
		customerViewModel.Keys = keySlice
		customerViewModel.CompanyTeamName = storedSession.CompanyTeamName
		customerViewModel.CompanyPlan = storedSession.CompanyPlan
		customerViewModel.AdminFirstName =storedSession.AdminFirstName
		customerViewModel.AdminLastName =storedSession.AdminLastName
		customerViewModel.ProfilePicture =storedSession.ProfilePicture
		c.Data["vm"] = customerViewModel
		c.Layout = "layout/layout.html"
		c.TplName = "template/customer-details.html"
	case false:
		log.Println(helpers.ServerConnectionError)
	}
}

// delete each customer using customer id
func (c *CustomerController) DeleteCustomer() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	ReadSession(w, r, companyTeamName)
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
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	customer := models.Customers{}
	customerId := c.Ctx.Input.Param(":customerid")
	log.Println("customerId",customerId)
	if r.Method == "POST" {
		customer.Info.CustomerName = c.GetString("customername")
		customer.Info.Address = c.GetString("address")
		customer.Info.ContactPerson = c.GetString("contactperson")
		customer.Info.Email= c.GetString("email")
		customer.Info.Phone = c.GetString("phone")
		customer.Info.ZipCode = c.GetString("zipcode")
		customer.Info.State = c.GetString("state")
		customer.Info.CompanyTeamName = storedSession.CompanyTeamName
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
			customerViewModel.CompanyTeamName = storedSession.CompanyTeamName
			customerViewModel.CompanyPlan = storedSession.CompanyPlan
			customerViewModel.AdminLastName =storedSession.AdminLastName
			customerViewModel.AdminFirstName =storedSession.AdminFirstName
			customerViewModel.ProfilePicture =storedSession.ProfilePicture
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

